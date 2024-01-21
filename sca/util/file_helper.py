import io
import tarfile
import zipfile
from abc import abstractmethod, ABCMeta
from typing import Callable, IO, Union

import chardet


class Archive(metaclass=ABCMeta):

    @abstractmethod
    def get_file_by_name(self, name: str) -> bytes:
        pass

    @abstractmethod
    def iter(self, func: Callable[[str], None]) -> None:
        pass


class ZipArchive(Archive):
    def __init__(self, f: IO[bytes]):
        self.f = zipfile.ZipFile(file=f, mode='r')

    def get_file_by_name(self, name: str) -> bytes:
        return self.f.read(name=name)

    def iter(self, func: Callable[[str], None]):
        for name in self.f.namelist():
            func(name)


class TarArchive(Archive):
    def __init__(self, f: IO[bytes]):
        self.f = tarfile.open(fileobj=f, mode='r')

    def get_file_by_name(self, name: str) -> bytes:
        return self.f.extractfile(member=name).read()

    def iter(self, func: Callable[[str], None]):
        for member in self.f.getmembers():
            func(member.name)


def resolve_archive(data: Union[bytes, IO[bytes]]) -> Archive:
    f: IO[bytes] = data
    if isinstance(data, bytes):
        f = io.BytesIO(data)
    archive = None
    # 如果是 zip 文件
    if zipfile.is_zipfile(f):
        archive = ZipArchive(f)
    # 如果是 tar 文件
    elif is_tarfile(f):
        archive = TarArchive(f)
    return archive


def is_tarfile(f: IO[bytes]) -> bool:
    try:
        f.seek(0)
        tarfile.open(fileobj=f, mode='r')
    except tarfile.TarError:
        return False
    return True


def is_text(b):
    return not chardet.detect(b)['encoding'] is None
