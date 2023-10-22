import {Tag} from "antd";
import React from "react";

export enum ParamType {
    String,
    Double,
    Int,
}

export enum AlgoType {
    rank,
    cluster,
    metrics,
}

export const ParamTypeOptions = [
    {
        label: '字符串',
        value: ParamType.String
    },
    {
        label: '浮点数',
        value: ParamType.Double
    },
    {
        label: '整数',
        value: ParamType.Int
    }
]

export const AlgoOptions = [
    {
        label: '中心度算法',
        color: 'blue',
        value: AlgoType.rank,
    },
    {
        label: '聚类算法',
        color: 'purple',
        value: AlgoType.cluster,

    },
    {
        label: '网络结构特征',
        color: 'green',
        value: AlgoType.metrics,
    },
]

AlgoOptions.map((o: any) => {
    return {
        ...o,
        text: o.label,
        status: o.value
    }
})

export const AlgoTypeMap: any = {}
AlgoOptions.forEach((o: any) => AlgoTypeMap[o.status] = o)


export function getAlgoTypeDesc(type: AlgoType) {
    const s = AlgoTypeMap[type]
    return <Tag color={s.color}>{s.text}</Tag>
}

export const getPrecise = (t: ParamType) => {
    switch (t) {
        case ParamType.Int:
            return {precision: 0}
        case ParamType.Double:
            return {precision: 4, step: 1e-4}
    }
}