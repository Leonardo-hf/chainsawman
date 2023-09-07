import {PlusOutlined} from '@ant-design/icons';
import {
    ActionType,
    DrawerForm,
    ProColumns,
    ProFormSelect,
    ProFormText,
    ProFormList,
    ProFormGroup, ProFormDependency, ProFormUploadButton
} from '@ant-design/pro-components';
import {PageContainer, ProTable} from '@ant-design/pro-components';
import {
    Button,
    Form,
    message, Space,
    Typography
} from 'antd';
import React, {CSSProperties, useRef} from 'react';
import ProCard from '@ant-design/pro-card';
import {createGraph, createGroup, dropGraph, getAllGraph, updateGraph} from '@/services/graph/graph';
import {useModel} from '@umijs/max';
import {upload} from '@/utils/oss';
import {splitGroupsGraph} from '@/models/global';

const {Title, Text} = Typography;

const HomePage: React.FC = (props) => {
    const ref = useRef<ActionType>();
    const {setGraphs, groups, setGroups} = useModel('global')
    // 创建图谱的drawer
    const getNewGraphModal = () => {
        // form提交处理函数
        const handleCreateGraph = async (value: Graph.CreateGraphRequest) => {
            const {graph, desc, groupId} = value;
            return await createGraph({groupId: groupId, graph: graph, desc: desc})
                .then(() => {
                    message.success('图谱创建中...')
                    ref.current?.reload()
                    return true
                }).catch(e => {
                    console.log(e)
                })
        }
        const [form] = Form.useForm<Graph.CreateGraphRequest>()
        return <DrawerForm<Graph.CreateGraphRequest>
            title='新建图谱'
            resize={{
                maxWidth: window.innerWidth * 0.6,
                minWidth: window.innerWidth * 0.3,
            }}
            form={form}
            trigger={
                <Button type='primary'>
                    <PlusOutlined/>
                    新建图谱
                </Button>
            }
            autoFocusFirstInput
            drawerProps={{
                destroyOnClose: true,
            }}
            onFinish={handleCreateGraph}
        >
            <ProFormText name='graph' label='名称' rules={[{required: true}]}/>
            <ProFormText name='desc' label='描述' rules={[{required: true}]}/>
            <ProFormSelect
                options={groups.map(g => {
                    return {
                        label: g.name,
                        value: g.id
                    }
                })}
                rules={[{required: true}]}
                name='groupId'
                label='策略组'
            />
            <ProFormDependency name={["groupId"]} key="d1">
                {({groupId}) => {
                    if (groupId) {
                        const crtGroup = groups.find(g => g.id === groupId)!
                        const filterGroup = {
                            name: crtGroup.name,
                            desc: crtGroup.desc,
                            nodes: crtGroup.nodeTypeList.map(n => {
                                return {
                                    name: n.name,
                                    desc: n.desc,
                                    attr: n.attrs.map(a => {
                                        return {
                                            name: a.name,
                                            desc: a.desc,
                                            primary: a.primary,
                                            type: a.type === 0 ? 'string' : 'number'
                                        }
                                    })
                                }
                            }),
                            edges: crtGroup.edgeTypeList.map(n => {
                                return {
                                    name: n.name,
                                    desc: n.desc,
                                    isDirect: n.edgeDirection,
                                    attr: n.attrs.map(a => {
                                        return {
                                            name: a.name,
                                            desc: a.desc,
                                            primary: a.primary,
                                            type: a.type === 0 ? 'string' : 'number'
                                        }
                                    })
                                }
                            }),
                        }
                        return <Typography>
                            <pre style={{border: 'none'}}>{JSON.stringify(filterGroup, null, 2)}</pre>
                        </Typography>
                    }
                }}
            </ProFormDependency>
        </DrawerForm>
    }
    // 更新图谱的drawer
    const getUpdateGraphModal = (graphId: number, groupId: number, extraStyle: React.CSSProperties | undefined) => {
        const crtGroup = groups.find(gr => gr.id === groupId)
        if (!crtGroup) {
            return <DrawerForm trigger={<a style={extraStyle}>更新</a>}/>
        }
        // form提交处理函数
        const handleUpdateGraph = async (value: any) => {
            let nodeFiles = [], edgeFiles = []
            // 读取当前组各个可能上传的文件
            for (const n of crtGroup.nodeTypeList) {
                if (value[n.name] && value[n.name].length) {
                    nodeFiles.push({
                        key: n.id.toString(),
                        value: await upload(value[n.name][0].originFileObj)
                    })
                }
            }
            for (const n of crtGroup.edgeTypeList) {
                if (value[n.name] && value[n.name].length) {
                    edgeFiles.push({
                        key: n.id.toString(),
                        value: await upload(value[n.name][0].originFileObj)
                    })
                }
            }
            return await updateGraph({edgeFileList: edgeFiles, graphId: graphId, nodeFileList: nodeFiles})
                .then(() => {
                    message.success('图谱更新中...')
                    ref.current?.reload()
                    return true
                }).catch(e => {
                    console.log(e)
                })
        }
        // 为当前组各个可能上传的文件生成上传栏
        const getUploadFormItem = (crtGroup: Graph.Group) => {
            let nodeItems: JSX.Element[] = []
            let edgeItems: JSX.Element[] = []
            crtGroup.nodeTypeList.forEach(t => {
                nodeItems.push(<ProFormUploadButton
                    max={1}
                    name={t.name}
                    label={t.name + '文件[可选]'}>
                </ProFormUploadButton>)
            })
            crtGroup.edgeTypeList.forEach(t => {
                edgeItems.push(<ProFormUploadButton
                    max={1}
                    name={t.name}
                    label={t.name + '文件[可选]'}>
                </ProFormUploadButton>)
            })
            return {nodeItems, edgeItems}
        }
        const {nodeItems, edgeItems} = getUploadFormItem(crtGroup)
        return <DrawerForm
            title='更新图谱'
            resize={{
                maxWidth: window.innerWidth * 0.4,
                minWidth: window.innerWidth * 0.2,
            }}
            trigger={
                <a style={extraStyle}>更新</a>
            }
            autoFocusFirstInput
            drawerProps={{
                destroyOnClose: true,
            }}
            onFinish={handleUpdateGraph}>
            <ProFormGroup title='节点文件'>
                <Space direction={"vertical"}>
                    {nodeItems}
                </Space>
            </ProFormGroup>
            <ProFormGroup title='边文件'>
                <Space direction={"vertical"}>
                    {edgeItems}
                </Space>
            </ProFormGroup>
        </DrawerForm>
    }
    // 创建策略组的drawer
    const getCreateGroupModal = () => {
        // form提交处理函数
        const handleCreateGroup = async (vs: FormData) => {
            let edgeTypeList: Graph.Structure[] = [], nodeTypeList: Graph.Structure[] = []
            for (let v of vs.entities) {
                if (v.type === 'node') {
                    nodeTypeList.push({
                        id: 0,
                        attrs: v.attrs,
                        desc: v.desc,
                        display: v.display,
                        edgeDirection: false,
                        name: v.name
                    })
                } else {
                    edgeTypeList.push({
                        id: 0,
                        attrs: v.attrs,
                        desc: v.desc,
                        display: v.display,
                        edgeDirection: v.direct,
                        name: v.name
                    })
                }
            }
            return await createGroup({
                name: vs.name,
                desc: vs.desc,
                edgeTypeList: edgeTypeList,
                nodeTypeList: nodeTypeList
            }).then(() => {
                message.success('策略组创建成功！')
                ref.current?.reload()
                return true
            }).catch(e => {
                console.log(e)
            })
        }
        type FormData = {
            name: string, desc: string, entities: {
                name: string, desc: string, type: string, display: string, direct: boolean,
                attrs: { name: string, desc: string, type: number, primary: boolean }[]
            }[]
        }
        const [form] = Form.useForm<FormData>()
        return <DrawerForm<FormData>
            title='新建策略组'
            resize={{
                maxWidth: window.innerWidth * 0.8,
                minWidth: window.innerWidth * 0.6,
            }}
            form={form}
            trigger={
                <Button type='primary'>
                    <PlusOutlined/>
                    新建策略组
                </Button>
            }
            autoFocusFirstInput
            drawerProps={{
                destroyOnClose: true,
            }}
            onFinish={handleCreateGroup}
        >
            <ProFormGroup title='策略组配置'>
                <ProFormText name='name' label='名称' rules={[{required: true}]}/>
                <ProFormText name='desc' label='描述' rules={[{required: true}]}/>
            </ProFormGroup>
            <ProFormList
                label={(<Text strong>实体组配置</Text>)}
                initialValue={[{
                    name: 'node_default',
                    desc: '默认节点',
                    type: 'node',
                    attrs: [{
                        name: 'name',
                        desc: '节点名称',
                        type: 0,
                        primary: true
                    }, {
                        name: 'desc',
                        desc: '节点描述',
                        type: 0
                    }]
                }, {
                    name: 'edge_default',
                    desc: '默认边',
                    type: 'edge',
                }]}
                name='entities'
                creatorButtonProps={{
                    creatorButtonText: '添加一个实体'
                }}
                itemRender={({listDom, action}, {record}) => {
                    return (
                        <ProCard
                            bordered
                            extra={action}
                            title={record?.name}
                            style={{
                                marginBlockEnd: 8,
                            }}
                        >
                            {listDom}
                        </ProCard>
                    );
                }}
            >
                <ProFormGroup>
                    <ProFormText name='name' label='名称'/>
                    <ProFormText name='desc' label='描述'/>
                    <ProFormSelect
                        initialValue={'node'}
                        options={[
                            {
                                label: '节点',
                                value: 'node'
                            }, {
                                label: '边',
                                value: 'edge'
                            }
                        ]}
                        name='type'
                        label='实体类型'
                    />
                    <ProFormDependency key="d2" name={['type']}>
                        {({type}) => {
                            if (type === 'node') {
                                return <ProFormSelect
                                    initialValue={'color'}
                                    options={[
                                        {
                                            label: '彩色节点',
                                            value: 'color'
                                        }, {
                                            label: '图标节点',
                                            value: 'icon'
                                        }
                                    ]}
                                    name='display'
                                    label='展示'
                                />
                            }
                            if (type === 'edge') {
                                return <Space size={"large"}>
                                    <ProFormSelect
                                        initialValue={true}
                                        options={[
                                            {
                                                label: '有向',
                                                value: true,
                                            }, {
                                                label: '无向',
                                                value: false
                                            }
                                        ]}
                                        name='direct'
                                        label='方向'
                                    />
                                    <ProFormSelect
                                        initialValue={'real'}
                                        options={[
                                            {
                                                label: '实线',
                                                value: 'real'
                                            }, {
                                                label: '虚线',
                                                value: 'dash'
                                            }
                                        ]}
                                        name='display'
                                        label='展示'
                                    />
                                </Space>
                            }
                        }}
                    </ProFormDependency>
                </ProFormGroup>
                <ProFormList
                    creatorButtonProps={{
                        creatorButtonText: '添加一个属性'
                    }}
                    name='attrs'
                    label='属性'
                    deleteIconProps={{
                        tooltipText: '删除本行',
                    }}
                >
                    <ProFormGroup key='group'>
                        <ProFormText name='name' label='名称'/>
                        <ProFormText name='desc' label='描述'/>
                        <ProFormSelect
                            initialValue={0}
                            options={[
                                {
                                    label: '字符串',
                                    value: 0
                                }, {
                                    label: '数值',
                                    value: 1
                                }
                            ]}
                            name='type'
                            label='类型'
                        />
                        <ProFormSelect
                            initialValue={false}
                            options={[
                                {
                                    label: '是',
                                    value: true
                                }, {
                                    label: '否',
                                    value: false
                                }
                            ]}
                            name='primary'
                            label='主属性'
                        />
                    </ProFormGroup>
                </ProFormList>
            </ProFormList>
        </DrawerForm>
    }
    // 表格列
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
            title: '组',
            dataIndex: 'groupId',
            copyable: true,
            ellipsis: true,
            hideInSearch: true,
            render: (text, record, _, action) => {
                return groups.find((g: Graph.Group) => g.id === record.groupId)?.name
            },
        },
        {
            title: '节点',
            dataIndex: 'numNode',
            sorter: {
                compare: (a, b) => a.numNode - b.numNode,
                multiple: 1
            },
            hideInSearch: true
        },
        {
            title: '边',
            dataIndex: 'numEdge',
            sorter: {
                compare: (a, b) => a.numEdge - b.numEdge,
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
                2: {
                    text: '更新中',
                    status: 'Updating',
                },
                1: {
                    text: '在线',
                    status: 'Online',
                },
                0: {
                    text: '创建中',
                    status: 'Offline',
                },
            },
        },
        {
            title: '创建时间',
            dataIndex: 'creatAt',
            valueType: 'date',
            sorter: {
                compare: (a, b) => a.creatAt - b.creatAt,
                multiple: 3
            },
            hideInSearch: true
        },
        {
            title: '更新时间',
            dataIndex: 'updateAt',
            valueType: 'date',
            sorter: {
                compare: (a, b) => a.updateAt - b.updateAt,
                multiple: 3
            },
            hideInSearch: true
        },
        {
            title: '操作',
            key: 'option',
            valueType: 'option',
            render: (text, record, _, action) => {
                const disable: CSSProperties | undefined = record.status !== 1 ? {
                    pointerEvents: 'none',
                    color: 'grey'
                } : undefined
                return <Space>
                    <a href={'/graph/' + record.groupId + '/' + record.id} style={disable}>
                        查看
                    </a>
                    {getUpdateGraphModal(record.id, record.groupId, disable)}
                    <a style={disable} onClick={() => {
                        dropGraph({
                            graphId: record.id
                        }).then(() => {
                            message.success('删除图`' + record.name + '`成功')
                            action?.reload()
                        })
                    }}>
                        删除
                    </a>
                </Space>
            },
        },
    ]

    // 查询图谱信息
    async function checkGraphs() {
        return await getAllGraph()
            .then(res => {
                if (!res.groups) {
                    return []
                }
                const {graphs, groups} = splitGroupsGraph(res.groups)
                // TODO: 可能反复触发重新渲染？
                setGroups(groups)
                setGraphs(graphs)
                // 如果存在状态不为完成的图，则尝试再次请求
                if (graphs.filter(a => a.status !== 1).length) {
                    setTimeout(_ => {
                        ref.current?.reload()
                    }, 5000)
                }
                return graphs
            })
    }

    return (
        <PageContainer>
            <ProTable<Graph.Graph>
                key='graphList'
                columns={columns}
                cardBordered
                actionRef={ref}
                request={async (params, sort, filter) => {
                    console.log('reload home')
                    const keyword = params.name ? params.name : ''
                    // 查询最新的图与组信息并更新
                    const graphs = await checkGraphs()
                    const fGraphs = graphs.filter(g => g.name.includes(keyword) && (!params.status || g.status == params.status))
                    return {
                        data: fGraphs,
                        success: true,
                        total: fGraphs.length
                    }
                }}
                columnsState={{
                    persistenceKey: 'graphs_columns_state',
                    persistenceType: 'localStorage',
                }}
                rowKey={(record) => {
                    return record.id
                }}
                search={{
                    labelWidth: 'auto',
                }}
                options={{
                    setting: {
                        listsHeight: 400,
                    },
                }}
                pagination={{
                    position: ['bottomCenter', 'bottomRight'],
                    pageSize: 10,
                }}
                dateFormatter='string'
                headerTitle={<Title level={4} style={{margin: '0 0 0 0'}}>图谱列表</Title>}
                toolBarRender={(action) => [
                    getCreateGroupModal(),
                    getNewGraphModal(),
                ]}
            />
        </PageContainer>
    );
};

export default HomePage;

