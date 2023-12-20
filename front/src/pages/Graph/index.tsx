import {PageContainer, ProFormSelect, ProList, QueryFilter} from '@ant-design/pro-components';
import Graphin, {Behaviors, Components, LegendChildrenProps} from '@antv/graphin';
import {
    Button,
    Card,
    Col,
    InputNumber,
    message,
    Popconfirm,
    Row,
    Space,
    Spin,
    Tooltip
} from 'antd';
import React, {SetStateAction, useEffect, useState} from 'react';
import {connect} from '@@/exports';
import {formatNumber} from '@/utils/format';
import {dropGraph, getMatchNodes} from '@/services/graph/graph';
import {history} from 'umi';
import {EdgeData, findByGid, getGraphName, isSubGraph, NodeData} from '@/models/graph';
import LayoutSelector from "@/components/LayoutSelector";
import {layouts, layoutsConfig} from "@/pages/Graph/_layout";
import {
    CloseOutlined,
    FileImageOutlined,
    QuestionCircleOutlined,
    SearchOutlined,
    UndoOutlined
} from "@ant-design/icons";
import {GraphRef2Group, TreeNodeGroup} from "@/models/global";

const {Legend} = Components;
const {Hoverable} = Behaviors;

type Props = {
    graph: GraphRef2Group,
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
    const [graphLayout, setGraphLayout] = useState<string>('graphin-force')
    const [tab, setTab] = useState<string>('attr')
    const [gid, setGID] = useState<string>(getGraphName(graph.id))
    // @ts-ignore
    const graphinRef = React.createRef()
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
    const group: TreeNodeGroup = graph.group
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
                    timer: timer,
                    group: group,
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
                    timer: timer,
                    group: group,
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
    ]
    // 图谱布局
    const layout = layoutsConfig.find(item => item.type === graphLayout)
    // 数据准备
    let data = {
        nodes: [],
        edges: [],
    }
    // 如果数据准备好，则加载数据
    if (!loading) {
        const graphData = findByGid(details, gid)
        data = {
            nodes: graphData.nodes.map((n: NodeData) => n.style),
            edges: graphData.edges.map((e: EdgeData) => e.style),
        }
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
        return <QueryFilter
            submitter={{
                submitButtonProps: {
                    icon: <SearchOutlined/>
                },
                searchConfig: {
                    submitText: '查询',
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
                        }).sort((a, b) => a.title < b.title ? 0 : 1)
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
                            description: data.nodes.length
                        },
                        {
                            title: '当前边数',
                            description: data.edges.length
                        },
                        {
                            title: <Tooltip title={'图谱展示节点上限'}><span>最大节点数</span></Tooltip>,
                            description: <InputNumber defaultValue={renderMax} formatter={formatNumber} min={1}
                                                      onChange={v => setRenderMax(v!)}/>
                        },
                        {
                            title: '操作',
                            description:
                                <Space size='large'>
                                    <Button icon={<UndoOutlined/>} type='primary' onClick={resetGraph}>重置</Button>
                                    <Popconfirm
                                        title={'确认删除图' + name + '?'}
                                        icon={<QuestionCircleOutlined style={{color: 'red'}}/>}
                                        onConfirm={() => {
                                            dropGraph({
                                                graphId: id
                                            }).then(() => {
                                                message.success('删除图`' + name + '`成功')
                                                history.push('/')
                                                window.location.reload()
                                            })
                                        }}
                                    >
                                        <Button icon={<CloseOutlined/>} type='primary' danger>删除</Button>
                                    </Popconfirm>
                                </Space>
                        },
                        {
                            title: '导出',
                            description: <Button icon={<FileImageOutlined/>} type='primary' onClick={() => {
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
                        <LayoutSelector options={layouts} value={graphLayout}
                                        onChange={(value: SetStateAction<string>) => {
                                            setGraphLayout(value)
                                        }}/>
                    }
                >
                    <Spin spinning={loading}>
                        <Graphin data={data} layout={layout} fitView={true}
                                 containerStyle={{height: '80vh'}}
                            // @ts-ignore
                                 ref={graphinRef}>
                            <Legend bindType="node" sortKey="tag">
                                {(renderProps: LegendChildrenProps) => {
                                    return <Legend.Node {...renderProps}/>
                                }}
                            </Legend>
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
