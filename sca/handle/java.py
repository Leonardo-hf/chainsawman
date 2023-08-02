import io

from sca.common.consts import RES_OK, RES_ILLEGAL_FILE, RES_NO_REQUIREMENTS, RES_EXTRA_REQUIREMENTS
from sca.util.file_helper import *
from sca.util.pom_helper import POM

MAVEN_DEPS = 'pom.xml'


def parse_java(data):
    file = io.BytesIO(data)
    res = RES_OK
    targets = []
    # 如果是 pom.xml
    if is_text(data):
        targets.append(data)
    # 如果是 zip 文件
    elif zipfile.is_zipfile(file):
        targets.extend(from_zip(file, MAVEN_DEPS))
    # 如果是 tar 文件
    else:
        try:
            file.seek(0)
            targets.extend(from_tar(file, MAVEN_DEPS))
        except:
            pass
    # 解析依赖文件
    if len(targets) == 0:
        return RES_NO_REQUIREMENTS
    elif len(targets) > 1:
        res = RES_EXTRA_REQUIREMENTS
    try:
        target = targets[0]
        pom = POM.from_string(target.decode())
        res['package'] = {
            'group': pom.get_group_id(),
            'artifact': pom.get_artifact(),
            'version': pom.get_version()
        }
        res['dependencies'] = list(map(lambda d: d.__dict__(), pom.get_dependencies()))
    except:
        return RES_ILLEGAL_FILE
    return res


def search_java(package: str):
    s1 = package.find('/')
    s2 = package.find(':')
    group = package[:s1].strip()
    artifact = package[s1 + 1:s2].strip()
    version = package[s2 + 1:].strip()
    res = RES_OK
    try:
        pom = POM.from_coordinate(artifact=artifact, group=group, version=version)
        res['package'] = {
            'group': group,
            'artifact': artifact,
            'version': version
        }
        deps = pom.get_dependencies()
        if deps is None:
            res['dependencies'] = []
        else:
            res['dependencies'] = list(map(lambda d: d.__dict__(), pom.get_dependencies()))
    except Exception as e:
        print(e)
        return RES_ILLEGAL_FILE
    return res
