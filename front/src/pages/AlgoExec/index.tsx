import {
    PageContainer,
    ProForm,
    ProFormDependency, ProFormGroup,
    ProFormInstance,
    ProFormSelect,
    ProList
} from '@ant-design/pro-components';
import {
    Col,
    message, Row,
    Space, Tag, theme,
    Typography
} from 'antd';
import React, {useRef, useState} from 'react';
import {useModel} from '@@/exports';
import {formatDate} from '@/utils/format';
import {algoExec, dropAlgoTask, getAlgoTask} from '@/services/graph/graph';
import {getTaskTypeDesc, TaskTypeMap} from './_task';
import RankTable from '@/components/RankTable';
import {getParamFormItem, getParamFormValue, getTooltip} from "@/constants/sp";
import {isAlgoIllegal} from "@/models/global";

const {Text} = Typography;


const Exec: React.FC = () => {
    const {initialState} = useModel('@@initialState')
    // @ts-ignore
    const {algos} = initialState
    const [extKeysTask, setExtKeysTask] = useState<any[]>([])
    const {graphs} = useModel('global')
    const taskListRef = React.createRef()
    const graphOptions = graphs.map(g => {
        return {
            label: g.name,
            value: g.id
        }
    })
    const graphEnums: any = {}
    graphOptions.forEach(op => graphEnums[op.value] = op.label)

    type FormData = {
        graphId: number,
        algoId: number,
    }
    const getExecForm = () => {
        const formRef = useRef<ProFormInstance>()
        const onFinish = async (params: FormData) => {
            const algoSelect = algos.find((a: Graph.Algo) => a.id === params.algoId)!
            const pairs = getParamFormValue(params, algoSelect)
            return await algoExec({
                graphId: params.graphId,
                algoId: params.algoId,
                params: pairs
            }).then(() => {
                formRef.current?.resetFields()
                message.success('算法已提交')
                // @ts-ignore
                taskListRef.current?.reload()
                return true
            })
        }
        const {token} = theme.useToken()
        const contentStyle: React.CSSProperties = {
            lineHeight: '260px',
            color: token.colorTextTertiary,
            backgroundColor: token.colorFillAlter,
            borderRadius: token.borderRadiusLG,
            border: `1px dashed ${token.colorBorder}`,
            marginTop: 16,
            padding: 16,
            height: '80vh'
        }
        return <div style={contentStyle}>
            <ProForm<FormData> onFinish={onFinish} formRef={formRef}
                               submitter={{
                                   resetButtonProps: {
                                       style: {
                                           display: "none"
                                       }
                                   },
                               }}>
                <ProFormSelect rules={[{required: true}]} label={'Step1：选择图谱'} name={'graphId'}
                               options={graphOptions}/>
                <ProFormDependency name={['graphId']}>
                    {({graphId}) => {
                        if (graphId) {
                            const graph = graphs.find(g => g.id === graphId)!
                            const limit = 50
                            const algoId = algos.filter((a: Graph.Algo) => isAlgoIllegal(graph, a)).map((a: Graph.Algo) => {
                                return {
                                    label: getTooltip(a.name, a.detail.length > limit ? a.detail.substring(0, limit) + '...' : a.detail),
                                    value: a.id
                                }
                            })
                            return <ProFormSelect
                                rules={[{required: true}]}
                                label={'Step2：选择算法'}
                                name={'algoId'}
                                options={algoId}
                            />
                        }
                    }}
                </ProFormDependency>
                <ProFormDependency name={['algoId']}>
                    {({algoId}) => {
                        if (algoId) {
                            const algoSelect = algos.find((a: Graph.Algo) => a.id === algoId)!
                            return getParamFormItem(algoSelect, 'Step3：设置 ')
                        }
                    }}
                </ProFormDependency>
            </ProForm>
        </div>
    }

    const getTaskList = () => {
        const getTaskContent = (task: Graph.AlgoTask) => {
            return <RankTable file={task.output}/>
        }
        const getTaskGraph = (task: Graph.AlgoTask) => {
            try {
                return <Tag color={'pink'}>{graphs.find(g => g.id === task.graphId)!.name}</Tag>
            } catch (e) {
                console.log(e)
                return
            }
        }

        return <ProList<Graph.AlgoTask>
            headerTitle='执行结果'
            key='taskProList'
            itemLayout='vertical'
            // @ts-ignore
            actionRef={taskListRef}
            rowKey='id'
            style={{
                height: '80vh',
                overflowY: 'scroll',
                overflowX: 'hidden',
                padding: 16
            }}
            expandable={{
                expandedRowKeys: extKeysTask, onExpandedRowsChange: (expandedKeys) => {
                    // @ts-ignore
                    setExtKeysTask(expandedKeys)
                }
            }}
            search={{
                filterType: 'light',
            }}
            request={async (params = {time: Date.now()}) => {
                const tasks: Graph.AlgoTask[] = (await getAlgoTask({graphId: params.graphId})).tasks
                return {
                    data: tasks.filter((t: { status: number; }) => !params.subTitle || t.status == params.subTitle),
                    success: true,
                    total: tasks.length
                }
            }}
            metas={{
                title: {
                    search: false,
                    render: (_, row) => <Text>{algos.find((a: Graph.Algo) => a.id === row.algoId)?.name}</Text>
                },
                subTitle: {
                    title: '类别',
                    render: (_, row) => {
                        return <Space size={0}>
                            {getTaskTypeDesc(row.status)}
                            {getTaskGraph(row)}
                        </Space>
                    },
                    valueType: 'select',
                    valueEnum: TaskTypeMap,
                },
                graphId: {
                    title: '图谱',
                    valueType: 'select',
                    valueEnum: graphEnums
                },
                extra: {
                    render: (_: any, row: { createTime: number; id: any; status: any; }) => {
                        return <Space direction={'vertical'}>
                            <Text type={'secondary'}>{formatDate(row.createTime)}</Text>
                            <a style={{float: 'right'}} onClick={() => {
                                dropAlgoTask({
                                    id: row.id
                                }).then(() => {
                                    // @ts-ignore
                                    taskListRef.current?.reload()
                                })
                            }}>{row.status ? '删除' : '终止'}</a>
                        </Space>
                    },
                    search: false,
                },
                content: {
                    search: false,
                    render: (_, row) => getTaskContent(row)
                },
            }}
        />
    }
    return <PageContainer>
        <Row gutter={16} style={{height: '100%'}}>
            <Col span={12}>{getExecForm()}</Col>
            <Col span={12}>{getTaskList()}</Col>
        </Row>
    </PageContainer>
}

export default Exec