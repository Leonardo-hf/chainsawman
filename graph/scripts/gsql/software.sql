INSERT INTO graph.group(id, name, `desc`, parentID) VALUES (2, 'software', '软件依赖图谱', 1);

INSERT INTO graph.node(id, groupID, name, `desc`, `primary`) VALUES (1, 2, 'library', '库', 'artifact');
INSERT INTO graph.nodeAttr(nodeID, name, `desc`, type) VALUES (1, 'artifact', '名称', 0);
INSERT INTO graph.nodeAttr(nodeID, name, `desc`, type) VALUES (1, 'desc', '描述', 0);
INSERT INTO graph.nodeAttr(nodeID, name, `desc`, type) VALUES (1, 'topic', '主题', 0);

INSERT INTO graph.node(id, groupID, name, `desc`, `primary`) VALUES (2, 2, 'release', '发行版本', 'idf');
INSERT INTO graph.nodeAttr(nodeID, name, `desc`, type) VALUES (2, 'idf', '标志符', 0);
INSERT INTO graph.nodeAttr(nodeID, name, `desc`, type) VALUES (2, 'artifact', '名称', 0);
INSERT INTO graph.nodeAttr(nodeID, name, `desc`, type) VALUES (2, 'version', '版本', 0);
INSERT INTO graph.nodeAttr(nodeID, name, `desc`, type) VALUES (2, 'createTime', '发布时间', 0);

INSERT INTO graph.edge(id, groupID, name, `desc`, `primary`, `direct`) VALUES (1, 2, 'depend', '依赖', null, 1);


INSERT INTO graph.edge(id, groupID, name, `desc`, `primary`, `direct`) VALUES (2, 2, 'belong2', '属于', null, 1);


INSERT INTO graph.algo(id, name, `desc`, detail, groupId, tag, jarPath, mainClass) VALUES (4, 'breadth', '广度排序算法，基于假设：软件的入度越大越重要。使用Voterank算法衡量节点的入度的相对大小。', '广度排序算法，基于假设：软件的入度越大越重要。使用Voterank算法衡量节点的入度的相对大小。', 2, '软件影响力', 's3a://lib/breadth-latest.jar', 'applerodite.breadth.Main');

INSERT INTO graph.algoParam(algoID, name, `desc`, type, `default`, `min`, `max`) VALUES (4, 'iter', '迭代次数，即返回前多少个高影响力软件', 2, '100', '1', '2147483647');

INSERT INTO graph.algo(id, name, `desc`, detail, groupId, tag, jarPath, mainClass) VALUES (5, 'depth', '深度排序算法，基于假设：在更多依赖路径中处于尾部的软件更重要。', '深度排序算法，基于假设：在更多依赖路径中处于尾部的软件更重要。', 2, '软件影响力', 's3a://lib/depth-latest.jar', 'applerodite.depth.Main');


INSERT INTO graph.algo(id, name, `desc`, detail, groupId, tag, jarPath, mainClass) VALUES (6, 'mediation', '中介度排序算法，基于假设：在更多依赖路径中处于中部的软件具有重要的集成作用。使用betweenness算法衡量软件的中介作用。', '中介度排序算法，基于假设：在更多依赖路径中处于中部的软件具有重要的集成作用。使用betweenness算法衡量软件的中介作用。', 2, '软件影响力', 's3a://lib/mediation-latest.jar', 'applerodite.mediation.Main');


INSERT INTO graph.algo(id, name, `desc`, detail, groupId, tag, jarPath, mainClass) VALUES (7, 'stability', '稳定性排序算法，基于假设：对软件开发生态稳定贡献更大的软件具有更高的影响力。使用基于最小渗流的collective influence算法衡量软件对生态稳定的贡献。', '稳定性排序算法，基于假设：对软件开发生态稳定贡献更大的软件具有更高的影响力。使用基于最小渗流的collective influence算法衡量软件对生态稳定的贡献。', 2, '软件影响力', 's3a://lib/stability-latest.jar', 'applerodite.stability.Main');

INSERT INTO graph.algoParam(algoID, name, `desc`, type, `default`, `min`, `max`) VALUES (7, 'radius', '半径，识别节点影响力时考虑与该节点距离为radius的其他节点，计算复杂度随radius增加而升高', 2, '3', '2', '5');

INSERT INTO graph.algo(id, name, `desc`, detail, groupId, tag, jarPath, mainClass) VALUES (8, 'hhi', '“hhi”指数是赫芬达尔—赫希曼指数 (Herfindahl-Hirschman Index)的简称，它是一种测量产业集中度的综合指数，被用于衡量同个主题下的软件的垄断程度。', '“hhi”指数是赫芬达尔—赫希曼指数 (Herfindahl-Hirschman Index)的简称，它是一种测量产业集中度的综合指数。它反映了一个行业中各市场竞争主体所占行业总收入或总资产百分比的平方和，用来计量市场份额的变化，即市场中厂商规模的离散度。它是经济学界和政府管制部门使用较多的指标。在此处，“hhi”指数被应用衡量同个主题下的软件的垄断程度。

“hhi”指数的公式如下：

$$HHI=\\sum_{i=1}^{n}{s_i^2}$$

其中，$n$表示主题下的软件数，$s_i$表示第$i$个软件的市场占有率，即其被依赖的次数占该主题下所有软件被依赖的总数。

“hhi”指数的内涵是，一个行业的集中度越高，即市场份额越集中在少数几家企业手中，那么“hhi”指数就越大，反之则越小。当市场处于完全垄断时，即只有一家企业，那么“hhi”指数等于1；当市场上有许多企业，且规模都相同时，那么“hhi”指数等于1/n，当n趋向于无穷大时，“hhi”指数就趋向于0。

“hhi”指数基于“结构—行为—绩效”理论，即市场结构影响到企业的行为，并最终决定市场的资源配置效率。随着市场份额的集中，企业会趋于采用相互勾结策略，最终制定出来的价格会偏离完全竞争市场价格，导致社会福利损失。因此，“hhi”指数可以在一定程度上反映市场的竞争程度和垄断力。

“hhi”指数可以用来分析不同行业的市场结构，找出存在垄断或寡头的行业，或者用它来评估企业的并购或兼并对市场竞争的影响，或者用它来制定反垄断的政策或法规。

', 2, '网络拓扑特性', 's3a://lib/hhi-latest.jar', 'applerodite.hhi.Main');


