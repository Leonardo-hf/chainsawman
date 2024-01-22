import {
    DrawerForm,
    PageContainer, ProColumns, ProFormDependency, ProFormDigit, ProFormDigitRange,
    ProFormGroup,
    ProFormList,
    ProFormSelect,
    ProFormText, ProFormTextArea, ProFormUploadButton,
    ProTable
} from "@ant-design/pro-components";
import {Button, Divider, Form, message, Popconfirm, Space, Tag, Typography} from "antd";
import {PlusOutlined, QuestionCircleOutlined} from "@ant-design/icons";
import {algoCreate, algoDrop} from "@/services/graph/graph";
import React from "react";
import ProCard from "@ant-design/pro-card";
import {useModel} from "@@/exports";
import {
    AlgoOptions,
    getPrecise,
    ParamType,
    ParamTypeOptions
} from "@/constants";
import {uploadLib} from "@/utils/oss";
import {genGroupOptions} from "@/models/global";
import 'katex/dist/katex.min.css'
import {limitStr} from "@/constants/sp";
import {history} from "umi";

const {Text} = Typography;

type AlgoMenuProps = {
    category: {
        algos: Graph.Algo[],
        tag: string
    }[]
}

const AlgoMenu: React.FC<AlgoMenuProps> = (props) => {
    //@ts-ignore
    const {groups} = useModel('global')

    // 创建策略组的drawer
    const getCreateAlgoModal = () => {
        // form提交处理函数
        const handleCreateAlgo = async (vs: FormData) => {
            const jar = await uploadLib(vs.jar[0].originFileObj)
            return await algoCreate({
                entryPoint: vs.entryPoint,
                jar: jar,
                algo: {
                    name: vs.name,
                    detail: vs.detail,
                    tag: vs.tag,
                    groupId: vs.groupId,
                    params: vs.params.map(p => {
                        return {
                            key: p.key,
                            keyDesc: p.keyDesc,
                            type: p.type,
                            max: p.range ? p.range[1].toString() : undefined,
                            min: p.range ? p.range[0].toString() : undefined,
                            initValue: p.initValue.toString()
                        }
                    })
                },
            }).then(() => {
                message.success('指标创建成功！')
                window.location.reload()
                return true
            }).catch(e => {
                console.log(e)
            })
        }
        type FormData = {
            name: string, desc: string, tag: string, groupId: number, detail: string,
            entryPoint: string, jar: any,
            params: { key: string, keyDesc: string, type: ParamType, initValue: number, range: number[] }[]
        }
        const [form] = Form.useForm<FormData>()
        const algoOptions = AlgoOptions.filter(op => {
            return props.category.map(c => c.tag).indexOf(op.label) >= 0
        })
        return <DrawerForm<FormData>
            title='新建指标'
            resize={{
                maxWidth: window.innerWidth * 0.8,
                minWidth: window.innerWidth * 0.6,
            }}
            form={form}
            trigger={
                <Button type='primary'>
                    <PlusOutlined/>
                    新建指标
                </Button>
            }
            autoFocusFirstInput
            drawerProps={{
                destroyOnClose: true,
            }}
            onFinish={handleCreateAlgo}
        >
            <ProFormGroup title='指标配置'>
                <ProFormText name='name' label='名称' rules={[{required: true}]}/>
                <ProFormSelect
                    initialValue={algoOptions[0].label}
                    options={algoOptions}
                    rules={[{required: true}]}
                    name='tag'
                    label='指标标签'
                />
                <ProFormSelect name='groupId' style={{width: '100%'}} label='适用策略组' rules={[{required: true}]}
                               options={genGroupOptions(groups)}/>
            </ProFormGroup>
            <ProFormTextArea name={'detail'} label={'详情'} rules={[{required: true}]}/>
            <ProFormGroup title='指标实现'>
                <ProFormText name='entryPoint' label='入口类' rules={[{required: true}]}/>
                <ProFormUploadButton rules={[{required: true}]}
                                     fieldProps={{
                                         customRequest: (option: any) => {
                                             option.onSuccess()
                                         }
                                     }}
                                     max={1}
                                     name='jar'
                                     label={'JAR'}/>
            </ProFormGroup>
            <ProFormList
                label={(<Text strong>参数列表</Text>)}
                name='params'
                creatorButtonProps={{
                    creatorButtonText: '添加一个参数'
                }}
                itemRender={({listDom, action}, {record}) => {
                    return (
                        <ProCard
                            bordered
                            extra={action}
                            title={record?.key}
                            style={{
                                marginBlockEnd: 8,
                            }}
                        >
                            {listDom}
                        </ProCard>
                    )
                }}
            >
                <ProFormGroup>
                    <ProFormText name='key' label='key' rules={[{required: true}]}/>
                    <ProFormText name='keyDesc' label='名称' rules={[{required: true}]}/>
                    <ProFormSelect
                        initialValue={ParamType.Int}
                        options={ParamTypeOptions}
                        name='type'
                        label='类型'
                    />
                    <ProFormDependency name={['type']}>
                        {({type}) => type == ParamType.Int || type == ParamType.Double && <Space size={"large"}>
                            <ProFormDigitRange name='range' fieldProps={getPrecise(type)} label='范围'/>
                            <ProFormDigit name='initValue' fieldProps={getPrecise(type)} label='默认值'/>
                        </Space>
                        }
                    </ProFormDependency>
                </ProFormGroup>
            </ProFormList>
        </DrawerForm>
    }

    const columns: ProColumns<Graph.Algo>[] = [
        {
            title: '指标名称',
            dataIndex: 'name',
            render: (_, v) => {
                return <a href={`#${history.location.pathname.replace('index', v.id!.toString())}`}>{v.name}</a>
            }
        },
        {
            title: '定义',
            dataIndex: 'detail',
            render: (_, v) => {
                return limitStr(v.detail)
            }
        },
        {
            title: '适用图谱',
            dataIndex: '',
            render: (_, v) => {
                return <Tag
                    color='#FFA54F'>{v.groupId === 1 ? '通用' : groups.find(g => g.id === v.groupId)!.desc}</Tag>
            }
        },
        {
            title: 'Actions',
            render: (_, v) => {
                return <Popconfirm
                    title="确认删除？"
                    icon={<QuestionCircleOutlined style={{color: 'red'}}/>}
                    onConfirm={() => {
                        algoDrop({algoId: v.id!}).then(_ => window.location.reload())
                    }}
                >
                    <Button danger>删除</Button>
                </Popconfirm>
            }
        }
    ]

    return <PageContainer>
        <Space direction={"vertical"} style={{width: '100%'}}>
            {
                props.category.map((c, id) => {
                    return <ProTable
                        key={id}
                        rowKey='id'
                        dataSource={c.algos}
                        columns={columns}
                        search={false}
                        options={false}
                        toolbar={{
                            title: c.tag
                        }}
                        toolBarRender={() => [
                            getCreateAlgoModal()
                        ]}
                        pagination={false}>

                    </ProTable>
                })
            }
        </Space>
    </PageContainer>
}

export default AlgoMenu