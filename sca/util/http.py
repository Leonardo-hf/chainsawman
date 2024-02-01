import requests

headers = {
    'User-Agent': 'Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.0.0 '
                  'Safari/537.36',
}


def spider(url):
    retry = 3
    while True:
        retry -= 1
        try:
            html = requests.get(url, headers=headers, timeout=10.05)
            return html
        except Exception as e:
            if retry <= 0:
                raise Exception(e)
