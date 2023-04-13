import {getAllGraph, getGraph} from "@/services/graph/graph";

export default {

    state: {
        graphs: [1, 2],
    },

    effects: {
        * queryGraphs({}, {call, put}) {
            const data: Promise<Graph.SearchAllGraphReply> = yield call(getAllGraph);
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
        setup({dispatch, history}) {
            history.listen(({location}) => {
                if (location.pathname === '/home') {
                    dispatch({
                        type: 'queryGraphs'
                    })
                }
            });
        }
    },
};