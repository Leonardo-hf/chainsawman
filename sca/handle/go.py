import io
import json
from ctypes import *

from common.consts import RES_OK, RES_NO_REQUIREMENTS, RES_EXTRA_REQUIREMENTS, RES_ILLEGAL_FILE
from util.file_helper import *
from util.http import spider

GOMOD_DEPS = 'go.mod'


def go_deps(content):
    content = content.encode()
    lib = CDLL("./lib/libmod.so")
    lib.Parse.restype = c_char_p
    res = lib.Parse(c_char_p(content))
    res = json.loads(res.decode())
    reqs = []
    if 'Require' in res and res['Require']:
        for r in res['Require']:
            reqs.append({'artifact': r['Mod']['Path'], 'version': r['Mod']['Version'],
                         'indirect': r['Indirect'], 'exclude': False})
    if 'Replace' in res and res['Replace']:
        for r in res['Replace']:
            old = {'artifact': r['Old']['Path'], 'version': r['Old'].setdefault('Version', '')}
            new = {'artifact': r['New']['Path'], 'version': r['New'].setdefault('Version', '')}
            index = -1
            for i in range(len(reqs)):
                if reqs[i]['artifact'] == old['artifact']:
                    index = i
                    break
            if index == -1:
                continue
            if new['artifact'] and new['artifact'].startswith('.'):
                del reqs[index]
                continue
            reqs[index]['artifact'] = new['artifact']
            reqs[index]['version'] = new['version']
    if 'Exclude' in res and res['Exclude']:
        for e in res['Exclude']:
            index = -1
            for i in range(len(reqs)):
                if reqs[i]['artifact'] == e['Mod']['Path']:
                    index = i
                    break
            indirect = False
            if index != -1:
                indirect = True
            reqs.append({'artifact': e['Mod']['Path'], 'version': e['Mod']['Version'],
                         'indirect': indirect, 'exclude': True})
    module = '?'
    if 'Module' in res and res['Module']:
        module = res['Module']['Mod']['Path']
    return module, reqs


def parse_go(data):
    file = io.BytesIO(data)
    res = RES_OK
    targets = []
    if is_text(data):
        targets.append(data)
    # 如果是 zip 文件
    elif zipfile.is_zipfile(file):
        targets.extend(from_zip(file, GOMOD_DEPS))
    # 如果是 tar 文件
    else:
        try:
            file.seek(0)
            targets.extend(from_tar(file, GOMOD_DEPS))
        except:
            pass
    # 解析依赖文件
    if len(targets) == 0:
        return RES_NO_REQUIREMENTS
    elif len(targets) > 1:
        res = RES_EXTRA_REQUIREMENTS
    try:
        target = targets[0]
        module, deps = go_deps(target)
        res['package'] = {
            'artifact': module,
            'version': '?'
        }
        res['dependencies'] = deps
    except:
        return RES_ILLEGAL_FILE
    return res


def search_go(package: str):
    res = RES_OK
    package = package.split(' ')[0]
    s = package.find(':')
    artifact = package[:s]
    version = package[s + 1:]
    url = 'https://goproxy.cn/{}/@v/{}.mod'.format(artifact, version)
    content = spider(url).text
    try:
        module, deps = go_deps(content)
        res['package'] = {
            'artifact': module,
        }
        res['dependencies'] = deps
        return res
    except:
        return RES_ILLEGAL_FILE
