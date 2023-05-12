import os
import time
from multiprocessing import Process, Queue
from typing import List

from client.python.CSMClient import CSMClient, EdgesBody, NodesBody
from example.connector.pypi.get_packages_v2 import get_packages_list, extract_package, get_desc


def get_all_dependencies(packages: List[str], out: Queue):
    s = Queue()
    for p in packages:
        s.put(p)
    pl = []
    num = os.cpu_count() * 2 - 1
    for i in range(0, num):
        p = Process(target=producer, args=(s, out))
        p.start()
        pl.append(p)
    # 等待所有进程执行结束
    for p in pl:
        p.join()


def producer(q: Queue, out: Queue):
    while not q.empty():
        name = q.get()
        out.put((name, extract_package(name)))


def consumer(out: Queue, graph: int, node_map, client: CSMClient):
    while True:
        if out.qsize() > 1000:
            edges = []
            for i in range(1000):
                name, targets = out.get(block=True)
                for t in targets:
                    if t in node_map:
                        edges.append([node_map[name], node_map[t]])
            client.insert_edges(graph, edges)
        else:
            time.sleep(10)


# TODO: 结合pypi
if __name__ == '__main__':
    packages = get_packages_list()
    client = CSMClient()
    graph_id = client.init_graph('test')
    # 写入存量节点数据
    nodes = []
    id = 0
    for p in packages:
        desc = get_desc(p)
        nodes.append(NodesBody.Node(node_id=id, name=p, desc=desc))
        id += 1
        if len(nodes) > 1000:
            client.creates(graph_id, nodes)
            nodes.clear()
    client.creates(graph_id, nodes)
    node_map = client.get_node_map(graph_id)
    # 为依赖数据创建消费者
    out = Queue()
    c = Process(target=consumer, args=(out, graph_id, node_map, client))
    c.daemon = True
    c.start()
    # 写入存量依赖数据
    get_all_dependencies(out)
    # 写入增量数据
    new_dependency = {'node1': {'node3'}}
    update_edges = []
    for k, v in new_dependency.items():
        k_id = node_map[k]
        t_ids = []
        for target in v:
            t_ids.append(node_map[target])
        update_edges.append(EdgesBody.Edge(source=k_id, target=t_ids))
    client.updates(graph_id, update_edges)
    # 1, 2, 3; 1 -> 3
