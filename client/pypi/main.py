import json

import feedparser
import requests
from lxml import etree
import os.path
import tarfile
import time
import zipfile

import uploadGraph
from requirements_analyze import get_packages_v2
from requirements_analyze import parse_v2
import requirements_analyze.requirements_detector
from uploadGraph import UpdateBody

pre = ''
headers = {
    'User-Agent': 'Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.0.0 '
                  'Safari/537.36',
}

if __name__ == '__main__':
    # d = feedparser.parse('https://pypi.org/rss/updates.xml')
    # requirements_analyze.get_packages_v2.setRoot('./pypi')
    # for e in d.entries:
    #     temp = e.title.split(' ')
    #     name = temp[0]
    #     ver = temp[1]
    #     url = e.link
    #     desc = e.summary
    #     requirements_analyze.get_packages_v2.extract_package(name)
    # requirements_analyze.parse_v2.parse()
    #
    dependency = {}
    # with open('files/f_requirements_pypi.csv', 'r') as f:
    #     for line in f:
    #         temp = line.split(',')
    #         package = temp[0]
    #         dependence = temp[1][0:-1]
    #         if package not in dependency:
    #             dependency[package] = set()
    #         dependency[package].add(dependence)
    dependency['b'] = set()
    dependency['b'].add('c')
    dependency['b'].add('e')
    id = requests.post('http://127.0.0.1:8888/api/graph/getGraphInfo', params={'name': 'test', 'id': 0}).json()['Id']
    #temp=requests.post('http://127.0.0.1:8888/api/graph/getNodesInfo',params={'id':id}).json()["Nodes"]

    temp = [{'name': 'e', 'Id': 5}, {'name': 'c', 'Id': 3}, {'name': 'd', 'Id': 4}, {'name': 'b', 'Id': 2},
            {'name': 'a', 'Id': 1}]

    record = {}
    for a in temp:
        record[a['name']] = a['Id']
    edges = {}
    for k, v in dependency.items():
        edges[record[k]] = list()
        for i in v:
            edges[record[k]].append(record[i])
    body = UpdateBody(id, edges)

    uploadGraph.initClient()
    uploadGraph.upload(body.toJson())
