import time
from typing import List, Optional, Tuple, Dict, Callable

import smart_open
from packageurl import PackageURL

from common import HttpStatus, PyLang
from util import spider, Singleton
from util.requirements_detector import from_setup_cfg, from_setup_py, from_requirements_txt, \
    from_pyproject_toml, DetectedRequirement
from vo import ModuleDeps, Dep, PackageDeps, ModuleMeta
from .index import DepsHandler, ArchiveDepsHandler


def _get_purl(artifact: str, version: Optional[str] = None) -> str:
    return PackageURL(type='pypi', name=artifact, version=version).to_string()


def _standardize(name: str) -> str:
    return name.lower().replace('-', '_')


@Singleton
class PyDepsHandler(PyLang, DepsHandler):
    PY_PARSE_MAP: Dict[str, Callable[[str], Tuple[str, str, List[DetectedRequirement]]]] = {
        PyLang.MODULE_SETUP_PY: from_setup_py,
        PyLang.MODULE_SETUP_CFG: from_setup_cfg,
        PyLang.MODULE_REQUIRES: from_requirements_txt,
        PyLang.MODULE_PYPROJECT: from_pyproject_toml,
    }

    def deps(self, module: str, data: bytes) -> Tuple[Optional[ModuleDeps], HttpStatus]:
        module_deps: List[Dep] = []
        for k in self.PY_PARSE_MAP:
            if module.endswith(k):
                artifact, version, reqs = self.PY_PARSE_MAP[k](data.decode())
                for r in reqs:
                    if len(r.version_specs) == 0:
                        module_deps.append(Dep(purl=_get_purl(r.name)))
                        continue
                    for v in r.version_specs:
                        module_deps.append(Dep(purl=_get_purl(r.name, v[1]),
                                               limit=v[0]))
                return ModuleDeps(lang=self.lang(), path=module,
                                  purl=_get_purl(artifact, version),
                                  dependencies=module_deps), HttpStatus.OK
        return None, HttpStatus.NOT_SUPPORT

    @staticmethod
    def get_python_package(artifact: str, version: Optional[str]) -> Optional[bytes]:
        if version is None:
            res = spider(f'https://pypi.org/pypi/{artifact}/json').json()
            version = res.get('info').get('version')
        res = spider(f'https://pypi.org/pypi/{artifact}/{version}/json').json()
        # 筛选源文件链接
        urls = list(filter(
            lambda r: not r.get('filename').endswith('whl') and
                      (r.get('filename').startswith(f'{_standardize(artifact)}-{version}.') or
                       r.get('filename').startswith(f'{artifact}-{version}.')),
            res.get('urls', [])))
        if len(urls) == 0:
            return None
        file_url = urls[0].get('url')
        while True:
            try:
                # 优先使用镜像
                with smart_open.open(
                        file_url.replace('https://files.pythonhosted.org/', 'https://pypi.tuna.tsinghua.edu.cn/'),
                        mode='rb') as f:
                    return f.read()
            except Exception:
                time.sleep(1)
            try:
                with smart_open.open(file_url, mode='rb') as f:
                    return f.read()
            except Exception:
                time.sleep(1)

    def meta(self, lang: str, purl: str) -> Tuple[Optional[ModuleMeta], HttpStatus]:
        def safe_get(j, k, default):
            v = j.get(k, default)
            if v is None:
                return default
            return v

        purl = PackageURL.from_string(purl)
        artifact = purl.name
        version = purl.version
        if version is None:
            res = spider(f'https://pypi.org/pypi/{artifact}/json').json()
            version = res.get('info').get('version')
        res = spider(f'https://pypi.org/pypi/{artifact}/{version}/json').json()
        # 获取元数据
        info = res.get('info', {})
        if info is None:
            return None, HttpStatus.NOT_FOUND
        # 1. 获取主页
        pages = sorted(list(filter(lambda url: url is not None,
                                   {*safe_get(info, 'project_urls', {}).values(), info.get('home_page'),
                                    info.get('project_url'),
                                    info.get('package_url')})),
                       key=lambda url: ('github.com' in url, 'pypi.org' in url, -len(url)), reverse=True)
        homepage = ''
        if len(pages) > 0:
            homepage = pages[0]
        # 2. 获取上传时间
        urls = res.get('urls', [])
        if len(urls) == 0:
            return None, HttpStatus.NOT_FOUND
        update_time = urls[0].get('upload_time_iso_8601')
        # 3. 获取摘要
        summary = info.get('summary', '')
        return ModuleMeta(desc=summary, homepage=homepage, upload_time=update_time), HttpStatus.OK

    def search(self, lang: str, package: str) -> Tuple[Optional[ModuleDeps], HttpStatus]:
        purl = PackageURL.from_string(package)
        artifact = purl.name
        version = purl.version
        data = self.get_python_package(artifact, version)
        # 查询不到源文件，直接返回
        if data is None:
            return None, HttpStatus.NOT_FOUND
        # 解析源文件
        package_dep, status = ArchiveDepsHandler.with_handlers([self]).deps('', data)
        if isinstance(package_dep, PackageDeps) and len(package_dep.modules) > 0:
            module_dep = package_dep.modules[0]
            module_dep.purl = package
            return module_dep, HttpStatus.OK
        return ModuleDeps(purl=package, lang=lang), HttpStatus.OK
