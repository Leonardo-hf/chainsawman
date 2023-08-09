import io
import json

import toml

from common.consts import RES_OK, RES_ILLEGAL_FILE, RES_NO_REQUIREMENTS, RES_EXTRA_REQUIREMENTS, RES_NOT_FOUND
from util.http import spider

CARGO_DEPS = 'cargo.toml'


def rust_v_spec(v):
    tvs = v.split(',')
    res = []
    for tv_specs in tvs:
        tv_specs = tv_specs.strip()
        if tv_specs.startswith('^') or tv_specs.startswith('~'):
            tv = tv_specs[1:]
            limit = tv_specs[0]
        elif ' ' in tv_specs:
            tv_specs = tv_specs.split(' ')
            tv = tv_specs[1]
            limit = tv_specs[0]
        else:
            tv = tv_specs
            limit = ''
        res.append((tv, limit))
    return res


def rust_deps(content):
    parsed = toml.loads(content)
    info = parsed.get('package', {})
    artifact = info.get('name', '?')
    version = info.get('version', '?')
    deps = parsed.get('dependencies', {})
    res = []
    for t, detail in deps.items():
        tv = 'latest'
        if isinstance(detail, str):
            tv = detail
        if isinstance(detail, dict):
            if 'path' in detail:
                continue
            elif 'git' in detail:
                t = detail.get('git')
                if 'rev' in detail:
                    tv = 'rev@' + detail.get('rev')
                elif 'tag' in detail:
                    tv = 'tag@' + detail.get('tag')
                else:
                    tv = 'latest@' + detail.get('branch', 'master')
            else:
                tv = detail.get('version', tv)
        tvs = rust_v_spec(tv)
        for tv_specs in tvs:
            res.append({
                'artifact': t,
                'version': tv_specs[0],
                'limit': tv_specs[1]
            })
    return {
        'artifact': artifact,
        'version': version
    }, res


def parse_rust(data):
    file = io.BytesIO(data)
    res = RES_OK
    targets = []
    # 如果是 文本 文件
    if is_text(data):
        targets.append(data)
    # 如果是 zip 文件
    elif zipfile.is_zipfile(file):
        targets.extend(from_zip(file, CARGO_DEPS))
    # 如果是 tar 文件
    else:
        try:
            file.seek(0)
            targets.extend(from_tar(file, CARGO_DEPS))
        except:
            pass
    # 解析依赖文件
    if len(targets) == 0:
        return RES_NO_REQUIREMENTS
    elif len(targets) > 1:
        res = RES_EXTRA_REQUIREMENTS
    try:
        target = targets[0]
        p, deps = rust_deps(target)
        res['package'] = p
        res['dependencies'] = deps
    except Exception as e:
        print(e)
        return RES_ILLEGAL_FILE
    return res


def search_rust(package: str):
    package = package[package.find(' ') + 1:]
    s = package.rfind(':')
    artifact = package[:s].strip()
    version = package[s + 1:]
    url = None
    res = RES_OK
    res['package'] = {
        'artifact': artifact,
        'version': version
    }
    res['dependencies'] = []
    if 1 <= len(artifact) <= 2:
        url = 'https://mirrors.tuna.tsinghua.edu.cn/crates.io-index/{}/{}'.format(len(artifact), artifact)
    elif len(artifact) == 3:
        url = 'https://mirrors.tuna.tsinghua.edu.cn/crates.io-index/3/{}/{}'.format(artifact[0], artifact)
    elif len(artifact) > 3:
        url = 'https://mirrors.tuna.tsinghua.edu.cn/crates.io-index/{}/{}/{}'.format(artifact[0:2], artifact[2:4],
                                                                                     artifact)
    if url is not None:
        content = spider(url).text
        content = content.strip().split('\n')
        target = None
        if version == 'latest':
            target = json.loads(content[-1].strip())
        else:
            for j in content:
                j = json.loads(j.strip())
                if j['vers'] == version:
                    target = j
                    break
        if target is None:
            return RES_NOT_FOUND
        deps = list(map(lambda d: (d['name'], rust_v_spec(d['req'])), target['deps']))
        for name, v_spec in deps:
            for v, limit in v_spec:
                res['dependencies'].append({
                    'artifact': name,
                    'version': v,
                    'limit': limit
                })
    return res
