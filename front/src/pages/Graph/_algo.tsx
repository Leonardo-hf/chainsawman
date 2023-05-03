import {
    algoAvgCC,
    algoBetweenness,
    algoCloseness,
    algoDegree,
    algoLouvain,
    algoPageRank,
    algoVoteRank
} from "@/services/graph/graph"

export type Algo = {
    id: number
    title: string
    description: string
    type: AlgoType
    action: any
    params: Param[]
}

type Param = {
    field: string
    text: string
    desc: string
    type: ParamType
    default: number
    min: number
    max: number
}

export enum ParamType {
    Int,
    Double
}

export enum AlgoType {
    rank,
    cluster,
    metrics,
}

export const AlgoTypeMap = {
    0: {
        text: '中心度算法',
        color: 'blue',
    },
    1: {
        text: '模块化算法',
        color: 'purple',
    },
    2: {
        text: '图属性',
        color: 'green',
    }
}

export function getAlgoTypeDesc(type: AlgoType) {
    return AlgoTypeMap[type]
}

export const algos: Algo[] = [
    {
        id: 1,
        title: 'degree',
        description: '度中心度算法，测量网络中一个节点与所有其它节点相联系的程度。',
        type: AlgoType.rank,
        action: algoDegree,
        params: []
    },
    {
        id: 2,
        title: 'pagerank',
        description: 'PageRank是Google使用的对其搜索引擎搜索结果中的网页进行排名的一种算法。能够衡量集合范围内某一元素的相关重要性。',
        type: AlgoType.rank,
        action: algoPageRank,
        params: [
            {
                field: 'iter',
                text: '迭代次数',
                desc: '',
                type: ParamType.Int,
                default: 3,
                min: 1,
                max: 100
            },
            {
                field: 'prob',
                text: '阻尼系数',
                desc: '',
                type: ParamType.Double,
                default: 0.85,
                min: 0.1,
                max: 1
            },
        ]
    },
    {
        id: 3,
        title: 'betweenness',
        description: '中介中心性用于衡量一个顶点出现在其他任意两个顶点对之间的最短路径的次数。',
        type: AlgoType.rank,
        action: algoBetweenness,
        params: []
    },
    {
        id: 4,
        title: 'closeness',
        description: '接近中心性反映在网络中某一节点与其他节点之间的接近程度。',
        type: AlgoType.rank,
        action: algoCloseness,
        params: []
    },

    {
        id: 5,
        title: 'voterank',
        description: '使用投票系统创建的基于邻居意见的排序算法。',
        type: AlgoType.rank,
        action: algoVoteRank,
        params: []
    },
    {
        id: 6,
        title: 'average clustering coefficient',
        description: '平均聚类系数。描图中的节点与其相连节点之间的聚集程度。',
        type: AlgoType.metrics,
        action: algoAvgCC,
        params: []
    },
    {
        id: 7,
        title: 'louvain',
        description: '一种基于模块度的社区发现算法。其基本思想是网络中节点尝试遍历所有邻居的社区标签，并选择最大化模块度增量的社区标签。',
        type: AlgoType.cluster,
        action: algoLouvain,
        params: [
            {
                field: 'maxIter',
                text: '外部迭代次数',
                desc: '社团合并后再次迭代的次数',
                type: ParamType.Int,
                default: 10,
                min: 1,
                max: 100
            },
            {
                field: 'internalIter',
                text: '内部迭代次数',
                desc: '',
                type: ParamType.Int,
                default: 5,
                min: 1,
                max: 50
            },
            {
                field: 'tol',
                text: '最小增加量',
                desc: '相对模块度最小增加量',
                type: ParamType.Double,
                default: 0.3,
                min: 0.1,
                max: 1
            },
        ]
    },
]