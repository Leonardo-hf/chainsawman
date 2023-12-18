import {Tag} from "antd";
import React from "react";
import {formatEnum2Options} from "@/utils/format";

export enum ParamType {
    String,
    Double,
    Int,
    StringList,
    DoubleList
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
    },
    {
        label: '字符串列表',
        value: ParamType.StringList
    },
    {
        label: '浮点数列表',
        value: ParamType.DoubleList
    }
]

export const AlgoTypeMap: any = {
    软件风险: {
        text: '软件风险',
        color: 'pink',
        status: '软件风险',
    },
    软件影响力: {
        text: '软件影响力',
        color: 'blue',
        status: '软件影响力',
    },
    社区发现: {
        text: '社区发现',
        color: 'purple',
        status: '社区发现',
    },
    网络拓扑性质: {
        text: '网络拓扑性质',
        color: 'green',
        status: '网络拓扑性质',
    },
    节点中心度: {
        text: '节点中心度',
        color: 'pink',
        status: '节点中心度',
    }
}
export const AlgoOptions = formatEnum2Options(AlgoTypeMap)

export function getAlgoTypeDesc(type: string) {
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

export const RootGroupID = 1