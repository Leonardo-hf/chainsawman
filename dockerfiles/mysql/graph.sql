create database if not exists graph;

use graph;

create table if not exists `group`
(
    id       int auto_increment
        primary key,
    name     varchar(255)  not null,
    `desc`   text          null,
    parentID int default 1 null comment '标识父策略组，子策略组继承父策略组的全部节点与边缘'
);

create index group_parentID_index
    on `group` (parentID);

create table if not exists algo
(
    id        int auto_increment
        primary key,
    name      varchar(255)  not null,
    define    varchar(1023)  not null,
    detail    text          null,
    jarPath   varchar(255)  null,
    mainClass varchar(255)  null,
    tag       varchar(255)  null,
    tagID     int           null,
    isTag     tinyint(1) default 0,
    groupId   int default 1 null comment '约束算法应用于某个策略组的图谱，此外：
1......应用于全部策略组',
    constraint algos_groups_id_fk
        foreign key (groupId) references `group` (id)
            on update cascade on delete set default
);

create table if not exists algoParam
(
    id        int auto_increment
        primary key,
    algoID    int           not null,
    name      varchar(255)  not null,
    `desc`    varchar(255)  null,
    type      int default 0 not null,
    `default` varchar(255)        null,
    `max`       varchar(255)        null,
    `min`       varchar(255)        null,
    constraint algoParam___id
        foreign key (algoID) references algo (id)
            on update cascade on delete cascade
);

create table if not exists edge
(
    id        int auto_increment
        primary key,
    groupID   int                         not null,
    name      varchar(255)                not null,
    `desc`    text                        null,
    direct    tinyint(1)   default 1      not null,
    display   varchar(255) default 'real' null,
    `primary` varchar(255)                null,
    constraint edges_groups_id_fk
        foreign key (groupID) references `group` (id)
            on update cascade on delete cascade
);

create table if not exists edgeAttr
(
    id     int auto_increment
        primary key,
    edgeID int           not null,
    name   varchar(255)  not null,
    `desc` text          null,
    type   int default 0 not null,
    constraint edges_attr_edges_id_fk
        foreign key (edgeID) references edge (id)
            on update cascade on delete cascade
);

create table if not exists graph
(
    id         int auto_increment
        primary key,
    name       varchar(255)                        not null,
    status     int       default 0                 not null,
    numNode    int       default 0                 not null,
    numEdge    int       default 0                 not null,
    groupID    int       default 0                 not null,
    createTime timestamp default CURRENT_TIMESTAMP null,
    updateTime timestamp default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP,
    constraint graphs_groups_id_fk
        foreign key (groupID) references `group` (id)
            on update cascade on delete cascade
);

create table if not exists `exec`
(
    id         int auto_increment
        primary key,
    status     int       default 0                 null,
    params     varchar(1024)                       null,
    algoID     int                                 not null,
    graphID    int                                 not null,
    output     varchar(255)                        null,
    appID        varchar(255)                      null,
    updateTime timestamp default CURRENT_TIMESTAMP not null,
    createTime timestamp default CURRENT_TIMESTAMP not null,
    constraint exec_algos_id_fk
        foreign key (algoID) references algo (id),
    constraint exec_graphs_id_fk
        foreign key (graphID) references graph (id)
)
    comment '算法执行结果';

create table if not exists node
(
    id        int auto_increment
        primary key,
    groupID   int                          not null,
    name      varchar(255)                 not null,
    `desc`    text                         null,
    display   varchar(255) default 'color' null,
    `primary` varchar(255)                 null,
    constraint nodes_groups_id_fk
        foreign key (groupID) references `group` (id)
            on update cascade on delete cascade
);

create table if not exists nodeAttr
(
    id     int auto_increment
        primary key,
    nodeID int           not null,
    name   varchar(255)  not null,
    `desc` text          null,
    type   int default 0 not null,
    constraint nodes_attr_nodes_id_fk
        foreign key (nodeID) references node (id)
            on update cascade on delete cascade
);INSERT INTO graph.group(id, name, `desc`, parentID) VALUES (2, 'normal', '标准图谱', 1);

INSERT INTO graph.node(id, groupID, name, `desc`, `primary`) VALUES (1, 2, 'normal', '标准节点', 'name');
INSERT INTO graph.nodeAttr(nodeID, name, `desc`, type) VALUES (1, 'name', '名称', 0);
INSERT INTO graph.nodeAttr(nodeID, name, `desc`, type) VALUES (1, 'desc', '描述', 0);

INSERT INTO graph.edge(id, groupID, name, `desc`, `primary`, `direct`) VALUES (1, 2, 'normal', '标准边', null, 1);


INSERT INTO graph.group(id, name, `desc`, parentID) VALUES (1, 'root', '根图谱', null);

INSERT INTO graph.algo(id, name, define, detail, groupID, tag, tagID, jarPath, mainClass) VALUES (4, 'degree', '“度中心度”算法是一种用于衡量网络中节点重要性的指标，它反映了一个节点与其他节点直接相连的程度。', '“度中心度”算法是一种用于衡量网络中节点重要性的指标，它反映了一个节点与其他节点直接相连的程度。

“度中心度”算法的公式如下：

$$C(i)=\\frac{k_i}{N−1}$$

其中，$N$表示网络中的节点总数，$k_i$表示节点$i$的度，即与它相连的边的数量。这个公式将节点$i$的度除以最大可能的度，即与其他$N−1$个节点都相连的情况，得到一个介于$0$和$1$之间的值，表示节点$i$与其他节点的连接比例。

“度中心度”认为一个节点的重要性取决于它的邻居节点的数量，而不考虑邻居节点的重要性或与其他节点的距离，它比较简单，时间复杂度为$O(N)$，但也有一些局限性，例如，它不能区分网络中的不同结构，如星形、环形或完全图，它们的节点度中心度都相同，但实际上它们的网络特性是不同的。

“度中心度”算法可以用它来分析社交网络中的用户影响力，找出与更多人有联系的用户，或者用它来分析互联网中的网站流量，找出与更多网站有链接的网站。还可以用它来分析交通网络中的节点重要性，找出与更多路线有连接的节点。', 1, '其他', 1, 's3a://lib/degree-latest.jar', 'applerodite.degree.Main');


INSERT INTO graph.algo(id, name, define, detail, groupID, tag, tagID, jarPath, mainClass) VALUES (5, 'pagerank', '“Pagerank”算法是一种用于衡量网络中节点重要性的指标，它反映了一个节点被其他节点引用的程度。', '“Pagerank”算法是一种用于衡量网络中节点重要性的指标，它反映了一个节点被其他节点引用的程度。

“Pagerank”算法的公式如下：

$$PR(i)=(1−d)+d*\\sum_{j∈M(i)}{\\frac{PR(j)}{L(j)}}$$

其中，$PR(i)$表示节点$i$的“Pagerank”值，$d$表示阻尼系数，一般取值为$0.85$，表示用户在某个页面继续点击链接的概率，$M(i)$表示链接到节点$i$的节点集合，$L(j)$表示节点$j$的出度，即链接到其他节点的数量。这个公式将节点$p_i$的“Pagerank”值定义为所有链接到它的节点的“Pagerank”值的加权和，再加上一个最小值$(1−d)$，用来处理没有出度的节点。

“Pagerank”算法认为一个节点的重要性取决于它的邻居节点的重要性和数量。这个算法比较复杂，其最优时间复杂度为$O(N+E)$，需要进行多次迭代计算，直到收敛为止。

“Pagerank”算法的使用场景有很多，例如，可以用它来分析互联网中的网页排名，找出被更多网页引用的网页，或者用它来分析社交网络中的用户影响力，找出被更多用户关注的用户。还可以用它来分析科学文献中的引文关系，找出被更多文献引用的文献。', 1, '其他', 1, 's3a://lib/pagerank-latest.jar', 'applerodite.pagerank.Main');

INSERT INTO graph.algoParam(algoID, name, `desc`, type, `default`, `min`, `max`) VALUES (5, 'iter', '迭代次数', 2, '10', '1', '100');
INSERT INTO graph.algoParam(algoID, name, `desc`, type, `default`, `min`, `max`) VALUES (5, 'prob', '阻尼系数', 1, '0.85', '0.1', '1');

INSERT INTO graph.algo(id, name, define, detail, groupID, tag, tagID, jarPath, mainClass) VALUES (6, 'louvain', '“louvain”算法是一种用于社区发现的算法，它可以在大规模的网络中快速地找出具有紧密连接的节点集合(即社区)', '“louvain”算法是一种用于社区发现的算法，它可以在大规模的网络中快速地找出具有紧密连接的节点集合(即社区)。它的基本思想是利用模块度作为评价指标，通过不断地调整节点的社区归属，使得模块度达到最大值。

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

“louvain”算法用它来分析社交网络中的用户群体，找出具有相似兴趣或行为的用户，或者用它来分析知识图谱中的实体关系，找出具有相似属性或语义的实体。还可以用它来分析生物网络中的基因或蛋白质，找出具有相似功能或结构的基因或蛋白质。', 1, '其他', 1, 's3a://lib/louvain-latest.jar', 'applerodite.louvain.Main');

INSERT INTO graph.algoParam(algoID, name, `desc`, type, `default`, `min`, `max`) VALUES (6, 'maxIter', '外部迭代次数', 2, '10', '1', '100');
INSERT INTO graph.algoParam(algoID, name, `desc`, type, `default`, `min`, `max`) VALUES (6, 'internalIter', '内部迭代次数', 2, '5', '1', '50');
INSERT INTO graph.algoParam(algoID, name, `desc`, type, `default`, `min`, `max`) VALUES (6, 'tol', '系数', 1, '0.5', '0', '1');

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


INSERT INTO graph.algo(id, name, define, detail, groupID, tag, tagID, jarPath, mainClass) VALUES (13, 'diversity', '团队研发人员国别占比差距是软件卡脖子风险中的一个指标，衡量了软件开发团队中国家政治因素导致的断供风险。', '# 团队研发人员国别占比差距

团队研发人员国别占比差距是软件卡脖子风险中的一个指标，衡量了软件开发团队中国家政治因素导致的断供风险。

## 定义

开源软件的开发与维护由社区的主要维护者主导，仅维护者具有更改仓库可见性、提交代码与其合并贡献者的分支的权限，因此组织在维护者团队中的占比越大，则其对该开源项目具有越大的话语权。同理，属于同一国籍的开发者在维护团队中的占比越高，甚至明显超过其他国籍开发者则可能面临国家命令导致的断供危险。我们通过使用软件的开发者中国别数据判断该组织国别的多元性，多元性越大则断供风险越低。

## 公式

$$
Risk(s)=(1-\\frac{Developer(s, China)}{\\sum_i{Developer(s, i)}})*(1-\\frac{\\sum_i{Developer(s, i)^2}}{(\\sum_i{Developer(s, i))^2}})\\tag{1}
$$

式中，$Developer(s, China)$表示软件$s$。公式由两部分组成，第一部分软件的开发团队中中国外的开发者占比，第二部分描述了开发团队中开发者国别的多元性。

## 参数

### 软件集合（libraries）

* 一组字符串

算法研究一组软件的开发团队的研发人员国别占比差距。

## 输出

| 参数名  | 类型   | 定义               |
| ------- | ------ | ------------------ |
| library | 字符串 | 目标软件           |
| s1      | 浮点数 | 中国开发者占比     |
| s2      | 浮点数 | 开发者国别的多元性 |
| score   | 浮点数 | 国别占比差距       |', 4, '软件卡脖子风险', 3, 's3a://lib/diversity-latest.jar', 'applerodite.diversity.Main');

INSERT INTO graph.algoParam(algoID, name, `desc`, type, `default`, `min`, `max`) VALUES (13, 'libraries', '待识别卡脖子风险的软件列表', 3, null, '1', null);

INSERT INTO graph.algo(id, name, define, detail, groupID, tag, tagID, jarPath, mainClass) VALUES (14, 'supply', '国家软件供给关系是软件卡脖子风险中的一个指标，衡量了软件供应关系导致的断供风险。', '# 国家软件供给关系

国家软件供给关系是软件卡脖子风险中的一个指标，衡量了软件供应关系导致的断供风险。

## 定义

供应风险与供应链的上下游的供给关系息息相关。若下游完全依赖于单一上游，则下游存在极大的断供风险，反之若下游依赖于多个上游，则分摊了供应风险，下游具有较多的自主性。若上下游间相互存在业务依赖，则提高了单方断供的成本，使断供风险进一步降低。为研究中国与其他国家在开源软件中的供给关系，我们通过计算中国对单个国家软件的依赖在其所有依赖中的比重评估中国对该国家的依赖程度，进一步地，我们研究了该国家对中国开源软件的依赖程度，并判断该国家对其他国家的依赖是单一还是多元。若中国依赖某个国家较多，且该国家对中国的依赖较少且自身自主性较强，则该国家对中国的断供风险极高。

## 公式

$$
Risk(A,B)=\\frac{Depends(B, A)}{\\sum_i{Depends(B, i)}}*(1-\\frac{Depends(A, B)}{\\sum_i{Depends(A, i)}})*(1-\\frac{\\sum_i{Depends(A, i)^2}}{(\\sum_i{Depends(A, i))^2}})\\tag{1}
$$

式中，$Depends(B, A)$表示$B$国依赖$A$国开源软件的数目。公式由三部分组成，第一部分描述了$B$国对$A$国开源软件的依赖程度，第二部分描述了$A$国对$B$国开源软件的依赖程度，第三部分则表示了$A$国开源软件的自主性。

## 参数

### 源国家（source）

* 类型：字符串

### 目标国家（targets）

* 类型：一组字符串

算法研究所有目标国家对于源国家的断供风险

## 输出

| 参数名  | 类型   | 定义                       |
| ------- | ------ | -------------------------- |
| country | 字符串 | 目标国家名称               |
| s1      | 浮点数 | 对目标国家的依赖程度       |
| s2      | 浮点数 | 目标国家对自身的依赖程度   |
| s3      | 浮点数 | 目标国家开源软件的自主性   |
| score   | 浮点数 | 目标国家对源国家的断供风险 |

', 4, '软件卡脖子风险', 3, 's3a://lib/diversity-latest.jar', 'applerodite.diversity.Main');

INSERT INTO graph.algoParam(algoID, name, `desc`, type, `default`, `min`, `max`) VALUES (14, 'source', '源国家', 3, null, null, null);
INSERT INTO graph.algoParam(algoID, name, `desc`, type, `default`, `min`, `max`) VALUES (14, 'targets', '目标国家列表', 3, null, null, null);

INSERT INTO graph.algo(id, name, define, detail, isTag) VALUES (1, '其他', '一系列从不同角度研究网络性质的算法。', null, 1);

INSERT INTO graph.algo(id, name, define, detail, isTag) VALUES (2, '软件研发影响力', '软件研发影响力是对一个软件对其他软件功能的影响大小的度量。', '软件研发影响力是对一个软件对其他软件功能的影响大小的度量。具体表现为开发者在实现软件功能的过程中是否须要依赖某些第三方软件以及更倾向于使用哪些第三方软件。具备高研发影响力的软件应当表现为被开发者广泛地使用或间接使用以开发其软件功能。一个具有高研发影响力的开源软件发生闭源或被废弃，将对大量其他软件的功能的可靠性造成负面影响。

进一步地，我们探讨软件在开发过程中如何被开发者使用。若开发者为了实现一个软件功能需要使用数据库，而访问数据库时则必须使用数据库驱动，数据库驱动对该软件的研发具有影响力，而开发者在开发过程中使用Lint工具来保障软件质量，这个工具对软件功能本身没有帮助，则可以认为Lint工具对软件的研发不具有影响力。我们认为具备研发影响力的软件可能存在四种性质，

1. 基础性质的软件，例如工具库等在功能开发中被普遍使用的软件
2. 底层性质的软件，例如网络协议等往往在被间接使用的软件
3. 框架性质的软件，例如开发框架等集成了众多功能的软件
4. 生态性质的软件，即为了使用软件的相关生态而须要使用的软件

假设一：具有基础、底层、框架、生态性质的软件相比其他软件更具备研发影响力

进一步地，我们基于依赖图谱，从软件被依赖的拓扑形式研究软件在开发中扮演的角色

假设二：入度越大的软件越具备基础性质

假设三：在更多的路径上处于尾部的软件具有更多的底层性质

假设四：在更多的路径上处于中间的软件具有更多的框架性质

假设五：对依赖网络稳定性贡献越大的软件具有更多的生态性质

进一步，基于专家经验AHP方法，我们认为：

假设五：具有基础性质的软件的研发影响力 > 具有底层性质的软件的研发影响力 > 具有生态性质的软件的研发影响力 > 具有框架性质的软件的研发影响力

进一步，对指标识别的高影响力软件，进一步分类，发现每种指标识别出的软件既在主观上对研发很有帮助（很常见），也符合开发者的认知（第三方评估指标），此外，还符合假设中提及的软件在开发中扮演的角色（比如基础类库属于基础性质，网络协议属于底层性质）。', 1);

INSERT INTO graph.algo(id, name, define, detail, isTag) VALUES (3, '软件卡脖子风险', '卡脖子软件指发生断供后会对下游依赖组织造成重大影响且该影响无法在短时间内消除的软件，软件卡脖子风险算法即计算开源软件成为行业内卡脖子软件的可能性。', '# 软件卡脖子风险

卡脖子软件指发生断供后会对下游依赖组织造成重大影响且该影响无法在短时间内消除的软件。为满足以上特征，我们认为一个卡脖子软件需要具备以下三种特征。

* 假设1：卡脖子软件需要具有高商业价值，通过使用该软件，组织可以提高自身在市场上的竞争力，具有高商业价值的软件断供会对竞争对手造成有力打击。
  * 假设1-1：代表着尖端技术、具有高产值的软件具有高商业价值。

  * 假设1-2：生产力工具、能够帮助提高生产效率的软件具有高商业价值。

* 假设2：卡脖子软件具有高开发成本，这会导致造成的危害得以持续。
  * 假设2-1：代码量大、组件结构复杂的软件具有高开发成本。

  * 假设2-2：有系列配套软件及活跃下游用户的软件具有高开发成本。

* 假设3：卡脖子软件需要具备实施断供的条件。

  * 假设3-1：符合供给关系的软件才可以卡脖子，即只有供应方对被供应方卡脖子。

  * 假设3-2：软件团队内部分人员或第三方组织有强话语权才能实施卡脖子。

![](/assets/strangle-procedure.png)', 1);

