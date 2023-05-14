import json
from enum import Enum
from typing import List

import redis
import requests


class ImportMqClient:

    def __init__(self, host='localhost', port=6379):
        pool = redis.ConnectionPool(
            host=host, port=port, decode_responses=True)
        self.rdc = redis.StrictRedis(connection_pool=pool)

    def create_group(self, stream, group):
        if self.rdc.exists(stream):
            self.rdc.delete(stream)
        self.rdc.xgroup_create(stream, group, id=0, mkstream=True)

    def send(self, stream, values):
        self.rdc.xadd(stream, values.to_dict())


class GraphAPI:
    def __init__(self, ip='localhost:8888'):
        self.ip = ip

    def get_graph_id(self, graph):
        return requests.get(self.ip + '/api/graph/getGraphInfo', params={'name': graph}).json()['GraphId']

    def get_node_map(self, graph):
        return requests.get(self.ip + '/api/graph/getNodesInfo', params={'id': graph}).json()["Nodes"]

    def init_graph(self, name, desc):
        return requests.post(self.ip + '/api/graph/createEmpty', params={'graph': name, 'desc': desc}).json()["Graph"]


class Body:
    def to_json(self):
        temp = self.__dict__
        return json.dumps(temp)


class EdgesBody(Body):
    class Edge:
        def __init__(self, source: int, target: List[int]):
            self.source = source
            self.target = target

    def __init__(self, edges: List[Edge]):
        super(Body, self).__init__()
        self.edges = edges


class NodesBody(Body):
    class Node:
        def __init__(self, node_id: int, name: str, desc: str = ''):
            self.node_id = node_id
            self.name = name
            self.desc = desc

    def __init__(self, nodes: List[Node]):
        super(Body, self).__init__()
        self.nodes = nodes


class SplitEdgesBody(Body):
    def __init__(self, edges: List[List[int]]):
        self.edges = edges


class Message:
    def __init__(self, graph_id, opt, body):
        self.opt = opt
        self.graph_id = graph_id
        self.body = body.to_json()

    def to_dict(self):
        return self.__dict__


class Opt(Enum):
    updates = 1
    creates = 2
    deletes = 3
    insertEdges = 4
    deleteEdges = 5


class Type(Enum):
    edges = 1
    nodes = 2
    splitEdges = 3


def body_factory(type_: Type, lists):
    if type_ == Type.edges:
        return EdgesBody(lists)
    elif type_ == Type.nodes:
        return NodesBody(lists)
    elif type_ == Type.splitEdges:
        return SplitEdgesBody(lists)


class CSMClient:
    def __init__(self, import_mq=ImportMqClient(), graph_api=GraphAPI(), stream='import', group='import_consumers',
                 batch=100):
        import_mq.create_group(stream, group)
        self.im = import_mq
        self.graph = graph_api
        self.stream = stream
        self.group = group
        self.batch = batch

    def send_in_array(self, graph_id, opt, array, type_):
        for i in range(0, len(array), self.batch):
            self.im.send(self.stream, Message(graph_id=graph_id, opt=opt,
                                              body=body_factory(type_, array[i:min(len(array), self.batch + i)])))

    def updates(self, graph: int, edges: List[EdgesBody.Edge]):
        self.send_in_array(graph_id=graph, opt=Opt.updates, array=edges, type_=Type.edges)

    def creates(self, graph: int, nodes: List[NodesBody.Node]):
        self.send_in_array(graph_id=graph, opt=Opt.creates, array=nodes, type_=Type.nodes)

    def deletes(self, graph: int, nodes: List[NodesBody.Node]):
        self.im.send(self.stream, Message(graph_id=graph, opt=Opt.deletes, body=NodesBody(nodes=nodes)))

    def insert_edges(self, graph: int, edges: List[List[int]]):
        self.send_in_array(graph_id=graph, opt=Opt.insertEdges, array=edges, type_=Type.splitEdges)

    def delete_edges(self, graph: int, edges: List[List[int]]):
        self.send_in_array(graph_id=graph, opt=Opt.deleteEdges, array=edges, type_=Type.splitEdges)

    def init_graph(self, name: str):
        return self.init_graph_with_desc(name, '')

    def init_graph_with_desc(self, name: str, desc: str):
        return self.graph.init_graph(name, desc)

    def get_graph_id(self, name: str):
        return self.graph.get_graph_id(name)

    def get_node_map(self, graph: int):
        return self.graph.get_node_map(graph)
