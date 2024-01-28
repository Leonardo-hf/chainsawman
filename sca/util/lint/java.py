import subprocess
from shlex import quote
from typing import Optional, Tuple

from common import HttpStatus, JavaLang
from util import Singleton
from vo import LangLint
from .index import LintHandler


@Singleton
class JavaLintHandler(JavaLang, LintHandler):
    def lint(self, path: str) -> Tuple[Optional[LangLint], HttpStatus]:
        ret = subprocess.run(['pmd', 'check', '-d', quote(path), '-R', 'config/pmd_mvn_ruleset.xml'],
                             capture_output=True)
        return LangLint(lang=self.lang(), out=ret.stdout.decode().replace(path + '/', ''),
                        err=ret.stderr.decode().replace(path + '/', '')), HttpStatus.OK
