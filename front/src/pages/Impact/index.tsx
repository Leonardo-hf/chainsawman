import {
    PageContainer,
    ProForm,
    ProFormInstance,
    ProFormSelect
} from "@ant-design/pro-components"
import {Button, message, Steps, theme} from "antd"
import {algoExec} from "@/services/graph/graph";
import React, {useRef, useState} from "react";
import {useModel} from "@@/exports";
import {isAlgoIllegal} from "@/models/global";
import ProCard from "@ant-design/pro-card";
import RankTable from "@/components/RankTable";
import {AlgoMulImpactId} from "@/constants/sp";
import {getWeights} from "@/components/InputWeights/InputWeights";
import InputWeights from "@/components/InputWeights";
import Loading from "@ant-design/pro-card/es/components/Loading";

const Impact: React.FC = () => {
    const {initialState} = useModel('@@initialState')
    //@ts-ignore
    const {algos} = initialState
    const algo: Graph.Algo = algos.find((a: Graph.Algo) => a.id === AlgoMulImpactId)
    const weights = algo.params![0]

    const {graphs} = useModel('global')
    const [lastTaskId, setLastTaskId] = useState<string>()
    const [lastFileName, setLastFileName] = useState<string>()
    const [lastTimer, setLastTimer] = useState<NodeJS.Timer>()
    const graphOptions = graphs.filter(g => isAlgoIllegal(g, algo)).map(g => {
        return {
            label: g.name,
            value: g.id
        }
    })

    type FormData = {
        graphId: number,
        weights: number[],
    }

    const impact = async () => {
        let params: FormData
        try {
            params = await formRef.current!.validateFieldsReturnFormatValue?.()
        } catch (e: any) {
            for (let f of e.errorFields) {
                if (f.errors.length) {
                    message.warning(f.errors[0])
                }
            }
            return
        }
        const req: Graph.ExecAlgoRequest = {
            graphId: params.graphId,
            params: [{
                key: weights.key,
                type: weights.type,
                listValue: getWeights(params.weights)
            }],
            algoId: algo.id!
        }
        return await algoExec(req).then((res) => {
            message.success('算法已提交')
            setLastTaskId(res.base.taskId)
            setLastFileName(undefined)
        })
    }

    // steps 表单
    const steps = [
        {
            title: '第一步',
            desc: '选择识别算法',
        },
        {
            title: '第二步',
            desc: '获得高影响力软件',
        },
    ]
    const [current, setCurrent] = useState(0)
    const items = steps.map((item) => ({key: item.title, title: item.title, description: item.desc}))
    const {token} = theme.useToken()
    const contentStyle: React.CSSProperties = {
        lineHeight: '260px',
        color: token.colorTextTertiary,
        backgroundColor: token.colorFillAlter,
        borderRadius: token.borderRadiusLG,
        border: `1px dashed ${token.colorBorder}`,
        marginTop: 16,
        padding: 16
    }
    const getStep1 = () => {
        return <div style={{display: current === 0 ? '' : 'none'}}>
            <ProFormSelect rules={[{required: true}]} label={'图谱'} name={'graphId'}
                           options={graphOptions}/>
            <InputWeights headers={['广度', '深度', '中介度', '稳定性']} innerProps={{
                name: 'weightsOnImpact',
                label: '影响力算法权重',
                rules: [{required: true}]
            }}/>
        </div>
    }

    // 提交算法后轮询算法结果
    if (current === steps.length - 1 && !lastTimer) {
        const timer = setInterval(() => {
            const params: Graph.ExecAlgoRequest = {algoId: 0, graphId: 0, taskId: lastTaskId}
            algoExec(params).then((res) => {
                if (res.file) {
                    setLastFileName(res.file)
                    clearInterval(timer)
                    setLastTimer(undefined)
                }
            }).catch(() => {
                clearInterval(timer)
                setLastTimer(undefined)
            })
        }, 10000)
        setLastTimer(timer)
    }

    const formRef = useRef<ProFormInstance>()

    return <PageContainer>

        <Steps current={current} items={items}/>
        <div style={contentStyle}>
            {
                current < steps.length - 1 && <ProForm<FormData>
                    formRef={formRef}
                    submitter={{
                        resetButtonProps: {
                            style: {
                                display: "none"
                            }
                        },
                        submitButtonProps: {
                            style: {
                                display: "none"
                            }
                        },
                    }}
                >
                    {/*step1 选择识别算法*/}
                    {
                        getStep1()
                    }
                </ProForm>
            }
            {/*step2 获得高影响力软件*/}
            {
                current === steps.length - 1 && <ProCard title={'预览'}>
                    {lastFileName ? <RankTable file={lastFileName}/> : <Loading/>}
                </ProCard>
            }
        </div>
        <div style={{marginTop: 24}}>
            {current === 0 && (
                <Button type="primary"
                        onClick={async () => {
                            await impact()
                            setCurrent(current + 1)
                        }}>
                    执行
                </Button>
            )}
            {current === 1 && (
                <Button disabled={!lastFileName && !lastTimer} style={{margin: '0 8px'}}
                        onClick={() => setCurrent(current - 1)}>
                    上一步
                </Button>
            )}
        </div>
    </PageContainer>
}

export default Impact