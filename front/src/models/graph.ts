import {getAllGraph, getGraph} from "@/services/graph/graph";
import {useState} from "react";
import { request } from '@umijs/max';
import {history} from 'umi';

export default {

    state: {
        graphs: [1,2],
    },

    effects: {
        * queryGraphs({}, {call, put}) {
            const data:Promise<API.SearchAllGraphReply> = yield call(getAllGraph);
            yield put({type: 'getGraph', payload: data});
        },
    },

    reducers: {
        getGraph(state, {payload}) {
            return {
                ...state,
                graphs: payload
            }
        },
    },

    subscriptions: {
        setup ({dispatch, history}) {
            history.listen(({action,location}) => {
                if (location.pathname==='/home')
                    dispatch({
                        type:'queryGraphs'
                    })
            });
        }
    },

    test(state) {
        console.log('test');
        return state;
    },
};