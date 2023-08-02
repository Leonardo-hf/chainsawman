import tarfile
import zipfile
from collections import defaultdict

import chardet


def from_zip(f, target):
    zip_file = zipfile.ZipFile(file=f, mode='r')
    members = list(filter(lambda n: n.lower().endswith(target), zip_file.namelist()))
    if len(members) == 0:
        return []
    return list(map(lambda m: zip_file.read(name=m).decode(), members))


def from_zips(f, targets: []):
    zip_file = zipfile.ZipFile(file=f, mode='r')
    res = defaultdict(list)
    for name in zip_file.namelist():
        for t in targets:
            if name.lower().endswith(t):
                res[t].append(zip_file.read(name=name).decode())
    return res


def from_tar(f, target):
    tar_file = tarfile.open(fileobj=f, mode='r')
    members = list(filter(lambda m: m.name.lower().endswith(target), tar_file.getmembers()))
    if len(members) == 0:
        return []
    return list(map(lambda m: tar_file.extractfile(member=m).read().decode(), members))


def from_tars(f, targets: []):
    tar_file = tarfile.open(fileobj=f, mode='r')
    res = defaultdict(list)
    for name in tar_file.getmembers():
        for t in targets:
            if name.lower().endswith(t):
                res[t].append(tar_file.extractfile(member=name).read().decode())
    return res


def is_text(b):
    return not chardet.detect(b)['encoding'] is None
