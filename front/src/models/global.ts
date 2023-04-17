import {useState} from 'react';


let init_graphs: Graph.Graph[] = []

export function setInitGraphs(graphs: Graph.Graph[]) {
    init_graphs = graphs
}

export default () => {
    const [graphs, setGraphs] = useState<Graph.Graph[]>(init_graphs)
    return {
        graphs,
        setGraphs
    }
}
