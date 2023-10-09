import {useState} from 'react';


let init_graphs: Graph.Graph[] = []

let init_groups: Graph.Group[] = []

export function setInit(graphs: Graph.Graph[], groups: Graph.Group[]) {
    init_graphs = graphs
    init_groups = groups
}

// 将groups分割为graph和group
export function splitGroupsGraph(g: Graph.Group[]) {
    let graphs: Graph.Graph[] = []
    let groups: Graph.Group[] = []
    g.forEach(g => {
        graphs = [...graphs, ...g.graphs]
        groups.push(g)
    })
    graphs = graphs.sort((a, b) => b.id - a.id)
    groups = groups.sort((a, b) => a.name > b.name ? 0 : 1)
    return {graphs, groups}
}

export default () => {
    const [graphs, setGraphs] = useState<Graph.Graph[]>(init_graphs)
    const [groups, setGroups] = useState<Graph.Group[]>(init_groups)

    return {
        graphs,
        setGraphs,
        groups,
        setGroups
    }
}
