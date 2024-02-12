INSERT INTO graph.group(id, name, `desc`, parentID) VALUES (3, 'software', '软件依赖图谱', 1);

INSERT INTO graph.node(id, groupID, name, `desc`, `primary`) VALUES (2, 3, 'library', '库', 'artifact');
INSERT INTO graph.nodeAttr(nodeID, name, `desc`, type) VALUES (2, 'artifact', '名称', 0);
INSERT INTO graph.nodeAttr(nodeID, name, `desc`, type) VALUES (2, 'desc', '描述', 0);
INSERT INTO graph.nodeAttr(nodeID, name, `desc`, type) VALUES (2, 'topic', '主题', 0);
INSERT INTO graph.nodeAttr(nodeID, name, `desc`, type) VALUES (2, 'home', '主页', 0);

INSERT INTO graph.node(id, groupID, name, `desc`, `primary`) VALUES (3, 3, 'release', '发行版本', 'idf');
INSERT INTO graph.nodeAttr(nodeID, name, `desc`, type) VALUES (3, 'idf', '标志符', 0);
INSERT INTO graph.nodeAttr(nodeID, name, `desc`, type) VALUES (3, 'artifact', '名称', 0);
INSERT INTO graph.nodeAttr(nodeID, name, `desc`, type) VALUES (3, 'version', '版本', 0);
INSERT INTO graph.nodeAttr(nodeID, name, `desc`, type) VALUES (3, 'createTime', '发布时间', 0);

INSERT INTO graph.edge(id, groupID, name, `desc`, `primary`, `direct`) VALUES (2, 3, 'depend', '依赖', null, 1);


INSERT INTO graph.edge(id, groupID, name, `desc`, `primary`, `direct`) VALUES (3, 3, 'belong2', '属于', null, 1);


INSERT INTO graph.algo(id, name, define, detail, groupID, tag, tagID, jarPath, mainClass) VALUES (7, 'breadth', '广度排序算法，基于假设：软件的入度（被依赖程度）越大越具有研发影响力。', '广度排序算法，基于假设：软件的入度（被依赖程度）越大越具有研发影响力。

依赖图谱中软件的入度直观地反映了软件在研发中被使用的程度，入度越高，则直接依赖于该软件的下游软件就越多。因此，本指标使用Voterank算法衡量节点的入度的相对大小。有公式如下：

$$C_{breadth}=\\sum_{p∈d(i)}^n{V_p}$$

其中，$d(i)$表示直接依赖软件$i$的软件集合。VoteRank算法首先为图谱中的每个软件赋予分值$Vp=1$。该算法分为多个轮次，在每一轮中，需要计算图谱中所有软件的$C_{breadth}$ ，并选取具有最高$C_{breadth}$的软件$F$，取$C_{breadth}(F)$为$F$在本指标中的得分。然后，将$F$的分值置为$0$，并对每个依赖于$F$的软件，将其分值减少$f$，且最低减少至$0$，此时进行下一轮计算，直至一轮中最高的$C_{breadth}$为$0$时，结束该算法。通常将$f$设置为图谱中软件的平均入度的倒数。', 3, '软件研发影响力', 2, 's3a://lib/breadth-latest.jar', 'applerodite.breadth.Main');

INSERT INTO graph.algoParam(algoID, name, `desc`, type, `default`, `min`, `max`) VALUES (7, 'iter', '迭代次数，即返回前多少个高影响力软件', 2, '100', '1', '2147483647');

INSERT INTO graph.algo(id, name, define, detail, groupID, tag, tagID, jarPath, mainClass) VALUES (8, 'depth', '深度排序算法，基于假设：在更多依赖路径中处于尾部的软件具有更高的研发影响力。', '深度排序算法，基于假设：在更多依赖路径中处于尾部的软件具有更高的研发影响力。

在软件依赖图谱中，对于一条由依赖（有向边）组成的路径，可以将路径中包含的软件（节点）依据与路径起点软件的距离由短至长划分为头部软件、中间软件，和尾部软件。其中尾部软件的更改会对路径上所有其余软件造成影响，因此本算法认为尾部软件更为重要。

算法提出软件的应用层级的概念以帮助度量软件间的相对位置，其公式如下：

$$D_i=\\sqrt{\\frac{\\sum_{p∈q(i)}^n{D_p^2}}{N}}+1 (1)$$

其中，$D_i$表示软件$i$的应用层级，$q(i)$表示软件$i$所直接依赖的软件集合。

使用该应用层级表示软件间相对位置而非直接采用它们在图谱上的距离，是因为软件通常依赖于多个软件进行开发，因此会处于多条路径中，采用递推的计算方法可以对不同路径上软件的位置加以综合。此外，软件倾向于依赖更多低应用层级的软件，采用平方平均数可以放大被依赖的高应用层级的软件的影响。


基于此，有深度指标如下：

$$C_{depth}=\\frac{\\sum_{p∈g(i)}^{n-1}{D_p-D_i}}{N-1} (2)$$

其中，$N$是图中软件的总数，$g(i)$表示直接或间接依赖于软件$i$的软件集合。', 3, '软件研发影响力', 2, 's3a://lib/depth-latest.jar', 'applerodite.depth.Main');


INSERT INTO graph.algo(id, name, define, detail, groupID, tag, tagID, jarPath, mainClass) VALUES (9, 'mediation', '中介度排序算法，基于假设：在更多依赖路径中处于中部的软件具有重要的集成作用，具有更高的研发影响力。', '中介度排序算法，基于假设：在更多依赖路径中处于中部的软件具有重要的集成作用，具有更高的研发影响力。

在软件依赖图谱中，对于一条由依赖（有向边）组成的路径，可以将路径中包含的软件（节点）依据与路径起点软件的距离由短至长划分为头部软件、中间软件，和尾部软件。在软件研发实践中，高度集成的开发工具或框架通常被广泛使用，因为它们能够帮助开发者更快速、便捷地解决问题，而这些软件便属于中间软件。

因此，指标使用betweenness算法寻找依赖图谱中的中间软件，并研究其中介作用。详情可见betweenness的介绍。', 3, '软件研发影响力', 2, 's3a://lib/mediation-latest.jar', 'applerodite.mediation.Main');


INSERT INTO graph.algo(id, name, define, detail, groupID, tag, tagID, jarPath, mainClass) VALUES (10, 'stability', '稳定性指标，基于假设：对软件开发生态稳定贡献更大的软件具有更高的研发影响力。', '稳定性指标，基于假设：对软件开发生态稳定贡献更大的软件具有更高的研发影响力。

算法认为软件的影响力与其对开发生态的贡献程度息息相关，而一方面健康的开发生态通常表征为其依赖网络具有良好的稳定性，另一方面软件依赖网络是一个无标度网络，其稳定性依赖于少数核心软件。这恰好与研究少部分软件的影响力一致，因此算法基于渗流理论，通过从网络中逐个去除软件来观察其对整个网络稳定性的影响，具体而言则使用基于最小渗流的协同影响力算法度量软件对生态稳定的贡献。其公式如下：

$$C_{stability}=(k_i-1)\\sum_{p∈δBall(i,l)}{(k_p-1} (1)$$

$$λ(l)=(\\frac{\\sum_i{C_{stability}(i)}}{N<k>})^{\\frac{1}{l+1}} (2)$$

其中，$k_i$表示软件$i$的度数，$Ball(i, l)$表示距离软件$i$最短距离为$l$的软件集合，$N$是节点的总数，$<k>$是图当前的平均度。算法首先计算依赖网络中所有软件的$C_{stability}$ ，并认为软件对网络的稳定性的贡献取决于$C_{stability}$ 的大小。接着算法计算图的$λ(l)$，若$λ(l)>1$，则将$C_(st)$最大的软件从网络中移除，并重新计算所有软件的$C_(stability)$，直到$λ(l)≤1$，此时认为图谱被破坏，不再重新计算软件的$C_{stability}$。', 3, '软件研发影响力', 2, 's3a://lib/stability-latest.jar', 'applerodite.stability.Main');

INSERT INTO graph.algoParam(algoID, name, `desc`, type, `default`, `min`, `max`) VALUES (10, 'radius', '半径，识别节点影响力时考虑与该节点距离为radius的其他节点，计算复杂度随radius增加而升高', 2, '3', '2', '5');

INSERT INTO graph.algo(id, name, define, detail, groupID, tag, tagID, jarPath, mainClass) VALUES (11, 'integrated', '综合影响力指标，基于专家意见的AHP方法调整breadth, depth, mediation, stability四种高研发影响力软件识别指标的权重。', '综合影响力指标，基于专家意见的AHP方法调整breadth, depth, mediation, stability四种高研发影响力软件识别指标的权重。其公式如下：

$$C_{integrated}(i)=a*C_{breadth}(i)+b*C_{depth}(i)+c*C_{mediation}(i)+d*C_{stability}(i)$$

其中，a,b,c,d为各个指标的权重，基于对Python软件的观察，其建议取值为$a=0.52887$，$b=0.21942$，$c=0.09656$，$d=0.15515$。对于其他语言，上述取值可能不适用。

综合影响力算法可以有效综合各指标的优势，使之达到互补的效果，与下载量名列前茅、受到Awesome Project的开发者广泛推荐的软件有更高的拟合度。', 3, '软件研发影响力', 2, 's3a://lib/integrated-latest.jar', 'applerodite.integrated.Main');

INSERT INTO graph.algoParam(algoID, name, `desc`, type, `default`, `min`, `max`) VALUES (11, 'weights', '集成的各个算法（breadth, depth, mediation, stability）的权重，范围为0~1', 4, null, '1', '4');

INSERT INTO graph.algo(id, name, define, detail, groupID, tag, tagID, jarPath, mainClass) VALUES (12, 'hhi', '“hhi”指数是赫芬达尔—赫希曼指数的简称，它是一种测量产业集中度的综合指数，可以反应同应用领域中头部软件在研发中的垄断程度。', '“hhi”指数是赫芬达尔—赫希曼指数的简称，它是一种测量产业集中度的综合指数，可以反应同应用领域中头部软件在研发中的垄断程度。它反映了一个行业中各市场竞争主体所占行业总收入或总资产百分比的平方和，用来计量市场份额的变化，即市场中厂商规模的离散度。它是经济学界和政府管制部门使用较多的指标。在此处，“hhi”指数被应用衡量同个主题下的软件的垄断程度。

“hhi”指数的公式如下：

$$HHI=\\sum_{i=1}^{n}{s_i^2}$$

其中，$n$表示主题下的软件数，$s_i$表示第$i$个软件的市场占有率，即其被依赖的次数占该主题下所有软件被依赖的总数。

“hhi”指数的内涵是，一个行业的集中度越高，即市场份额越集中在少数几家企业手中，那么“hhi”指数就越大，反之则越小。当市场处于完全垄断时，即只有一家企业，那么“hhi”指数等于1；当市场上有许多企业，且规模都相同时，那么“hhi”指数等于1/n，当n趋向于无穷大时，“hhi”指数就趋向于0。

“hhi”指数基于“结构—行为—绩效”理论，即市场结构影响到企业的行为，并最终决定市场的资源配置效率。随着市场份额的集中，企业会趋于采用相互勾结策略，最终制定出来的价格会偏离完全竞争市场价格，导致社会福利损失。因此，“hhi”指数可以在一定程度上反映市场的竞争程度和垄断力。

“hhi”指数可以用来分析不同行业的市场结构，找出存在垄断或寡头的行业，或者用它来评估企业的并购或兼并对市场竞争的影响，或者用它来制定反垄断的政策或法规。

', 3, '软件研发影响力', 2, 's3a://lib/hhi-latest.jar', 'applerodite.hhi.Main');


