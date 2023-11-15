import {
    PageContainer,
    ProForm,
    ProFormDependency,
    ProFormInstance,
    ProFormRadio,
    ProFormSelect,
} from "@ant-design/pro-components"
import {Button, message, Spin, Steps, theme} from "antd"
import {algoExec, getMatchNodesByTag} from "@/services/graph/graph";
import React, {useRef, useState} from "react";
import {useModel} from "@@/exports";
import {isAlgoIllegal} from "@/models/global";
import ProCard from "@ant-design/pro-card";
import RankTable from "@/components/RankTable";
import {AlgoCountryStrangle, AlgoMulStrangleId} from "@/constants/sp";
import InputWeights from "@/components/InputWeights";
import {getWeights} from "@/components/InputWeights/InputWeights";


const Strangle: React.FC = () => {
    // 获得两种算法的详细信息
    const {initialState} = useModel('@@initialState')
    //@ts-ignore
    const {algos} = initialState
    const algoSingle: Graph.Algo = algos.find((a: Graph.Algo) => a.id === AlgoCountryStrangle)
    const algoMul: Graph.Algo = algos.find((a: Graph.Algo) => a.id === AlgoMulStrangleId)
    const libraries = algoSingle.params![0]
    const weightsOnImpact = algoMul.params![0]
    const weightsOnStrangle = algoMul.params![1]

    const {graphs} = useModel('global')
    const [lastTaskId, setLastTaskId] = useState<string>()
    const [lastFileName, setLastFileName] = useState<string>()
    const [lastTimer, setLastTimer] = useState<NodeJS.Timer>()

    // 执行算法的目标图谱
    const graphOptions = graphs.filter(g => isAlgoIllegal(g, algoSingle)).map(g => {
        return {
            label: g.name,
            value: g.id
        }
    })

    type Source = number
    const sourceManual: Source = 0, sourceImpact: Source = 1
    const sourceOptions = [{
        label: '人工录入',
        value: sourceManual
    }, {
        label: '高影响力软件',
        value: sourceImpact
    }]

    // 算法表单相关
    type FormData = {
        source: Source,
        graphId: number,
        libraries: string[],
        weightsOnImpact: number[],
        weightsOnStrangle: number[]
    }

    // 表单提交执行算法请求
    const strangle = async () => {
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
            params: [],
            algoId: 0
        }
        if (params.source === sourceManual) {
            req.algoId = algoSingle.id!
            req.params!.push({
                key: libraries.key,
                type: libraries.type,
                listValue: params.libraries
            })
        } else {
            req.algoId = algoMul.id!
            req.params!.push({
                key: weightsOnImpact.key,
                type: weightsOnImpact.type,
                listValue: getWeights(params.weightsOnImpact)
            }, {
                key: weightsOnStrangle.key,
                type: weightsOnStrangle.type,
                listValue: getWeights(params.weightsOnStrangle)
            })
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
            desc: '选择待识别软件',
        },
        {
            title: '第二步',
            desc: '选择卡脖子风险算法',
        },
        {
            title: '第三步',
            desc: '获得识别结果',
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
            <ProFormRadio.Group rules={[{required: true}]} label={'软件源'} name={'source'}
                                options={sourceOptions} initialValue={sourceManual}/>
            <ProFormDependency name={['graphId', 'source']}>
                {({graphId, source}) => {
                    if (source === sourceImpact) {
                        return <InputWeights headers={['广度', '深度', '中介度', '稳定性']} innerProps={{
                            name: 'weightsOnImpact',
                            label: '影响力算法权重',
                            rules: [{required: true}]
                        }}/>
                    }
                    if (source === sourceManual) {
                        return <ProFormSelect
                            label={'软件名单'}
                            name={'libraries'}
                            rules={[{required: true}]}
                            showSearch
                            debounceTime={300}
                            fieldProps={{
                                filterOption: () => {
                                    return true
                                },
                                mode: "multiple",
                                allowClear: true
                            }}
                            request={async (v) => {
                                let packs: any[] = []
                                if (!v.keyWords) {
                                    return packs
                                }
                                await getMatchNodesByTag({
                                    graphId: graphId,
                                    keywords: v.keyWords,
                                    nodeId: 4,
                                }).then(res => {
                                    packs.push({
                                        label: "library",
                                        options: res.matchNodes.map(m => {
                                            return {
                                                label: m.primaryAttr,
                                                value: m.primaryAttr
                                            }
                                        })
                                    })
                                })
                                return packs
                            }}/>
                    }
                }
                }
            </ProFormDependency>
        </div>
    }
    const getStep2 = () => {
        return <div style={{display: current === 1 ? '' : 'none'}}>
            <InputWeights headers={['国别']} innerProps={{
                name: 'weightsOnStrangle',
                label: '卡脖子风险算法权重',
                rules: [{required: true}]
            }}/>
        </div>
    }

    // 提交算法后轮询算法结果
    if (current === steps.length - 1 && lastTaskId && !lastTimer) {
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
                    {/*step1 选择待识别软件*/}
                    {
                        getStep1()
                    }
                    {/*step2 选择卡脖子风险算法*/}
                    {
                        getStep2()
                    }
                </ProForm>
            }
            {/*step3 获得识别结果*/}
            {
                current === steps.length - 1 && <ProCard title={'预览'} loading={!lastFileName}>
                    {lastFileName && <RankTable file={lastFileName}/>}
                </ProCard>
            }
        </div>
        <div style={{marginTop: 24}}>
            {current < steps.length - 2 && (
                <Button type="primary" onClick={() => {
                    setCurrent(current + 1)
                }}>
                    下一步
                </Button>
            )}
            {current === steps.length - 2 && (
                <Button type="primary"
                        onClick={async () => {
                            await strangle()
                            setCurrent(current + 1)
                        }}>
                    执行
                </Button>
            )}
            {current > 0 && (
                <Button disabled={current === steps.length - 1 && !lastFileName && !lastTimer} style={{margin: '0 8px'}}
                        onClick={() => setCurrent(current - 1)}>
                    上一步
                </Button>
            )}
        </div>
    </PageContainer>
}

export default Strangle