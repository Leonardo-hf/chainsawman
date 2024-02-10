import {
    PageContainer
} from "@ant-design/pro-components";
import {Typography} from "antd";
import React from "react";
import Markdown from 'react-markdown'
import remarkMath from "remark-math";
import remarkGfm from 'remark-gfm'
import rehypeKatex from 'rehype-katex'
import styles from './juejin.css'
// import 'katex/dist/katex.min.css'

const {Title} = Typography;

const AlgoDoc: React.FC<{ algoId: string }> = (props) => {
    return <PageContainer>
        <Markdown className={styles.markdown} remarkPlugins={[remarkGfm, remarkMath]}
            // @ts-ignore
                  rehypePlugins={[rehypeKatex]}>{`
# 社区服务与支撑

该模型用于评估开发者在贡献过程中，直接感知到的社区提供的服务和支撑做得如何。之所以强调直接感知，是因为社区提供的许多底层服务，例如开发涉及的Devops基础设施同样是构建社区服务的关键元素，但社区参与者很难有直观感受，同时缺乏通识性的评估手段。我们使用在社区式开发中，参与者所能感知到的指标维度，来间接性的评估社区整个“开源贡献马拉松的后勤保障系统”。需要注意的是，这并不意味着只做到模型中提及的指标就足够了，模型为了保证指标间的独立性，做了强相关性指标降维处理；所以如果想长期保持该项模型的长期积极发展，社区付出的努力要远远超过当前指标包含的内容。

## 评估模型中的指标

### 更新 Issue 数量[](https://oss-compass.org/zh/docs/metrics-models/collaboration/productivity/community-service-and-support#更新-issue-数量)

- 定义：过去 90 天 Issue 更新的数量。
- 权重：19.721%
- 阈值：2000

有两个原因促使我们选择使用Issue更新的数量而不是统计关闭或者解决Issue的数量。首先，Issue有很多不同的类型，比如bug、功能需求、用户咨询和CVEs。只有特定类型的问题必须很快得到解决，比如CVEs。对于其余类型的问题, 并不追求Issue的快速解决，我们需要与问题创建者进行多次沟通，以更好地了解详细信息。如果是功能需求，从接受到解决，是按照发布计划进行的，这类场景可能也需要几个月的时间。其次，从Issue更新的数量来看，我们可以监控Issue处理的活跃度。问题更新还可以包括重开问题，表明对问题理解的变化的关注度。

### 关闭 PR 数量[](https://oss-compass.org/zh/docs/metrics-models/collaboration/productivity/community-service-and-support#关闭-pr-数量)

- 定义：过去 90 天内合并和拒绝的 PR 数量。
- 权重：19.721%
- 阈值：4500

代码贡献越多，需要关闭(接受或拒绝)的PR请求就越多。这表明社区正在积极地处理PR请求。我们将*关闭 PR 数量* 与 *更新 Issue 数量* 作为该模型的结果性指标，用来总体观察社区服务与支撑力度。

### Issue 首次响应时间[](https://oss-compass.org/zh/docs/metrics-models/collaboration/productivity/community-service-and-support#issue-首次响应时间)

- 定义：过去 90 天新建 Issue 首次响应时间的均值和中位数（天）。这不包括机器人响应、创建者自己的评论或 Issue 的分配动作（action）。如果 Issue 一直未被响应，该 Issue 不被算入统计。
- 权重：-14.372%
- 阈值：15 天

我们用这项指标来感知“社区温度”，对于加入社区的贡献者来说，他的提问如果能得到社区的及时回复，将会很大几率留存并持续贡献社区([依据Mozilla研究报告](https://docs.google.com/presentation/d/1hsJLv1ieSqtXBzd5YZusY-mB8e1VJzaeOmh8Q4VeMio/edit#slide=id.g43d857af8_0177))。 同时我们发现近年来越来越多的机器人被用来辅助Issue处理，所以我们排除了机器人的干扰，专注人的回复。

### Bug类Issue处理时间[](https://oss-compass.org/zh/docs/metrics-models/collaboration/productivity/community-service-and-support#bug类issue处理时间)

- 定义：过去 90 天新建的 Bug 类 Issue 处理时间的均值和中位数（天），包含已经关闭的 Issue 以及未解决的 Issue。
- 权重：-12.88%
- 阈值：60 天
- 注：标记为 Bug 类的 Issue。

Bug类Issue代表了社区对需要快速解决的Issue的处理效率。我们选择使用Bug类型的Issue来代表这类Issue，当然也具有一定的局限性，因为并不是所有Bug都是高优先级处理的Bug，但相比于不区分issue类型， 该指标具备一定代表性。

### PR 处理时间[](https://oss-compass.org/zh/docs/metrics-models/collaboration/productivity/community-service-and-support#pr-处理时间)

- 定义：过去 90 天新建 PR 的处理时间的均值和中位数（天），包含已经关闭的 PR 以及未解决的 PR
- 权重：-12.88%
- 阈值：30 天

我们追求PR的快速关闭，包括代码合并或拒绝。否则，解决PR请求所需的时间越长，发生合并冲突的风险就越大，而依赖于这个PR代码的其他代码提交请求也将堆积。

### Issue 评论频率[](https://oss-compass.org/zh/docs/metrics-models/collaboration/productivity/community-service-and-support#issue-评论频率)

- 定义：过去 90 天内新建 Issue 的评论平均数（不包含机器人和 Issue 作者本人评论）。
- 权重：10.217%
- 阈值：5

我们希望看到社区鼓励参与者围绕具体的Bug或者需求，通过Issue的方式进行公开和透明的讨论。这样Issue的相应结论也可以做为知识储备积累下来，同时为更多人所能看到。

### 代码审查评论频率[](https://oss-compass.org/zh/docs/metrics-models/collaboration/productivity/community-service-and-support#代码审查评论频率)

- 定义：过去 90 天内新建 PR 的评论平均数量（不包含机器人和 PR 作者本人评论）。
- 权重：10.217%
- 阈值：8

我们希望代码审查能够通过PR review的方式展示出来，让大家看到社区对于代码质量，安全方面管理的重视程度，同时为新人快速成长提供帮助。

## 评估模型算法

### 权重[](https://oss-compass.org/zh/docs/metrics-models/collaboration/productivity/community-service-and-support#权重)

我们使用 [AHP](https://en.wikipedia.org/wiki/Analytic_hierarchy_process) 来计算每个指标的权重。

#### AHP 输入数据[](https://oss-compass.org/zh/docs/metrics-models/collaboration/productivity/community-service-and-support#ahp-输入数据)


| 指标名称              | 更新 Issue 数量 | 关闭 PR 数量 | Issue 首次响应时间 | Bug 类 Issue 处理时间 | PR 处理时间 | Issue 评论频率 | 代码审查评论频率 |
| --------------------- | --------------- | ------------ | ------------------ | --------------------- | ----------- | -------------- | ---------------- |
| 更新 Issue 数量       | 1.000           | 1.000        | 1.333              | 1.500                 | 1.500       | 2.000          | 2.000            |
| 关闭 PR 数量          | 1.000           | 1.000        | 1.333              | 1.500                 | 1.500       | 2.000          | 2.000            |
| Issue 首次响应时间    | 0.750           | 0.750        | 1.000              | 1.143                 | 1.143       | 1.333          | 1.333            |
| Bug 类 Issue 处理时间 | 0.667           | 0.667        | 0.875              | 1.000                 | 1.000       | 1.250          | 1.250            |
| PR 处理时间           | 0.667           | 0.667        | 0.875              | 1.000                 | 1.000       | 1.250          | 1.250            |
| Issue 评论频率        | 0.500           | 0.500        | 0.750              | 0.800                 | 0.800       | 1.000          | 1.000            |
| 代码审查评论频率      | 0.500           | 0.500        | 0.750              | 0.800                 | 0.800       | 1.000          | 1.000            |

#### AHP 分析结果[](https://oss-compass.org/zh/docs/metrics-models/collaboration/productivity/community-service-and-support#ahp-分析结果)

| 指标名称              | 特征向量 | 权重     |
| --------------------- | -------- | -------- |
| 更新 Issue 数量       | 1.380    | 19.721%  |
| 关闭 PR 数量          | 1.380    | 19.721%  |
| Issue 首次响应时间    | 1.006    | -14.372% |
| Bug 类 Issue 处理时间 | 0.901    | -12.876% |
| PR 处理时间           | 0.901    | -12.876% |
| Issue 评论频率        | 0.715    | 10.217%  |
| 代码审查评论频率      | 0.715    | 10.217%  |

#### 一致性检验结果[](https://oss-compass.org/zh/docs/metrics-models/collaboration/productivity/community-service-and-support#一致性检验结果)

| 最大特征根 | CI 值 | RI 值 | CR 值 | 一致性检验结果 |
| ---------- | ----- | ----- | ----- | -------------- |
| 7.002      | 0.000 | 1.360 | 0.000 | PASS           |

### 阈值[](https://oss-compass.org/zh/docs/metrics-models/collaboration/productivity/community-service-and-support#阈值)

我们选择的阈值是基于不同类型开源项目的大数据观测。

## 参考文献

- [CHAOSS 度量模型：社区服务与支撑](https://chaoss.community/kb/metrics-model-community-service-and-support/)

## 贡献者

### 前端[](https://oss-compass.org/zh/docs/metrics-models/collaboration/productivity/community-service-and-support#前端)

- Shengxiang Zhang
- Feng Zhong
- Chaoqun Huang
- Huatian Qin
- Xingyou Lai

### 后端[](https://oss-compass.org/zh/docs/metrics-models/collaboration/productivity/community-service-and-support#后端)

- Yehui Wang
- Chenqi Shan
- Shengbao Li
- Huatian Qin

### 评估模型[](https://oss-compass.org/zh/docs/metrics-models/collaboration/productivity/community-service-and-support#评估模型)

- Yehui Wang
- Liang Wang
- Chenqi Shan
- Shengbao Li
- Matt Germonprez
- Sean Goggins 
                 
                  `}</Markdown>
    </PageContainer>
}

export default AlgoDoc