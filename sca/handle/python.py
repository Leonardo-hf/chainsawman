import io

from common.consts import RES_OK, RES_ILLEGAL_FILE, RES_NOT_FOUND, RES_NO_REQUIREMENTS
from requirements_detector.methods import from_setup_cfg, from_setup_py, from_requirements_txt, \
    from_pyproject_toml
from util.http import get_python_package

FILE_SETUP_PY = 'setup.py'
FILE_SETUP_CFG = 'setup.cfg'
FILE_REQUIRES = 'requires.txt'
FILE_PYPROJECT = 'pyproject.toml'
FILE_PACKAGE = 'pkg-info'
PY_PARSE_MAP = {
    FILE_SETUP_PY: from_setup_py,
    FILE_SETUP_CFG: from_setup_cfg,
    FILE_REQUIRES: from_requirements_txt,
    FILE_PYPROJECT: from_pyproject_toml,
}


def parse_python(data):
    file = io.BytesIO(data)
    res = RES_OK
    reqs = set()
    p, version = '?', '?'
    targets = {}
    # 如果是 zip 文件
    if zipfile.is_zipfile(file):
        targets = from_zips(file, [FILE_SETUP_PY, FILE_SETUP_CFG, FILE_REQUIRES, FILE_PYPROJECT, FILE_PACKAGE])
    # 检查是否是 tar 文件
    else:
        try:
            file.seek(0)
            targets = from_tars(file, [FILE_SETUP_PY, FILE_SETUP_CFG, FILE_REQUIRES, FILE_PYPROJECT, FILE_PACKAGE])
        except Exception as e:
            print(e)
    if len(targets) == 0 or (len(targets) == 1 and FILE_PACKAGE in targets):
        res = RES_NO_REQUIREMENTS
    for k, vs in targets.items():
        if k == FILE_PACKAGE:
            v = vs[0]
            lines = v.split('\n')
            for line in lines:
                if line.startswith('Name:'):
                    p = line[5:].strip()
                elif line.startswith('Version:'):
                    version = line[8:].strip()
        else:
            for v in vs:
                parse = PY_PARSE_MAP[k]
                try:
                    reqs = reqs.union(parse(v))
                except:
                    return RES_ILLEGAL_FILE
    res['package'] = {
        'artifact': p,
        'version': version
    }
    res['dependencies'] = []
    for r in reqs:
        if len(r.version_specs) == 0:
            res['dependencies'].append({
                'artifact': r.name,
                'version': 'latest',
                'limit': ''
            })
            continue
        for v in r.version_specs:
            res['dependencies'].append({
                'artifact': r.name,
                'version': v[1],
                'limit': v[0]
            })
    return res


def search_python(package: str):
    package = package[package.find(' ') + 1:]
    s = package.rfind(':')
    artifact = package[:s].strip()
    version = package[s + 1:].strip()
    data = get_python_package(artifact, version)
    if data is None:
        return RES_NOT_FOUND
    return parse_python(data)
