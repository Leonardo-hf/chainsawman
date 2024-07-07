import {DataNode} from "antd/es/tree"
import {UUID} from "@/utils/format"
import {Space, Typography} from "antd";
import {ReactNode} from "react";

const {Text} = Typography
export type OSV = {
    id: string,
    aliases?: string,
    summary: string,
    details: string,
    cwe?: string,
    severity?: string,
    ref?: string,
    affect?: string
}

export type Dep = {
    purl: string,
    osv?: OSV[],
    limit?: string
    indirect?: boolean
    exclude?: boolean
    optional?: boolean
    scope?: boolean
}

export type ModuleDep = {
    lang: string,
    path: string,
    purl: string,
    dependencies: Dep[]
}

export type DepExtractor = {
    value: string,
    label: string,
    // info: string,
    tree: (p: Dep) => DataNode,
    treeDep: (p: Dep[]) => DataNode[],
    graph: (p: Dep, d: Dep[]) => { nodes: any, edges: any }
}

export const getExtractor = (v: string) => {
    for (let e of extractors) {
        if (e.value == v) {
            return e
        }
    }
}


type tDecorate = (p: Dep) => ReactNode

const build: tDecorate = (p: Dep) => <Text>{p.purl}</Text>

const buildWithLimit = (decorate?: tDecorate) =>
    (p: Dep) => <Space>
        <Text>{p.limit}</Text>
        {decorate && decorate(p)}
    </Space>


const buildWithScope = (decorate?: tDecorate) =>
    (p: Dep) => <Space>
        {decorate && decorate(p)}
        <Text italic>{(p.optional || p.scope) && `<${p.optional ? 'optional' : p.scope}>`}</Text>
    </Space>

const buildWithComment = (decorate?: tDecorate) =>
    (p: Dep) => <Space>
        {decorate && decorate(p)}
        <Text italic>{p.indirect ? '//indirect' : p.exclude ? '//exclude' : ''}</Text>
    </Space>


const buildTree = (b: tDecorate) =>
    (p: Dep) => {
        return {
            title: b(p),
            purl: p.purl,
            key: UUID()
        }
    }

const buildTreeDep = (b: tDecorate) =>
    (p: Dep[]) => p.map(pi => buildTree(b)(pi))


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

type gnDecorate = (d: Dep) => {
    keyshape: {
        fill: string
    },
    label: {
        value: string
    }
}

type geDecorate = (d: Dep) => {
    label: {
        value: string,
        fill: string
    } | undefined,
    keyshape: {
        stroke: string
    } | undefined
}

const buildGN: gnDecorate = (d: Dep) => {
    return {
        keyshape: {
            fill: d.optional ? 'green' : 'yellow'
        },
        label: {
            value: d.purl
        }
    }
}

const buildGE: geDecorate = (d: Dep) => {
    const tag = d.limit ?? d.exclude ?? d.indirect
    if (typeof tag == 'string') {
        const color = getEdgeColor(tag)
        return {
            keyshape: {
                stroke: color
            },
            label: {
                value: tag,
                fill: color
            }
        }
    }
    return {label: undefined, keyshape: undefined}
}


const buildGraph = (node: gnDecorate, edge?: geDecorate) => (source: Dep, depends: Dep[]) => {
    const root = {id: source.purl, style: node(source)}
    const nodes = [root, ...depends.map(d => {
        return {
            id: d.purl,
            style: node(d)
        }
    })]
    const edges = depends.map(d => {
        return {
            source: source.purl,
            target: d.purl,
            style: edge && edge(d)
        }
    })
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
        tree: buildTree(build),
        treeDep: buildTreeDep(buildWithScope(build)),
        graph: buildGraph(buildGN),
    },
    {
        value: 'python',
        label: 'python',
        // info: '支持上传携带 setup.py, setup.cfg, pyproject.toml, requires.txt 文件的 tar/zip 压缩包',
        tree: buildTree(build),
        treeDep: buildTreeDep(buildWithLimit(build)),
        graph: buildGraph(buildGN, buildGE),
    },
    {
        value: 'go',
        label: 'go',
        // info: '支持上传 go.mod 管理的 tar/zip 压缩包',
        tree: buildTree(build),
        treeDep: buildTreeDep(buildWithComment(build)),
        graph: buildGraph(buildGN, buildGE),
    },
    {
        value: 'rust',
        label: 'rust',
        // info: '支持上传携带 Cargo.toml 文件的 tar/zip 压缩包',
        tree: buildTree(build),
        treeDep: buildTreeDep(buildWithScope(buildWithLimit(build))),
        graph: buildGraph(buildGN, buildGE),
    },
]