import {uploadSource} from "@/utils/oss";
import {InboxOutlined} from "@ant-design/icons";
import {PageContainer} from "@ant-design/pro-components"
import {request} from "@umijs/max";
import {Button, Card, Divider, Empty, message, Space, Tabs, TabsProps, Typography} from "antd"
import {UploadFile} from "antd/es/upload";
import Dragger from "antd/es/upload/Dragger";
import React, {useEffect, useRef, useState} from "react";

const {Text, Paragraph} = Typography

type LintMsg = {
    lang: string,
    out: string,
    err: string
}

type DisplayLintProps = {
    lint: LintMsg
}

const DisplayLint: React.FC<DisplayLintProps> = (props: DisplayLintProps) => {
    const {lint} = props
    const {out, err} = lint
    return <Space direction={"horizontal"} style={{width: '100%', alignItems: 'flex-start'}}>
        {
            out && <Card title={<Text type={'warning'}>Warning</Text>}>
                <Paragraph style={{whiteSpace: 'pre-wrap'}}>{out}</Paragraph>
            </Card>
        }
        {
            err && <Card title={<Text type={'danger'}>Error</Text>}>
                <Paragraph style={{whiteSpace: 'pre-wrap'}}>{err}</Paragraph>
            </Card>
        }
    </Space>
}

const Lint: React.FC = () => {
    const [tabs, setTabs] = useState<TabsProps["items"]>()
    const [fileList, setFileList] = useState<UploadFile[]>([])
    const endRef = useRef(null)

    useEffect(() => {
        //@ts-ignore
        endRef.current.scrollIntoView({behavior: 'smooth'});
    });

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
                langLint: {
                    lints: LintMsg[]
                }
            }) => {
                message.success('检查完毕')
                const msgs = res.langLint.lints
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
            <Typography.Text strong>解析结果：</Typography.Text>
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