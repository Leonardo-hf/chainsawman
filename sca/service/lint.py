from abc import ABCMeta, abstractmethod

from vo import LintsResponse, LintsRequest


class LintService(metaclass=ABCMeta):
    @abstractmethod
    def lint(self, req: LintsRequest) -> LintsResponse:
        pass
