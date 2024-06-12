import time
from typing import List, Optional, Tuple, Dict, Callable

import smart_open
from packageurl import PackageURL

from common import HttpStatus, PyLang
from util import spider, Singleton
from util.requirements_detector import from_setup_cfg, from_setup_py, from_requirements_txt, \
    from_pyproject_toml, DetectedRequirement
from vo import ModuleDeps, Dep, PackageDeps
from .index import DepsHandler, ArchiveDepsHandler


def _get_purl(artifact: str, version: Optional[str] = None) -> str:
    return PackageURL(type='pypi', name=artifact, version=version).to_string()


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
        def standardize(name: str) -> str:
            return name.lower().replace('-', '_')

        if version is None:
            res = spider(f'https://pypi.org/pypi/{artifact}/json').json()
            version = res.get('info').get('version')
        res = spider(f'https://pypi.org/pypi/{artifact}/{version}/json').json()
        print(res.get('urls'))
        urls = list(filter(
            lambda r: not r.get('filename').endswith('whl') and
                      (r.get('filename').startswith(f'{standardize(artifact)}-{version}.') or
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

    def search(self, lang: str, package: str) -> Tuple[Optional[ModuleDeps], HttpStatus]:
        purl = PackageURL.from_string(package)
        artifact = purl.name
        version = purl.version
        data = self.get_python_package(artifact, version)
        if data is None:
            return None, HttpStatus.NOT_FOUND
        package_dep, status = ArchiveDepsHandler.with_handlers([self]).deps('', data)
        if isinstance(package_dep, PackageDeps) and len(package_dep.modules) > 0:
            module_dep = package_dep.modules[0]
            module_dep.purl = package
            return module_dep, HttpStatus.OK
        return ModuleDeps(purl=package, lang=lang), HttpStatus.OK
