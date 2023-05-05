import {EllipsisOutlined, PlusOutlined, UploadOutlined} from '@ant-design/icons';
import type {ActionType, ProColumns} from '@ant-design/pro-components';
import {PageContainer, ProTable} from '@ant-design/pro-components';
import {Button, Dropdown, Form, Input, message, Modal, Typography, Upload, UploadProps} from 'antd';
import React, {useRef, useState} from 'react';
import FormItem from "antd/es/form/FormItem";
import ProCard from "@ant-design/pro-card";
import {upload} from "@/services/file/file";
import {createGraph, getAllGraph} from "@/services/graph/graph";
import {
    CSSProperties
} from "../../../../../../../../usr/local/src/GoLand-2022.3.1/plugins/javascript-impl/jsLanguageServicesImpl/external/react";
import {useModel} from '@umijs/max';

const {Title} = Typography;

const columns: ProColumns<Graph.Graph>[] = [
    {
        title: '名称',
        dataIndex: 'name',
        copyable: true,
        fixed: 'left'
    },
    {
        title: '描述',
        dataIndex: 'desc',
        copyable: true,
        ellipsis: true,
        hideInSearch: true
    },
    {
        title: '节点',
        dataIndex: 'nodes',
        sorter: {
            compare: (a, b) => a.nodes - b.nodes,
            multiple: 1
        },
        hideInSearch: true
    },
    {
        title: '边',
        dataIndex: 'edges',
        sorter: {
            compare: (a, b) => a.edges - b.edges,
            multiple: 2
        },
        hideInSearch: true
    },
    {
        disable: true,
        title: '状态',
        dataIndex: 'status',
        filters: true,
        onFilter: true,
        valueType: 'select',
        valueEnum: {
            1: {
                text: '在线',
                status: 'Online',
            },
            0: {
                text: '离线',
                status: 'Offline',
            },
        },
    },
    {
        title: '创建时间',
        dataIndex: 'createdAt',
        valueType: 'date',
        sorter: {
            compare: (a, b) => a.createdAt - b.createdAt,
            multiple: 3
        },
        hideInSearch: true
    },
    {
        title: '更新时间',
        dataIndex: 'updatedAt',
        valueType: 'date',
        sorter: {
            compare: (a, b) => a.updatedAt - b.updatedAt,
            multiple: 3
        },
        hideInSearch: true
    },
    {
        title: '操作',
        valueType: 'option',
        key: 'option',
        render: (text, record, _, action) => {
            const disable: CSSProperties | undefined = record.status === 0 ? {
                pointerEvents: 'none',
                color: 'grey'
            } : undefined
            return [
                <a href={'/graph/' + record.id} style={disable}>
                    查看
                </a>,
                <a style={disable}>
                    删除
                </a>,
            ]
        },
    },
];


const HomePage: React.FC = (props) => {
    const ref = useRef<ActionType>();
    const {setGraphs} = useModel('global')
    const [modalOpen, setModalOpen] = useState(false)
    const getModal = () => {
        const uploadProps: UploadProps = {
            beforeUpload: (file) => {
                return false;
            },
            maxCount: 1
        };

        const normFile = (e: any) => {
            return e?.fileList;
        };

        const handleFinish = async (value: any) => {
            const {name, desc, node, edge} = value;
            const nodeData = new FormData();
            if (!name) {
                message.error('必须输入一个图名称')
                return
            }
            if (!desc) {
                message.error('必须输入图的描述')
                return
            }
            if (!node) {
                message.error('必须上传一个节点文件')
                return
            }
            if (!edge) {
                message.error('必须上传一个边文件')
                return
            }
            nodeData.append('file', node[0].originFileObj)
            //console.log(nodeData.get('node'))
            const edgeData = new FormData();
            edgeData.append('file', edge[0].originFileObj)
            let nodeId: string = '', edgeId: string = '';
            await upload(
                {
                    headers: {
                        'Content-Type': 'multipart/form-data',
                    },
                    data: nodeData
                }
            ).then(res => {
                nodeId = res.id
            }).catch(e => console.log(e))
            await upload(
                {
                    headers: {
                        'Content-Type': 'multipart/form-data',
                    },
                    data: edgeData
                }
            ).then(res => {
                edgeId = res.id
            }).catch(e => console.log(e))
            createGraph({taskId: 0, nodeId: nodeId, edgeId: edgeId, graph: name, desc: desc})
                .then((res) => {
                    setModalOpen(false)
                    message.success("文件上传成功")
                    // let haha: Graph.Graph[] = graphs
                    // haha.push({id: res.graph.id, name: name, desc: desc, nodes: 0, edges: 0, status: 0})
                    // setGraphs(haha)
                    ref.current?.reload()
                }).catch(e=>{console.log(e)})
        }
        return <Modal open={modalOpen} footer={null} onCancel={()=>setModalOpen(false)}>
            <ProCard style={{height: "fit-content"}}>
                <Form onFinish={handleFinish}>
                    <FormItem
                        name='name'
                        label={"图名称"}
                    >
                        <Input/>
                    </FormItem>
                    <FormItem
                        name='desc'
                        label={"图描述"}
                    >
                        <Input/>
                    </FormItem>
                    <FormItem
                        name="node"
                        valuePropName="file"
                        getValueFromEvent={normFile}
                        label="节点文件"
                    >
                        <Upload {...uploadProps}>
                            <Button icon={<UploadOutlined/>}>Click to Upload</Button>
                        </Upload>
                    </FormItem>
                    <FormItem
                        name="edge"
                        valuePropName="file"
                        getValueFromEvent={normFile}
                        label="边文件"
                    >
                        <Upload {...uploadProps}>
                            <Button icon={<UploadOutlined/>}>Click to Upload</Button>
                        </Upload>
                    </FormItem>

                    <FormItem wrapperCol={{span: 12, offset: 6}}>
                        <Button type="primary" htmlType="submit">
                            确认发布
                        </Button>
                    </FormItem>
                </Form>
            </ProCard>
        </Modal>
    }

    async function checkGraphs() {
        let graphs: Graph.Graph[] = []
        await getAllGraph()
            .then(res => {
                graphs = res.graphs
                setGraphs(graphs)
                if (graphs.filter(a => a.status === 0).length) {
                    setTimeout(_ => {
                        ref.current?.reload()
                    }, 5000)
                }
            })
        return graphs
    }

    return (
        <PageContainer>
            {getModal()}
            {/*<button onClick={test}></button>*/}
            <ProTable<Graph.Graph>
                columns={columns}
                cardBordered
                actionRef={ref}
                request={async (params, sort, filter) => {
                    // TODO: graphs是打开主页发起查询获得的，感觉这里有可能会查不到graphs
                    // TODO: 根据status查询
                    // 读以下global的图，如果有效直接返回
                    // 如果无效（新增了），发起查询，看有没有正在创建的，有就个五秒
                    console.log('reload')
                    const keyword = params.name ? params.name : ''
                    const graphs = await checkGraphs()
                    const fGraphs = graphs.filter(g => g.name.includes(keyword) && (!params.status || g.status == params.status))
                    return {
                        data: fGraphs,
                        success: true,
                        total: graphs.length
                    }
                }}
                columnsState={{
                    persistenceKey: 'graphs_columns_state',
                    persistenceType: 'localStorage',
                    // onChange(value) {
                    //     console.log('value: ', value);
                    // },
                }}
                rowKey={record => record.name}
                search={{
                    labelWidth: 'auto',
                }}
                options={{
                    setting: {
                        listsHeight: 400,
                    },
                }}
                // form={{
                //     // 由于配置了 transform，提交的参与与定义的不同这里需要转化一下
                //     syncToUrl: (values, type) => {
                //         if (type === 'get') {
                //             return {
                //                 ...values,
                //                 created_at: [values.startTime, values.endTime],
                //             };
                //         }
                //         return values;
                //     },
                // }}
                pagination={{
                    position: ['bottomCenter', 'bottomRight'],
                    pageSize: 10,
                    // onChange: (page) => console.log(page),
                }}
                dateFormatter="string"
                headerTitle={<Title level={4} style={{margin: '0 0 0 0'}}>图谱列表</Title>}
                toolBarRender={(action) => [
                    <Button
                        key="button"
                        icon={<PlusOutlined/>}
                        onClick={() => {
                            setModalOpen(true)
                        }}
                        type="primary"
                    >
                        新建
                    </Button>,
                    <Dropdown
                        key="menu"
                        menu={{
                            items: [
                                {
                                    label: '1st item',
                                    key: '1',
                                },
                                {
                                    label: '2nd item',
                                    key: '2',
                                },
                                {
                                    label: '3rd item',
                                    key: '3',
                                },
                            ],
                        }}
                    >
                        <Button>
                            <EllipsisOutlined/>
                        </Button>
                    </Dropdown>,
                ]}
            />
        </PageContainer>
    );
};

export default HomePage;

