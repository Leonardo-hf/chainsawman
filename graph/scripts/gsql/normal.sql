INSERT INTO graph.group(id, name, `desc`, parentID) VALUES (4, 'normal', '标准图谱', 1);

INSERT INTO graph.node(id, groupID, name, `desc`, `primary`) VALUES (7, 4, 'normal', '标准节点', 'name');
INSERT INTO graph.nodeAttr(nodeID, name, `desc`, type) VALUES (7, 'name', '名称', 0);
INSERT INTO graph.nodeAttr(nodeID, name, `desc`, type) VALUES (7, 'desc', '描述', 0);

INSERT INTO graph.edge(id, groupID, name, `desc`, `primary`, `direct`) VALUES (8, 4, 'normal', '标准边', null, 1);


