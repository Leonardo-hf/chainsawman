from enum import Enum

from attr import dataclass


class HttpStatus(Enum):
    @dataclass
    class V:
        status: int
        msg: str

    OK = V(status=2000, msg='成功')
    ILLEGAL_FILE = V(status=4000, msg='文件格式不合法')
    NOT_SUPPORT = V(status=4001, msg='不支持此编程语言或包管理工具')
    NOT_FOUND = V(status=4002, msg='无法从源中获取该软件依赖')
