from collections import defaultdict
from functools import reduce

from common import Client
from util.deps import AllDepsHandler, PyDepsHandler, GoDepsHandler, JavaDepsHandler, RustDepsHandler
from util import resolve_archive
from vo import SearchDepsRequest, SearchDepsResponse, DepsRequest, DepsResponse, LanguageCount
from .deps import DepsService


class DepsServiceImpl(DepsService):

    def __init__(self):
        self.dh = AllDepsHandler.with_handlers(
            [PyDepsHandler(), GoDepsHandler(), JavaDepsHandler(), RustDepsHandler()])

    def deps(self, req: DepsRequest) -> DepsResponse:
        oss = Client.get_oss()
        data = oss.fetch(req.file_id)
        res, status = self.dh.deps(req.filename, data)
        archive = resolve_archive(data)
        count = defaultdict(int)

        def count_file(name: str):
            for h in [PyDepsHandler(), GoDepsHandler(), JavaDepsHandler(), RustDepsHandler()]:
                if reduce(lambda a, b: a or b, map(lambda ext: name.endswith(ext), h.exts())):
                    count[h.lang()] += 1

        archive.iter(count_file)
        return DepsResponse(base=status.value, packages=res,
                            counts=list(map(lambda it: LanguageCount(lang=it[0], cnt=it[1]), count.items())))

    def search(self, req: SearchDepsRequest) -> SearchDepsResponse:
        res, status = self.dh.search(req.lang, req.package)
        return SearchDepsResponse(base=status.value, deps=res)
