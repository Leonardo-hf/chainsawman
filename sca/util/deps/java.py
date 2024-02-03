from typing import Optional, Tuple

from packageurl import PackageURL

from common import HttpStatus, JavaLang
from util import POM, Singleton
from vo import Dep, ModuleDeps
from .index import DepsHandler


def _get_purl(group, artifact, version) -> str:
    return PackageURL(type='maven', namespace=group, name=artifact, version=version).to_string()


@Singleton
class JavaDepsHandler(JavaLang, DepsHandler):

    def deps(self, module: str, data: bytes) -> Tuple[Optional[ModuleDeps], HttpStatus]:
        if module.endswith(self.MODULE_POM):
            try:
                pom = POM.from_string(data.decode())
                deps = pom.get_dependencies()
                if deps is not None:
                    deps = list(
                        map(lambda d: Dep(
                            purl=_get_purl(d.group, d.artifact, d.version),
                            scope=d.scope, optional=d.optional), deps))
                return ModuleDeps(lang=self.lang(), path=module,
                                  purl=_get_purl(pom.get_group_id(), pom.get_artifact(), pom.get_version()),
                                  dependencies=deps), HttpStatus.OK
            except Exception as e:
                print(e)
        return None, HttpStatus.NOT_FOUND

    def search(self, lang: str, package: str) -> Tuple[Optional[ModuleDeps], HttpStatus]:
        try:
            purl = PackageURL.from_string(package)
            group = purl.namespace
            artifact = purl.name
            version = purl.version
            pom = POM.from_coordinate(artifact=artifact, group=group, version=version)
            deps = pom.get_dependencies()
            if deps is not None:
                deps = list(
                    map(lambda d: Dep(
                        purl=_get_purl(d.group, d.artifact, d.version),
                        scope=d.scope, optional=d.optional), deps))
            return ModuleDeps(lang=self.lang(),
                              purl=_get_purl(pom.get_group_id(), pom.get_artifact(), pom.get_version()),
                              dependencies=deps), HttpStatus.OK
        except Exception:
            return None, HttpStatus.ILLEGAL_FILE
