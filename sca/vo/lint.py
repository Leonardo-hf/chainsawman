from typing import List, Optional

from attr import dataclass

# @dataclass
# class Lint:
#     file: str
#     pos: str
#     tip: str
from common import HttpStatus


@dataclass
class LangLint:
    lang: str = ''
    out: Optional[str] = None
    err: Optional[str] = None
    # lints: List[Lint]


@dataclass
class LangLints(LangLint):
    lints: List[LangLint] = []


@dataclass
class LintsResponse:
    base: HttpStatus
    langLint: LangLints = []


@dataclass
class LintsRequest:
    file_id: str
