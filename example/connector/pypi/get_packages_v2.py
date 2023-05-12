import os.path
import tarfile
import time
import zipfile
from shutil import rmtree

import requests
from lxml import etree

from example.connector.pypi.requirements_detector import detect

office_url = 'https://pypi.org/simple'
tsinghua_url = 'https://pypi.tuna.tsinghua.edu.cn/simple'
aliyun_url = 'https://mirrors.aliyun.com/pypi/simple/'
base_url = office_url

headers = {
    'User-Agent': 'Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.0.0 '
                  'Safari/537.36',
}


def spider(url):
    while True:
        try:
            html = requests.get(url, headers=headers)
            return html
        except:
            time.sleep(3)
            pass


def download(url, path, chunk_s=1024):
    while True:
        try:
            req = requests.get(url, stream=True, headers=headers)
            with open(path, 'wb') as fh:
                for chunk in req.iter_content(chunk_size=chunk_s):
                    if chunk:
                        fh.write(chunk)
            return
        except:
            time.sleep(3)
            pass


def get_packages_list():
    html = spider(base_url).text
    content = etree.HTML(html)
    packages = list(map(lambda p: pformat(p), content.xpath('/html/body/a/text()')))
    return packages


def pformat(p: str):
    return p.replace('_', '-').lower()


def _extract_tar_files(package_file, path):
    try:
        tar_file = tarfile.open(name=package_file, mode='r:*', encoding='utf-8')
        ensure_dir(path)
        for member in tar_file.getmembers():
            if not member.isfile():
                continue
            f_name = member.name[member.name.rfind('/') + 1:]
            if f_name in ('setup.py', 'setup.cfg', 'requires.txt', 'pyproject.toml') or 'requirements' in f_name:
                with open('{}/{}'.format(path, f_name), 'wb') as file:
                    with tar_file.extractfile(member.name) as w:
                        file.write(w.read())
            paths = member.path.split('/')
            if len(paths) > 1 and paths[-2] == 'requirements':
                ensure_dir('{}/requirements'.format(path))
                with open('{}/requirements/{}'.format(path, f_name), 'wb') as file:
                    with tar_file.extractfile(member.name) as w:
                        file.write(w.read())
    except Exception as e:
        print('extract error on {} : {}'.format(package_file, str(e)))


def _extract_zip_files(package_file, path):
    try:
        z_file = zipfile.ZipFile(package_file, "r")
        ensure_dir(path)
        for name in z_file.namelist():
            if name.endswith('/'):
                continue
            f_name = name[name.rfind('/') + 1:]
            if f_name in ('setup.py', 'setup.cfg', 'requires.txt', 'pyproject.toml') or 'requirements' in f_name:
                with open('{}/{}'.format(path, f_name), 'wb') as file:
                    with z_file.open(name, 'r') as w:
                        file.write(w.read())
            paths = name.split('/')
            if len(paths) > 1 and paths[-2] == 'requirements':
                ensure_dir('{}/requirements'.format(path))
                with open('{}/requirements/{}'.format(path, f_name), 'wb') as file:
                    with z_file.open(name, 'r') as w:
                        file.write(w.read())
    except Exception as e:
        print('extract error on {} : {}'.format(package_file, str(e)))


def get_desc(name):
    url = 'https://pypi.org/project/{}'.format(name)
    content = etree.HTML(spider(url).text)
    desc = content.xpath('//p[@class=\'package-description__summary\']/text()')
    if len(desc) == 0:
        return desc[0]
    return ''


def extract_package(name):
    out_file = '/tmp/{}'.format(name)
    out_dir = '/tmp/dir_{}'.format(name)
    file_url = ''
    # 获得file_url
    url = 'https://pypi.org/project/{}'.format(name)
    content = etree.HTML(spider(url).text)
    edition = content.xpath('//h1/text()')
    if len(edition) == 0:
        return
    edition = pformat(''.join(edition[0].strip().split(' ')[1:]))
    repo = etree.HTML(spider('{}/{}'.format(base_url, name)).text)
    poss_file_url = repo.xpath('/html/body/a')
    for a in poss_file_url:
        desc = pformat(str(a.text))
        if edition in desc and (is_zip(desc) or is_tar(desc)):
            file_url = a.attrib.get('href')
            file_url = file_url[0:file_url.rfind('#')]
            break
    # 下载源代码文件
    if file_url != '':
        return
    if is_tar(file_url):
        download(file_url, out_file)
        _extract_tar_files(out_file, path=out_dir)
    elif is_zip(file_url):
        download(file_url, out_file)
        _extract_zip_files(out_file, path=out_dir)
    # 解析文件
    requirements = parse(out_dir)
    # 删除文件
    if os.path.exists(out_file):
        os.remove(out_file)
    if os.path.exists(out_dir):
        rmtree(out_dir)
    return requirements


def is_zip(s):
    return s.endswith('zip') or s.endswith('egg')


def is_tar(s):
    return s.endswith('tar.gz') or s.endswith('tar.bz2')


def ensure_dir(dirs):
    if not os.path.exists(dirs):
        os.makedirs(dirs)
        return False
    return True


def parse(p):
    try:
        requirements = detect.find_requirements(p)
        requirements = set(filter(lambda r: 'unknown' not in r, map(lambda r: pformat(r), requirements)))
        return requirements
    except:
        return {}
