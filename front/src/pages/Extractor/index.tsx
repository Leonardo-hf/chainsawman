import {upload} from "@/utils/oss";
import {InboxOutlined} from "@ant-design/icons";
import {PageContainer} from "@ant-design/pro-components"
import Graphin from "@antv/graphin";
import {request} from "@umijs/max";
import {Button, Checkbox, Divider, Empty, message, Select, Space, Tree, Typography} from "antd"
import {DataNode} from "antd/es/tree";
import {UploadFile} from "antd/es/upload";
import Dragger from "antd/es/upload/Dragger";
import {useState} from "react";
import {extractors, getExtractor} from "./_extractor";

const Extractor: React.FC = (props) => {
    const [language, setLanguage] = useState('java')
    const [visual, setVisual] = useState(['tree'])
    const [fileList, setFileList] = useState<UploadFile[]>([])
    const [p, setPackage] = useState()
    const [dependencies, setDependencies] = useState([])
    const [treeData, setTreeData] = useState<DataNode[]>([])
    const [graphData, setGraphData] = useState<any>([])
    const visualOptions = [
        {label: '图形', value: 'graph'},
        {label: '树', value: 'tree'},
    ];
    let extractor = getExtractor(language)
    const commit = async () => {
        if (fileList.length === 0) {
            message.error('必须上传一个包文件')
            return
        }
        const fileId = await upload(fileList[0].originFileObj!)
        request('/api/util/parse', {
            timeout: 20000,
            method: 'get',
            params: {
                'fileId': fileId,
                'lang': extractor!.value
            }
        }).then(res => {
            console.log(res)
            if (res.base.status >= 4000) {
                message.error(res.base.msg)
                return
            } else if (res.base.status >= 3000) {
                message.warning(res.base.msg)
            }
            message.success('解析成功')
            // 接受依赖解析结果
            setPackage(res.package)
            setDependencies(res.dependencies)
            setTreeData(getTreeData(res.package, res.dependencies))
            setGraphData(getGraphData(res.package, res.dependencies))
            // 清理上传框
            setFileList([])
        })
    }
    const isTreeVisable = () => {
        return p !== undefined && visual.includes('tree')
    }
    const getTreeData = (p: any, dependencies: any[]) => {
        let root = extractor!.tree(p)
        root.children = extractor!.treeDep(dependencies)
        return [root]
    }
    const isGraphVisable = () => {
        return p !== undefined && visual.includes('graph')
    }
    const getGraphData = (p: any, dependencies: any[]) => {
        return extractor!.graph(p, dependencies)
    }

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

    const onLoadData = ({key, title, children}: any) => new Promise<void>(
        resolve => {
            if (children) {
                resolve()
                return
            }
            request('/api/util/search', {
                timeout: 20000,
                method: 'get',
                params: {
                    'package': title,
                    'lang': extractor!.value
                }
            }).then(res => {
                console.log(res)
                if (res.base.status >= 4000) {
                    message.error(res.base.msg)
                    resolve()
                    return
                } else if (res.base.status >= 3000) {
                    message.warning(res.base.msg)
                }
                // 接受依赖解析结果
                setTreeData((origin) =>
                    updateTreeData(origin, key, extractor!.treeDep(res.dependencies)),
                )
                const ng = extractor!.graph(res.package, res.dependencies)
                setGraphData((origin) => {
                    const unique = (arr: any[]) => {
                        const res = new Map()
                        return arr.filter((a) => !res.has(a.id) && res.set(a.id, 1))
                    }
                    const a = {
                        nodes: unique(origin.nodes.concat(ng.nodes)),
                        edges: origin.edges.concat(ng.edges)
                    }
                    console.log(a)
                    return a
                })
                resolve()
            })
        })
    return <PageContainer>
        <Space direction="vertical" size="middle" style={{display: 'flex'}}>
            <div>
                <Typography.Text strong>解析对象：</Typography.Text>
                <Select
                    showSearch
                    placeholder="选择一种语言"
                    optionFilterProp="children"
                    onChange={(v) => {
                        setLanguage(v)
                        setPackage(undefined)
                    }}
                    filterOption={(input, option) =>
                        (option?.label ?? '').toLowerCase().includes(input.toLowerCase())
                    }
                    defaultValue={language}
                    style={{width: 100}}
                    options={extractors}/>
            </div>
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
                <p className="ant-upload-hint">{extractor!.info}</p>
            </Dragger>
            <Space size="middle">
                <Button type="primary" onClick={commit}>解析</Button>
                <Button onClick={() => {
                    setPackage(undefined)
                }}>清除</Button>
            </Space>
            <Divider/>
            <Checkbox.Group options={visualOptions} defaultValue={visual} onChange={(options) => {
                setVisual(options)
            }}/>
            {!isTreeVisable() && !isGraphVisable() && <Empty/>}
            {
                isTreeVisable() && <Tree
                    showLine={true}
                    showIcon={true}
                    loadData={onLoadData}
                    treeData={treeData}/>
            }
            {
                isGraphVisable() && <Graphin
                    data={graphData}
                    layout={{type: 'circular'}}
                    fitView={true}/>
            }
        </Space>
    </PageContainer>
}

export default Extractor