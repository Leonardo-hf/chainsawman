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
import {Select, Row, Col, Card, Spin, Divider} from 'antd';
import React, {SetStateAction} from "react";
import {connect} from "@@/exports";
import {Space, Typography} from 'antd';

const {Text, Paragraph} = Typography;
const {Hoverable} = Behaviors;

const tabListNoTitle = [
    {
        key: 'attr',
        tab: '属性',
    },
    {
        key: 'sys',
        tab: '设置',
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


@connect(mapStateToProps)
class Graph extends React.Component<{ graph: { name: string, desc: string, nodes: number, edges: number }, details: any, dispatch: any }, any> {

    constructor(props: { graph: { name: string, desc: string, nodes: number, edges: number }, details: any, dispatch: any }) {
        super(props)
        const graphinRef = React.createRef();
        this.state = {
            type: 'graphin-force',
            tab: 'attr',
            graphinRef: graphinRef,
            nodeSelect: false,
            nodeName: '',
            nodeDesc: '',
            nodeDeg: 0,
        }
    }

    handleNodeClick = (e: { item: { get: (arg0: string) => string; }; }) => {
        // console.log(e.item)
        const id = e.item.get('id')
        const node = this.props.details[this.props.graph.name].nodes.find((n: { name: string; }) => n.name === id);
        this.setState({
            nodeSelect: true,
            nodeName: id,
            nodeDesc: node.desc,
            nodeDeg: node.deg,
        });
    };

    componentDidMount() {
        // console.log('mount')
        const {graph} = this.state.graphinRef.current;
        graph.on('node:click', this.handleNodeClick);
        graph.on('canvas:click', () => {
            this.setState({
                nodeSelect: false
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
        const {name} = graph
        const loading = details[name] === undefined || details[name].status === 0
        const nodes: { id: string, style: { keyshape: { size: number, fill: string }, label: { value: string } } }[] = []
        const edges: { source: string, target: string }[] = []
        const data = {
            nodes: nodes,
            edges: edges,
        }

        const handleChange = (value: SetStateAction<string>) => {
            this.setState({
                type: value
            })
        };
        const layout = layouts.find(item => item.type === this.state.type);
        if (details[name] === undefined) {
            const queryDetail = async () => {
                let taskId = 0
                if (this.props.details[name] !== undefined) {
                    taskId = this.props.details[name].taskId
                }
                this.props.dispatch({
                    type: 'graph/queryDetail',
                    payload: {
                        graph: name,
                        taskId: taskId,
                        min: 0,
                    }
                }).then(taskStatus => {
                    if (taskStatus === 1) {
                        clearInterval(timer)
                    }
                })
                // console.log(this.props.details)
                // if (this.props.details[name].status === 1) {
                //     clearInterval(timer)
                // }
            }
            queryDetail()
            const timer = setInterval(queryDetail, 3000)
        }
        if (!loading) {
            details[name].nodes.forEach((n: { name: string; deg: number; color: string; }) => nodes.push({
                id: n.name,
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
            details[name].edges.forEach((e: { source: string; target: string; }) => edges.push({
                source: e.source,
                target: e.target,
            }))
        }

        const getTabContent = () => {
            if (this.state.tab == 'attr') {
                return this.state.nodeSelect ?
                    <Space direction="vertical">
                        <Space><Text strong>节点：</Text><Text>{this.state.nodeName}</Text></Space>
                        <Space><Text strong>描述：</Text><Paragraph ellipsis={{
                            rows: 3,
                            expandable: true,
                            symbol: 'more'
                        }} style={{marginBottom: 0}}>{this.state.nodeDesc}</Paragraph></Space>
                        <Space><Text strong>度数：</Text><Text>{this.state.nodeDeg}</Text></Space>
                    </Space> :
                    <Space direction="vertical">
                        <Space><Text strong>图：</Text><Text>{name}</Text></Space>
                        <Space><Text strong>描述：</Text><Paragraph ellipsis={{
                            rows: 3,
                            expandable: true,
                            symbol: 'more'
                        }} style={{marginBottom: 0}}>{this.props.graph.desc}</Paragraph></Space>
                        <Space><Text strong>节点数：</Text><Text>{this.props.graph.nodes}</Text></Space>
                        <Space><Text strong>边数：</Text><Text>{this.props.graph.edges}</Text></Space>
                    </Space>
            }
            return <Space direction={"vertical"}>
                <Space><Text strong>来源：</Text><Text>文件导入</Text></Space>
            </Space>
        }


        return (
            <PageContainer header={{title: ''}} content={
                <Row gutter={16} style={{height: '100%'}}>
                    <Col span={18}>
                        <Card
                            title={name}
                            style={{height: '100%'}}
                            bodyStyle={{padding: '0 0 0 0'}}
                            extra={<LayoutSelector options={layouts} value={this.state.type} onChange={handleChange}/>}
                        >
                            <Spin spinning={loading}>
                                <Graphin data={data} layout={layout} fitView={true} containerStyle={{height: '80vh'}}
                                         ref={this.state.graphinRef}>
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