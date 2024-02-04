import {uploadSource} from "@/utils/oss";
import {InboxOutlined} from "@ant-design/icons";
import {PageContainer, ProList} from "@ant-design/pro-components"
import Graphin from "@antv/graphin";
import {request} from "@umijs/max";
import {Button, Checkbox, Divider, Empty, message, Space, Tabs, TabsProps, Tag, Tree, Typography} from "antd"
import {DataNode} from "antd/es/tree";
import {UploadFile} from "antd/es/upload";
import Dragger from "antd/es/upload/Dragger";
import React, {Key, memo, useEffect, useRef, useState} from "react";
import {DepExtractor, getExtractor, ModuleDep, Dep, OSV} from "./_extractor"
import {Pie} from "@ant-design/plots";


type DisplayProps = {
    module: ModuleDep,
    extractor: DepExtractor
}
const Display: React.FC<DisplayProps> = memo((props: DisplayProps) => {
    const {module, extractor} = props
    const visualOptions = [
        {label: '依赖树', value: 'tree'},
        {label: '依赖图', value: 'graph'},
    ]
    const [visual, setVisual] = useState(['tree'])
    const [treeData, setTreeData] = useState<DataNode[]>()
    const [graphData, setGraphData] = useState<any>()
    const [osvs, setOSVs] = useState<OSV[]>([])

    useEffect(() => {
        const getTreeData = (p: Dep, dependencies: Dep[]) => {
            let root = extractor.tree(p)
            root.children = extractor!.treeDep(dependencies)
            return [root]
        }

        const getGraphData = (p: ModuleDep, dependencies: Dep[]) => {
            return extractor!.graph({purl: p.purl}, dependencies)
        }
        setTreeData(getTreeData(module, module.dependencies))
        setGraphData(getGraphData(module, module.dependencies))
        setOSVs(module.dependencies.flatMap(d => (d.osv ?? []).map(o => {
            return {...o, affect: d.purl}
        })))
    }, [])

    // It's just a simple demo. You can use tree map to optimize update perf.
    const updateTreeData = (list: DataNode[], key: React.Key, children: DataNode[]): DataNode[] =>
        list.map((node) => {
            if (node.key === key) {
                return {
                    ...node,
                    children,
                }
            }
            if (node.children) {
                return {
                    ...node,
                    children: updateTreeData(node.children, key, children),
                }
            }
            return node
        })

    const onLoadData = ({key, purl, children}: any) => new Promise<void>(
        resolve => {
            if (children) {
                resolve()
                return
            }
            request('/api/util/search', {
                timeout: 30000,
                method: 'get',
                params: {
                    'package': purl,
                    'lang': extractor.value
                }
            }).then((res: {
                base: { status: number, msg: string },
                deps?: ModuleDep
            }) => {
                const module = res.deps
                if (!module) {
                    return
                }
                // 接受依赖解析结果
                setTreeData((origin: any) =>
                    updateTreeData(origin, key, extractor.treeDep(module.dependencies!))
                )
                const ng = extractor!.graph(module, module.dependencies)
                setGraphData((origin: any) => {
                    const unique = (arr: any[]) => {
                        const res = new Map()
                        return arr.filter((a) => !res.has(a.id) && res.set(a.id, 1))
                    }
                    return {
                        nodes: unique(origin.nodes.concat(ng.nodes)),
                        edges: [...origin.edges, ng.edges]
                    }
                })
            }).finally(() => resolve())
        })

    return <Space direction={"vertical"} style={{width: '100%'}}>
        <Checkbox.Group options={visualOptions} defaultValue={visual} onChange={(options) => {
            // @ts-ignore
            setVisual(options)
        }}/>
        {visual.length == 0 && <Empty/>}
        <div style={{width: '100%', display: 'flex'}}>
            {
                visual.includes('tree') &&
                <div style={{flexGrow: 1, flexBasis: '30%', flexShrink: 0}}>
                    <Tree
                        showLine={true}
                        showIcon={true}
                        loadData={onLoadData}
                        defaultExpandParent
                        treeData={treeData}/>
                </div>
            }
            {visual.includes('tree') && visual.includes('graph') &&
            <Divider type={'vertical'} style={{height: 'auto'}}/>}
            {
                visual.includes('graph') && <Graphin
                    style={{flexGrow: 1, flexBasis: '70%', flexShrink: 1}}
                    data={graphData}
                    layout={{type: 'circular'}}
                    fitView={true}/>
            }
        </div>
        {osvs.length > 0 && <VulDisplay osvs={osvs}/>}
    </Space>
})

const VulDisplay: React.FC<{ osvs: OSV[] }> = memo(({osvs}) => {
    const getSeverity = (severity: string | undefined) => {
        switch (severity) {
            case 'CRITICAL':
                return <Tag color={'magenta'}>极危</Tag>
            case 'HIGH':
                return <Tag color={'red'}>高危</Tag>
            case 'MEDIUM':
                return <Tag color={'magenta'}>中危</Tag>
            case 'LOW':
                return <Tag color={'magenta'}>低危</Tag>
            default:
                return
        }
    }
    return <ProList<OSV>
        style={{width: '100%'}}
        headerTitle={<Typography.Text strong type={'warning'}>漏洞列表</Typography.Text>}
        metas={{
            title: {
                render: (_, d) => <a href={d.ref}>{d.id}</a>
            },
            subTitle: {
                render: (_, d) => <Space>
                    {d.aliases && <Tag color={'purple'}>{d.aliases}</Tag>}
                    <Tag>{d.affect!}</Tag>
                    {d.cwe && <Tag color={'green'}>{d.cwe}</Tag>}
                </Space>
            },
            description: {
                dataIndex: 'summary',
            },
            content: {
                dataIndex: 'details'
            },
            actions: {
                render: (_, d) => getSeverity(d.severity)
            },
        }}
        dataSource={osvs}
    />
})

const Extractor: React.FC = () => {
    const [tabs, setTabs] = useState<TabsProps["items"]>()
    const [counts, setCounts] = useState<{ type: string, value: number }[]>()
    const [fileList, setFileList] = useState<UploadFile[]>([])
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
        const filename = fileList[0].fileName
        const fileId = await uploadSource(fileList[0].originFileObj!)
        request('/api/util/parse', {
            timeout: 20000,
            method: 'get',
            params: {
                'fileId': fileId,
                'filename': filename!
            }
        }).then((res: {
                base: { status: number, msg: string },
                packages?: { modules?: ModuleDep[] }
                counts: { type: string, value: number }[]
            }) => {
                message.success('解析成功')
                const modules = res.packages?.modules
                if (!modules) {
                    return
                }
                // 接受依赖解析结果
                setCounts(res.counts)
                setTabs(modules.map(m => {
                    return {
                        key: m.path,
                        label: m.path,
                        children: <Display extractor={getExtractor(m.lang)!} module={m}/>
                    }
                }))
            }
        ).finally(() => {
            // 清理上传框
            setFileList([])
        })
    }

    const getStatistic = (data: { type: string; value: number; }[]) => {
        const config = {
            data,
            angleField: 'value',
            colorField: 'type',
            label: {
                text: 'value',
                style: {
                    fontWeight: 'bold',
                },
            },
            legend: {
                color: {
                    title: false,
                    position: 'right',
                    rowPadding: 5,
                },
            },
        }
        return <Space direction={'vertical'} style={{width: '100%'}}>
            {data.length > 1 && <Pie {...config} />}
            <Typography.Paragraph strong>检测到{data.map(d => `${d.value}个${d.type}文件`).join(', ')}</Typography.Paragraph>
        </Space>
    }

    return <PageContainer>
        <Space direction="vertical" size="middle" style={{display: 'flex'}}>
            <Typography.Text strong>解析对象：</Typography.Text>
            <Dragger style={{width: '100%'}} maxCount={1} beforeUpload={() => false} fileList={fileList}
                     onChange={(e) => {
                         let newFileList = [...e.fileList];
                         newFileList = newFileList.slice(-1);
                         setFileList(newFileList)
                     }}>
                <p className="ant-upload-drag-icon">
                    <InboxOutlined/>
                </p>
                <p className="ant-upload-text">拖拽文件到这里上传</p>
                <p className="ant-upload-hint">支持上传使用 java, python, go, rust 编程语言的代码包</p>
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
                    {getStatistic(counts!)}
                    <Tabs items={tabs}/>
                </Space>
                : <Empty/>}
        </Space>
        <div ref={endRef}/>
    </PageContainer>
}

export default Extractor