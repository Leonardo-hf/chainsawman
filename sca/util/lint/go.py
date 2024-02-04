import re
import subprocess
from shlex import quote
from typing import Optional, Tuple, List

from common import HttpStatus, GoLang
from util import Singleton
from vo import Lint
from .index import LintHandler


@Singleton
class GoLintHandler(GoLang, LintHandler):
    def __init__(self):
        self._re = re.compile(
            r'(?P<path>[\S]*.go):(?P<pos>(?:[0-9]*:?)+):[\s]+(?P<msg>[\s\S]*?)[\s]+\((?P<lint>[\S]*)\)\n')

    def lint(self, path: str) -> Tuple[List[Lint], str]:
        cd = f'cd {quote(path)}'
        lint = 'golangci-lint run'
        ret = subprocess.Popen(';'.join([cd, lint]), shell=True,
                               stdout=subprocess.PIPE, stderr=subprocess.PIPE)
        out = ret.stdout.read().decode().replace(path + '/', '')
        out = self._re.finditer(out)
        lints = list(map(lambda o: Lint(path=o.group('path'), pos=o.group('pos'),
                                        msg=o.group('msg'), lint=o.group('lint')), out))
        return lints, ret.stderr.read().decode().replace(path + '/', '')
