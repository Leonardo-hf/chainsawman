from collections import defaultdict
from functools import reduce
from typing import Optional, Dict, List

from packageurl import PackageURL

from client import Client
from util.deps import ArchiveDepsHandler, PyDepsHandler, GoDepsHandler, JavaDepsHandler, RustDepsHandler
from util import resolve_archive
from vo import SearchDepsRequest, SearchDepsResponse, DepsRequest, DepsResponse, LanguageCount, PackageDeps, OSV, Dep, \
    SearchMetaRequest, SearchMetaResponse
from .deps import DepsService


class DepsServiceImpl(DepsService):

    def __init__(self):
        self._inter_handlers = [PyDepsHandler(), GoDepsHandler(), JavaDepsHandler(), RustDepsHandler()]
        self.dh = ArchiveDepsHandler.with_handlers(self._inter_handlers)

    @staticmethod
    def get_osv_for_deps(res: PackageDeps):
        purl2deps: Dict[str, List[Dep]] = defaultdict(list)
        for m in res.modules:
            for d in m.dependencies:
                purl2deps[d.purl].append(d)
        purls = list(purl2deps.keys())[:2]

        def safe_from_str(p: str) -> PackageURL:
            try:
                return PackageURL.from_string(p)
            except Exception:
                return PackageURL(type='generic', name='generic')

        osvs = Client.get_osv().query_vul_by_purls(list(map(lambda p: safe_from_str(p), purls)))
        for i in range(len(purls)):
            purl = purls[i]
            for d in purl2deps[purl]:
                d.osv = osvs[i]

    def deps(self, req: DepsRequest) -> DepsResponse:
        oss = Client.get_oss()
        data = oss.fetch(req.file_id)
        res, status = self.dh.deps(req.filename, data)
        if isinstance(res, PackageDeps):
            self.get_osv_for_deps(res)
        archive = resolve_archive(data)
        count = defaultdict(int)

        def count_file(name: str):
            for h in self._inter_handlers:
                if reduce(lambda a, b: a or b, map(lambda ext: name.endswith(ext), h.exts())):
                    count[h.lang()] += 1
                    break

        archive.iter(count_file)
        return DepsResponse(base=status.value, packages=res,
                            counts=list(map(lambda it: LanguageCount(type=it[0], value=it[1]), count.items())))

    def search(self, req: SearchDepsRequest) -> SearchDepsResponse:
        deps, status = self.dh.search(req.lang, req.purl)
        return SearchDepsResponse(base=status.value, deps=deps)

    def meta(self, req: SearchMetaRequest) -> SearchMetaResponse:
        meta, status = self.dh.meta(req.lang, req.purl)
        return SearchMetaResponse(base=status.value, meta=meta)
