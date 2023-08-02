import time

import requests
import smart_open
from lxml import etree

headers = {
    'User-Agent': 'Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.0.0 '
                  'Safari/537.36',
}


def spider(url):
    if url.startswith('file://'):
        url = url[7:]
        with open(url, 'r') as f:
            return f.read()
    while True:
        try:
            html = requests.get(url, headers=headers)
            return html
        except:
            time.sleep(3)


def is_zip(s):
    return s.endswith('zip') or s.endswith('egg')


def is_tar(s):
    return s.endswith('tar.gz') or s.endswith('tar.bz2')


def get_python_package(artifact, version):
    if version == 'latest':
        url = 'https://pypi.org/pypi/{}/json'.format(artifact)
        version = spider(url).json()['info']['version']
    repo = etree.HTML(spider('{}/{}'.format('https://pypi.org/simple', artifact)).text)
    file_url = repo.xpath('/html/body/a')
    for a in file_url:
        desc = str(a.text).lower().replace('_', '-')
        if not is_zip(desc) and not is_tar(desc):
            continue
        edition = desc
        for suffix in [artifact + '-', '.zip', '.egg', '.tar.gz', '.tar.bz2']:
            edition = edition.replace(suffix, '')
        if edition == version:
            file_url = a.attrib.get('href')
            file_url = file_url[:file_url.rfind('#')]
            while True:
                try:
                    with smart_open.open(file_url, mode='rb') as f:
                        return f.read()
                except:
                    time.sleep(1)
    return None
