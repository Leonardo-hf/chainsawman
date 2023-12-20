import {PlusOutlined} from '@ant-design/icons';
import {
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
import React, {CSSProperties} from 'react';
import {createGraph, dropGraph, updateGraph} from '@/services/graph/graph';
import {useModel} from '@umijs/max';
import {uploadSource} from '@/utils/oss';
import {GraphRef2Group, TreeNodeGroup} from '@/models/global';


const GTable: React.FC = () => {
    const {graphs, groups} = useModel('global')
    // 创建图谱的drawer
    const getNewGraphModal = () => {
        // form提交处理函数
        const handleCreateGraph = async (value: Graph.CreateGraphRequest) => {
            const {graph, groupId} = value;
            return await createGraph({groupId: groupId, graph: graph})
                .then(() => {
                    message.success('图谱创建中...')
                    window.location.reload()
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
                                    primary: n.primary,
                                    attr: n.attrs?.map(a => {
                                        return {
                                            name: a.name,
                                            desc: a.desc,
                                            type: a.type === 0 ? 'string' : 'number'
                                        }
                                    })
                                }
                            }),
                            edges: crtGroup.edgeTypeList.map(n => {
                                return {
                                    name: n.name,
                                    desc: n.desc,
                                    primary: n.primary,
                                    isDirect: n.edgeDirection,
                                    attr: n.attrs?.map(a => {
                                        return {
                                            name: a.name,
                                            desc: a.desc,
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
                    window.location.reload()
                    return true
                }).catch(e => {
                    console.log(e)
                })
        }

        // 为当前组各个可能上传的文件生成上传栏
        const getUploadFormItem = (crtGroup: TreeNodeGroup) => {
            let nodeItems: JSX.Element[] = crtGroup.nodeTypeList.map(t => <ProFormUploadButton
                    max={1}
                    fieldProps={{
                        customRequest: (option: any) => {
                            option.onSuccess()
                        }
                    }}
                    key={t.name}
                    name={t.name}
                    label={t.name + '文件[可选]'}>
                </ProFormUploadButton>)
            let edgeItems: JSX.Element[] = crtGroup.edgeTypeList.map(t => <ProFormUploadButton
                    fieldProps={{
                        customRequest: (option: any) => {
                            option.onSuccess()
                        }
                    }}
                    max={1}
                    key={t.name}
                    name={t.name}
                    label={t.name + '文件[可选]'}>
                </ProFormUploadButton>)
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
        text: g.desc,
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
            title: '图结构',
            copyable: true,
            ellipsis: true,
            valueType: 'select',
            valueEnum: groupsEnum,
            filters: true,
            onFilter: true,
            tooltip: '对图中节点、边结构的一组定义',
            render: (_, d) => d.group.desc
        },
        {
            title: '节点数目',
            dataIndex: 'numNode',
        },
        {
            title: '边数目',
            dataIndex: 'numEdge',
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
            hideInSearch: true
        },
        {
            title: '更新时间',
            dataIndex: 'updateAt',
            valueType: 'date',
            sorter: {
                compare: (a, b) => a.updateAt - b.updateAt,
                multiple: 1
            },
            hideInSearch: true
        },
        {
            title: '操作',
            key: 'option',
            valueType: 'option',
            render: (text, record) => {
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
                            window.location.reload()
                        })
                    }}>
                        删除
                    </a>
                </Space>
            },
        },
    ]

    return (
        <PageContainer title={false}>
            <ProTable<GraphRef2Group>
                key='id'
                columns={columns}
                cardBordered
                dataSource={graphs}
                columnsState={{
                    persistenceKey: 'graphs_columns_state',
                    persistenceType: 'localStorage',
                }}
                rowKey={(record) => {
                    return record.id
                }}
                options={{
                    setting: {
                        listsHeight: 400,
                    },
                }}
                pagination={false}
                search={false}
                dateFormatter='string'
                headerTitle={'图谱列表'}
                toolBarRender={() => [
                    getNewGraphModal(),
                ]}
            />
        </PageContainer>
    );
};

export default GTable;

