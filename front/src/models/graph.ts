import {getAllGraph, getGraph} from "@/services/graph/graph";

export default {
    state: {
        graphs: [],
    },

    effects: {
        * queryGraphs({}, {call, put}) {
            const {data} = yield call(getAllGraph);
            yield put({type: 'queryGraphsSuccess', payload: data.graphs});
        },
    },

    reducers: {
        queryGraphsSuccess(state, {payload}) {
            return {
                ...state,
                graphs: payload
            }
        },
    },
};