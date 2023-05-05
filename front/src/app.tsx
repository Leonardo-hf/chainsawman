import {RequestConfig} from "@@/plugin-request/request";
import React from "react";
import {useModel} from "@umijs/max";
import Graph from "./pages/Graph"
import { setInitGraphs } from "./models/global";
import { getAllGraph } from "./services/graph/graph";
import { Utils } from "@antv/graphin";

let graphs: any


export function render(oldRender: () => void) {
    getAllGraph().then(data => {
        const routes: { path: string; element: JSX.Element; name: string; }[] = []
        if (data.graphs) {
            setInitGraphs(data.graphs)
            graphs = data.graphs
            data.graphs.forEach((graph) => routes.push({
                path: '/graph/' + graph.id,
                element: <Graph graph={graph} key={graph.id}/>,
                name: graph.name,
            }))
        }
        graphs = routes
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

export async function getInitialState(): Promise<{ name: string }> {
    return {name: '@umijs/max'};
}

//
// export function rootContainer(container) {
//     return React.createElement(RoutesProvider, null, container);
// }
//
// const RoutesProvider: React.FC = (props) => {
//     const clone = Object.assign({}, props.children)
//     console.log(clone.props)
//     const cloneProps = {...clone.props}
//     cloneProps.routes = []
//     clone.props = cloneProps
//     console.log(clone.props)
//     return <div>{clone}</div>
// }


export const request: RequestConfig = {
    timeout: 1000,
};

export const layout = () => {
    return {
        logo: require('@/assets/title.png'),
        menu: {
            locale: false,
            // request: () => {
            //     return []
            // }
        },
    };
};