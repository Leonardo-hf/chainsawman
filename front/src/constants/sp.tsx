import {ParamType} from "@/constants/index";
import {ProFormDigit, ProFormSelect} from "@ant-design/pro-components";
import React from "react";
import {Space, Tooltip} from "antd";
import {InfoCircleOutlined} from "@ant-design/icons";

export const AlgoCountryStrangle = 11

export const AlgoStrangleIds = [11]

export const AlgoImpactIds = [7, 8, 9, 10]

export const AlgoMulImpactId = 12

export const AlgoMulStrangleId = 13

export const getTooltip = (title: string, tip: string) => {
    return <Space>
        <span>{title}</span>
        <Tooltip title={tip}>
            <InfoCircleOutlined/>
        </Tooltip>
    </Space>

}

export const getParamFormItem = (algo: Graph.Algo, pre: string = '') => {
    return algo.params?.map(p => {
            switch (p.type) {
                case ParamType.Int:
                    return <ProFormDigit rules={[{required: true}]} name={p.key} fieldProps={{precision: 0}}
                                         initialValue={p.initValue}
                                         label={getTooltip(pre + p.key, p.keyDesc)}
                                         max={p.max ? p.max : Number.MAX_SAFE_INTEGER}
                                         min={p.min ? p.min : Number.MIN_SAFE_INTEGER}/>
                case ParamType.Double:
                    return <ProFormDigit rules={[{required: true}]} name={p.key} fieldProps={{precision: 4, step: 1e-4}}
                                         initialValue={p.initValue}
                                         label={getTooltip(pre + p.key, p.keyDesc)}
                                         max={p.max ? p.max : Number.MAX_SAFE_INTEGER}
                                         min={p.min ? p.min : Number.MIN_SAFE_INTEGER}/>
                case ParamType.StringList:
                case ParamType.DoubleList:
                    return <ProFormSelect rules={[{required: true}]} name={p.key} label={getTooltip(pre + p.key, p.keyDesc)}
                                          fieldProps={{mode: "tags"}}/>
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