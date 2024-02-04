import json

from util.lint import PyLintHandler, GoLintHandler, JavaLintHandler, ArchiveLintHandler
from vo import LangLint

archive_lint = ArchiveLintHandler.with_handlers([PyLintHandler(), GoLintHandler(), JavaLintHandler()])


def test_py():
    path = 'test/cases/requirements-detector.zip'
    res, _ = archive_lint.lint(path)
    assert len(res) == 1
    r1 = res[0]
    assert isinstance(r1, LangLint) and r1.lang == 'python' and len(r1.lints) == 24


def test_go():
    path = 'test/cases/gorm.zip'
    res, _ = archive_lint.lint(path)
    assert len(res) == 1
    r1 = res[0]
    assert isinstance(r1, LangLint) and r1.lang == 'go' and len(r1.lints) == 19


def test_java():
    path = 'test/cases/log4j-core.tar.gz'
    res, _ = archive_lint.lint(path)
    assert len(res) == 1
    r1 = res[0]
    assert isinstance(r1, LangLint) and r1.lang == 'java' and len(r1.lints) == 302
