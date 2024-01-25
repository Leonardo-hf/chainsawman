from common import HttpStatus
from util.lint import PyLintHandler, GoLintHandler, JavaLintHandler, ArchiveLintHandler

archive_lint = ArchiveLintHandler.with_handlers([PyLintHandler(), GoLintHandler(), JavaLintHandler()])


def test_py():
    path = 'test/cases/requirements-detector.zip'
    res, status = archive_lint.lint(path)
    assert status == HttpStatus.OK


def test_go():
    path = 'test/cases/gorm.zip'
    res, status = archive_lint.lint(path)
    assert status == HttpStatus.OK


def test_java():
    path = 'test/cases/log4j-core.tar.gz'
    res, status = archive_lint.lint(path)
    assert status == HttpStatus.OK