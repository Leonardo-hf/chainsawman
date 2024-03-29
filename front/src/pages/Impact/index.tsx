import {
    PageContainer,
    ProForm, ProFormDependency,
    ProFormInstance, ProFormRadio,
    ProFormSelect
} from "@ant-design/pro-components"
import {Button, message, Steps, theme} from "antd"
import {algoExec, getAlgoTaskByID} from "@/services/graph/graph";
import React, {useRef, useState} from "react";
import {useModel} from "@@/exports";
import {isAlgoIllegal} from "@/models/global";
import ProCard from "@ant-design/pro-card";
import RankTable from "@/components/RankTable";
import {AlgoImpactNames, AlgoMulImpactName, getParamFormItem, getParamFormValue} from "@/constants/sp";
import InputWeights from "@/components/InputWeights";
import Loading from "@ant-design/pro-card/es/components/Loading";

const Impact: React.FC = () => {
    const {initialState} = useModel('@@initialState')
    //@ts-ignore
    const {algos} = initialState
    const mulAlgo: Graph.Algo = algos.find((a: Graph.Algo) => a.name == AlgoMulImpactName)

    const {graphs} = useModel('global')
    const [lastTaskId, setLastTaskId] = useState<number>()
    const [lastFileName, setLastFileName] = useState<string>()
    const [lastTimer, setLastTimer] = useState<NodeJS.Timer>()
    const graphOptions = graphs.filter(g => isAlgoIllegal(g, mulAlgo)).map(g => {
        return {
            label: g.name,
            value: g.id
        }
    })

    type FormData = {
        graphId: number,
        algoId: number,
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
            algoId: params.algoId,
            params: getParamFormValue(params, algos.find((a: Graph.Algo) => a.id === params.algoId))
        }
        return await algoExec(req).then((res) => {
            message.success('算法已提交')
            setLastTaskId(res.task.id)
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
        const algoSelect = AlgoImpactNames.map(name => {
            const a = algos.find((a: Graph.Algo) => a.name == name)!
            return {
                label: a.name,
                value: a.id
            }
        }).concat({label: mulAlgo.name, value: mulAlgo.id})
        return <div style={{display: current === 0 ? '' : 'none'}}>
            <ProFormSelect rules={[{required: true}]} label={'图谱'} name={'graphId'}
                           options={graphOptions}/>
            <ProFormRadio.Group rules={[{required: true}]} label={'算法'} name={'algoId'}
                                options={algoSelect} initialValue={mulAlgo.id}/>
            <ProFormDependency name={['algoId']}>
                {({algoId}) => {
                    if (algoId === mulAlgo.id) {
                        return <InputWeights headers={AlgoImpactNames} innerProps={{
                            name: 'weights',
                            label: '影响力算法权重',
                            rules: [{required: true}]
                        }}/>
                    } else {
                        return getParamFormItem(algos.find((a: Graph.Algo) => a.id == algoId)!)
                    }
                }}
            </ProFormDependency>
        </div>
    }

    // 提交算法后轮询算法结果
    if (current === steps.length - 1 && lastTaskId && !lastTimer) {
        const timer = setInterval(() => {
            const params: Graph.GetAlgoTaskRequest = {id: lastTaskId}
            getAlgoTaskByID(params).then((res) => {
                if (res.task.status) {
                    setLastFileName(res.task.output)
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
                current === steps.length - 1 && <ProCard title={'预览'} loading={!lastFileName}>
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