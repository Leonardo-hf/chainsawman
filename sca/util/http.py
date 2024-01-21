import time

import requests

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