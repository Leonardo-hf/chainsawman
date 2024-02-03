import json
from ctypes import CDLL, c_char_p
from typing import List, Optional, Tuple

from packageurl import PackageURL

from common import HttpStatus, GoLang
from util import spider, Singleton
from vo import Dep, ModuleDeps
from .index import DepsHandler


def _get_purl(name, version) -> str:
    return 'pkg:golang/{}@{}'.format(name, version.lstrip('v'))


@Singleton
class GoDepsHandler(GoLang, DepsHandler):

    def deps(self, module: str, data: bytes) -> Tuple[Optional[ModuleDeps], HttpStatus]:
        if module.endswith(self.MODULE_GOMOD):
            lib = CDLL('lib/libmod.so')
            lib.Parse.restype = c_char_p
            res = lib.Parse(c_char_p(data))
            res = json.loads(res.decode())
            reqs: List[Dep] = []
            if 'Require' in res and res['Require']:
                for r in res['Require']:
                    reqs.append(Dep(purl=_get_purl(r['Mod']['Path'], r['Mod']['Version']), indirect=r['Indirect'],
                                    exclude=False))
            if 'Replace' in res and res['Replace']:
                for r in res['Replace']:
                    old = r['Old']['Path']
                    new, new_v = r['New']['Path'], r['New'].setdefault('Version', '')
                    index = -1
                    for i in range(len(reqs)):
                        if old in reqs[i].purl:
                            index = i
                            break
                    if index == -1:
                        continue
                    if new and new.startswith('.'):
                        del reqs[index]
                        continue
                    reqs[index].purl = _get_purl(new, new_v)
            if 'Exclude' in res and res['Exclude']:
                for e in res['Exclude']:
                    index = -1
                    for i in range(len(reqs)):
                        if e['Mod']['Path'] in reqs[i].purl:
                            index = i
                            break
                    indirect = False
                    if index != -1:
                        indirect = True
                    reqs.append(Dep(purl=_get_purl(e['Mod']['Path'], e['Mod']['Version']),
                                    indirect=indirect, exclude=True))
            artifact = res.get('Module', {'Mod': {'Path': '?'}})['Mod']['Path']
            return ModuleDeps(lang=self.lang(), purl=_get_purl(artifact, '?'), path=module,
                              dependencies=reqs), HttpStatus.OK
        return None, HttpStatus.NOT_SUPPORT

    def search(self, lang: str, package: str) -> Tuple[Optional[ModuleDeps], HttpStatus]:
        purl = PackageURL.from_string(package)
        url = 'https://goproxy.cn/{}/{}/@v/v{}.mod'.format(purl.namespace, purl.name, purl.version)
        res = spider(url)
        if res.status_code != 200:
            return None, HttpStatus.NOT_FOUND
        module_dep, status = self.deps(module=self.MODULE_GOMOD, data=res.content)
        module_dep.purl = package
        if module_dep is None:
            return None, HttpStatus.ILLEGAL_FILE
        return module_dep, HttpStatus.OK
