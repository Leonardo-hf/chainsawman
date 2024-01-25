import json
from ctypes import CDLL, c_char_p
from typing import List, Optional, Tuple

from common import HttpStatus, GoLang
from util import spider, Singleton
from vo import Dep, ModuleDeps
from .index import DepsHandler


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
                    reqs.append(
                        Dep(artifact=r['Mod']['Path'], version=r['Mod']['Version'], indirect=r['Indirect'],
                            exclude=False))
            if 'Replace' in res and res['Replace']:
                for r in res['Replace']:
                    old = Dep(artifact=r['Old']['Path'], version=r['Old'].setdefault('Version', ''))
                    new = Dep(artifact=r['New']['Path'], version=r['New'].setdefault('Version', ''))
                    index = -1
                    for i in range(len(reqs)):
                        if reqs[i].artifact == old.artifact:
                            index = i
                            break
                    if index == -1:
                        continue
                    if new.artifact and new.artifact.startswith('.'):
                        del reqs[index]
                        continue
                    reqs[index].artifact = new.artifact
                    reqs[index].version = new.version
            if 'Exclude' in res and res['Exclude']:
                for e in res['Exclude']:
                    index = -1
                    for i in range(len(reqs)):
                        if reqs[i].artifact == e['Mod']['Path']:
                            index = i
                            break
                    indirect = False
                    if index != -1:
                        indirect = True
                    reqs.append(
                        Dep(artifact=e['Mod']['Path'], version=e['Mod']['Version'], indirect=indirect, exclude=True))
            artifact = '?'
            if 'Module' in res and res['Module']:
                artifact = res['Module']['Mod']['Path']
            return ModuleDeps(lang=self.lang(), artifact=artifact, path=module, dependencies=reqs), HttpStatus.OK
        return None, HttpStatus.NOT_SUPPORT

    def search(self, lang: str, package: str) -> Tuple[Optional[ModuleDeps], HttpStatus]:
        package = package.split(' ')[0]
        s = package.find(':')
        artifact = package[:s]
        version = package[s + 1:]
        url = 'https://goproxy.cn/{}/@v/{}.mod'.format(artifact, version)
        res = spider(url)
        if res.status_code != 200:
            return None, HttpStatus.NOT_FOUND
        module_dep, status = self.deps(module=self.MODULE_GOMOD, data=res.content)
        module_dep.artifact = artifact
        module_dep.version = version
        if module_dep is None:
            return None, HttpStatus.ILLEGAL_FILE
        return module_dep, HttpStatus.OK
