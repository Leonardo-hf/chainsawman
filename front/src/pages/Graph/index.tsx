import {PageContainer, ProList} from "@ant-design/pro-components";
import {
    ApartmentOutlined,
    AppstoreFilled,
    BranchesOutlined,
    ChromeFilled,
    CopyrightCircleFilled,
    CustomerServiceFilled,
    ShareAltOutlined,
    TrademarkCircleFilled,
} from '@ant-design/icons';
import Graphin, {Behaviors} from '@antv/graphin';
import {
    Button,
    Card,
    Col,
    Divider,
    Form,
    Input,
    InputNumber,
    message,
    Row,
    Select,
    Space,
    Spin,
    Tag,
    Tooltip,
    Typography
} from 'antd';
import React, {SetStateAction} from "react";
import {connect} from "@@/exports";
import {formatDate, formatNumber, getTag} from "@/utils/format";
import {Algo, algos, AlgoType, AlgoTypeMap, getAlgoTypeDesc, ParamType} from './_algo';
import {dropGraph, dropTask, getGraphTasks} from "@/services/graph/graph";
import {getTaskTypeDesc, TaskTypeMap} from "./_task";
import MetricTable from "@/components/MetricTable";
import RankTable from "@/components/RankTable";
import {history} from 'umi';

const {Text, Paragraph} = Typography;
const {Hoverable} = Behaviors;
const {Search} = Input

const iconMap = {
    'graphin-force': <ShareAltOutlined/>,
    random: <TrademarkCircleFilled/>,
    concentric: <ChromeFilled/>,
    circle: <BranchesOutlined/>,
    force: <AppstoreFilled/>,
    dagre: <ApartmentOutlined/>,
    grid: <CopyrightCircleFilled/>,
    radial: <ShareAltOutlined/>,
};

const mapStateToProps = (state: any) => {
    const {details} = state.graph
    return {
        details: details
    }
}

const defaultNameProps = {
    hint: '目标查询节点',
    status: undefined,
    min: 0,
    distance: 1,
}

@connect(mapStateToProps)
class Graph extends React.Component<{ graph: { id: number, name: string, desc: string, nodes: number, edges: number }, details: any, dispatch: any }, any> {

    constructor(props: { graph: { id: number, name: string, desc: string, nodes: number, edges: number }, details: any, dispatch: any }) {
        super(props)

        this.state = {
            graph: {
                min: 10,
                kmin: 10
            },
            extKeysAlgo: [],
            extKeysTask: [],
            type: 'graphin-force',
            tab: 'attr',
            graphinRef: React.createRef(),
            taskListRef: React.createRef(),
            select: {
                ok: false,
                name: '',
                desc: '',
                deg: '',
            },
            query: {
                ok: false,
                nameHint: defaultNameProps.hint,
                nameStatus: defaultNameProps.status,
                // 绑定input
                name: '',
                id: 0,
                min: defaultNameProps.min,
                distance: defaultNameProps.distance,
                // 记录搜索成功后的input值
                kname: '',
                kmin: defaultNameProps.min,
                kdistance: defaultNameProps.distance,
            }
        }
    }

    handleNodeClick = (e: { item: { get: (arg0: string) => string; }; }) => {
        // console.log(e.item)
        const id = e.item.get('id')
        const target = this.state.query.ok ? getTag(this.props.graph.id, this.state.query.id) : this.props.graph.id
        // 图模式和节点模式需要查询不同的图
        const node = this.props.details[target].nodes.find((n: { id: number; }) => n.id.toString() === id);
        this.setState({
            select: {
                ok: true,
                name: node.name,
                desc: node.desc,
                deg: node.deg,
            }
        });
    };

    componentDidMount() {
        // console.log('mount')
        const {graph} = this.state.graphinRef.current;
        graph.on('node:click', this.handleNodeClick);
        graph.on('canvas:click', () => {
            this.setState({
                select: {
                    ...this.state.select,
                    ok: false
                }
            })
        })
    }

    componentWillUnmount() {
        // console.log('unmount')
        const {graph} = this.state.graphinRef.current;
        graph.off('node:click');
        graph.off('canvas:click')
    }

    render() {
        // console.log('render')
        const {graph, details} = this.props
        let {id, name} = graph
        const {tab, type, graphinRef, taskListRef, select, query} = this.state
        const {min, kmin} = this.state.graph
        const tag = getTag(id, query.id)
        // 判断图有没有加载好
        const loading = !details[id] || !details[id].status ||
            (query.ok && (!details[tag] || !details[tag].status))
        const nodes: { id: string, style: { keyshape: { size: number, fill: string }, label: { value: string } } }[] = []
        const edges: { source: string, target: string }[] = []
        const data = {
            nodes: nodes,
            edges: edges,
        }
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
                disabled: !details[id] || !details[id].status
            },
        ];

        const layout = layouts.find(item => item.type === type);
        // 图模式，如果graph state没有就发起查询
        if (!query.ok && !details[id]) {
            const queryDetail = async () => {
                let taskId = ""
                if (this.props.details[id]) {
                    taskId = this.props.details[id].taskId
                }
                this.props.dispatch({
                    type: 'graph/queryDetail',
                    payload: {
                        graphId: id,
                        taskId: taskId,
                        min: min,
                    }
                }).then((taskStatus: number) => {
                    if (taskStatus === 1) {
                        clearInterval(timer)
                    }
                })
                // if (this.props.details[name].status === 1) {
                //     clearInterval(timer)
                // }
            }
            queryDetail()
            const timer = setInterval(queryDetail, 5000)
        }
        // 节点模式
        if (query.ok && !details[tag]) {
            const queryNeibors = async () => {
                let taskId = ""
                if (this.props.details[tag]) {
                    taskId = this.props.details[tag].taskId
                }
                console.log(query)
                this.props.dispatch({
                    type: 'graph/queryNeibors',
                    payload: {
                        graphId: id,
                        taskId: taskId,
                        nodeId: query.id,
                        min: query.min,
                        distance: query.distance
                    }
                }).then((taskStatus: number) => {
                    this.state.query.kmin = this.state.query.min
                    this.state.query.kname = this.state.query.name
                    this.state.query.kdistance = this.state.query.distance
                    if (taskStatus === 1) {
                        clearInterval(timer)
                    }
                })
                // if (this.props.details[name].status === 1) {
                //     clearInterval(timer)
                // }
            }
            queryNeibors()
            const timer = setInterval(queryNeibors, 5000)
        }

        if (!loading) {
            const target = query.ok ? tag : id
            details[target].nodes?.forEach((n: { id: number, name: string; deg: number; color: string; }) => nodes.push({
                id: n.id.toString(),
                style: {
                    keyshape:
                        {
                            fill: n.color,
                            size: Math.floor((Math.log(n.deg + 1) + 1) * 10),
                        },
                    label: {
                        value: n.name
                    }
                }
            }))
            details[target].edges?.forEach((e: { source: number; target: number; }) => edges.push({
                source: e.source.toString(),
                target: e.target.toString(),
            }))
        }

        const getSearch = () => {
            const resetGraph = () => {
                // console.log('search')
                if (!query.name) {
                    this.setState({
                        query: {
                            ...this.state.query,
                            nameHint: '必须设置目标查询节点',
                            nameStatus: 'error'
                        }
                    })
                    return
                }
                const nodeID = details[id].nodes.find((node: { name: string; }) => node.name === query.name)?.id
                if (nodeID === undefined) {
                    message.error('图中没有该节点')
                    return;
                }
                this.setState({
                    query: {
                        ...this.state.query,
                        id: nodeID,
                        ok: true,
                        nameHint: defaultNameProps.hint,
                        nameStatus: defaultNameProps.status
                    }
                })
            }
            return <Space direction="horizontal">
                <Search addonBefore='目标节点' placeholder={query.nameHint}
                        status={query.nameStatus}
                        disabled={loading}
                        onSearch={resetGraph}
                        onChange={(v) => {
                            this.state.query.name = v.target.value
                        }}/>
                <InputNumber addonBefore={
                    <Tooltip title="子图仅展示与目标节点距离低于`最大距离`的节点">
                        <span>最大距离</span>
                    </Tooltip>
                } defaultValue={defaultNameProps.distance} formatter={formatNumber}
                             onChange={(v) => this.state.query.distance = v}/>
                <InputNumber addonBefore={
                    <Tooltip title="子图仅展示高于`最低度数`的节点">
                        <span>最低度数</span>
                    </Tooltip>
                } defaultValue={defaultNameProps.min} formatter={formatNumber}
                             onChange={(v) => this.state.query.min = v}/>
            </Space>
        }
        const getTabContent = () => {
            switch (tab) {
                case 'attr':
                    let dataSource = []
                    if (select.ok) {
                        dataSource = [
                            {
                                title: '节点',
                                description: select.name
                            },
                            {
                                title: '描述',
                                description: select.desc
                            },
                            {
                                title: '度数',
                                description: select.deg
                            },
                        ]
                    } else {
                        dataSource = [
                            {
                                title: '图',
                                description: name
                            },
                            {
                                title: '总节点数',
                                description: graph.nodes
                            },
                            {
                                title: '总边数',
                                description: graph.edges
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
                                title: '来源',
                                description: '文件导入'
                            },
                            {
                                title: <Tooltip
                                    title={'大图仅展示高于`最低度数`的节点'}><span>图最低度数</span></Tooltip>,
                                description: <InputNumber defaultValue={min} formatter={formatNumber} min={0}
                                                          onChange={v => {
                                                              this.state.graph.min = v
                                                          }}/>
                            },
                            {
                                title: '操作',
                                description: <Button danger type="primary" onClick={() => {
                                    dropGraph({
                                        graphId: id
                                    }).then(res => {
                                        message.success('删除图`' + name + '`成功')
                                        history.push('/')
                                        // setTimeout(() => window.location.reload(), 1000)
                                    })
                                }
                                }>删除图</Button>
                            }
                        ]
                        if (query.ok) {
                            dataSource.push({
                                title: '子节点',
                                description: this.state.query.kname
                            })
                            dataSource.push({
                                title: '最大距离',
                                description: this.state.query.kdistance
                            })
                            dataSource.push({
                                title: '最低度数',
                                description: this.state.query.kmin
                            })
                        }
                    }
                    return <ProList
                        key="attrListGraph"
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
                        const onFinish = (params: any) => {
                            algo.action({
                                graphId: id,
                                ...params
                            }).then(() => {
                                message.success('算法已提交')
                                this.setState({
                                    extKeysAlgo: [],
                                    tab: 'task'
                                })
                            })
                        }
                        return <Space direction={'vertical'}>
                            <Text type={'secondary'}>{algo.description}</Text>
                            <Divider/>
                            <Form
                                name='basic'
                                onFinish={onFinish}
                                autoComplete='off'
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
                                            return <Form.Item name={p.field}>
                                                <InputNumber addonBefore={p.text} max={p.max} min={p.min}
                                                             formatter={formatNumber}/>
                                            </Form.Item>
                                        case ParamType.Double:
                                            return <Form.Item name={p.field}>
                                                <InputNumber addonBefore={p.text} max={p.max} min={p.min}
                                                             precision={4} step={1e-4}/>
                                            </Form.Item>
                                    }
                                })
                            }
                                <Form.Item>
                                    <Button htmlType="submit" style={{float: 'right'}}>执行</Button>
                                </Form.Item>
                            </Form>
                        </Space>
                    }
                    return <ProList<Algo>
                        key="algoProList"
                        rowKey={(row, index) => row.title}
                        toolBarRender={() => {
                            return [
                                <Button key="3" type="primary">
                                    新建
                                </Button>,
                            ];
                        }}
                        style={{
                            height: '80vh',
                            overflowY: 'scroll',
                        }}
                        expandable={{
                            expandedRowKeys: this.state.extKeysAlgo, onExpandedRowsChange: (expandedKeys) => {
                                this.setState({
                                    extKeysAlgo: expandedKeys
                                })
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
                                const res = JSON.parse(sres)
                                if (res?.score) {
                                    return <MetricTable score={res.score}/>
                                }
                                if (res?.ranks) {
                                    return <RankTable file={res.file} rows={
                                        res.ranks.map((r: { nodeId: number; score: any; }) => {
                                            return {
                                                node: details[id].nodes.find((n: { id: number; }) => r.nodeId == n.id)?.name,
                                                rank: r.score,
                                            }
                                        })}/>
                                }
                            } catch (e) {
                                return
                            }
                        }
                        return getTaskResult(task.res)
                    }
                    return <ProList<Graph.Task>
                        key="taskProList"
                        itemLayout="vertical"
                        actionRef={taskListRef}
                        rowKey="id"
                        style={{
                            height: '80vh',
                            overflowY: 'scroll',
                            overflowX: 'hidden'
                        }}
                        expandable={{
                            // rowExpandable: (row) => {
                            //     return row.status == TaskType.finished
                            // },
                            expandedRowKeys: this.state.extKeysTask, onExpandedRowsChange: (expandedKeys) => {
                                this.setState({
                                    extKeysTask: expandedKeys
                                })
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
                                render: (_, row) => <Text>{row.desc}</Text>
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
                                render: (_, row) => {
                                    return <Space direction={'vertical'}>
                                        <Text type={'secondary'}>{formatDate(row.createTime)}</Text>
                                        <a style={{float: 'right'}} onClick={() => {
                                            console.log(row.id)
                                            dropTask({
                                                taskId: row.id
                                            }).then(() => {
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
        return <PageContainer header={{
            title: ''
        }} content={
            <Row gutter={16} style={{height: '100%'}}>
                <Col span={18}>
                    <Card
                        title={getSearch()}
                        style={{height: '100%'}}
                        bodyStyle={{padding: '0 0 0 0'}}
                        extra={
                            <Space>
                                <Button type={'primary'} onClick={() => {
                                    const {graph} = graphinRef.current
                                    graph.emit('canvas:click')
                                    this.setState({
                                        query: {
                                            ...this.state.query,
                                            ok: false
                                        }
                                    })
                                    if (this.state.graph.kmin !== this.state.graph.min) {
                                        this.state.graph.kmin = this.state.graph.min
                                        this.props.dispatch({
                                            type: 'graph/resetGraph',
                                            payload: {
                                                graphID: id
                                            }
                                        })
                                    }
                                }}>
                                    重置
                                </Button>
                                <LayoutSelector options={layouts} value={type}
                                                onChange={(value: SetStateAction<string>) => {
                                                    this.setState({
                                                        type: value
                                                    })
                                                }}/>
                            </Space>
                        }
                    >
                        <Spin spinning={loading}>
                            <Graphin data={data} layout={layout} fitView={true}
                                     containerStyle={{height: '80vh'}}
                                     ref={graphinRef}>
                                <Hoverable bindType="node"/>
                                <Hoverable bindType="edge"/>
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
                            this.setState({tab: key})
                        }}>
                        {getTabContent()}
                    </Card>
                </Col>
            </Row>
        }>
        </PageContainer>
    }
}

export default Graph;
const SelectOption = Select.Option;
const LayoutSelector = (props: { value: any; onChange: any; options: any; }) => {
    const {value, onChange, options} = props;
    return (
        <div
            // style={{ position: 'absolute', top: 10, left: 10 }}
        >
            <Select style={{width: '120px'}} value={value} onChange={onChange}>
                {options.map((item: { type: any; }) => {
                    const {type} = item;
                    // @ts-ignore
                    const iconComponent = iconMap[type] || <CustomerServiceFilled/>;
                    return (
                        <SelectOption key={type} value={type}>
                            {iconComponent} &nbsp;
                            {type}
                        </SelectOption>
                    );
                })}
            </Select>
        </div>
    );
};

const
    layouts = [
        {
            type: 'graphin-force'
        },
        {
            type: 'grid',
            // begin: [0, 0], // 可选，
            // preventOverlap: true, // 可选，必须配合 nodeSize
            // preventOverlapPdding: 20, // 可选
            // nodeSize: 30, // 可选
            // condense: false, // 可选
            // rows: 5, // 可选
            // cols: 5, // 可选
            // sortBy: 'degree', // 可选
            // workerEnabled: false, // 可选，开启 web-worker
        },
        {
            type: 'circular',
            // center: [200, 200], // 可选，默认为图的中心
            // radius: null, // 可选
            // startRadius: 10, // 可选
            // endRadius: 100, // 可选
            // clockwise: false, // 可选
            // divisions: 5, // 可选
            // ordering: 'degree', // 可选
            // angleRatio: 1, // 可选
        },
        {
            type: 'radial',
            // center: [200, 200], // 可选，默认为图的中心
            // linkDistance: 50, // 可选，边长
            // maxIteration: 1000, // 可选
            // focusNode: 'node11', // 可选
            // unitRadius: 100, // 可选
            // preventOverlap: true, // 可选，必须配合 nodeSize
            // nodeSize: 30, // 可选
            // strictRadial: false, // 可选
            // workerEnabled: false, // 可选，开启 web-worker
        },
        {
            type: 'force',
            preventOverlap: true,
            // center: [200, 200], // 可选，默认为图的中心
            linkDistance: 50, // 可选，边长
            nodeStrength: 30, // 可选
            edgeStrength: 0.8, // 可选
            collideStrength: 0.8, // 可选
            nodeSize: 30, // 可选
            alpha: 0.9, // 可选
            alphaDecay: 0.3, // 可选
            alphaMin: 0.01, // 可选
            forceSimulation: null, // 可选
            onTick: () => {
                // 可选
                console.log('ticking');
            },
            onLayoutEnd: () => {
                // 可选
                console.log('force layout done');
            },
        },
        {
            type: 'gForce',
            linkDistance: 150, // 可选，边长
            nodeStrength: 30, // 可选
            edgeStrength: 0.1, // 可选
            nodeSize: 30, // 可选
            onTick: () => {
                // 可选
                console.log('ticking');
            },
            onLayoutEnd: () => {
                // 可选
                console.log('force layout done');
            },
            workerEnabled: false, // 可选，开启 web-worker
            gpuEnabled: false, // 可选，开启 GPU 并行计算，G6 4.0 支持
        },
        {
            type: 'concentric',
            maxLevelDiff: 0.5,
            sortBy: 'degree',
            // center: [200, 200], // 可选，

            // linkDistance: 50, // 可选，边长
            // preventOverlap: true, // 可选，必须配合 nodeSize
            // nodeSize: 30, // 可选
            // sweep: 10, // 可选
            // equidistant: false, // 可选
            // startAngle: 0, // 可选
            // clockwise: false, // 可选
            // maxLevelDiff: 10, // 可选
            // sortBy: 'degree', // 可选
            // workerEnabled: false, // 可选，开启 web-worker
        },
        {
            type: 'dagre',
            rankdir: 'LR', // 可选，默认为图的中心
            // align: 'DL', // 可选
            // nodesep: 20, // 可选
            // ranksep: 50, // 可选
            // controlPoints: true, // 可选
        },
        {
            type: 'fruchterman',
            // center: [200, 200], // 可选，默认为图的中心
            // gravity: 20, // 可选
            // speed: 2, // 可选
            // clustering: true, // 可选
            // clusterGravity: 30, // 可选
            // maxIteration: 2000, // 可选，迭代次数
            // workerEnabled: false, // 可选，开启 web-worker
            // gpuEnabled: false, // 可选，开启 GPU 并行计算，G6 4.0 支持
        },
        {
            type: 'mds',
            workerEnabled: false, // 可选，开启 web-worker
        },
        {
            type: 'comboForce',
            // // center: [200, 200], // 可选，默认为图的中心
            // linkDistance: 50, // 可选，边长
            // nodeStrength: 30, // 可选
            // edgeStrength: 0.1, // 可选
            // onTick: () => {
            //   // 可选
            //   console.log('ticking');
            // },
            // onLayoutEnd: () => {
            //   // 可选
            //   console.log('combo force layout done');
            // },
        },
    ];
