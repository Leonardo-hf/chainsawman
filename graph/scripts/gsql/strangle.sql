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

