import os
import shutil
import uuid

from common import Client
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
        res, status = self.dh.lint(p)
        print(res)
        shutil.rmtree(tmp_dir)
        return LintsResponse(langLint=res, base=status.value)
