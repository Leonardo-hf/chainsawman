import {getGraph, getNeighbors} from "@/services/graph/graph";
import {getRandomColor} from "@/utils/format";
import {sum} from "@antfu/utils";


export const getGraphName = function (graphID: number, tgtNodeID: number | undefined = undefined) {
    if (tgtNodeID) {
        return 'Graph@' + graphID.toString() + '/' + tgtNodeID.toString()
    }
    return 'Graph@' + graphID.toString()
}

export const isSubGraph = function (gid: string) {
    return gid.indexOf('/') !== -1;
}

export const findByGid = function (state: any[], gid: string) {
    return state.find(s => s.gid === gid)
}

export const filterByGraph = function (state: any[], graphId: number) {
    return state.filter(s => s.gid.startsWith(getGraphName(graphId)))
}


// @ts-ignore
export default {
    state: [],
    effects: {
        // @ts-ignore
        * queryGraph({payload, timer, group}, {call, put}) {
            const data: Graph.GetGraphDetailReply = yield call(getGraph, payload)
            yield put(
                {
                    type: 'add',
                    payload: {
                        gid: getGraphName(payload.graphId),
                        timer: timer,
                        group: group,
                        ...data
                    }
                })
            return data.base.taskStatus
        },

        // @ts-ignore
        * queryNeighbors({payload, timer, group}, {call, put}) {
            const data: Graph.GetGraphDetailReply = yield call(getNeighbors, payload);
            yield put(
                {
                    type: 'add',
                    payload: {
                        gid: getGraphName(payload.graphId, payload.nodeId),
                        timer: timer,
                        group: group,
                        ...data
                    }
                })
            return data.base.taskStatus
        }
    },

    reducers: {
        // @ts-ignore
        clearSubGraph(state, {payload}) {
            return state.filter((d: { gid: string; }) => !d.gid.startsWith(getGraphName(payload) + '/'))
        },
        // @ts-ignore
        resetGraph(state, {payload}) {
            // 删除graphID对应的图和子图数据
            return state.filter((d: { gid: string; }) => !d.gid.startsWith(getGraphName(payload)))
        },
        // @ts-ignore
        add(state, {payload}) {
            const nodes: NodeData[] = [],
                edges: EdgeData[] = []
            const nState = [...state]
            // TODO: 过滤有不存在的节点的边
            const nodeSet = new Set()
            payload.nodePacks?.forEach((np: Graph.NodePack) => {
                // TODO：display
                // 选取33%节点展示label
                const border = np.nodes.map(n => n.deg).sort()[Math.floor(np.nodes.length * 2 / 3)]
                const nodeType = payload.group.nodeTypeList.find((nt: Graph.Structure) => nt.name == np.tag)!
                const display = nodeType.display
                const labelAttr = nodeType.attrs?.find((a: Graph.Attr) => a.primary)
                // 为每一类型的节点设置一种颜色
                const color = getRandomColor()
                np.nodes.forEach(n => {
                    let label = ''
                    if (labelAttr && n.deg >= border) {
                        label = n.attrs.find(a => a.key === labelAttr.name)!.value
                        // if (label.length >= 10) {
                        //     label = label.substring(0, 10) + '...'
                        // }
                    }
                    nodeSet.add(n.id)
                    nodes.push({
                        id: n.id,
                        tag: np.tag,
                        attrs: n.attrs,
                        deg: n.deg,
                        type: 'graphin-circle',
                        style: {
                            tag: np.tag,
                            id: n.id.toString(),
                            style: {
                                keyshape: {
                                    size: Math.floor((Math.log(n.deg + 1) + 1) * 10),
                                    fill: color
                                },
                                label: {
                                    value: label
                                }
                            }
                        }
                    })
                })
            })
            payload.edgePacks?.forEach((ep: Graph.EdgePack) => {
                // TODO: 过滤有不存在的节点的边
                const edgeType = payload.group.edgeTypeList.find((et: Graph.Structure) => et.name === ep.tag)!
                const display = edgeType.display
                const labelAttr = edgeType.attrs?.find((a: Graph.Attr) => a.primary)
                ep.edges.forEach(e => {
                    if (!nodeSet.has(e.source) || !nodeSet.has(e.target)) {
                        return
                    }
                    let label = ''
                    if (labelAttr) {
                        label = e.attrs.find(a => a.key === labelAttr.name)!.value
                        // if (label.length >= 10) {
                        //     label = label.substring(0, 10) + '...'
                        // }
                    }
                    const edgeStyle: any = {
                        source: e.source.toString(),
                        target: e.target.toString(),
                        style: {
                            label: {
                                value: label
                            }
                        }
                    }
                    if (display === 'dash') {
                        edgeStyle.style = {
                            ...edgeStyle.style,
                            keyshape: {
                                lineDash: [4, 4],
                            }
                        }
                    }
                    edges.push({
                        source: e.source,
                        target: e.target,
                        tag: ep.tag,
                        attrs: e.attrs,
                        style: edgeStyle
                    })
                })
            })
            const i = nState.findIndex(s => s.gid == payload.gid)
            const v = {
                gid: payload.gid,
                taskId: payload.base.taskId,
                status: payload.base.taskStatus,
                nodes: nodes,
                edges: edges,
                timer: payload.timer
            }
            if (i === -1) {
                nState.push(v)
            } else {
                nState[i] = v
            }
            // 新增图或子图数据
            return nState
        },
    },
}

export type NodeData = { id: number; tag: string; attrs: Graph.Pair[]; deg: number; style: any; type: string }
export type EdgeData = { source: number; target: number; tag: string; attrs: Graph.Pair[]; style: any }