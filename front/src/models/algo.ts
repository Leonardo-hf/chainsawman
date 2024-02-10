// import {getGraph, getNeighbors} from "@/services/graph/graph";
// import {getNRandomColor} from "@/utils/format";
// import {getGraphName} from "@/models/graph";
//
// export default {
//     state: [],
//     effects: {
//         // @ts-ignore
//         * queryAlgoDoc({payload}, {call, put}) {
//             const data: Graph.GetGraphDetailReply = yield call(getAlgoDoc, payload)
//             yield put(
//                 {
//                     type: 'add',
//                     payload: {
//                         gid: getGraphName(payload.graphId),
//                         timer: timer,
//                         group: group,
//                         ...data
//                     }
//                 })
//             return data.base.taskStatus
//         },
//
//         // @ts-ignore
//         * queryNeighbors({payload, timer, group}, {call, put}) {
//             const data: Graph.GetGraphDetailReply = yield call(getNeighbors, payload);
//             yield put(
//                 {
//                     type: 'add',
//                     payload: {
//                         gid: getGraphName(payload.graphId, payload.nodeId),
//                         timer: timer,
//                         group: group,
//                         ...data
//                     }
//                 })
//             return data.base.taskStatus
//         }
//     },
//
//     reducers: {
//         // @ts-ignore
//         clearSubGraph(state, {payload}) {
//             return state.filter((d: { gid: string; }) => !d.gid.startsWith(getGraphName(payload) + '/'))
//         },
//         // @ts-ignore
//         resetGraph(state, {payload}) {
//             // 删除graphID对应的图和子图数据
//             return state.filter((d: { gid: string; }) => !d.gid.startsWith(getGraphName(payload)))
//         },
//         // @ts-ignore
//         add(state, {payload}) {
//             const nodes: NodeData[] = [],
//                 edges: EdgeData[] = []
//             const nState = [...state]
//             // TODO: 过滤有不存在的节点的边
//             const nodeSet = new Set()
//             const colors = getNRandomColor(payload.nodePacks?.length)
//             payload.nodePacks?.forEach((np: Graph.NodePack) => {
//                 // TODO：display
//                 // 选取33%节点展示label
//                 const border = np.nodes.map(n => n.deg).sort()[Math.floor(np.nodes.length * 2 / 3)]
//                 const nodeType = payload.group.nodeTypeList.find((nt: Graph.Structure) => nt.name == np.tag)!
//                 const display = nodeType.display
//                 const labelAttr = nodeType.primary
//                 // 为每一类型的节点设置一种颜色
//                 const color = colors.pop()
//                 np.nodes.forEach(n => {
//                     let label = ''
//                     if (labelAttr && n.deg >= border) {
//                         label = n.attrs.find(a => a.key === labelAttr)!.value
//                     }
//                     nodeSet.add(n.id)
//                     nodes.push({
//                         id: n.id,
//                         tag: np.tag,
//                         attrs: n.attrs,
//                         deg: n.deg,
//                         type: 'graphin-circle',
//                         style: {
//                             tag: np.tag,
//                             id: n.id.toString(),
//                             style: {
//                                 keyshape: {
//                                     size: Math.floor((Math.log(n.deg + 1) + 2) * 10),
//                                     fill: color,
//                                 },
//                                 label: {
//                                     value: label
//                                 }
//                             }
//                         }
//                     })
//                 })
//             })
//             payload.edgePacks?.forEach((ep: Graph.EdgePack) => {
//                 // TODO: 过滤有不存在的节点的边
//                 const edgeType = payload.group.edgeTypeList.find((et: Graph.Structure) => et.name === ep.tag)!
//                 const display = edgeType.display
//                 const labelAttr = edgeType.primary
//                 ep.edges.forEach(e => {
//                     if (!nodeSet.has(e.source) || !nodeSet.has(e.target)) {
//                         return
//                     }
//                     let label = ''
//                     if (labelAttr) {
//                         label = e.attrs.find(a => a.key === labelAttr)!.value
//                     }
//                     const edgeStyle: any = {
//                         source: e.source.toString(),
//                         target: e.target.toString(),
//                         style: {
//                             label: {
//                                 value: label
//                             }
//                         }
//                     }
//                     if (display === 'dash') {
//                         edgeStyle.style = {
//                             ...edgeStyle.style,
//                             keyshape: {
//                                 lineDash: [4, 4],
//                             }
//                         }
//                     }
//                     edges.push({
//                         source: e.source,
//                         target: e.target,
//                         tag: ep.tag,
//                         attrs: e.attrs,
//                         style: edgeStyle
//                     })
//                 })
//             })
//             const i = nState.findIndex(s => s.gid == payload.gid)
//             const v = {
//                 gid: payload.gid,
//                 taskId: payload.base.taskId,
//                 status: payload.base.taskStatus,
//                 nodes: nodes,
//                 edges: edges,
//                 timer: payload.timer
//             }
//             if (i === -1) {
//                 nState.push(v)
//             } else {
//                 nState[i] = v
//             }
//             // 新增图或子图数据
//             return nState
//         },
//     },
// }