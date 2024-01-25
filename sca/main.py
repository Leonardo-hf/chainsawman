import json

from attrs import asdict
from flask import Flask, request
from gevent import pywsgi

from common import Client
from service import DepsService, DepsServiceImpl, LintServiceImpl, LintService
from vo import DepsRequest, SearchDepsRequest, LintsRequest

app = Flask(__name__)

deps_service: DepsService = DepsServiceImpl()
lint_service: LintService = LintServiceImpl()


@app.route('/parse', methods=['GET'])
def parse():
    file_id = request.values.get('fileId')
    filename = request.values.get('filename')
    req = DepsRequest(file_id=file_id, filename=filename)
    return asdict(deps_service.deps(req))


@app.route('/search', methods=['GET'])
def search():
    package = request.values.get('package')
    lang = request.values.get('lang')
    req = SearchDepsRequest(lang=lang, package=package)
    return asdict(deps_service.search(req))


@app.route('/lint', methods=['GET'])
def lint():
    file_id = request.values.get('fileId')
    req = LintsRequest(file_id=file_id, )
    return asdict(lint_service.lint(req))


if __name__ == '__main__':
    # 获取配置
    app.config.setdefault('CHS_ENV', '')
    app.config.from_prefixed_env()
    if app.config.get('CHS_ENV') == 'pre':
        app.config.from_file('config/client-pre.json', load=json.load)
    else:
        app.config.from_file('config/client.json', load=json.load)

    # 启动客户端
    Client.init(app.config)
    # 启动服务
    host = app.config['HOST']
    port = int(app.config['PORT'])
    if app.config.get('CHS_ENV') == 'pre':
        server = pywsgi.WSGIServer((host, port), app)
        server.serve_forever()
    else:
        app.run(host=app.config['HOST'], port=app.config['PORT'])
