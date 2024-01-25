from abc import abstractmethod, ABCMeta
from typing import List


class Language(metaclass=ABCMeta):
    @abstractmethod
    def exts(self) -> List[str]:
        pass

    @abstractmethod
    def modules(self) -> List[str]:
        pass

    @abstractmethod
    def lang(self) -> str:
        pass


class GoLang(Language):
    MODULE_GOMOD = 'go.mod'

    def lang(self) -> str:
        return 'go'

    def exts(self) -> List[str]:
        return ['.go']

    def modules(self) -> List[str]:
        return [self.MODULE_GOMOD]


class JavaLang(Language):
    MODULE_POM = 'pom.xml'

    def lang(self) -> str:
        return 'java'

    def exts(self) -> List[str]:
        return ['.java', '.scala', '.class']

    def modules(self) -> List[str]:
        return [self.MODULE_POM]


class RustLang(Language):
    MODULE_CARGO = 'cargo.toml'

    def lang(self) -> str:
        return 'rust'

    def exts(self) -> List[str]:
        return ['.rs']

    def modules(self) -> List[str]:
        return [self.MODULE_CARGO]


class PyLang(Language):
    MODULE_SETUP_PY = 'setup.py'
    MODULE_SETUP_CFG = 'setup.cfg'
    MODULE_REQUIRES = 'requires.txt'
    MODULE_PYPROJECT = 'pyproject.toml'

    def lang(self) -> str:
        return 'python'

    def exts(self) -> List[str]:
        return ['.py']

    def modules(self) -> List[str]:
        return [self.MODULE_SETUP_PY, self.MODULE_SETUP_CFG,
                self.MODULE_REQUIRES, self.MODULE_PYPROJECT]
