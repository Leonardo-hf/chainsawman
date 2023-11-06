import {PageContainer, ProForm, ProFormDependency, ProFormGroup, ProFormSelect} from "@ant-design/pro-components"
import {Divider, message, Space, Tooltip} from "antd"
import {algoExec, getGraphTasks, getMatchNodesByTag} from "@/services/graph/graph";
import React, {useState} from "react";
import {useModel} from "@@/exports";
import {isAlgoIllegal} from "@/models/global";
import ProCard from "@ant-design/pro-card";
import RankTable from "@/components/RankTable";


const Strangle: React.FC = (props) => {
    const {initialState} = useModel('@@initialState')

    //@ts-ignore
    const {algos} = initialState
    const algo: Graph.Algo = algos.find((a: Graph.Algo) => a.name === 'strangle risk')
    const libraries = algo.params![0]
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
        libraries: string[]
    }

    const strangle = async (params: FormData) => {
        const req: Graph.ExecAlgoRequest = {
            graphId: params.graphId,
            params: [{
                key: libraries.key,
                type: libraries.type,
                listValue: params.libraries
            }],
            algoId: algo.id!
        }
        console.log(req)
        return await algoExec(req).then((res) => {
            message.success('算法已提交')
            setLastTaskId(res.base.taskId)
            setLastFileName(undefined)
        })
    }


    console.log(lastTaskId, lastFileName, lastTimer, !!(lastTaskId && !lastFileName && lastTimer))

    if (lastTaskId && !lastFileName && !lastTimer) {
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

    // const genAlgoFileList = async (graphId: number) => {
    //     const tasks = await getGraphTasks({graphId: graphId})
    //     return tasks.tasks.filter(t=>t.)
    // }

    return <PageContainer>
        <Space direction={"vertical"} style={{width: '100%'}}>
            <ProForm<FormData>
                disabled={!!(lastTaskId && !lastFileName && lastTimer)}
                onFinish={strangle}
                submitter={{
                    searchConfig: {
                        resetText: '重置',
                        submitText: '执行',
                    }
                }}
            >
                <ProFormSelect rules={[{required: true}]} label='图谱' name={'graphId'} options={graphOptions}/>
                <ProFormDependency name={['graphId']} >
                    {({graphId}) =>
                            <ProFormSelect
                                label={
                                    <Tooltip title={libraries.keyDesc}>
                                        <span>{libraries.key}</span>
                                    </Tooltip>
                                }
                                disabled={!graphId}
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
                </ProFormDependency>
                {/*<ProFormSelect name={'algoFile'} label={'从算法结果获取库'} options={genAlgoFileList(graphId)}/>*/}
                {/*<ProFormDependency name={['algoFile']}></ProFormDependency>*/}
            </ProForm>
            <Divider/>
            {
                lastTaskId && <ProCard title={'预览'}>
                    {lastFileName && <RankTable file={lastFileName}/>}
                </ProCard>
            }
        </Space>
    </PageContainer>
}

export default Strangle