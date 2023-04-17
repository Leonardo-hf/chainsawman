
import {EllipsisOutlined, PlusOutlined} from '@ant-design/icons';
import type {ActionType, ProColumns} from '@ant-design/pro-components';
import {PageContainer, ProTable} from '@ant-design/pro-components';
import {Button, Dropdown, Typography} from 'antd';
import {useRef} from 'react';
import {useModel} from "@@/exports";

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
            online: {
                text: '在线',
                status: 'Online',
            },
            offline: {
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
        render: (text, record, _, action) => [
            // <a
            //     key="editable"
            //     onClick={() => {
            //         action?.startEditable?.(record.id);
            //     }}
            // >
            //     编辑
            // </a>,
            <a href={'/graph/' + record.name}>
                查看
            </a>,
            <a>
                删除
            </a>,
            // <TableDropdown
            //     key="actionGroup"
            //     onSelect={() => action?.reload()}
            //     menus={[
            //         {key: 'copy', name: '复制'},
            //         {key: 'delete', name: '删除'},
            //     ]}
            // />,
        ],
    },
];

const HomePage: React.FC = (props) => {
    const actionRef = useRef<ActionType>();
    const {graphs} = useModel('global')
    return (
        <PageContainer>
            <ProTable<Graph.Graph>
                columns={columns}
                actionRef={actionRef}
                cardBordered
                request={async (params, sort, filter) => {
                    // TODO: graphs是打开主页发起查询获得的，感觉这里有可能会查不到graphs
                    // TODO: 根据status查询
                    const keyword = params.name ? params.name : ''
                    const fGraphs = graphs.filter(g => g.name.includes(keyword))
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
                toolBarRender={() => [
                    <Button
                        key="button"
                        icon={<PlusOutlined/>}
                        onClick={() => {
                            actionRef.current?.reload();
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

