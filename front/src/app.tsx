import {RequestConfig} from "@@/plugin-request/request";
import React from "react";
import {RunTimeLayoutConfig} from "@umijs/max";
import Graph from "./pages/Graph"

import {algoGetAll, getAllGraph} from "./services/graph/graph";
import {setInit, splitGroupsGraph} from "@/models/global";

let newRoutes: any[] = []

// 查询组信息，生成路由
export function render(oldRender: () => void) {
    getAllGraph().then(data => {
        const gs = data.groups
        if (gs) {
            // 初始化图与组信息
            const {graphs, groups} = splitGroupsGraph(gs)
            setInit(graphs, groups)
            // 生成路由
            gs.forEach((g) => {
                newRoutes.push({
                    path: '/graph/' + g.id,
                    name: g.name,
                    children: g.graphs.map(graph => {
                        return {
                            path: '/graph/' + g.id + '/' + graph.id,
                            name: graph.name,
                            element: <Graph graph={graph} key={graph.id}/>,
                        }
                    })
                })
            })
        }
        oldRender()
    })
}


export function patchClientRoutes({routes}) {
    // TODO: 这个是直接根据路由的序号找的，扩展性差
    let menu = routes[0].children[2]
    // menu["routes"] = []
    menu["children"] = []
    newRoutes.forEach((r: any) => {
        // menu.routes.push(graph)
        menu.children.push(r)
    })
}

export async function getInitialState(): Promise<{ algos: Graph.Algo[] }> {
    const algosRes = await algoGetAll()
    return {
        algos: algosRes.algos
    }
}

export const request: RequestConfig = {
    timeout: 10 * 1000,
}


export const layout: RunTimeLayoutConfig = () => {
    return {
        logo: require('@/assets/title.png'),
    }
}