import json

from flask import Flask, request

from sca.common.consts import LANG_PY, RES_NOT_SUPPORT, LANG_JAVA, LANG_GO, LANG_RUST
from sca.handle.go import search_go, parse_go
from sca.handle.java import parse_java, search_java
from sca.handle.python import parse_python, search_python
from sca.handle.rust import parse_rust, search_rust
from sca.util.minio_helper import MinioHelper

app = Flask(__name__)

PARSE = 'parse'
SEARCH = 'search'

PARSE_MAP = {
    LANG_GO: {
        PARSE: parse_go,
        SEARCH: search_go
    },
    LANG_JAVA: {
        PARSE: parse_java,
        SEARCH: search_java
    },
    LANG_RUST: {
        PARSE: parse_rust,
        SEARCH: search_rust
    },
    LANG_PY: {
        PARSE: parse_python,
        SEARCH: search_python
    },
}


@app.route('/parse', methods=['GET'])
def parse():
    id = request.values.get('fileId')
    data = app.config['MINIO_CLIENT'].fetch(id)
    lang = request.values.get('lang')
    res = RES_NOT_SUPPORT
    if lang in PARSE_MAP:
        res = PARSE_MAP[lang][PARSE](data)
    print(res)
    return res


@app.route('/search', methods=['GET'])
def search():
    package = request.values.get('package')
    lang = request.values.get('lang')
    res = RES_NOT_SUPPORT
    if lang in PARSE_MAP:
        res = PARSE_MAP[lang][SEARCH](package)
    print(res)
    return res


if __name__ == '__main__':
    app.config.setdefault('CHS_ENV', '')
    app.config.from_envvar('CHS_ENV', silent=True)
    if app.config.get('CHS_ENV') == 'pre':
        app.config.from_file('config/client-pre.json', load=json.load)
    else:
        app.config.from_file('config/client.json', load=json.load)
    app.config.setdefault('MINIO_CLIENT', MinioHelper(app.config['MINIO']))
    app.run(host=app.config['HOST'], port=app.config['PORT'])
