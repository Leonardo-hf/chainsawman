import {PlusOutlined} from '@ant-design/icons';
import {
    ActionType,
    DrawerForm,
    PageContainer,
    ProColumns,
    ProFormDependency,
    ProFormGroup,
    ProFormSelect,
    ProFormText,
    ProFormUploadButton,
    ProTable
} from '@ant-design/pro-components';
import {Button, Form, message, Space, Typography} from 'antd';
import React, {CSSProperties, useRef} from 'react';
import {createGraph, dropGraph, getAllGraph, updateGraph} from '@/services/graph/graph';
import {useModel} from '@umijs/max';
import {uploadSource} from '@/utils/oss';
import {GraphRef2Group, parseGroups, TreeNodeGroup} from '@/models/global';

const {Title} = Typography;

const HomePage: React.FC = () => {
    const ref = useRef<ActionType>();
    const {setGraphs, groups, setGroups} = useModel('global')
    // 创建图谱的drawer
    const getNewGraphModal = () => {
        // form提交处理函数
        const handleCreateGraph = async (value: Graph.CreateGraphRequest) => {
            const {graph, groupId} = value;
            return await createGraph({groupId: groupId, graph: graph})
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
            <ProFormSelect
                options={groups.map(g => {
                    return {
                        label: g.desc,
                        value: g.id
                    }
                })}
                rules={[{required: true}]}
                name='groupId'
                label='图结构'
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
                                    attr: n.attrs?.map(a => {
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
                                    attr: n.attrs?.map(a => {
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
    const getUpdateGraphModal = (graphId: number, group: TreeNodeGroup, extraStyle: React.CSSProperties | undefined) => {
        if (!group) {
            return <DrawerForm trigger={<a style={extraStyle}>更新</a>}/>
        }
        // form提交处理函数
        const handleUpdateGraph = async (value: any) => {
            let nodeFiles = [], edgeFiles = []
            // 读取当前组各个可能上传的文件
            for (const n of group.nodeTypeList) {
                if (value[n.name] && value[n.name].length) {
                    nodeFiles.push({
                        key: n.id.toString(),
                        value: await uploadSource(value[n.name][0].originFileObj)
                    })
                }
            }
            for (const n of group.edgeTypeList) {
                if (value[n.name] && value[n.name].length) {
                    edgeFiles.push({
                        key: n.id.toString(),
                        value: await uploadSource(value[n.name][0].originFileObj)
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
        const getUploadFormItem = (crtGroup: TreeNodeGroup) => {
            let nodeItems: JSX.Element[] = []
            let edgeItems: JSX.Element[] = []
            crtGroup.nodeTypeList.forEach(t => {
                nodeItems.push(<ProFormUploadButton
                    max={1}
                    fieldProps={{
                        customRequest: (option: any) => {
                            option.onSuccess()
                        }
                    }}
                    name={t.name}
                    label={t.name + '文件[可选]'}>
                </ProFormUploadButton>)
            })
            crtGroup.edgeTypeList.forEach(t => {
                edgeItems.push(<ProFormUploadButton
                    fieldProps={{
                        customRequest: (option: any) => {
                            option.onSuccess()
                        }
                    }}
                    max={1}
                    name={t.name}
                    label={t.name + '文件[可选]'}>
                </ProFormUploadButton>)
            })
            return {nodeItems, edgeItems}
        }
        const {nodeItems, edgeItems} = getUploadFormItem(group)
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

    const groupsEnum: any = {}
    groups.forEach(g => groupsEnum[g.id] = {
        text: g.name,
    })

    const getGraphRoute = (g: GraphRef2Group) => {
        let href = '/#/graph'
        let group = g.group
        let groupHref: string = group.id.toString()
        while (group.parentId > 1) {
            group = group.parentGroup!
            groupHref = group.id + '/' + groupHref
        }
        return href + '/' + groupHref + '/s/' + g.id
    }
    // 表格列
    const columns: ProColumns<GraphRef2Group>[] = [
        {
            title: '名称',
            dataIndex: 'name',
            copyable: true,
            fixed: 'left',
        },
        {
            title: '组',
            copyable: true,
            ellipsis: true,
            valueType: 'select',
            valueEnum: groupsEnum,
            filters: true,
            onFilter: true,
            hideInSearch: true,
            tooltip: '对图中节点、边结构的一组定义',
            render: (_,d)=>d.group.desc
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
            hideInSearch: true,
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
                    <a href={getGraphRoute(record)} style={disable}>
                        查看
                    </a>
                    {getUpdateGraphModal(record.id, record.group, disable)}
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
                const {graphs, groups} = parseGroups(res.groups)
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
            <ProTable<GraphRef2Group>
                key='graphList'
                columns={columns}
                cardBordered
                actionRef={ref}
                request={async (params) => {
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
                pagination={false}
                dateFormatter='string'
                headerTitle={<Title level={4} style={{margin: '0 0 0 0'}}>图谱列表</Title>}
                toolBarRender={() => [
                    getNewGraphModal(),
                ]}
            />
        </PageContainer>
    );
};

export default HomePage;

