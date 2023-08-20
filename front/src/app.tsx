import {RequestConfig} from "@@/plugin-request/request";
import React from "react";
import {RunTimeLayoutConfig} from "@umijs/max";
import Graph from "./pages/Graph"
import {getAllGraph} from "./services/graph/graph";

let graphs: any[] = []

export function render(oldRender: () => void) {
    getAllGraph().then(data => {
        if (data.graphs) {
            // setInitGraphs(data.graphs)
            graphs = data.graphs
            if (graphs != null) {
                data.graphs.forEach((graph) => graphs.push({
                    path: '/graph/' + graph.id,
                    element: <Graph graph={graph} key={graph.id}/>,
                    name: graph.name,
                }))
            }
        }
        oldRender()
    })
}


export function patchClientRoutes({routes}) {
    // TODO: 这个是直接根据路由的序号找的，扩展性差
    let menu = routes[0].children[2]
    // menu["routes"] = []
    menu["children"] = []
    graphs.forEach(graph => {
        // menu.routes.push(graph)
        menu.children.push(graph)
    })
    // console.log(routes)
}

// export async function getInitialState(): Promise<{ name: string }> {
//     return {name: '@umijs/max'};
// }

export const request: RequestConfig = {
    timeout: 10000,
}


export const layout: RunTimeLayoutConfig = () => {
    return {
        logo: require('@/assets/title.png'),
    }
}