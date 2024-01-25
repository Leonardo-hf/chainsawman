from typing import Optional, Tuple

from common import HttpStatus, RustLang
from util import Singleton
from vo import LangLint
from .index import LintHandler


@Singleton
class RustLintHandler(RustLang, LintHandler):
    def lint(self, path: str) -> Tuple[Optional[LangLint], HttpStatus]:
        pass
