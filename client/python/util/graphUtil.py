import requests


def getGraphId(name, ip='http://127.0.0.1:8888'):
    return requests.post(ip + '/api/graph/getGraphInfo', params={'name': name, 'id': 0}).json()['Id']


def getNodesInfo(id, ip='http://127.0.0.1:8888'):
    return requests.post(ip + '/api/graph/getNodesInfo', params={'id': id}).json()["Nodes"]
