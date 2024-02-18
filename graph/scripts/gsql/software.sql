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


INSERT INTO graph.algo(id, name, define, detail, groupID, tag, tagID, jarPath, mainClass) VALUES (7, 'breadth', '广度排序算法是软件研发影响力的一种，其通过度量软件的基础程度衡量其研发影响力。', '# 广度排序算法

广度排序算法是软件研发影响力的一种，其基于假设：基础性质的软件，例如工具库等在功能开发中被普遍使用的软件对研发更有帮助。而入度（被依赖程度）越大的软件越具备基础性质。广度排序算法试图发现**软件的基础程度**。

## 定义

依赖图谱中软件的入度直观地反映了软件在研发中被使用的程度，入度越高，则直接依赖于该软件的下游软件就越多，该软件的通用性越强，更符合基础软件的特征。为了衡量软件被直接依赖的程度，本指标使用基于投票的$Voterank$算法衡量节点的入度的相对大小，相比于直接使用入度有如下优势：

* 第三方软件通常被组合使用以实现其完整功能，$Voterank$能够增加被组合使用的软件的区分度，识别出其中更重要的软件。

## 公式

$$
C_{breadth}=\\sum_{p∈d(i)}^n{V_p}\\tag{1}
$$

式中，$d(i)$表示直接依赖软件$i$的软件集合，$V_p$为软件$p$在算法中被赋予的投票能力，$VoteRank$算法为图谱中的每个软件赋予初始投票能力$V_p=1$。该算法分为多个轮次，在每一轮中，需要计算图谱中每个软件的$C_{breadth}$ ，即所有软件向自己直接依赖的软件进行一次投票，并选取具有最高$C_{breadth}$的软件$F$，此时$C_{breadth}(F)$即为$F$的基础程度。然后，将$V_F$置为$0$，并对每个依赖于$F$的软件，将其投票能力减少损失因子$f$，且最低减少至$0$，此时进行下一轮计算，直至一轮中最高的$C_{breadth}$为$0$时（即全部软件已经失去投票能力），结束该算法。

### 复杂度

* 时间复杂度$O(N*E)$

## 参数

### 损失因子（f）

* 范围：$(0,1]$

* 默认值：通常将$f$设置为图谱中软件的平均入度的倒数。

当进行投票的软件成功选举出当轮次下得分最高的软件时，这些进行投票的软件的投票能力被减弱$f$，这保障了投票的公平性，即不会有多个高得分的软件被同一批软件投票胜出。

### 迭代次数（iter）

* 范围：$[1,N]$

* 默认值：100

$VoteRank$算法是多轮次的算法，并在每轮此中筛选出当前轮次下得分最高的软件，通过设置迭代次数可以提前结束算法。将迭代次数设置为$N$，即为计算排名前$N$的软件的分值。

## 引用

* [Zhang, J. X., Chen, D. B., Dong, Q., & Zhao, Z. D. (2016). Identifying a set of influential spreaders in complex networks. *Scientific reports*, *6*(1), 27823.](https://www.nature.com/articles/srep27823)

', 3, '软件研发影响力', 2, 's3a://lib/breadth-latest.jar', 'applerodite.breadth.Main');

INSERT INTO graph.algoParam(algoID, name, `desc`, type, `default`, `min`, `max`) VALUES (7, 'iter', '迭代次数，即返回前多少个高影响力软件', 2, '100', '1', '2147483647');

INSERT INTO graph.algo(id, name, define, detail, groupID, tag, tagID, jarPath, mainClass) VALUES (8, 'depth', '深度排序算法是软件研发影响力的一种，其通过度量软件的底层程度衡量其研发影响力。', '# 深度排序算法

深度排序算法是软件研发影响力的一种，其基于假设：底层性质的软件，例如网络协议等常常被间接引入的软件的变更将对大量软件的功能造成影响，因此对研发更具影响力。而在更多依赖路径中处于尾部的软件往往更容易被间接引入，更具备底层性质。深度排序算法试图发现**软件的底层程度**。

## 定义

在软件依赖图谱中，对于一条由依赖（有向边）组成的路径，可以将路径中包含的软件（节点）依据与路径起点软件的距离由短至长划分为头部软件、中间软件，和尾部软件。其中尾部软件经由中间软件而被头部软件间接引入，被间接引入的最多，最具有底层性质。本指标提出一种基于软件的“应用层级”的方法以帮助度量软件间的相对位置。

## 公式

$$
D_i=\\sqrt{\\frac{\\sum_{p∈q(i)}^n{D_p^2}}{N}}+1 \\tag{1}
$$

$$
C_{depth}=\\frac{\\sum_{p∈g(i)}^{n-1}{D_p-D_i}}{N-1}\\tag{2}
$$

式$(1)$中，$D_i$表示软件$i$的应用层级，$q(i)$表示软件$i$所直接依赖的软件集合，应用层级通过递推的方式计算软件间的相对位置。式$(2)$中，$N$是图中软件的总数，$g(i)$表示直接或间接依赖于软件$i$的软件集合，$C_{depth}$统计软件与所有直接或间接依赖该软件的应用层级的差值。使用应用层级而非直接使用软件在路径上的距离有以下好处：

* 软件通常依赖于多个第三方软件进行开发，因此会处于多条路径中，采用递推的方法可以对不同路径上软件的位置加以综合。

* 软件研发倾向于依赖更多低应用层级的软件，采用平方平均数可以放大被依赖的高应用层级的软件的影响。

### 复杂度

* 时间复杂度$O(N+E)$

', 3, '软件研发影响力', 2, 's3a://lib/depth-latest.jar', 'applerodite.depth.Main');


INSERT INTO graph.algo(id, name, define, detail, groupID, tag, tagID, jarPath, mainClass) VALUES (9, 'mediation', '中介度排序算法是软件研发影响力的一种，其通过度量软件的框架性质衡量其研发影响力。', '# 中介度排序算法

中介度排序算法是软件研发影响力的一种，其基于假设：框架性质的软件，例如开发框架等集成了众多功能的软件，对研发更具影响力。而中间软件通常帮助头部软件集成了大量尾部软件，因此更具备框架性质。中介度排序算法试图发现**软件的框架性质**。

## 定义

在软件依赖图谱中，对于一条由依赖（有向边）组成的路径，可以将路径中包含的软件（节点）依据与路径起点软件的距离由短至长划分为头部软件、中间软件，和尾部软件。其中中间软件作为头部软件和底层软件的中介，通常集成了多数底层软件的功能，能够帮助开发者更快速、便捷地解决问题，具备框架性质。本指标使用$betweenness$算法寻找依赖图谱中的中间软件，该算法被广泛应用于描述中间节点对头部及尾部节点的中介作用。

## 公式

$$
C_{mediation}=\\sum_{s!=i!=t}{\\frac{σ_{s,t}(i)}{σ_{s,t}}}\\tag{1}
$$

式中，$σ_{s,t}$表示软件$s$和软件$t$之间的最短路径数目，$σ_{s,t}(i)$则表示这些最短路径经过软件$i$的次数。$C_{mediation}$即计算图谱中的所有最短路径中经过目标软件的比例。

### 复杂度

* 时间复杂度$O(N*E)$

## 引用

* [Freeman L C. A set of measures of centrality based on betweenness[J]. Sociometry, 1977: 35-41. ](https://www.jstor.org/stable/3033543)
* [Brandes, U. (2001). A faster algorithm for betweenness centrality. *Journal of mathematical sociology*, *25*(2), 163-177.](https://pdodds.w3.uvm.edu/research/papers/others/2001/brandes2001a.pdf)', 3, '软件研发影响力', 2, 's3a://lib/mediation-latest.jar', 'applerodite.mediation.Main');


INSERT INTO graph.algo(id, name, define, detail, groupID, tag, tagID, jarPath, mainClass) VALUES (10, 'stability', '生态排序算法是软件研发影响力的一种，其通过度量软件的生态性质衡量其研发影响力。', '# 生态排序算法

生态排序算法是软件研发影响力的一种，其基于假设：生态性质的软件，即为了使用软件的相关生态而须要使用的软件。软件的相关软件越多，软件的生态性质就越强，研发影响力越高。生态排序算法试图发现**软件的生态性质**。

## 定义

在软件依赖图谱中，软件与软件之间通过依赖关系相互关联，而依赖通常表现为二次开发或引用，在二次开发关系中，软件改进第三方软件的功能从而纵向地扩展软件生态；而在引用关系中软件仅仅使用第三方软件的部分功能，关注的功能往往超过第三方软件原本的功能领域，是对生态横向地扩展。二次开发后的软件可能会被再次二次开发或引用，可以通过间接依赖对其展开描述，而引用的软件通常代表了完整的功能不再被依赖，可以通过直接依赖对其展开描述。为了探讨这些依赖如何反映软件的生态，我们参考Morone F等人提出的协同影响力算法，该算法从网络整体稳定性的角度出发描述了节点间的关联。

## 公式

$$
C_{stability}=(k_i-1)\\sum_{p∈δBall(i,l)}{(k_p-1)}\\tag{1}
$$

$$
λ(l)=(\\frac{\\sum_i{C_{stability}(i)}}{N<k>})^{\\frac{1}{l+1}}\\tag{2}
$$

式$(1)$中，$k_i$表示软件$i$的度数，$Ball(i, l)$表示距离软件$i$最短距离为$l$的软件集合，$C_{stability}$兼顾了软件的直接依赖和间接依赖。式$(2)$中，$N$是节点的总数，$<k>$是图当前的平均度，$λ(l)$是一个阈值，被用于判断网络中是否仍然具备生态。算法首先计算依赖网络中所有软件的$C_{stability}$ ，接着算法计算$λ(l)$，若$λ(l)>1$，则将$C_{stability}$最大的软件从网络中删除，并重新计算剩余所有软件的$C_{stability}$，直到$λ(l)≤1$，此时认为网络中不再存在具有生态的软件，不再重新计算软件的$C_{stability}$。

### 复杂度

* 时间复杂度$O(N*logN)$

## 参数

### 距离（l）

* 范围：$[2,INF]$

* 默认值：3。

$C_{stability}$统计与目标软件距离为$l$的软件的度数，算法的时间复杂度随距离的增加而增加。注意到软件间的依赖通常不会有过多的嵌套，当$l$取值过大时，可能找不到对应的软件。

## 引用

* [Morone, F., & Makse, H. A. (2015). Influence maximization in complex networks through optimal percolation. *Nature*, *524*(7563), 65-68.](https://www.nature.com/articles/nature14604)
* [Morone, F., Min, B., Bo, L., Mari, R., & Makse, H. A. (2016). Collective influence algorithm to find influencers via optimal percolation in massively large social media. *Scientific reports*, *6*(1), 30062.](https://www.nature.com/articles/srep30062)
', 3, '软件研发影响力', 2, 's3a://lib/stability-latest.jar', 'applerodite.stability.Main');

INSERT INTO graph.algoParam(algoID, name, `desc`, type, `default`, `min`, `max`) VALUES (10, 'radius', '半径，识别节点影响力时考虑与该节点距离为radius的其他节点，计算复杂度随radius增加而升高', 2, '3', '2', '5');

INSERT INTO graph.algo(id, name, define, detail, groupID, tag, tagID, jarPath, mainClass) VALUES (11, 'integrated', '综合排序算法是软件研发影响力的一种，其基于AHP方法融合breadth, depth, mediation, stability四种研发影响力指标。', '# 综合排序算法

综合排序算法是软件研发影响力的一种，其基于$AHP$方法对$breadth, depth, mediation, stability$四种软件研发影响力算法加以融合。综合排序算法可以有效综合各指标的优势，使之达到互补的效果

## AHP矩阵

$AHP$是一种基于专家意见的主观调权方法。基于对Python软件的观察，经过专家讨论，得到了以下意见：

> 表中，1表示同样重要，3表示稍微重要，5表示明显重要，2、4为中值。例如，基础性质相比底层性质和生态性质稍微重要，而相比框架性质明显重要

| 指标     | 基础性质 | 底层性质 | 框架性质 | 生态性质 |
| -------- | -------- | -------- | -------- | -------- |
| 基础性质 | 1        | 3        | 5        | 3        |
| 底层性质 | 1/3      | 1        | 2        | 2        |
| 框架性质 | 1/5      | 1/2      | 1        | 1/2      |
| 生态性质 | 1/3      | 1/2      | 2        | 1        |

## 公式

$$
C_{integrated}=a*C_{breadth}+b*C_{depth}+c*C_{mediation}+d*C_{stability}\\tag{1}
$$

根据$AHP$矩阵，建议取值为$a=0.52887$，$b=0.21942$，$c=0.09656$，$d=0.15515$

## 引用

* [Saaty, Thomas L. "Analytic hierarchy process. Encyclopedia of Biostatistics." *New York: McGraw-Hill, doi* 10.0470011815 (2005): b2a4a002.](https://onlinelibrary.wiley.com/doi/abs/10.1002/0470011815.b2a4a002)

', 3, '软件研发影响力', 2, 's3a://lib/integrated-latest.jar', 'applerodite.integrated.Main');

INSERT INTO graph.algoParam(algoID, name, `desc`, type, `default`, `min`, `max`) VALUES (11, 'weights', '集成的各个算法（breadth, depth, mediation, stability）的权重，范围为0~1', 4, null, '1', '4');

INSERT INTO graph.algo(id, name, define, detail, groupID, tag, tagID, jarPath, mainClass) VALUES (12, 'hhi', '“hhi”指数是赫芬达尔—赫希曼指数的简称，它是一种测量产业集中度的综合指数，可以反应同应用领域中头部软件在研发中的垄断程度。', '“hhi”指数是赫芬达尔—赫希曼指数的简称，它是一种测量产业集中度的综合指数，可以反应同应用领域中头部软件在研发中的垄断程度。它反映了一个行业中各市场竞争主体所占行业总收入或总资产百分比的平方和，用来计量市场份额的变化，即市场中厂商规模的离散度。它是经济学界和政府管制部门使用较多的指标。在此处，“hhi”指数被应用衡量同个主题下的软件的垄断程度。

“hhi”指数的公式如下：

$$HHI=\\sum_{i=1}^{n}{s_i^2}$$

其中，$n$表示主题下的软件数，$s_i$表示第$i$个软件的市场占有率，即其被依赖的次数占该主题下所有软件被依赖的总数。

“hhi”指数的内涵是，一个行业的集中度越高，即市场份额越集中在少数几家企业手中，那么“hhi”指数就越大，反之则越小。当市场处于完全垄断时，即只有一家企业，那么“hhi”指数等于1；当市场上有许多企业，且规模都相同时，那么“hhi”指数等于1/n，当n趋向于无穷大时，“hhi”指数就趋向于0。

“hhi”指数基于“结构—行为—绩效”理论，即市场结构影响到企业的行为，并最终决定市场的资源配置效率。随着市场份额的集中，企业会趋于采用相互勾结策略，最终制定出来的价格会偏离完全竞争市场价格，导致社会福利损失。因此，“hhi”指数可以在一定程度上反映市场的竞争程度和垄断力。

“hhi”指数可以用来分析不同行业的市场结构，找出存在垄断或寡头的行业，或者用它来评估企业的并购或兼并对市场竞争的影响，或者用它来制定反垄断的政策或法规。

', 3, '软件研发影响力', 2, 's3a://lib/hhi-latest.jar', 'applerodite.hhi.Main');


