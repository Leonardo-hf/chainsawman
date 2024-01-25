import os
import shutil
from abc import ABCMeta, abstractmethod
from functools import reduce
from typing import Tuple, Optional, List, Set

from common import HttpStatus, Language
from util import resolve_archive
from vo import LangLint, LangLints


class LintHandler(Language, metaclass=ABCMeta):

    # 解析依赖文件，返回 ModuleDeps 或 None(解析失败)
    @abstractmethod
    def lint(self, path: str) -> Tuple[Optional[LangLint], HttpStatus]:
        pass

    def __hash__(self):
        return self.lang().__hash__()

    def __eq__(self, other):
        if isinstance(other, LintHandler):
            return other.lang() == self.lang()
        return False


class ArchiveLintHandler(LintHandler):

    def exts(self) -> List[str]:
        pass

    def modules(self) -> List[str]:
        pass

    def lang(self) -> str:
        pass

    def __init__(self, handlers: List[LintHandler]):
        self._handlers: List[LintHandler] = handlers

    @classmethod
    def with_handlers(cls, handlers: List[LintHandler]):
        return ArchiveLintHandler(handlers)

    def lint(self, path: str) -> Tuple[Optional[LangLint], HttpStatus]:
        with open(path, 'rb') as f:
            archive = resolve_archive(f.read())

            # 记录使用到的 lint
            used_handlers: Set[LintHandler] = set()

            def get_temp_dir(lang: str) -> str:
                return '{}/{}'.format(os.path.dirname(path), lang)

            def filter_save(name: str):
                for _h in self._handlers:
                    # 将不同的代码文件相关文件解压到对应的目录下
                    if reduce(lambda a, b: a or b,
                              map(lambda ext: name.lower().endswith(ext), _h.exts() + _h.modules())):
                        out_path = get_temp_dir(_h.lang())
                        used_handlers.add(_h)
                        os.makedirs(out_path, exist_ok=True)
                        archive.decompress_by_name(name, out_path)

            archive.iter(filter_save)

            res: List[LangLint] = []
            for h in used_handlers:
                p = get_temp_dir(h.lang())
                # 获得Lint
                lints, status = h.lint(p)
                if status != HttpStatus.OK:
                    continue
                res.append(lints)
                # 删除文件
                shutil.rmtree(p)
            return LangLints(lints=res), HttpStatus.OK
