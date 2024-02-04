import json
from collections import defaultdict
from functools import reduce

from packageurl import PackageURL

from client import Client
from common import HttpStatus
from service import DepsServiceImpl
from util import resolve_archive
from util.deps import ArchiveDepsHandler, PyDepsHandler, GoDepsHandler, JavaDepsHandler, RustDepsHandler
from vo import PackageDeps

dh = ArchiveDepsHandler.with_handlers(
    [PyDepsHandler(), GoDepsHandler(), JavaDepsHandler(), RustDepsHandler()])


# def test_osv():
#     Client.init({'OSV_API': ''})
#     path = 'test/cases/htmlunit-2.17.zip'
#     with open(path, 'rb') as f:
#         ps, status = dh.deps(path, f)
#         assert isinstance(ps, PackageDeps)
#         DepsServiceImpl.get_osv_for_deps(ps)
#         print(ps)


def test_deps_go():
    path = 'test/cases/gorm.zip'
    with open(path, 'rb') as f:
        ps, status = dh.deps(path, f)
        assert status == HttpStatus.OK
        assert isinstance(ps, PackageDeps) and len(ps.modules) == 2 and 'gorm.io/gorm' in ps.modules[0].purl
        assert len(ps.modules[0].dependencies) == 2
        d = ps.modules[1].dependencies[-1]
        assert d.purl == 'pkg:golang/golang.org/x/text@0.14.0' and d.indirect


def test_search_go():
    name = 'pkg:golang/gorm.io/gorm@1.25.5'
    ps, status = dh.search('go', name)
    assert status == HttpStatus.OK
    assert len(ps.dependencies) == 2 and ps.purl == name
    d = ps.dependencies[-1]
    assert d.purl == 'pkg:golang/github.com/jinzhu/now@1.1.5'


def test_deps_java():
    path = 'test/cases/log4j-1.2.17.jar'
    with open(path, 'rb') as f:
        ps, status = dh.deps(path, f)
        assert status == HttpStatus.OK
        assert status == HttpStatus.OK
        assert isinstance(ps, PackageDeps) and len(ps.modules) == 1 and \
               PackageURL.from_string(ps.modules[0].purl).namespace == 'log4j' and \
               PackageURL.from_string(ps.modules[0].purl).name == 'log4j'
        assert len(ps.modules[0].dependencies) == 5
        d = ps.modules[0].dependencies[-1]
        assert d.purl == 'pkg:maven/org.apache.geronimo.specs/geronimo-jms_1.1_spec@1.0' and d.scope == 'compile'


def test_search_java():
    name = 'pkg:maven/log4j/log4j@1.2.17'
    ps, status = dh.search('java', name)
    assert status == HttpStatus.OK
    assert len(ps.dependencies) == 5 and PackageURL.from_string(ps.purl).namespace == 'log4j' and \
           PackageURL.from_string(ps.purl).name == 'log4j'
    d = ps.dependencies[-1]
    assert d.purl == 'pkg:maven/org.apache.geronimo.specs/geronimo-jms_1.1_spec@1.0' and d.scope == 'compile'


def test_deps_python():
    path = 'test/cases/requirements-detector.zip'
    with open(path, 'rb') as f:
        ps, status = dh.deps(path, f)
        assert status == HttpStatus.OK
        assert isinstance(ps, PackageDeps) and len(ps.modules) == 4 and PackageURL.from_string(
            ps.modules[0].purl).name == 'requirements-detector'
        assert len(ps.modules[0].dependencies) == 7
        d = ps.modules[-1].dependencies[-1]
        assert d.purl == 'pkg:pypi/ssh://git%40github.com/myorg/myrepo.git@rev:1.2.3'


def test_search_python():
    name = 'pkg:pypi/django-redis@5.4.0'
    ps, status = dh.search('python', name)
    assert status == HttpStatus.OK
    assert len(ps.dependencies) == 4 and PackageURL.from_string(ps.purl).name == 'django-redis'
    d = ps.dependencies[0]
    assert PackageURL.from_string(d.purl).name == 'django' and PackageURL.from_string(
        d.purl).version == '3.2' and d.limit == '>='


def test_deps_rust():
    path = 'test/cases/rustix.zip'
    with open(path, 'rb') as f:
        ps, status = dh.deps(path, f)
        assert status == HttpStatus.OK
        assert isinstance(ps, PackageDeps) and len(ps.modules) == 4 and PackageURL.from_string(
            ps.modules[0].purl).name == 'rustix'
        assert len(ps.modules[0].dependencies) == 12
        d = ps.modules[0].dependencies[-1]
        assert d.purl == 'pkg:cargo/static_assertions@1.1.0' and d.scope == 'dev'


def test_search_rust():
    name = 'pkg:cargo/rustix@0.38.30'
    ps, status = dh.search('rust', name)
    assert status == HttpStatus.OK
    assert len(ps.dependencies) == 23 and PackageURL.from_string(ps.purl).name == 'rustix'
    d = ps.dependencies[-1]
    assert d.purl == 'pkg:cargo/windows-sys@0.52.0' and d.scope == ''


def test_deps_multiple():
    path = 'test/cases/all.zip'
    with open(path, 'rb') as f:
        ps, status = dh.deps(path, f)
        assert status == HttpStatus.OK
        assert isinstance(ps, PackageDeps) and len(ps.modules) == 11


def test_count_multiple():
    path = 'test/cases/all.zip'
    with open(path, 'rb') as f:
        archive = resolve_archive(f)
        count = defaultdict(int)

        def count_file(name: str):
            for h in [PyDepsHandler(), GoDepsHandler(), JavaDepsHandler(), RustDepsHandler()]:
                if reduce(lambda a, b: a or b, map(lambda ext: name.endswith(ext), h.exts())):
                    count[h.lang()] += 1

        archive.iter(count_file)
        assert count['go'] == 147 and count['python'] == 35 and count['rust'] == 431 and count['java'] == 314
