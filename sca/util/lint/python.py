import subprocess
from shlex import quote
from typing import Optional, Tuple

from common import HttpStatus, PyLang
from util import Singleton
from vo import LangLint
from .index import LintHandler


@Singleton
class PyLintHandler(LintHandler, PyLang):
    def lint(self, path: str) -> Tuple[Optional[LangLint], HttpStatus]:
        ret = subprocess.run(['ruff', 'check', quote(path)], capture_output=True)
        return LangLint(lang=self.lang(), out=ret.stdout.decode().replace(path + '/', ''),
                        err=ret.stderr.decode()), HttpStatus.OK
        # out = ret.stdout
        # err = ret.stderr
        # lints = []
        # if out is not None and len(out):
        #     lines = out.decode().split('\n')
        #     for line in lines:
        #         line = line.strip()
        #         # 输出 Found xxx errors. 时，结束解析
        #         if line.startswith('Found'):
        #             break
        #         msgs = line.split(': ')
        #         pos = msgs[0].replace(path + '/', '')
        #         i = pos.find(':')
        #         lints.append(Lint(file=pos[:i], pos=pos[i + 1:], tip=msgs[1]))
        # elif err is not None and len(err):
        #     err = err.decode()
        #     if err.startswith('warning: No Python files found under the given path(s)'):
        #         pass
        #     else:
        #         print(err)
        #         return None, HttpStatus.ILLEGAL_FILE
        # return LangLint(lang=self.lang(), lints=lints), HttpStatus.OK
