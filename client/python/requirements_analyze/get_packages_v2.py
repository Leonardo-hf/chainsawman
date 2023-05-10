import os.path
import tarfile
import time
import zipfile
from multiprocessing import Queue, Process

import requests
import tqdm
from lxml import etree

office_url = 'https://pypi.org/simple'
tsinghua_url = 'https://pypi.tuna.tsinghua.edu.cn/simple'
aliyun_url = 'https://mirrors.aliyun.com/pypi/simple/'
base_url = office_url

headers = {
    'User-Agent': 'Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.0.0 '
                  'Safari/537.36',
}

root = '../../pypi'


def setRoot(path):
    global root
    root = path


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
    p = 'files/packages.txt'
    if os.path.exists(p):
        with open(p, 'r') as f:
            return list(map(lambda x: x.strip(), f.readlines()))
    html = spider(base_url).text
    content = etree.HTML(html)
    packages = content.xpath('/html/body/a/text()')
    with open(p, "w") as f:
        for i in tqdm.tqdm(range(len(packages)), total=len(packages), desc="获得包列表进度"):
            f.write(str(packages[i]) + "\n")
    return packages


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


def extract_package(name):
    outdir = '{}/{}'.format(root, name)
    if os.path.exists(outdir):
        return
    url = 'https://pypi.org/project/{}/#files'.format(name)
    html = spider(url).text
    content = etree.HTML(html)
    file_url = content.xpath('//*[@id="files"]/div[1]/div[2]/a[1]/@href')
    if len(file_url) == 0:
        return
    file_url = file_url[0]
    tmp_path = '/tmp/{}'.format(file_url[str(file_url).rfind('/') + 1:])

    if not is_tar(file_url) and not is_zip(file_url):
        edition = content.xpath('//h1/text()')
        if len(edition) == 0:
            print('fail: ' + name)
            return
        edition = ''.join(edition[0].strip().split(' ')[1:]).replace('_', '-')
        repo = etree.HTML(spider('{}/{}'.format(base_url, name)).text)
        poss_file_url = repo.xpath('/html/body/a')
        for a in poss_file_url:
            desc = str(a.text).replace('_', '-')
            if edition in desc and (is_zip(desc) or is_tar(desc)):
                file_url = a.attrib.get('href')
                file_url = file_url[0:file_url.rfind('#')]
                print('get: ' + name + ', ' + file_url)
                break

    if is_tar(file_url):
        download(file_url, tmp_path)
        _extract_tar_files(tmp_path, path=outdir)
        if os.path.exists(tmp_path):
            os.remove(tmp_path)
    elif is_zip(file_url):
        download(file_url, tmp_path)
        _extract_zip_files(tmp_path, path=outdir)
        if os.path.exists(tmp_path):
            os.remove(tmp_path)


def is_zip(s):
    return s.endswith('zip') or s.endswith('egg')


def is_tar(s):
    return s.endswith('tar.gz') or s.endswith('tar.bz2')


def ensure_dir(dirs):
    if not os.path.exists(dirs):
        os.makedirs(dirs)
        return False
    return True


def get(q):
    while not q.empty():
        name = q.get()
        x = q.qsize()
        if x % 1000 == 0:
            print(x)
        extract_package(name)


if __name__ == '__main__':
    ensure_dir('files')
    packages = get_packages_list()
    s = Queue()
    for p in packages:
        if os.path.exists('{}/{}'.format(root, p)):
            continue
        s.put(p)
    # for i in tqdm.tqdm(range(len(packages)), total=len(packages), desc="下载包进度"):
    #     extract_package(packages[i])
    pl = []
    num = os.cpu_count() * 2
    for i in range(0, num):
        p = Process(target=get, args=(s,))
        p.start()
        pl.append(p)
