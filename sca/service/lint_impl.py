import os
import shutil
import uuid

from client import Client
from common import HttpException, HttpStatus
from util.lint import GoLintHandler, JavaLintHandler, PyLintHandler, ArchiveLintHandler
from vo import LintsRequest, LintsResponse
from .lint import LintService


class LintServiceImpl(LintService):

    def __init__(self):
        self._inter_handlers = [GoLintHandler(), JavaLintHandler(), PyLintHandler()]
        self.dh = ArchiveLintHandler.with_handlers(self._inter_handlers)

    def lint(self, req: LintsRequest) -> LintsResponse:
        oss = Client.get_oss()
        data = oss.fetch(req.file_id)
        tmp_dir = '/tmp/{}'.format(uuid.uuid4())
        os.makedirs(tmp_dir)
        p = '{}/{}'.format(tmp_dir, req.file_id)
        with open(p, 'wb') as f:
            f.write(data)
        try:
            res, _ = self.dh.lint(p)
        except HttpException as e:
            return LintsResponse(langLints=[], base=e.status())
        finally:
            shutil.rmtree(tmp_dir)
        return LintsResponse(langLints=res, base=HttpStatus.OK)
