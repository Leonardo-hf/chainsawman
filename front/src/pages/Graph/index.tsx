import {PageContainer, ProForm, ProFormDigit, ProFormSelect, ProList, QueryFilter} from '@ant-design/pro-components';
import Graphin, {Behaviors} from '@antv/graphin';
import {Button, Card, Col, Divider, InputNumber, message, Row, Space, Spin, Tag, Tooltip, Typography} from 'antd';
import React, {SetStateAction, useEffect, useState} from 'react';
import {connect, useModel} from '@@/exports';
import {formatDate, formatNumber, getRandomColor} from '@/utils/format';
import {Algo, algos, AlgoType, AlgoTypeMap, getAlgoTypeDesc, ParamType} from './_algo';
import {dropGraph, dropTask, getGraphTasks, getMatchNodes} from '@/services/graph/graph';
import {getTaskTypeDesc, TaskTypeMap} from './_task';
import MetricTable from '@/components/MetricTable';
import RankTable from '@/components/RankTable';
import {history} from 'umi';
import {findByGid, getGraphName, isSubGraph} from '@/models/graph';
import LayoutSelector from "@/components/LayoutSelector";
import {layouts, layoutsConfig} from "@/pages/Graph/_layout";


const {Text} = Typography;
const {Hoverable} = Behaviors;

type Props = {
    graph: Graph.Graph,
    details: any,
    dispatch: any
}

const Graph: React.FC<Props> = (props) => {
    console.log('reload graph')
    const {graph, details, dispatch} = props
    const [renderMax, setRenderMax] = useState<number>(1000)
    const [renderTop, setRenderTop] = useState<number>(10)
    const [renderDirection, setRenderDirection] = useState<string>(' ')
    const [renderSrcNode, setRenderSrcNode] = useState<{ id: number, primaryAttr: string, tag: string }>()
    const [select, setSelect] = useState<boolean>(false)
    const [selectNode, setSelectNode] = useState<{ id: number, deg: number, tag: string, attrs: Graph.Pair[] }>()
    const [extKeysAlgo, setExtKeysAlgo] = useState<any[]>([])
    const [extKeysTask, setExtKeysTask] = useState<any[]>([])
    const [graphLayout, setGraphLayout] = useState<string>('graphin-force')
    const [tab, setTab] = useState<string>('attr')
    const [gid, setGID] = useState<string>(getGraphName(graph.id))

    const graphinRef = React.createRef(), taskListRef = React.createRef()
    const graphDetail = findByGid(details, gid)

    useEffect(() => {
        const handleNodeClick = (e: { item: { get: (arg0: string) => string; }; }) => {
            const id = e.item.get('id')
            const node = graphDetail.nodes.find((n: { id: number; }) => n.id.toString() === id)
            setSelect(true)
            setSelectNode(node)
        }
        // @ts-ignore
        const {graph} = graphinRef.current
        graph.on('node:click', handleNodeClick);
        graph.on('canvas:click', () => {
            setSelect(false)
        })
        return () => {
            graph.off('node:click');
            graph.off('canvas:click')
        };
    }, [graphDetail])

    const timerManager: NodeJS.Timer[] = []
    const {groups} = useModel('global')
    const {id, name} = graph
    // 判断是否为子图
    const isSub = isSubGraph(gid)
    // 判断是否数据加载完毕
    const loading = !graphDetail?.status
    const hasTimer = graphDetail?.timer
    if (loading && !hasTimer) {
        // 图模式
        if (!isSub) {
            let timer: NodeJS.Timer
            const queryDetail = async () => {
                dispatch({
                    type: 'graph/queryGraph',
                    payload: {
                        graphId: id,
                        taskId: graphDetail?.taskId,
                        top: renderTop,
                        max: renderMax,
                    },
                    timer: timer
                }).then((taskStatus: number) => {
                    if (taskStatus) {
                        clearInterval(timer)
                    }
                })
            }
            if (graphDetail) {
                timer = setInterval(queryDetail, 5000)
                timerManager.push(timer)
            } else {
                queryDetail().then()
            }
        } else {
            let timer: NodeJS.Timer
            const queryNeibors = async () => {
                dispatch({
                    type: 'graph/queryNeighbors',
                    payload: {
                        graphId: id,
                        taskId: findByGid(details, gid)?.taskId,
                        nodeId: renderSrcNode!.id,
                        direction: renderDirection,
                        max: renderMax,
                    },
                    timer: timer
                }).then((taskStatus: number) => {
                    if (taskStatus) {
                        clearInterval(timer)
                    }
                })
            }
            if (graphDetail) {
                timer = setInterval(queryNeibors, 5000)
                timerManager.push(timer)
            } else {
                queryNeibors().then()
            }
        }
    }
    // Tab栏
    const tabListNoTitle = [
        {
            key: 'attr',
            tab: '属性',
        },
        {
            key: 'algo',
            tab: '算法',
        },
        {
            key: 'task',
            tab: '任务',
        },
    ]
    // 图谱布局
    const layout = layoutsConfig.find(item => item.type === graphLayout)
    // 数据准备
    const nodes: any[] = []
    const edges: any[] = []
    const data = {
        nodes: nodes,
        edges: edges,
    }
    // 如果数据准备好，则加载数据
    if (!loading) {
        const group: Graph.Group = groups.find((g: Graph.Group) => g.id === graph.groupId)!
        const data = findByGid(details, gid)
        // 设置节点数据
        data.nodes.forEach((n: { tag: string; attrs: Graph.Pair[]; id: number; deg: number }) => {
            const tag = n.tag
            const nodeType = group.nodeTypeList.find(nt => nt.name === tag)!
            // TODO：display
            const display = nodeType.display
            const labelAttr = nodeType.attrs?.find(a => a.primary)
            let v = ''
            if (labelAttr) {
                v = n.attrs.find(a => a.key === labelAttr.name)!.value
            }
            const node = {
                id: n.id.toString(),
                style: {
                    keyshape: {
                        size: Math.floor((Math.log(n.deg + 1) + 1) * 10),
                        fill: getRandomColor()
                    },
                    label: {
                        value: v
                    }
                }
            }
            nodes.push(node)
        })
        // 设置边数据
        data.edges.forEach((n: { tag: string; attrs: Graph.Pair[]; source: number; target: number }) => {
            const tag = n.tag
            const edgeType = group.edgeTypeList.find(nt => nt.name === tag)!
            const display = edgeType.display
            const labelAttr = edgeType.attrs?.find(a => a.primary)
            let v = ''
            if (labelAttr) {
                v = n.attrs.find(a => a.key === labelAttr.name)!.value
            }
            let edge: any = {
                source: n.source.toString(),
                target: n.target.toString(),
                style: {
                    label: {
                        value: v
                    }
                }
            }
            if (display === 'dash') {
                edge.style = {
                    ...edge.style,
                    keyshape: {
                        lineDash: [4, 4],
                    }
                }
            }
            edges.push(edge)
        })
    }
    const resetGraph = () => {
        // 重置图谱点击
        // @ts-ignore
        graphinRef.current.graph.emit('canvas:click')
        // 删除图谱数据
        dispatch({
            type: 'graph/resetGraph',
            payload: id
        })
        // 重置图谱ID
        setGID(getGraphName(id))
        // 重置定时器
        for (const timer of timerManager) {
            clearInterval(timer)
        }
        timerManager.length = 0
    }
    const getSearch = () => {
        return <QueryFilter span={6}
                            submitter={{
                                searchConfig: {
                                    submitText: '查询子节点',
                                },
                                resetButtonProps: {
                                    style: {
                                        display: 'none',
                                    },
                                },
                            }}
                            onFinish={
                                async (v) => {
                                    // 清除原有的子节点查询记录
                                    dispatch({
                                        type: 'graph/clearSubGraph',
                                        payload: id
                                    })
                                    const src = JSON.parse(v.src)
                                    console.log(v)
                                    setGID(getGraphName(id, src.id))
                                    setRenderDirection(v.direct)
                                    setRenderSrcNode(src)
                                }
                            } disabled={loading}>
            <ProFormSelect label='目标节点' name='src' rules={[{required: true}]}
                           showSearch
                           debounceTime={300}
                           fieldProps={{
                               filterOption: () => {
                                   return true
                               }
                           }}
                           request={async (v) => {
                               let packs: any[] = []
                               if (!v.keyWords) {
                                   return packs
                               }
                               await getMatchNodes({
                                   graphId: id,
                                   keywords: v.keyWords
                               }).then(res => {
                                   res.matchNodePacks.forEach(mp => packs.push({
                                           label: mp.tag,
                                           options: mp.match.map(m => {
                                               return {
                                                   label: m.primaryAttr,
                                                   value: JSON.stringify({
                                                       tag: mp.tag,
                                                       primaryAttr: m.primaryAttr,
                                                       id: m.id
                                                   })
                                               }
                                           })
                                       })
                                   )
                               })
                               return packs
                           }}/>
            <ProFormSelect label='方向' name='direct' rules={[{required: true}]} initialValue={'REVERSELY'}
                           options={[{
                               label: '正向',
                               value: ' '
                           }, {
                               label: '反向',
                               value: 'REVERSELY'
                           }, {
                               label: '双向',
                               value: 'BIDIRECT'
                           },]}/>
        </QueryFilter>
    }
    const getTabContent = () => {
        switch (tab) {
            case 'attr':
                let dataSource = []
                if (select) {
                    dataSource = [
                        {
                            title: 'ID',
                            description: selectNode!.id
                        },
                        {
                            title: '类别',
                            description: selectNode!.tag
                        },
                        {
                            title: '度数',
                            description: selectNode!.deg
                        },
                        ...selectNode!.attrs.map((a: Graph.Pair) => {
                            return {
                                title: a.key,
                                description: a.value,
                            }
                        })
                    ]
                } else {
                    dataSource = [
                        {
                            title: '图',
                            description: name
                        },
                        {
                            title: '总节点数',
                            description: graph.numNode
                        },
                        {
                            title: '总边数',
                            description: graph.numEdge
                        },
                        {
                            title: '当前节点数',
                            description: nodes.length
                        },
                        {
                            title: '当前边数',
                            description: edges.length
                        },
                        {
                            title: <Tooltip title={'图谱展示节点上限'}><span>最大节点数</span></Tooltip>,
                            description: <InputNumber defaultValue={renderMax} formatter={formatNumber} min={1}
                                                      onChange={v => setRenderMax(v!)}/>
                        },
                        {
                            title: '操作',
                            description: <Button danger type='primary' onClick={() => {
                                dropGraph({
                                    graphId: id
                                }).then(res => {
                                    message.success('删除图`' + name + '`成功')
                                    history.push('/')
                                })
                            }
                            }>删除图谱</Button>
                        },
                        {
                            title: '导出',
                            description: <Button type='primary' onClick={() => {
                                // @ts-ignore
                                graphinRef.current.graph.downloadFullImage()
                            }}>
                                PNG
                            </Button>
                        }
                    ]
                    if (isSub) {
                        let direct = '正向'
                        if (renderDirection == 'REVERSELY')
                            direct = '反向'
                        else if (renderDirection == 'BIDIRECT')
                            direct = '双向'
                        dataSource.push(...[{
                            title: '目标节点',
                            description: renderSrcNode!.primaryAttr
                        }, {
                            title: '类型',
                            description: renderSrcNode!.tag
                        }, {
                            title: '方向',
                            description: direct
                        },])
                    }
                }
                return <ProList
                    key='attrListGraph'
                    rowKey='title'
                    dataSource={dataSource}
                    metas={{
                        title: {},
                        description: {},
                    }}/>
            case 'algo':
                const getAlgoTag = (type: AlgoType) => {
                    const desc = getAlgoTypeDesc(type)
                    return <Tag color={desc.color}>{desc.text}</Tag>
                }
                const getAlgoContent = (algo: Algo) => {
                    const onFinish = async (params: any) => {
                        return await algo.action({
                            graphId: id,
                            ...params
                        }).then(() => {
                            message.success('算法已提交')
                            setExtKeysAlgo([])
                            setTab('task')
                            return true
                        })
                    }
                    return <Space direction={'vertical'}>
                        <Text type={'secondary'}>{algo.description}</Text>
                        <Divider/>
                        <ProForm
                            onFinish={onFinish}
                            submitter={{
                                searchConfig: {
                                    resetText: '重置',
                                    submitText: '执行',
                                }
                            }}
                            initialValues={{
                                // pagerank
                                iter: 3,
                                prob: 0.85,
                                // louvain
                                maxIter: 10,
                                internalIter: 5,
                                tol: 0.3,
                            }}
                        >{
                            algo.params.map(p => {
                                switch (p.type) {
                                    case ParamType.Int:
                                        return <ProFormDigit name={p.field} fieldProps={{precision: 0}}
                                                             label={p.text} max={p.max} min={p.min}/>
                                    case ParamType.Double:
                                        return <ProFormDigit name={p.field} fieldProps={{precision: 4, step: 1e-4}}
                                                             label={p.text} max={p.max} min={p.min}/>
                                }
                            })
                        }
                        </ProForm>
                    </Space>
                }
                return <ProList<Algo>
                    key='algoProList'
                    rowKey={(row, index) => row.id}
                    toolBarRender={() => {
                        return [
                            <Button key='3' type='primary'>
                                新建
                            </Button>,
                        ];
                    }}
                    style={{
                        height: '80vh',
                        overflowY: 'scroll',
                    }}
                    expandable={{
                        expandedRowKeys: extKeysAlgo, onExpandedRowsChange: (expandedKeys) => {
                            // @ts-ignore
                            setExtKeysAlgo(expandedKeys)
                        }
                    }}
                    search={{
                        filterType: 'light',
                    }}
                    request={
                        async (params = {time: Date.now()}) => {
                            return {
                                data: algos.filter(a => !params.subTitle || a.type == params.subTitle),
                                success: true,
                                total: algos.length
                            }
                        }}
                    metas={{
                        title: {
                            search: false
                        },
                        subTitle: {
                            title: '类别',
                            render: (_, row) => {
                                return <Space size={0}>
                                    {getAlgoTag(row.type)}
                                </Space>
                            },
                            valueType: 'select',
                            valueEnum: AlgoTypeMap,
                        },
                        description: {
                            search: false,
                            render: (_, row) => getAlgoContent(row)
                        },
                    }}
                />
            case 'task':
                const getTaskTag = (status: number) => {
                    const s = getTaskTypeDesc(status)
                    return <Tag color={s.color}>{s.text}</Tag>
                }
                const getTaskContent = (task: Graph.Task) => {
                    const getTaskResult = (sres: string) => {
                        try {
                            if (!sres) {
                                return
                            }
                            const res = JSON.parse(sres)
                            if (res?.score) {
                                return <MetricTable score={res.score}/>
                            }
                            if (res?.ranks) {
                                console.log(res)
                                return <RankTable file={res.file} rows={
                                    res.ranks.map((r: { tag: string, node: Graph.Node; score: any; }) => {
                                        let node = r.tag + '(' + r.node.id + ','
                                        for (let p of r.node.attrs) {
                                            node = node + p.key + "=" + p.value + ','
                                        }
                                        node = node.substring(0, node.length - 1) + ')'
                                        return {
                                            node: node,
                                            rank: r.score,
                                        }
                                    })}/>
                            }
                        } catch (e) {
                            console.log(e)
                            return
                        }
                    }
                    return getTaskResult(task.res)
                }
                return <ProList<Graph.Task>
                    key='taskProList'
                    itemLayout='vertical'
                    // @ts-ignore
                    actionRef={taskListRef}
                    rowKey='id'
                    style={{
                        height: '80vh',
                        overflowY: 'scroll',
                        overflowX: 'hidden'
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
                        const tasks = await getGraphTasks({
                            graphId: id
                        })
                        return {
                            data: tasks.tasks?.filter((t: { status: number; }) => !params.subTitle || t.status == params.subTitle),
                            success: true,
                            total: tasks.tasks.length
                        }
                    }}
                    metas={{
                        title: {
                            search: false,
                            render: (_, row) => <Text>{row.idf}</Text>
                        },
                        subTitle: {
                            title: '类别',
                            render: (_, row) => {
                                return <Space size={0}>
                                    {getTaskTag(row.status)}
                                </Space>
                            },
                            valueType: 'select',
                            valueEnum: TaskTypeMap,
                        },
                        extra: {
                            render: (_: any, row: { createTime: number; id: any; status: any; }) => {
                                return <Space direction={'vertical'}>
                                    <Text type={'secondary'}>{formatDate(row.createTime)}</Text>
                                    <a style={{float: 'right'}} onClick={() => {
                                        dropTask({
                                            taskId: row.id
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
    }
    return <PageContainer header={{title: ''}} content={
        <Row gutter={16} style={{height: '100%'}}>
            <Col span={18}>
                <Card
                    title={getSearch()}
                    style={{height: '100%'}}
                    bodyStyle={{padding: '0 0 0 0'}}
                    extra={
                        <Space>
                            <Button danger type={'primary'} onClick={resetGraph}>重置图谱</Button>
                            <LayoutSelector options={layouts} value={graphLayout}
                                            onChange={(value: SetStateAction<string>) => {
                                                setGraphLayout(value)
                                            }}/>
                        </Space>
                    }
                >
                    <Spin spinning={loading}>
                        <Graphin data={data} layout={layout} fitView={true}
                                 containerStyle={{height: '80vh'}}
                            // @ts-ignore
                                 ref={graphinRef}>
                            <Hoverable bindType='node'/>
                            <Hoverable bindType='edge'/>
                        </Graphin>
                    </Spin>
                </Card>
            </Col>
            <Col span={6} style={{height: '100%'}}>
                <Card
                    activeTabKey={tab}
                    style={{height: '100%'}}
                    bodyStyle={{padding: 0}}
                    tabList={tabListNoTitle}
                    onTabChange={key => {
                        setTab(key)
                    }}>
                    {getTabContent()}
                </Card>
            </Col>
        </Row>
    }
    >
    </PageContainer>


}

const mapStateToProps = (state: any) => {
    return {
        details: state.graph
    }
}

export default connect(mapStateToProps, null)(Graph)
