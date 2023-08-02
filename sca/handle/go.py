import io

from sca.common.consts import RES_OK, RES_NO_REQUIREMENTS, RES_EXTRA_REQUIREMENTS, RES_ILLEGAL_FILE
from sca.util.file_helper import *
from sca.util.http import spider

GOMOD_DEPS = 'go.mod'


def go_deps(content):
    content = content.split('\n')
    require = False
    exclude = False
    replace = False
    module = '?'
    deps = {}

    def handle_require(l, exclude=False):
        l = l.split(' ')
        ta = l[0]
        tv = l[1]
        indirect = len(l) >= 4 and l[3] == 'indirect'
        deps[ta + '@' + tv] = {
            'artifact': ta,
            'version': tv,
            'indirect': indirect,
            'exclude': exclude
        }

    def handle_replace(l):
        l = l.split(' ')
        oa = l[0]
        tod = ''
        if len(l) == 3:
            for key in deps.keys():
                if key.startswith(oa + '@'):
                    tod = key
                    break
        elif len(l) == 4:
            na = l[2]
            nv = l[3]
            for key in deps.keys():
                if key.startswith(oa + '@'):
                    deps[na + '@' + nv] = {
                        'artifact': na,
                        'version': nv,
                        'indirect': deps[key]['indirect'],
                        'exclude': False
                    }
                    tod = key
                    break
        elif len(l) == 5:
            ov = l[1]
            tod = oa + '@' + ov
            na = l[3]
            if not na.startswith('.'):
                nv = l[4]
                deps[na + '@' + nv] = {
                    'artifact': na,
                    'version': nv,
                    'indirect': deps[tod]['indirect'],
                    'exclude': False
                }
        if tod in deps:
            del deps[tod]

    for line in content:
        line = line.strip()
        if len(line) == 0 or (len(line) == 1 and line[0] == ')'):
            require = False
            exclude = False
            replace = False
        elif line.startswith('require ('):
            require = True
        elif line.startswith('require'):
            handle_require(line[8:])
        elif line.startswith('exclude ('):
            exclude = True
        elif line.startswith('exclude'):
            handle_require(line[8:], True)
        elif require or exclude:
            handle_require(line, exclude)
        elif line.startswith('replace ('):
            replace = True
        elif line.startswith('replace'):
            handle_replace(line[8:])
        elif replace:
            handle_replace(line)
        elif line.startswith('module'):
            module = line[6:].strip()
    return module, list(deps.values())


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
