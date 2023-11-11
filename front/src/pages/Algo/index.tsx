import {
    DrawerForm,
    PageContainer, ProDescriptions, ProFormDependency, ProFormDigit, ProFormDigitRange,
    ProFormGroup,
    ProFormList,
    ProFormSelect,
    ProFormText, ProFormUploadButton,
    ProList
} from "@ant-design/pro-components";
import {Button, Divider, Form, message, Popconfirm, Space, Tag, Typography} from "antd";
import {CloseOutlined, PlusOutlined, QuestionCircleOutlined} from "@ant-design/icons";
import {algoCreate, algoDrop} from "@/services/graph/graph";
import React, {Key, useState} from "react";
import ProCard from "@ant-design/pro-card";
import {useModel} from "@@/exports";
import {
    AlgoOptions,
    AlgoType,
    AlgoTypeMap,
    getAlgoTypeDesc,
    getPrecise,
    ParamType,
    ParamTypeOptions
} from "@/constants";
import {uploadLib} from "@/utils/oss";
import {genGroupOptions} from "@/models/global";

const {Text} = Typography;

const Algo: React.FC = (props) => {

    const [expandedRowKeys, setExpandedRowKeys] = useState<readonly Key[]>([]);
    const {initialState} = useModel('@@initialState')
    //@ts-ignore
    const {algos} = initialState
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
                    isCustom: true,
                    name: vs.name,
                    desc: vs.desc,
                    type: vs.type,
                    groupId: vs.groupId,
                    params: vs.params.map(p => {
                        return {
                            key: p.key,
                            keyDesc: p.keyDesc,
                            type: p.type,
                            max: p.range ? p.range[1] : undefined,
                            min: p.range ? p.range[0] : undefined,
                            initValue: p.initValue
                        }
                    })
                },
            }).then(() => {
                message.success('算法创建成功！')
                window.location.reload()
                return true
            }).catch(e => {
                console.log(e)
            })
        }
        type FormData = {
            name: string, desc: string, type: AlgoType, groupId: number,
            entryPoint: string, jar: any,
            params: { key: string, keyDesc: string, type: ParamType, initValue: number, range: number[] }[]
        }
        const [form] = Form.useForm<FormData>()
        return <DrawerForm<FormData>
            title='新建算法'
            resize={{
                maxWidth: window.innerWidth * 0.8,
                minWidth: window.innerWidth * 0.6,
            }}
            form={form}
            trigger={
                <Button type='primary'>
                    <PlusOutlined/>
                    新建算法
                </Button>
            }
            autoFocusFirstInput
            drawerProps={{
                destroyOnClose: true,
            }}
            onFinish={handleCreateAlgo}
        >
            <ProFormGroup title='算法配置'>
                <ProFormText name='name' label='名称' rules={[{required: true}]}/>
                <ProFormText name='desc' label='描述' rules={[{required: true}]}/>
                <ProFormSelect
                    initialValue={AlgoType.rank}
                    options={AlgoOptions}
                    rules={[{required: true}]}
                    name='type'
                    label='算法分类'
                />
                <ProFormSelect name='groupId' style={{width: '100%'}} label='适用策略组' options={genGroupOptions(groups)}/>
            </ProFormGroup>
            <ProFormGroup title='算法实现'>
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


    const getAlgoContent = (algo: Graph.Algo) => {
        // 设置算法的初始值
        const initValues: any = {}
        algo.params?.forEach(p => {
            if (p.initValue) {
                initValues[p.key] = p.initValue
            }
        })
        return <Space direction={'vertical'}>
            <Text type={'secondary'}>{algo.desc}</Text>
            {algo.params?.length! > 0 && <Divider/>}
            {
                algo.params?.map(p =>
                    <>
                        <ProDescriptions dataSource={p}>
                            <ProDescriptions.Item dataIndex={'key'} label='key' valueType={'text'}/>
                            <ProDescriptions.Item dataIndex={'keyDesc'} label='名称' valueType={'text'}/>
                            <ProDescriptions.Item dataIndex={'type'} label='类型' valueEnum={ParamType}/>
                            {p.initValue &&
                            <ProDescriptions.Item dataIndex={'initValue'} label='默认值' valueType={'text'}/>}
                            {p.max && <ProDescriptions.Item dataIndex={'max'} label='最大值' valueType={'text'}/>}
                            {p.min && <ProDescriptions.Item dataIndex={'min'} label='最小值' valueType={'text'}/>}
                        </ProDescriptions>
                        <Divider/>
                    </>
                )
            }

        </Space>
    }
    return <PageContainer>
        <ProList<Graph.Algo>
            rowKey="id"
            headerTitle="算法列表"
            toolBarRender={() => {
                return [
                    getCreateAlgoModal(),
                ];
            }}
            search={{
                filterType: 'light',
            }}
            request={
                async (params = {time: Date.now()}) => {
                    const data = algos.filter((a: Graph.Algo) => !params.subTitle || a.type == params.subTitle)
                    return {
                        data: data,
                        success: true,
                        total: data.length
                    }
                }}
            expandable={{expandedRowKeys, onExpandedRowsChange: setExpandedRowKeys}}
            metas={{
                title: {
                    dataIndex: 'name', search: false,
                },
                subTitle: {
                    title: '类别',
                    render: (_, row) => {
                        return <Space size={0}>
                            {getAlgoTypeDesc(row.type)}
                            <Tag
                                color='#FFA54F'>{row.groupId === 1 ? '通用' : groups.find(g => g.id === row.groupId)!.desc}</Tag>
                        </Space>
                    },
                    valueType: 'select',
                    valueEnum: AlgoTypeMap,
                },
                description: {
                    search: false,
                    render: (_, row) => getAlgoContent(row)
                },
                actions: {
                    render: (_, a) => {
                        return a.isCustom && <Popconfirm
                            title="确认删除？"
                            icon={<QuestionCircleOutlined style={{color: 'red'}}/>}
                            onConfirm={() => {
                                algoDrop({algoId: a.id!}).then(_ => window.location.reload())
                            }}
                        >
                            <Button danger>删除</Button>
                        </Popconfirm>
                    },
                },
            }}/>
    </PageContainer>
}

export default Algo