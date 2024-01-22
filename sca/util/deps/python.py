import time
from functools import reduce
from typing import List, Optional, Tuple, Dict, Callable

import smart_open
from lxml import etree

from common import HttpStatus
from requirements_detector import from_setup_cfg, from_setup_py, from_requirements_txt, \
    from_pyproject_toml, DetectedRequirement
from util import spider, Singleton
from vo import ModuleDeps, Dep, PackageDeps
from .index import DepsHandler, AllDepsHandler


@Singleton
class PyDepsHandler(DepsHandler):
    MODULE_SETUP_PY = 'setup.py'
    MODULE_SETUP_CFG = 'setup.cfg'
    MODULE_REQUIRES = 'requires.txt'
    MODULE_PYPROJECT = 'pyproject.toml'
    PY_PARSE_MAP: Dict[str, Callable[[str], Tuple[str, str, List[DetectedRequirement]]]] = {
        MODULE_SETUP_PY: from_setup_py,
        MODULE_SETUP_CFG: from_setup_cfg,
        MODULE_REQUIRES: from_requirements_txt,
        MODULE_PYPROJECT: from_pyproject_toml,
    }

    def lang(self) -> str:
        return 'python'

    def exts(self) -> List[str]:
        return ['.py']

    def modules(self) -> List[str]:
        return [self.MODULE_SETUP_PY, self.MODULE_SETUP_CFG,
                self.MODULE_REQUIRES, self.MODULE_PYPROJECT]

    def deps(self, module: str, data: bytes) -> Tuple[Optional[ModuleDeps], HttpStatus]:
        module_deps: List[Dep] = []
        for k in self.PY_PARSE_MAP:
            if module.endswith(k):
                artifact, version, reqs = self.PY_PARSE_MAP[k](data.decode())
                for r in reqs:
                    if len(r.version_specs) == 0:
                        module_deps.append(Dep(artifact=r.name))
                        continue
                    for v in r.version_specs:
                        module_deps.append(Dep(artifact=r.name, version=v[1], limit=v[0]))
                return ModuleDeps(lang=self.lang(), path=module, artifact=artifact, version=version,
                                  dependencies=module_deps), HttpStatus.OK
        return None, HttpStatus.NOT_SUPPORT

    @staticmethod
    def get_python_package(artifact, version) -> Optional[bytes]:
        if version == 'latest':
            url = 'https://pypi.org/pypi/{}/json'.format(artifact)
            version = spider(url).json()['info']['version']
        repo = etree.HTML(spider('{}/{}'.format('https://pypi.org/simple', artifact)).text)
        file_url = repo.xpath('/html/body/a')
        for a in file_url:
            v = str(a.text).lower().replace('_', '-')
            archive_suffix = ['.zip', '.egg', '.tar.gz', '.tar.bz2']
            if not reduce(lambda l, r: l or r, map(lambda s: v.endswith(s), archive_suffix)):
                continue
            archive_suffix.append(artifact + '-')
            for s in archive_suffix:
                v = v.replace(s, '')
            if v == version:
                file_url = a.attrib.get('href')
                file_url = file_url[:file_url.rfind('#')]
                while True:
                    try:
                        with smart_open.open(file_url, mode='rb') as f:
                            return f.read()
                    except:
                        time.sleep(1)
        return None

    def search(self, lang: str, package: str) -> Tuple[Optional[ModuleDeps], HttpStatus]:
        package = package[package.find(' ') + 1:]
        s = package.rfind(':')
        artifact = package[:s].strip()
        version = package[s + 1:].strip()
        data = self.get_python_package(artifact, version)
        if data is None:
            return None, HttpStatus.NOT_FOUND
        package_dep, status = AllDepsHandler.with_handlers([self]).deps('', data)
        if isinstance(package_dep, PackageDeps) and len(package_dep.modules) > 0:
            module_dep = package_dep.modules[0]
            module_dep.artifact = artifact
            module_dep.version = version
            return module_dep, HttpStatus.OK
        return None, HttpStatus.NOT_FOUND
