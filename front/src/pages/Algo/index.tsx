import {
    PageContainer, ProDescriptions
} from "@ant-design/pro-components";
import {Divider, Space, Typography} from "antd";
import React from "react";
import {
    getAlgoTypeDesc,
    ParamType,
    ParamTypeOptions
} from "@/constants";
import Markdown from 'react-markdown'
import remarkMath from "remark-math";
import rehypeKatex from 'rehype-katex'
import 'katex/dist/katex.min.css'

const {Title} = Typography;

type AlgoProps = {
    algo: Graph.Algo
}

const Algo: React.FC<AlgoProps> = (props) => {
    const getAlgoContent = (algo: Graph.Algo) => {
        // 设置指标的初始值
        const initValues: any = {}
        algo.params?.forEach(p => {
            if (p.initValue) {
                initValues[p.key] = p.initValue
            }
        })
        const getMaxMin = (type: number) => {
            switch (ParamTypeOptions[type].value) {
                case ParamType.DoubleList:
                case ParamType.StringList:
                    return {
                        maxLabel: '最大长度',
                        minLabel: '最小长度'
                    }
                default:
                    return {
                        maxLabel: '最大值',
                        minLabel: '最小值'
                    }
            }
        }
        return <Space direction={'vertical'}>
            <Space>
                <Title>{props.algo.name}</Title>
                {getAlgoTypeDesc(algo.tag)}
            </Space>
            <Markdown remarkPlugins={[remarkMath]}
                // @ts-ignore
                      rehypePlugins={[rehypeKatex]}>{`定义: ${algo.detail}`}</Markdown>
            {algo.params?.length! > 0 && <Divider/>}
            {
                algo.params?.map((p, i) => {
                        const {maxLabel, minLabel} = getMaxMin(p.type)
                        return <div key={i}>
                            <ProDescriptions dataSource={p} column={4}>
                                <ProDescriptions.Item dataIndex={'key'} label={`参数${i + 1}`} valueType={'text'}/>
                                <ProDescriptions.Item dataIndex={'keyDesc'} span={3} label='描述' valueType={'text'}/>
                                <ProDescriptions.Item label='类型' render={(_, t) => ParamTypeOptions[t['type']].label}/>
                                {p.initValue &&
                                <ProDescriptions.Item dataIndex={'initValue'} label='默认值' valueType={'text'}/>}
                                {p.max && <ProDescriptions.Item dataIndex={'max'} label={maxLabel} valueType={'text'}/>}
                                {p.min && <ProDescriptions.Item dataIndex={'min'} label={minLabel} valueType={'text'}/>}
                            </ProDescriptions>
                            <Divider/>
                        </div>
                    }
                )
            }
        </Space>
    }
    return <PageContainer>
        {getAlgoContent(props.algo)}
    </PageContainer>
}

export default Algo