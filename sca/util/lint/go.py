import subprocess
from shlex import quote
from typing import Optional, Tuple

from common import HttpStatus, GoLang
from util import Singleton
from vo import LangLint
from .index import LintHandler


@Singleton
class GoLintHandler(GoLang, LintHandler):
    def lint(self, path: str) -> Tuple[Optional[LangLint], HttpStatus]:
        cd = f'cd {quote(path)}'
        lint = 'golangci-lint run'
        ret = subprocess.Popen(';'.join([cd, lint]), shell=True,
                               stdout=subprocess.PIPE, stderr=subprocess.PIPE)
        return LangLint(lang=self.lang(), out=ret.stdout.read().decode().replace(path + '/', ''),
                        err=ret.stderr.read().decode().replace(path + '/', '')), HttpStatus.OK
