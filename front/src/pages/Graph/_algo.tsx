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