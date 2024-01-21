from abc import ABCMeta, abstractmethod

from vo import DepsRequest, SearchDepsRequest, DepsResponse, SearchDepsResponse


class DepsService(metaclass=ABCMeta):
    @abstractmethod
    def deps(self, req: DepsRequest) -> DepsResponse:
        pass

    @abstractmethod
    def search(self, req: SearchDepsRequest) -> SearchDepsResponse:
        pass
