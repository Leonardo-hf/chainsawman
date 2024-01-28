from collections import defaultdict
from functools import reduce

from common import HttpStatus
from util import resolve_archive
from util.deps import ArchiveDepsHandler, PyDepsHandler, GoDepsHandler, JavaDepsHandler, RustDepsHandler
from vo import PackageDeps

dh = ArchiveDepsHandler.with_handlers(
    [PyDepsHandler(), GoDepsHandler(), JavaDepsHandler(), RustDepsHandler()])


def test_deps_go():
    path = 'test/cases/gorm.zip'
    with open(path, 'rb') as f:
        ps, status = dh.deps(path, f)
        assert status == HttpStatus.OK
        assert isinstance(ps, PackageDeps) and len(ps.modules) == 2 and ps.modules[0].artifact == 'gorm.io/gorm'
        assert len(ps.modules[0].dependencies) == 2
        d = ps.modules[1].dependencies[-1]
        assert d.artifact == 'golang.org/x/text' and d.version == 'v0.14.0' and d.indirect

#
# def test_search_go():
#     name = 'gorm.io/gorm:v1.25.5'
#     ps, status = dh.search('go', name)
#     assert status == HttpStatus.OK
#     assert len(ps.dependencies) == 2 and ps.artifact == 'gorm.io/gorm'
#     d = ps.dependencies[-1]
#     assert d.artifact == 'github.com/jinzhu/now' and d.version == 'v1.1.5'
#
#
# def test_deps_java():
#     path = 'test/cases/log4j-1.2.17.jar'
#     with open(path, 'rb') as f:
#         ps, status = dh.deps(path, f)
#         assert status == HttpStatus.OK
#         assert status == HttpStatus.OK
#         assert isinstance(ps, PackageDeps) and len(ps.modules) == 1 and ps.modules[0].artifact == 'log4j' and \
#                ps.modules[0].group == 'log4j'
#         assert len(ps.modules[0].dependencies) == 5
#         d = ps.modules[0].dependencies[-1]
#         assert d.artifact == 'geronimo-jms_1.1_spec' and d.version == '1.0' and d.scope == 'provided'
#
#
# def test_search_java():
#     name = 'log4j/log4j:1.2.17'
#     ps, status = dh.search('java', name)
#     assert status == HttpStatus.OK
#     assert len(ps.dependencies) == 5 and ps.artifact == 'log4j' and ps.group == 'log4j'
#     d = ps.dependencies[-1]
#     assert d.artifact == 'geronimo-jms_1.1_spec' and d.version == '1.0' and d.scope == 'provided'
#
#
# def test_deps_python():
#     path = 'test/cases/requirements-detector.zip'
#     with open(path, 'rb') as f:
#         ps, status = dh.deps(path, f)
#         assert status == HttpStatus.OK
#         assert isinstance(ps, PackageDeps) and len(ps.modules) == 4 and ps.modules[
#             0].artifact == 'requirements-detector'
#         assert len(ps.modules[0].dependencies) == 7
#         d = ps.modules[-1].dependencies[-1]
#         assert d.artifact == 'ssh://git@github.com/MyOrg/MyRepo.git' and d.version == 'rev@1.2.3'
#
#
# def test_search_python():
#     name = 'django-redis:5.4.0'
#     ps, status = dh.search('python', name)
#     print(ps, len(ps.dependencies), ps.dependencies[-1])
#     assert status == HttpStatus.OK
#     assert len(ps.dependencies) == 4 and ps.artifact == 'django-redis'
#     d = ps.dependencies[0]
#     assert d.artifact == 'Django' and d.version == '3.2' and d.limit == '>='
#
#
# def test_deps_rust():
#     path = 'test/cases/rustix.zip'
#     with open(path, 'rb') as f:
#         ps, status = dh.deps(path, f)
#         assert status == HttpStatus.OK
#         assert isinstance(ps, PackageDeps) and len(ps.modules) == 4 and ps.modules[0].artifact == 'rustix'
#         assert len(ps.modules[0].dependencies) == 12
#         d = ps.modules[0].dependencies[-1]
#         assert d.artifact == 'static_assertions' and d.version == '1.1.0' and d.scope == 'dev'
#
#
# def test_search_rust():
#     name = 'rustix:0.38.30'
#     ps, status = dh.search('rust', name)
#     assert status == HttpStatus.OK
#     assert len(ps.dependencies) == 23 and ps.artifact == 'rustix'
#     d = ps.dependencies[-1]
#     assert d.artifact == 'windows-sys' and d.version == '0.52.0' and d.scope == ''
#
#
# def test_deps_multiple():
#     path = 'test/cases/all.zip'
#     with open(path, 'rb') as f:
#         ps, status = dh.deps(path, f)
#         assert status == HttpStatus.OK
#         assert isinstance(ps, PackageDeps) and len(ps.modules) == 11
#
#
# def test_count_multiple():
#     path = 'test/cases/all.zip'
#     with open(path, 'rb') as f:
#         archive = resolve_archive(f)
#         count = defaultdict(int)
#
#         def count_file(name: str):
#             for h in [PyDepsHandler(), GoDepsHandler(), JavaDepsHandler(), RustDepsHandler()]:
#                 if reduce(lambda a, b: a or b, map(lambda ext: name.endswith(ext), h.exts())):
#                     count[h.lang()] += 1
#
#         archive.iter(count_file)
#         assert count['go'] == 147 and count['python'] == 35 and count['rust'] == 431 and count['java'] == 314
