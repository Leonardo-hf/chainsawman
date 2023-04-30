import {PageContainer} from "@ant-design/pro-components";
import {
    TrademarkCircleFilled,
    ChromeFilled,
    BranchesOutlined,
    ApartmentOutlined,
    AppstoreFilled,
    CopyrightCircleFilled,
    CustomerServiceFilled,
    ShareAltOutlined,
} from '@ant-design/icons';
import Graphin, {Behaviors} from '@antv/graphin';
import {Select, Row, Col, Card, Spin, Input, Button, InputNumber, message} from 'antd';
import React, {SetStateAction} from "react";
import {connect} from "@@/exports";
import {Space, Typography} from 'antd';
import {getTag} from "@/utils/format";

const {Text, Paragraph} = Typography;
const {Hoverable} = Behaviors;
const {Search} = Input

const tabListNoTitle = [
    {
        key: 'attr',
        tab: '属性',
    },
    {
        key: 'sys',
        tab: '设置',
    },
    {
        key: 'algo',
        tab: '算法',
    },
    {
        key: 'task',
        tab: '任务',
    },
];

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
    min: 2,
    distance: 1,
}

@connect(mapStateToProps)
class Graph extends React.Component<{ graph: { id: number, name: string, desc: string, nodes: number, edges: number }, details: any, dispatch: any }, any> {

    constructor(props: { graph: { id: number, name: string, desc: string, nodes: number, edges: number }, details: any, dispatch: any }) {
        super(props)
        const graphinRef = React.createRef();
        this.state = {
            type: 'graphin-force',
            tab: 'attr',
            graphinRef: graphinRef,
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
                min: 0,
                distance: 0,
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
                name: id,
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
        const {tab, type, graphinRef, select, query} = this.state
        const tag = getTag(id, query.id)
        // 判断图有没有加载好
        const loading = (details[id] === undefined || details[id].status === 0) ||
            (query.ok && (details[tag] === undefined || details[tag].status === 0))
        const nodes: { id: string, style: { keyshape: { size: number, fill: string }, label: { value: string } } }[] = []
        const edges: { source: string, target: string }[] = []
        const data = {
            nodes: nodes,
            edges: edges,
        }

        const layout = layouts.find(item => item.type === type);
        // 图模式，如果graph state没有就发起查询
        if (!query.ok && details[id] === undefined) {
            const queryDetail = async () => {
                let taskId = 0
                if (this.props.details[id] !== undefined) {
                    taskId = this.props.details[id].taskId
                }
                this.props.dispatch({
                    type: 'graph/queryDetail',
                    payload: {
                        graphId: id,
                        taskId: taskId,
                        min: 0,
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
            const timer = setInterval(queryDetail, 3000)
        }
        // 节点模式
        if (query.ok && details[tag] === undefined) {
            const queryNeibors = async () => {
                let taskId = 0
                if (this.props.details[tag] !== undefined) {
                    taskId = this.props.details[tag].taskId
                }
                this.props.dispatch({
                    type: 'graph/queryNeibors',
                    payload: {
                        graphId: id,
                        taskId: taskId,
                        node: query.id,
                        min: query.min,
                        distance: query.distance
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
            queryNeibors()
            const timer = setInterval(queryNeibors, 3000)
        }

        if (!loading) {
            const target = query.ok ? tag : id
            details[target].nodes.forEach((n: { id: number, name: string; deg: number; color: string; }) => nodes.push({
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
            details[target].edges.forEach((e: { source: number; target: number; }) => edges.push({
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
            const format = (v: any) => Math.floor(v).toString()
            return <Space direction="horizontal">
                <Search addonBefore={<Text strong>{graph.name}</Text>} placeholder={query.nameHint}
                        status={query.nameStatus}
                        disabled={loading}
                        onSearch={resetGraph}
                        onChange={(v) => {
                            this.state.query.name = v.target.value
                        }}/>
                <InputNumber addonBefore='距离' defaultValue={defaultNameProps.distance} formatter={format}
                             onChange={(v) => this.state.query.distance = v}/>
                <InputNumber addonBefore='最小度数' defaultValue={defaultNameProps.min} formatter={format}
                             onChange={(v) => this.state.query.min = v}/>
            </Space>
        }

        const getTabContent = () => {
            switch (tab) {
                case 'attr':
                    return select.ok ?
                        <Space direction="vertical">
                            <Space><Text strong>节点：</Text><Text>{select.name}</Text></Space>
                            <Space><Text strong>描述：</Text><Paragraph ellipsis={{
                                rows: 3,
                                expandable: true,
                                symbol: 'more'
                            }} style={{marginBottom: 0}}>{select.desc}</Paragraph></Space>
                            <Space><Text strong>度数：</Text><Text>{select.deg}</Text></Space>
                        </Space> :
                        <Space direction="vertical">
                            <Space><Text strong>图：</Text><Text>{name}</Text></Space>
                            {
                                query.ok && <Space direction="vertical">
                                    <Space><Text strong>子节点：</Text><Text>{query.name}</Text></Space>
                                    <Space><Text strong>距离：</Text><Text>{query.distance}</Text></Space>
                                    <Space><Text strong>最小度数：</Text><Text>{query.min}</Text></Space>
                                </Space>
                            }
                            <Space><Text strong>描述：</Text><Paragraph ellipsis={{
                                rows: 3,
                                expandable: true,
                                symbol: 'more'
                            }} style={{marginBottom: 0}}>{graph.desc}</Paragraph></Space>
                            <Space><Text strong>总节点数：</Text><Text>{graph.nodes}</Text></Space>
                            <Space><Text strong>总边数：</Text><Text>{graph.edges}</Text></Space>
                            <Space><Text strong>当前节点数：</Text><Text>{nodes.length}</Text></Space>
                            <Space><Text strong>当前边数：</Text><Text>{edges.length}</Text></Space>
                        </Space>
                case 'sys':
                    return <Space direction={"vertical"}>
                        <Space><Text strong>来源：</Text><Text>文件导入</Text></Space>
                    </Space>
                case 'algo':
                case 'task':
                //TODO: finish
            }
        }


        return (
            <PageContainer header={{title: ''}} content={
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
                                <Graphin data={data} layout={layout} fitView={true} containerStyle={{height: '80vh'}}
                                         ref={graphinRef}>
                                    <Hoverable bindType="node"/>
                                    <Hoverable bindType="edge"/>
                                </Graphin>
                            </Spin>
                        </Card>
                    </Col>
                    <Col span={6} style={{height: '100%'}}>
                        <Card
                            style={{height: '100%'}}
                            tabList={tabListNoTitle}
                            onTabChange={key => this.setState({tab: key})}>
                            {getTabContent()}
                        </Card>
                    </Col>
                </Row>
            }>
            </PageContainer>
        );
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

const layouts = [
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