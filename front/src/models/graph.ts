import {getAllGraph, getGraph} from "@/services/graph/graph";
import {useState} from "react";
import { request } from '@umijs/max';
import {history} from 'umi';


const getRandomColor = function () {
    const text = '00000' + (Math.random() * 0x1000000 << 0).toString(16)
    return '#' + text.substring(text.length - 6);
}

export default {
    state: {

        graphs: [1, 2],
        details: {}
    },

    effects: {
        * queryGraphs({}, {call, put}) {
            const data: Promise<Graph.SearchAllGraphReply> = yield call(getAllGraph);
            yield put({type: 'getGraphs', payload: data});
        },

        * queryDetail({payload}, {call, put}) {
            const data: Promise<Graph.SearchGraphDetailReply> = yield call(getGraph, payload);
            data['graph'] = payload.graph;
            yield put({type: 'getDetail', payload: data});
            return data.base.taskStatus
        },
    },

    reducers: {
        getGraphs(state, {payload}) {
            return {
                ...state,
                graphs: payload
            }
        },
        getDetail(state, {payload}) {
            const details = {}
            for (let attr in state.details) {
                details[attr] = state.details[attr]
            }
            details[payload.graph] = {
                taskId: payload.base.taskId,
                status: payload.base.taskStatus,
                nodes: payload.nodes,
                edges: payload.edges,
            }
            for (let i = 0; i < details[payload.graph].nodes.length; i++) {
                details[payload.graph].nodes[i]['color'] = getRandomColor()
            }
            return {
                ...state,
                details
            }
        },
    },


    // subscriptions: {
    //     setup({dispatch, history}) {
    //         history.listen(({location}) => {
    //             if (location.pathname === '/home') {
    //                 dispatch({
    //                     type: 'queryGraphs'
    //                 })
    //             }
    //         });
    //     }
    // },
    //
}
