import {algoGetDoc} from "@/services/graph/graph";

export default {
    state: {},
    effects: {
        * queryAlgoDoc({algoID}: { algoID: number }, {call, put}) {
            const res: Graph.GetAlgoDocReply = yield call(algoGetDoc, {algoId: algoID})
            yield put(
                {
                    type: 'add',
                    algoID: algoID,
                    doc: res.doc
                })
        },
    },

    reducers: {
        add(state, {algoID, doc}: { algoID: number, doc: string }) {
            state[algoID] = doc
            return {...state}
        },
    },
}