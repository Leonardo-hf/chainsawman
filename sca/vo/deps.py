from typing import List, Optional

from attr import dataclass

from common import HttpStatus


@dataclass
class DepsRequest:
    file_id: str
    filename: str


@dataclass
class Dep:
    artifact: str
    version: str = "latest"
    group: str = ''
    limit: str = ''
    indirect: bool = False
    exclude: bool = False
    optional: bool = False
    scope: str = ''


@dataclass
class ModuleDeps:
    lang: str = ''
    path: str = ''
    group: str = ''
    artifact: str = ''
    version: str = ''
    dependencies: List[Dep] = []


@dataclass
class PackageDeps(ModuleDeps):
    modules: List[ModuleDeps] = []


@dataclass
class LanguageCount:
    type: str
    value: int


@dataclass
class DepsResponse:
    base: HttpStatus
    packages: Optional[ModuleDeps] = None
    counts: List[LanguageCount] = []


@dataclass
class SearchDepsRequest:
    package: str
    lang: str


@dataclass
class SearchDepsResponse:
    base: HttpStatus
    deps: Optional[ModuleDeps] = None
