from abc import ABCMeta, abstractmethod

from vo import DepsRequest, SearchDepsRequest, DepsResponse, SearchDepsResponse, SearchMetaRequest, SearchMetaResponse


class DepsService(metaclass=ABCMeta):
    @abstractmethod
    def deps(self, req: DepsRequest) -> DepsResponse:
        pass

    @abstractmethod
    def search(self, req: SearchDepsRequest) -> SearchDepsResponse:
        pass

    @abstractmethod
    def meta(self, req: SearchMetaRequest) -> SearchMetaResponse:
        pass
