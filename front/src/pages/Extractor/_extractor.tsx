import {DataNode} from "antd/es/tree"
import {UUID} from "@/utils/format"

export type Extractor = {
    value: string,
    label: string,
    info: string,
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

const buildJava = (p: any) => p.group + '/' + p.artifact + ':' + p.version

const buildPython = (p: any) => p.artifact + ':' + p.version
export const extractors: Extractor[] = [
    {
        value: 'java',
        label: 'java',
        info: '支持上传 *pom.xml 或携带该文件的 tar/zip 压缩包',
        tree: (p: any) => {
            return {
                title: buildJava(p),
                key: UUID(),
            }
        },
        treeDep: (p: any[]) => {
            const res = []
            for (let pi of p) {
                res.push({
                    title: buildJava(pi),
                    key: UUID(),
                })
            }
            return res
        },
        graph: (p: any, d: any[]) => {
            let nodes = [], edges = []
            const v = buildJava(p)
            const root = {
                id: v
                , style: {label: {value: v}}
            }
            nodes.push(root)
            for (let di of d) {
                const vi = buildJava(di)
                const target = {
                    id: vi,
                    style: {
                        keyshape: {fill: di.optional ? 'yellow' : 'green'},
                        label: {value: vi}
                    }
                }
                nodes.push(target)
                edges.push({
                    source: root.id,
                    target: target.id,
                })
            }
            return {
                nodes: nodes,
                edges: edges
            }
        },
    },
    {
        value: 'python',
        label: 'python',
        info: '支持上传 setup.py, setup.cfg, pyproject.toml, requires.txt 或携带以上文件的 tar/zip 压缩包',
        tree: (p: any) => {
            return {
                title: buildPython(p),
                key: UUID(),
            }
        },
        treeDep: (p: any[]) => {
            const res = []
            for (let pi of p) {
                res.push({
                    title: pi.limit + ' ' + buildPython(pi),
                    key: UUID(),
                })
            }
            return res
        },
        graph: (p: any, d: any[]) => {
            let nodes = [], edges = []
            const v = buildPython(p)
            const root = {id: v, style: {label: {value: v}}}
            nodes.push(root)
            for (let di of d) {
                const vi = buildPython(di)
                const target = {
                    id: vi,
                    style: {label: {value: vi}}
                }
                nodes.push(target)
                edges.push({
                    source: root.id,
                    target: target.id,
                })
            }
            return {
                nodes: nodes,
                edges: edges,
            }
        },
    },
]