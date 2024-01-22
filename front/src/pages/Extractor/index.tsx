import {uploadSource} from "@/utils/oss";
import {InboxOutlined} from "@ant-design/icons";
import {PageContainer} from "@ant-design/pro-components"
import Graphin from "@antv/graphin";
import {request} from "@umijs/max";
import {Button, Checkbox, Divider, Empty, message, Select, Space, Tabs, TabsProps, Tree, Typography} from "antd"
import {DataNode} from "antd/es/tree";
import {UploadFile} from "antd/es/upload";
import Dragger from "antd/es/upload/Dragger";
import React, {useEffect, useRef, useState} from "react";
import {DepExtractor, getExtractor} from "./_extractor"
import {Pie} from "@ant-design/plots";

type ModuleDep = {
    lang: string,
    path: string,
    group: string,
    artifact: string,
    version: string,
    dependencies: {
        artifact: string
        version: string
        group: string
        limit: string
        indirect: boolean
        exclude: boolean
        optional: boolean
        scope: boolean
    }[]
}
type DisplayProps = {
    module: ModuleDep,
    extractor: DepExtractor
}
const Display: React.FC<DisplayProps> = (props: DisplayProps) => {
    const {module, extractor} = props

    const getTreeData = (p: any, dependencies: any[]) => {
        let root = extractor.tree(p)
        root.children = extractor!.treeDep(dependencies)
        return [root]
    }

    const getGraphData = (p: any, dependencies: any[]) => {
        return extractor!.graph(p, dependencies)
    }
    const visualOptions = [
        {label: '依赖树', value: 'tree'},
        {label: '依赖图', value: 'graph'},
    ];

    const [visual, setVisual] = useState(['tree'])
    const [treeData, setTreeData] = useState<DataNode[]>(getTreeData(module, module.dependencies))
    const [graphData, setGraphData] = useState<any>(getGraphData(module, module.dependencies))

    // It's just a simple demo. You can use tree map to optimize update perf.
    const updateTreeData = (list: DataNode[], key: React.Key, children: DataNode[]): DataNode[] =>
        list.map((node) => {
            if (node.key === key) {
                return {
                    ...node,
                    children,
                };
            }
            if (node.children) {
                return {
                    ...node,
                    children: updateTreeData(node.children, key, children),
                };
            }
            return node;
        })

    const onLoadData = ({key, origin, children}: any) => new Promise<void>(
        resolve => {
            if (children) {
                resolve()
                return
            }
            request('/api/util/search', {
                timeout: 20000,
                method: 'get',
                params: {
                    'package': extractor.tree(origin).title,
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
                        edges: origin.edges.concat(ng.edges)
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
        {
            visual.includes('tree') && <Tree
                showLine={true}
                showIcon={true}
                loadData={onLoadData}
                defaultExpandParent
                treeData={treeData}/>
        }
        {
            visual.includes('graph') && <Graphin
                data={graphData}
                layout={{type: 'circular'}}
                fitView={true}/>
        }
    </Space>
}

const Extractor: React.FC = () => {
    const [tabs, setTabs] = useState<TabsProps["items"]>()
    const [counts, setCounts] = useState<{ type: string, value: number }[]>()
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

                // 清理上传框
                setFileList([])
            }
        )
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
        return <Pie {...config} />
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
                <p className="ant-upload-hint">支持上传使用 java, python, go, rust 编程语言的代码包</p>
            </Dragger>
            <Space size="middle">
                <Button type="primary" onClick={commit}>解析</Button>
                <Button onClick={() => {
                    setTabs(undefined)
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