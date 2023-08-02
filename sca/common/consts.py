RES_OK = {
    'base': {
        'status': 2000,
        'msg': '成功'
    },
}

RES_EXTRA_REQUIREMENTS = {
    'base': {
        'status': 3000,
        'msg': '压缩包中有多个 依赖描述 文件，使用第一个文件解析依赖关系'
    },
}

RES_NO_REQUIREMENTS = {
    'base': {
        'status': 4000,
        'msg': '压缩包中没有必须的`依赖描述文件`'
    },
}

RES_ILLEGAL_FILE = {
    'base': {
        'status': 4001,
        'msg': '文件格式不合法'
    },
}

RES_NOT_SUPPORT = {
    'base': {
        'status': 4002,
        'msg': '不支持的语言类型'
    },
}

RES_NOT_FOUND = {
    'base': {
        'status': 4003,
        'msg': '无法从源中获取该软件依赖'
    },
}

LANG_PY = 'python'
LANG_JAVA = 'java'
LANG_GO = 'go'
LANG_RUST = 'rust'
