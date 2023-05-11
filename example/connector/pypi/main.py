import time

from client.python.CSMClient import CSMClient, EdgesBody, NodesBody

# TODO: 结合pypi
if __name__ == '__main__':
    client = CSMClient()
    graph_id = client.init_graph('test')
    init_nodes = [NodesBody.Node(node_id=1, name='node1'), NodesBody.Node(node_id=2, name='node2'),
                  NodesBody.Node(node_id=3, name='node3')]
    init_edges = [EdgesBody.Edge(source=1, target=[2])]
    client.creates(graph_id, init_nodes)
    # 1, 2, 3
    time.sleep(10)  # 等待创建完成
    client.updates(graph_id, init_edges)
    # 1, 2, 3; 1 -> 2
    time.sleep(10)  # 等待创建完成
    node_map = client.get_node_map(graph_id)
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
