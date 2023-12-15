INSERT INTO graph.group(id, name, `desc`, parentID) VALUES (4, 'normal', '标准图谱', 1);

INSERT INTO graph.node(id, groupID, name, `desc`, `primary`) VALUES (7, 4, 'normal', '标准节点', 'name');
INSERT INTO graph.nodeAttr(nodeID, name, `desc`, type) VALUES (7, 'name', '名称', 0);
INSERT INTO graph.nodeAttr(nodeID, name, `desc`, type) VALUES (7, 'desc', '描述', 0);

INSERT INTO graph.edge(id, groupID, name, `desc`, `primary`, `direct`) VALUES (8, 4, 'normal', '标准边', null, 1);


INSERT INTO graph.algo(id, name, `desc`, detail, groupId, tag, jarPath, mainClass) VALUES (7, 'degree', '“度中心度”算法是一种用于衡量网络中节点重要性的指标，它反映了一个节点与其他节点直接相连的程度。', '“度中心度”算法是一种用于衡量网络中节点重要性的指标，它反映了一个节点与其他节点直接相连的程度。
“度中心度”算法的公式如下：
$$C(i)=\\frac{k_i}{N−1}$$其中，$N$表示网络中的节点总数，$k_i$表示节点$i$的度，即与它相连的边的数量。这个公式将节点$i$的度除以最大可能的度，即与其他$N−1$个节点都相连的情况，得到一个介于$0$和$1$之间的值，表示节点$i$与其他节点的连接比例。
“度中心度”认为一个节点的重要性取决于它的邻居节点的数量，而不考虑邻居节点的重要性或与其他节点的距离，它比较简单，时间复杂度为$O(N)$，但也有一些局限性，例如，它不能区分网络中的不同结构，如星形、环形或完全图，它们的节点度中心度都相同，但实际上它们的网络特性是不同的。
“度中心度”算法可以用它来分析社交网络中的用户影响力，找出与更多人有联系的用户，或者用它来分析互联网中的网站流量，找出与更多网站有链接的网站。还可以用它来分析交通网络中的节点重要性，找出与更多路线有连接的节点。', 4, '节点中心度', 's3a://lib/degree-latest.jar', 'applerodite.Main');


INSERT INTO graph.algo(id, name, `desc`, detail, groupId, tag, jarPath, mainClass) VALUES (8, 'pagerank', '“Pagerank”算法是一种用于衡量网络中节点重要性的指标，它反映了一个节点被其他节点引用的程度。', '“Pagerank”算法是一种用于衡量网络中节点重要性的指标，它反映了一个节点被其他节点引用的程度。
“Pagerank”算法的公式如下：
$$PR(i)=(1−d)+d*\\sum_{j∈M(i)}{\\frac{PR(j)}{L(j)}}$$
其中，$PR(i)$表示节点$i$的“Pagerank”值，$d$表示阻尼系数，一般取值为$0.85$，表示用户在某个页面继续点击链接的概率，$M(i)$表示链接到节点$i$的节点集合，$L(j)$表示节点$j$的出度，即链接到其他节点的数量。这个公式将节点$p_i$的“Pagerank”值定义为所有链接到它的节点的“Pagerank”值的加权和，再加上一个最小值$(1−d)$，用来处理没有出度的节点。
“Pagerank”算法认为一个节点的重要性取决于它的邻居节点的重要性和数量。这个算法比较复杂，其最优时间复杂度为$O(N+E)$，需要进行多次迭代计算，直到收敛为止。
“Pagerank”算法的使用场景有很多，例如，可以用它来分析互联网中的网页排名，找出被更多网页引用的网页，或者用它来分析社交网络中的用户影响力，找出被更多用户关注的用户。还可以用它来分析科学文献中的引文关系，找出被更多文献引用的文献。', 4, '节点中心度', 's3a://lib/pagerank-latest.jar', 'applerodite.Main');

INSERT INTO graph.algoParam(algoID, name, `desc`, type, `default`, `min`, `max`) VALUES (8, 'iter', '迭代次数', 2, '10', '1', '100');
INSERT INTO graph.algoParam(algoID, name, `desc`, type, `default`, `min`, `max`) VALUES (8, 'prob', '阻尼系数', 1, '0.85', '0.1', '1');

INSERT INTO graph.algo(id, name, `desc`, detail, groupId, tag, jarPath, mainClass) VALUES (9, 'louvain', 'louvain”算法是一种用于社区发现的算法，它可以在大规模的网络中快速地找出具有紧密连接的节点集合，即社区。', '“louvain”算法是一种用于社区发现的算法，它可以在大规模的网络中快速地找出具有紧密连接的节点集合，即社区。它的基本思想是利用模块度作为评价指标，通过不断地调整节点的社区归属，使得模块度达到最大值。

模块度是一种衡量网络划分质量的指标，它反映了网络中的社区结构与随机网络的差异程度。一个好的社区划分应该使得同一个社区内的节点连接更紧密，而不同社区间的节点连接更稀疏。模块度的计算公式如下：

$$Q=\\frac{1}{2m}\\sum_{i,j}{((A_{ij}−\\frac{k_ik_j}{2m})δ(ci,cj))}$$

其中，$A_{ij}$表示节点i和节点j之间的边的权重，$k_i$和$k_j$表示节点i和节点j的度，即与它们相连的边的权重之和，$m$表示网络中所有边的权重之和，$c_i$和$c_j$表示节点i和节点j所属的社区，δ(ci,cj)表示克罗内克函数，当$c_i=c_j$时为$1$，否则为$0$。这个公式将网络中每一对节点的实际连接权重与随机连接权重的差值进行累加，如果两个节点属于同一个社区，那么这个差值会对模块度有正向的贡献，否则会有负向的贡献。模块度的取值范围是$[−1,1]$，一般认为模块度在$[0.3,0.7]$之间属于较好的社区划分结果。

“louvain”算法的原理是通过两个阶段的迭代来优化模块度，第一个阶段是节点移动，第二个阶段是社区合并。具体的步骤如下：

- 初始化每个节点为一个单独的社区，计算网络的初始模块度$Q_0$。
- 对网络中的每个节点，遍历它的所有邻居节点，计算将该节点移动到邻居节点所属的社区后的模块度增益$ΔQ$，选择最大化$ΔQ$的社区作为该节点的新社区，如果没有正的$ΔQ$，则保持原社区不变。
- 重复上一步，直到所有节点的社区归属不再变化，或者模块度不再增加，得到第一个阶段的最终社区划分$C_1$，以及对应的模块度$Q_1$。
- 将$C_1$中的每个社区看作一个超级节点，将社区间的边权重定义为原来属于不同社区的节点间的边权重之和，得到一个新的网络$G_1$。
- 对$G_1$重复第二步和第三步，得到第二个阶段的最终社区划分$C_2$，以及对应的模块度$Q_2$。
- 重复上一步，直到模块度不再增加，或者网络中只剩下一个节点，得到最终的社区划分$C_n$，以及对应的模块度$C_n$。

“louvain”算法用它来分析社交网络中的用户群体，找出具有相似兴趣或行为的用户，或者用它来分析知识图谱中的实体关系，找出具有相似属性或语义的实体。还可以用它来分析生物网络中的基因或蛋白质，找出具有相似功能或结构的基因或蛋白质。', 4, '社区发现', 's3a://lib/pagerank-latest.jar', 'applerodite.Main');

INSERT INTO graph.algoParam(algoID, name, `desc`, type, `default`, `min`, `max`) VALUES (9, 'maxIter', '外部迭代次数', 2, '10', '1', '100');
INSERT INTO graph.algoParam(algoID, name, `desc`, type, `default`, `min`, `max`) VALUES (9, 'internalIter', '内部迭代次数', 1, '5', '1', '50');

