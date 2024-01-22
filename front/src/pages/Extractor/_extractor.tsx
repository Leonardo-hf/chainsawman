import {DataNode} from "antd/es/tree"
import {UUID} from "@/utils/format"

export type DepExtractor = {
    value: string,
    label: string,
    // info: string,
    tree: (p: any) => DataNode,
    treeDep: (p: any[]) => DataNode[],
    graph: (p: any, d: any[]) => { nodes: any, edges: any }
}

export const getExtractor = (v: string) => {
    for (let e of extractors) {
        if (e.value == v) {
            return e
        }
    }
}
const build = (p: any) => p.artifact + ':' + p.version
const buildWithGroup = (p: any) => p.group + '/' + build(p)
const buildWithLimit = (p: any) => p.limit + ' ' + build(p)
const buildWithComment = (p: any) => {
    return build(p) + (p.indirect ? '  //indirect' : p.exclude ? '  //exclude' : '')
}
const buildTree = (b: any) =>
    (p: any) => {
        return {
            title: b(p),
            origin: p,
            key: UUID()
        }
    }

const buildTreeDep = (b: any) =>
    (p: any[]) => p.map(pi => {
        return {
            title: b(pi),
            origin: pi,
            key: UUID()
        }
    })


const getEdgeColor = (tag: string) => {
    switch (tag) {
        case '>':
            return '#104E8B'
        case '>=':
            return '#1874CD'
        case '<':
            return '#228B22'
        case '<=':
            return '#6B8E23'
        case '~':
            return '#8B668B'
        case '^':
            return '#7A378B'
        case 'indirect':
            return '#9C9C9C'
        case '!=':
        case 'exclude':
            return '#B22222'
    }
    return '#363636'
}

const buildGraph = (b: any) => buildGraphWithEdgeTag(b, p => '')

const buildGraphWithEdgeTag = (b: any, getTag: (p: any) => string) => (p: any, d: any[]) => {
    let nodes = [], edges = []
    const v = b(p)
    const root = {id: v, style: {label: {value: v}}}
    nodes.push(root)
    for (let di of d) {
        const vi = b(di)
        const target = {
            id: vi,
            style: {
                keyshape: {fill: di.optional ? 'yellow' : 'green'},
                label: {value: vi}
            }
        }
        nodes.push(target)
        const tag = getTag(di)
        edges.push({
            source: root.id,
            target: target.id,
            style: {
                label: {
                    value: tag,
                    fill: getEdgeColor(tag),
                }
            }
        })
    }
    return {
        nodes: nodes,
        edges: edges,
    }
}

export const extractors: DepExtractor[] = [
    {
        value: 'java',
        label: 'java',
        // info: '支持上传 pom.xml 或携带该文件的 tar/zip 压缩包',
        tree: buildTree(buildWithGroup),
        treeDep: buildTreeDep(buildWithGroup),
        graph: buildGraph(buildWithGroup),
    },
    {
        value: 'python',
        label: 'python',
        // info: '支持上传携带 setup.py, setup.cfg, pyproject.toml, requires.txt 文件的 tar/zip 压缩包',
        tree: buildTree(build),
        treeDep: buildTreeDep(buildWithLimit),
        graph: buildGraphWithEdgeTag(build, p => p.limit),
    },
    {
        value: 'go',
        label: 'go',
        // info: '支持上传 go.mod 管理的 tar/zip 压缩包',
        tree: buildTree(build),
        treeDep: buildTreeDep(buildWithComment),
        graph: buildGraphWithEdgeTag(build, p => p.indirect ? 'indirect' : p.exclude ? 'exclude' : ''),
    },
    {
        value: 'rust',
        label: 'rust',
        // info: '支持上传携带 Cargo.toml 文件的 tar/zip 压缩包',
        tree: buildTree(build),
        treeDep: buildTreeDep(buildWithLimit),
        graph: buildGraphWithEdgeTag(build, p => p.limit),
    },
]