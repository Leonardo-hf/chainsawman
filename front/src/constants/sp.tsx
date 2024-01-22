import {ParamType} from "@/constants/index";
import {ProFormDigit, ProFormSelect} from "@ant-design/pro-components";
import React from "react";
import {Space, Tooltip} from "antd";
import {InfoCircleOutlined} from "@ant-design/icons";

export const AlgoStrangleNames = ["diversity"]

export const AlgoImpactNames = ["breadth", "depth", "mediation", "stability"]

export const AlgoMulImpactName = "integrated"

const limit = 50

export const limitStr = (s: string) => {
    const lines = s.split('\n')
    if (lines) {
        return lines[0]
    }
    return s.length > limit ? s.substring(0, limit) + '...' : s
}

export const getTooltip = (title: string, tip: string) => {
    const shortTip = limitStr(tip)
    return <Space>
        <span>{title}</span>
        <Tooltip title={shortTip}>
            <InfoCircleOutlined/>
        </Tooltip>
    </Space>

}

export const getParamFormItem = (algo: Graph.Algo, pre: string = '') => {
    return algo.params?.map(p => {
            switch (p.type) {
                case ParamType.Int:
                    return <ProFormDigit rules={[{required: true}]} name={p.key} fieldProps={{precision: 0}}
                                         initialValue={p.initValue} key={p.key}
                                         label={getTooltip(pre + p.key, p.keyDesc)}
                                         max={p.max ? p.max : Number.MAX_SAFE_INTEGER}
                                         min={p.min ? p.min : Number.MIN_SAFE_INTEGER}/>
                case ParamType.Double:
                    return <ProFormDigit rules={[{required: true}]} name={p.key} fieldProps={{precision: 4, step: 1e-4}}
                                         initialValue={p.initValue} key={p.key}
                                         label={getTooltip(pre + p.key, p.keyDesc)}
                                         max={p.max ? p.max : Number.MAX_SAFE_INTEGER}
                                         min={p.min ? p.min : Number.MIN_SAFE_INTEGER}/>
                case ParamType.StringList:
                case ParamType.DoubleList:
                    return <ProFormSelect rules={[{required: true}]} name={p.key} label={getTooltip(pre + p.key, p.keyDesc)}
                                          fieldProps={{mode: "tags"}} key={p.key} options={[]}/>
            }
        }
    )
}

export const getParamFormValue = (form: any, algo: Graph.Algo) => {
    const pairs: Graph.Param[] = []
    if (algo.params) {
        for (let k in form) {
            const type = algo.params!.find(p => p.key === k)?.type
            if (type === undefined) {
                continue
            }
            switch (type) {
                case ParamType.Int:
                case ParamType.Double:
                    pairs.push({type: type, key: k, value: form[k].toString()})
                    break
                case ParamType.StringList:
                case ParamType.DoubleList:
                    pairs.push({type: type, key: k, listValue: form[k].toString()})
            }
        }
    }
    return pairs
}