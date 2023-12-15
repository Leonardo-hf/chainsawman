INSERT INTO graph.group(id, name, `desc`, parentID) VALUES (3, 'strangle', '卡脖子软件识别', 2);

INSERT INTO graph.node(id, groupID, name, `desc`, `primary`) VALUES (3, 3, 'organization', '开发团队', 'artifact');
INSERT INTO graph.nodeAttr(nodeID, name, `desc`, type) VALUES (3, 'name', '名称', 0);
INSERT INTO graph.nodeAttr(nodeID, name, `desc`, type) VALUES (3, 'home', '主页', 0);

INSERT INTO graph.node(id, groupID, name, `desc`, `primary`) VALUES (4, 3, 'developer', '开发者', 'idf');
INSERT INTO graph.nodeAttr(nodeID, name, `desc`, type) VALUES (4, 'name', '用户名', 0);
INSERT INTO graph.nodeAttr(nodeID, name, `desc`, type) VALUES (4, 'avator', '头像', 0);
INSERT INTO graph.nodeAttr(nodeID, name, `desc`, type) VALUES (4, 'blog', '博客', 0);
INSERT INTO graph.nodeAttr(nodeID, name, `desc`, type) VALUES (4, 'email', '邮箱', 0);
INSERT INTO graph.nodeAttr(nodeID, name, `desc`, type) VALUES (4, 'location', '常住地', 0);
INSERT INTO graph.nodeAttr(nodeID, name, `desc`, type) VALUES (4, 'company', '公司', 0);

INSERT INTO graph.node(id, groupID, name, `desc`, `primary`) VALUES (5, 3, 'library', '库', 'artifact');
INSERT INTO graph.nodeAttr(nodeID, name, `desc`, type) VALUES (5, 'artifact', '名称', 0);
INSERT INTO graph.nodeAttr(nodeID, name, `desc`, type) VALUES (5, 'desc', '描述', 0);
INSERT INTO graph.nodeAttr(nodeID, name, `desc`, type) VALUES (5, 'topic', '主题', 0);

INSERT INTO graph.node(id, groupID, name, `desc`, `primary`) VALUES (6, 3, 'release', '发行版本', 'idf');
INSERT INTO graph.nodeAttr(nodeID, name, `desc`, type) VALUES (6, 'idf', '标志符', 0);
INSERT INTO graph.nodeAttr(nodeID, name, `desc`, type) VALUES (6, 'artifact', '名称', 0);
INSERT INTO graph.nodeAttr(nodeID, name, `desc`, type) VALUES (6, 'version', '版本', 0);
INSERT INTO graph.nodeAttr(nodeID, name, `desc`, type) VALUES (6, 'createTime', '发布时间', 0);

INSERT INTO graph.edge(id, groupID, name, `desc`, `primary`, `direct`) VALUES (3, 3, 'maintain', '维护', null, 1);


INSERT INTO graph.edge(id, groupID, name, `desc`, `primary`, `direct`) VALUES (4, 3, 'contribute', '贡献', 'commits', 1);

INSERT INTO graph.edgeAttr(edgeID, name, `desc`, type) VALUES (4, 'commits', '贡献量', 2);

INSERT INTO graph.edge(id, groupID, name, `desc`, `primary`, `direct`) VALUES (5, 3, 'host', '主持', null, 1);


INSERT INTO graph.edge(id, groupID, name, `desc`, `primary`, `direct`) VALUES (6, 3, 'depend', '依赖', null, 1);


INSERT INTO graph.edge(id, groupID, name, `desc`, `primary`, `direct`) VALUES (7, 3, 'belong2', '属于', null, 1);


INSERT INTO graph.algo(id, name, `desc`, detail, groupId, tag, jarPath, mainClass) VALUES (6, 'Team Diversity', '团队多元性，软件的开发团队中贡献者和维护者的国别越多样，该软件的团队多元性越高。', '团队多元性，软件的开发团队中贡献者和维护者的国别越多样，该软件的团队多元性越高。', 3, '软件风险', 's3a://lib/diversity-latest.jar', 'applerodite.Main');

INSERT INTO graph.algoParam(algoID, name, `desc`, type, `default`, `min`, `max`) VALUES (6, 'libraries', '待识别卡脖子风险的软件列表', 3, null, '1', null);

