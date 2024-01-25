from typing import Optional, Tuple

from common import HttpStatus, JavaLang
from util import POM, Singleton
from vo import Dep, ModuleDeps
from .index import DepsHandler


@Singleton
class JavaDepsHandler(JavaLang, DepsHandler):

    def deps(self, module: str, data: bytes) -> Tuple[Optional[ModuleDeps], HttpStatus]:
        if module.endswith(self.MODULE_POM):
            try:
                pom = POM.from_string(data.decode())
                deps = pom.get_dependencies()
                if deps is not None:
                    deps = list(
                        map(lambda d: Dep(group=d.group, artifact=d.artifact, version=d.version, scope=d.scope,
                                          optional=d.optional), deps))
                return ModuleDeps(lang=self.lang(), path=module, group=pom.get_group_id(), artifact=pom.get_artifact(),
                                  version=pom.get_version(), dependencies=deps), HttpStatus.OK
            except Exception:
                pass
        return None, HttpStatus.NOT_FOUND

    def search(self, lang: str, package: str) -> Tuple[Optional[ModuleDeps], HttpStatus]:
        try:
            s1 = package.find('/')
            s2 = package.find(':')
            group = package[:s1].strip()
            artifact = package[s1 + 1:s2].strip()
            version = package[s2 + 1:].strip()
            pom = POM.from_coordinate(artifact=artifact, group=group, version=version)
            deps = pom.get_dependencies()
            if deps is not None:
                deps = list(
                    map(lambda d: Dep(group=d.group, artifact=d.artifact, version=d.version, scope=d.scope,
                                      optional=d.optional), deps))
            return ModuleDeps(lang=self.lang(), group=pom.get_group_id(), artifact=pom.get_artifact(),
                              version=pom.get_version(), dependencies=deps), HttpStatus.OK
        except Exception:
            return None, HttpStatus.ILLEGAL_FILE
