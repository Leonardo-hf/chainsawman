import {getGraph, getNeighbors} from "@/services/graph/graph";


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
        * queryGraph({payload, timer}, {call, put}) {
            const data: Graph.GetGraphDetailReply = yield call(getGraph, payload)
            yield put(
                {
                    type: 'add',
                    payload: {
                        gid: getGraphName(payload.graphId),
                        timer: timer,
                        ...data
                    }
                })
            return data.base.taskStatus
        },

        // @ts-ignore
        * queryNeighbors({payload, timer}, {call, put}) {
            const data: Graph.GetGraphDetailReply = yield call(getNeighbors, payload);
            yield put(
                {
                    type: 'add',
                    payload: {
                        gid: getGraphName(payload.graphId, payload.nodeId),
                        timer: timer,
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
            const nodes: { id: number; tag: string; attrs: Graph.Pair[]; deg: number }[] = [],
                edges: { source: number; target: number; tag: string; attrs: Graph.Pair[]; }[] = []
            const nState = [...state]
            payload.nodePacks?.forEach((np: Graph.NodePack) => {
                np.nodes.forEach(n => {
                    nodes.push({
                        id: n.id,
                        tag: np.tag,
                        attrs: n.attrs,
                        deg: n.deg
                    })
                })
            })
            payload.edgePacks?.forEach((np: Graph.EdgePack) => {
                np.edges.forEach(n => {
                    edges.push({
                        source: n.source,
                        target: n.target,
                        tag: np.tag,
                        attrs: n.attrs
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
