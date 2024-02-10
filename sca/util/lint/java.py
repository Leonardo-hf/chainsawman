import re
import subprocess
from shlex import quote
from typing import List, Tuple

from common import JavaLang
from util import Singleton
from vo import Lint
from .index import LintHandler


@Singleton
class JavaLintHandler(JavaLang, LintHandler):
    def __init__(self):
        self._re = re.compile(
            r'(?P<path>[\S]*?.java):(?P<pos>(?:[0-9]*:?)+):[\s]+(?P<lint>[\S]+):[\s]+(?P<msg>[\S\s]*?)\n')

    def lint(self, path: str) -> Tuple[List[Lint], str]:
        ret = subprocess.run(['pmd', 'check', '-d', quote(path), '-R', 'config/pmd_mvn_ruleset.xml'],
                             capture_output=True)
        out = ret.stdout.decode().replace(path + '/', '')
        out = self._re.finditer(out)
        lints = list(map(lambda o: Lint(path=o.group('path'), pos=o.group('pos'),
                                        msg=o.group('msg'), lint=o.group('lint')), out))
        return lints, ret.stderr.decode().replace(path + '/', '')