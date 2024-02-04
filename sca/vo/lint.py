from typing import List, Optional

from attr import dataclass

from common import HttpStatus


@dataclass
class Lint:
    path: str = ''
    pos: str = ''
    msg: str = ''
    lint: str = ''


@dataclass
class LangLint(Lint):
    lang: str = ''
    lints: List[Lint] = []
    err: str = ''


@dataclass
class LintsResponse:
    base: HttpStatus
    langLints: List[LangLint]


@dataclass
class LintsRequest:
    file_id: str
