INSERT INTO graph.group(id, name, `desc`, parentID) VALUES (4, 'strangle', '卡脖子软件识别', 3);

INSERT INTO graph.node(id, groupID, name, `desc`, `primary`) VALUES (4, 4, 'organization', '开发团队', 'name');
INSERT INTO graph.nodeAttr(nodeID, name, `desc`, type) VALUES (4, 'name', '名称', 0);
INSERT INTO graph.nodeAttr(nodeID, name, `desc`, type) VALUES (4, 'home', '主页', 0);

INSERT INTO graph.node(id, groupID, name, `desc`, `primary`) VALUES (5, 4, 'developer', '开发者', 'name');
INSERT INTO graph.nodeAttr(nodeID, name, `desc`, type) VALUES (5, 'name', '用户名', 0);
INSERT INTO graph.nodeAttr(nodeID, name, `desc`, type) VALUES (5, 'avator', '头像', 0);
INSERT INTO graph.nodeAttr(nodeID, name, `desc`, type) VALUES (5, 'blog', '博客', 0);
INSERT INTO graph.nodeAttr(nodeID, name, `desc`, type) VALUES (5, 'email', '邮箱', 0);
INSERT INTO graph.nodeAttr(nodeID, name, `desc`, type) VALUES (5, 'location', '常住地', 0);
INSERT INTO graph.nodeAttr(nodeID, name, `desc`, type) VALUES (5, 'company', '公司', 0);

INSERT INTO graph.node(id, groupID, name, `desc`, `primary`) VALUES (6, 4, 'library', '库', 'artifact');
INSERT INTO graph.nodeAttr(nodeID, name, `desc`, type) VALUES (6, 'artifact', '名称', 0);
INSERT INTO graph.nodeAttr(nodeID, name, `desc`, type) VALUES (6, 'desc', '描述', 0);
INSERT INTO graph.nodeAttr(nodeID, name, `desc`, type) VALUES (6, 'topic', '主题', 0);
INSERT INTO graph.nodeAttr(nodeID, name, `desc`, type) VALUES (6, 'home', '主页', 0);

INSERT INTO graph.node(id, groupID, name, `desc`, `primary`) VALUES (7, 4, 'release', '发行版本', 'idf');
INSERT INTO graph.nodeAttr(nodeID, name, `desc`, type) VALUES (7, 'idf', '标志符', 0);
INSERT INTO graph.nodeAttr(nodeID, name, `desc`, type) VALUES (7, 'artifact', '名称', 0);
INSERT INTO graph.nodeAttr(nodeID, name, `desc`, type) VALUES (7, 'version', '版本', 0);
INSERT INTO graph.nodeAttr(nodeID, name, `desc`, type) VALUES (7, 'createTime', '发布时间', 0);

INSERT INTO graph.edge(id, groupID, name, `desc`, `primary`, `direct`) VALUES (4, 4, 'maintain', '维护', null, 1);


INSERT INTO graph.edge(id, groupID, name, `desc`, `primary`, `direct`) VALUES (5, 4, 'contribute', '贡献', 'commits', 1);

INSERT INTO graph.edgeAttr(edgeID, name, `desc`, type) VALUES (5, 'commits', '贡献量', 2);

INSERT INTO graph.edge(id, groupID, name, `desc`, `primary`, `direct`) VALUES (6, 4, 'host', '主持', null, 1);


INSERT INTO graph.edge(id, groupID, name, `desc`, `primary`, `direct`) VALUES (7, 4, 'depend', '依赖', null, 1);


INSERT INTO graph.edge(id, groupID, name, `desc`, `primary`, `direct`) VALUES (8, 4, 'belong2', '属于', null, 1);


INSERT INTO graph.algo(id, name, define, detail, groupID, tag, tagID, jarPath, mainClass) VALUES (13, 'diversity', '团队多元性，软件的开发团队中贡献者和维护者的国别越多样，该软件的团队多元性越高。', '团队多元性，软件的开发团队中贡献者和维护者的国别越多样，该软件的团队多元性越高。', 4, '软件卡脖子风险', 0, 's3a://lib/diversity-latest.jar', 'applerodite.diversity.Main');

INSERT INTO graph.algoParam(algoID, name, `desc`, type, `default`, `min`, `max`) VALUES (13, 'libraries', '待识别卡脖子风险的软件列表', 3, null, '1', null);

