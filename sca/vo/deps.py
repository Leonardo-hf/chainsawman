from typing import List, Optional

from attr import dataclass

from common import HttpStatus


@dataclass
class DepsRequest:
    file_id: str
    filename: str


@dataclass
class OSV:
    id: str
    aliases: Optional[str]
    summary: str
    details: str
    cwe: Optional[str]
    severity: Optional[str]
    ref: Optional[str]


@dataclass
class Dep:
    purl: str
    osv: List[OSV] = []
    limit: str = ''
    indirect: bool = False
    exclude: bool = False
    optional: bool = False
    scope: str = ''


@dataclass
class ModuleDeps:
    purl: str = ''
    lang: str = ''
    path: str = ''
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
    purl: str
    lang: str


@dataclass
class SearchDepsResponse:
    base: HttpStatus
    deps: Optional[ModuleDeps] = None
