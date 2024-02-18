import {uploadSource} from "@/utils/oss";
import {InboxOutlined} from "@ant-design/icons";
import {PageContainer, ProList} from "@ant-design/pro-components"
import {request} from "@umijs/max";
import {Button, Card, Divider, Empty, message, Space, Tabs, TabsProps, Tag, Typography} from "antd"
import {UploadFile} from "antd/es/upload";
import Dragger from "antd/es/upload/Dragger";
import React, {memo, useEffect, useRef, useState} from "react";
import {exportPlain} from "@/utils/format";

const {Text, Paragraph} = Typography

type Lint = {
    path: string
    pos: string
    msg: string
    lint: string
}

type LangLint = {
    lang: string,
    lints: Lint[],
    err: string
}

type DisplayLintProps = {
    lint: LangLint
}

const DisplayLint: React.FC<DisplayLintProps> = memo((props: DisplayLintProps) => {
    const {lint} = props
    const {lints, err} = lint
    const [lintsEnum, setLintsEnum] = useState({})
    useEffect(() => {
        setLintsEnum(lints.reduce((a: any, b) => {
            if (!a[b.lint]) {
                a[b.lint] = {
                    text: b.lint,
                    status: b.lint,
                }
            }
            return a
        }, {}))
    }, [])
    return <div style={{display: 'flex', flexDirection: 'column', width: '100%', alignItems: 'flex-start'}}>
        {
            err && <Card style={{width: '100%', marginBottom: '16px'}} title={<Text type={'danger'}>Error</Text>}>
                <Paragraph style={{whiteSpace: 'pre-wrap'}}>{err}</Paragraph>
            </Card>
        }
        {
            lints.length > 0 && <ProList<Lint>
                style={{width: '100%'}}
                headerTitle={<Text type={'warning'}>Warning</Text>}
                pagination={{
                    defaultPageSize: 10,
                    showSizeChanger: true,
                }}
                search={{filterType: "light"}}
                request={async ({actions}) => {
                    const res = lints.filter(l => !actions || l.lint == actions)
                    return {
                        data: res,
                        success: true,
                        total: res.length,
                    }
                }}
                metas={{
                    title: {
                        dataIndex: 'path',
                        search: false
                    },
                    subTitle: {
                        render: (_, d) => <Tag color={'geekblue'}>{d.pos}</Tag>,
                        search: false
                    },
                    description: {
                        dataIndex: 'msg',
                        search: false
                    },
                    actions: {
                        title: <Text strong>LINT规则</Text>,
                        valueType: 'select',
                        render: (_, d) => <Tag color={'orange'}>{d.lint}</Tag>,
                        valueEnum: lintsEnum
                    },
                }}
            />
        }
    </div>
})

const Lint: React.FC = () => {
    const [tabs, setTabs] = useState<TabsProps["items"]>()
    const [fileList, setFileList] = useState<UploadFile[]>([])
    const [resJSON, saveRes] = useState<string>('')
    const endRef = useRef(null)

    useEffect(() => {
        //@ts-ignore
        endRef.current.scrollIntoView({behavior: 'smooth'})
    }, [tabs])

    const commit = async () => {
        if (fileList.length === 0) {
            message.error('必须上传一个包文件')
            return
        }
        const fileId = await uploadSource(fileList[0].originFileObj!)
        request('/api/util/lint', {
            timeout: 20000,
            method: 'get',
            params: {
                'fileId': fileId,
            }
        }).then((res: {
                base: { status: number, msg: string },
                langLints: LangLint[]
            }) => {
                message.success('检查完毕')
                const msgs = res.langLints
                saveRes(JSON.stringify(res.langLints))
                setTabs(msgs.map((m, i) => {
                    return {
                        key: m.lang,
                        label: m.lang,
                        children: <DisplayLint lint={m}/>
                    }
                }))
            }
        ).finally(() => {
            // 清理上传框
            setFileList([])
        })
    }


    return <PageContainer>
        <Space direction="vertical" size="middle" style={{display: 'flex'}}>
            <Typography.Text strong>解析对象：</Typography.Text>
            <Dragger style={{width: '100%'}} maxCount={1} beforeUpload={(file) => false} fileList={fileList}
                     onChange={(e) => {
                         let newFileList = [...e.fileList];
                         newFileList = newFileList.slice(-1);
                         setFileList(newFileList)
                     }}>
                <p className="ant-upload-drag-icon">
                    <InboxOutlined/>
                </p>
                <p className="ant-upload-text">拖拽文件到这里上传</p>
                <p className="ant-upload-hint">支持上传使用 java, python, go 编程语言的代码包</p>
            </Dragger>
            <Space size="middle">
                <Button type="primary" onClick={commit}>解析</Button>
                <Button onClick={() => {
                    setTabs(undefined)
                    setFileList([])
                }}>清除</Button>
            </Space>
            <Divider/>
            <div style={{display: 'inline-flex', justifyContent: 'space-between', width: '100%'}}>
                <Typography.Text strong>解析结果：</Typography.Text>
                {tabs && tabs.length && <Button
                    onClick={() => exportPlain(`lint-${new Date().toISOString()}.json`, resJSON)}>导出</Button>}
            </div>
            {tabs && tabs.length ?
                <Space style={{width: '100%'}} direction={'vertical'}>
                    <Tabs items={tabs}/>
                </Space>
                : <Empty/>}
        </Space>
        <div ref={endRef}/>
    </PageContainer>
}

export default Lint