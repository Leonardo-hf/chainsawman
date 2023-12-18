import {useEffect, useState} from 'react';
import {RootGroupID} from "@/constants";
import {getAllGraph} from "@/services/graph/graph";


let init_graphs: GraphRef2Group[] = []

let init_groups: TreeNodeGroup[] = []

export type GraphRef2Group = { id: number, status: number, group: TreeNodeGroup, name: string, numNode: number, numEdge: number, creatAt: number, updateAt: number }

export type TreeNodeGroup = {
    id: number, name: string, desc: string, parentId: number,
    nodeTypeList: Graph.Structure[], edgeTypeList: Graph.Structure[], parentGroup: TreeNodeGroup | undefined
}

export function setInit(graphs: GraphRef2Group[], groups: TreeNodeGroup[]) {
    init_graphs = graphs
    init_groups = groups
}

// 将 groups 解析为前端使用的 graphs 和 groups
export function parseGroups(g: Graph.Group[]) {
    let graphs: GraphRef2Group[] = []
    let groups: TreeNodeGroup[] = []
    g.forEach(g => {
        const treeNodeGroup = {
            id: g.id,
            nodeTypeList: g.nodeTypeList,
            edgeTypeList: g.edgeTypeList,
            name: g.name,
            desc: g.desc,
            parentId: g.parentId,
            parentGroup: undefined
        }
        groups.push(treeNodeGroup)
        // 使 graph 引用 group
        graphs = [...graphs, ...g.graphs.map(g2 => {
            return {
                id: g2.id,
                status: g2.status,
                group: treeNodeGroup,
                name: g2.name,
                numNode: g2.numNode,
                numEdge: g2.numEdge,
                creatAt: g2.creatAt,
                updateAt: g2.updateAt
            }
        })]
    })
    // 寻找每个组的父策略组
    groups.filter(g => g.id !== RootGroupID).forEach(g => g.parentGroup = groups.find(g2 => g2.id === g.parentId))
    // 滤除根策略组，对结果排序
    groups = groups.filter(g => g.id !== RootGroupID).sort((a, b) => a.name > b.name ? 0 : 1)
    graphs = graphs.sort((a, b) => b.id - a.id)
    return {graphs, groups}
}

// 查看算法是否对于策略组合法
export function isAlgoIllegal(g: GraphRef2Group, a: Graph.Algo) {
    if (g.group.id == a.groupId) {
        return true
    }
    let parent = g.group.parentGroup
    while (parent) {
        if (parent.id == a.groupId) {
            return true
        }
        parent = parent.parentGroup
    }
    return false
}

// 生成策略组列表（Options）
export function genGroupOptions(gs: TreeNodeGroup[]) {
    return [{label: '通用', value: RootGroupID}, ...gs.map(g => {
        if (g.parentId > RootGroupID) {
            return {
                label: g.desc + ' -> ' + g.parentGroup!.desc!,
                value: g.id
            }
        }
        return {
            label: g.desc,
            value: g.id
        }
    })]
}


export default () => {
    const [graphs, setGraphs] = useState<GraphRef2Group[]>([])
    const [groups, setGroups] = useState<TreeNodeGroup[]>([])

    useEffect(() => {
        const update = () => {
            getAllGraph()
                .then(res => {
                    if (!res.groups) {
                        return
                    }
                    const {graphs, groups} = parseGroups(res.groups)
                    setGroups(groups)
                    setGraphs(graphs)
                    // 如果存在状态不为完成的图，则尝试再次请求
                    if (graphs.filter(a => a.status !== 1).length) {
                        setTimeout(_ => {
                            update()
                        }, 10000)
                    }
                    return
                })
        }
        update()
    }, [])
    return {
        graphs,
        groups,
    }
}
