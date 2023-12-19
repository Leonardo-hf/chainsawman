import {RequestConfig} from "@@/plugin-request/request";
import React from "react";
import {RunTimeLayoutConfig} from "@umijs/max";
import Graph from "./pages/Graph"

import {algoGetAll, getAllGraph, getHHI, getHot} from "./services/graph/graph";
import {parseGroups} from "@/models/global";
import {message} from "antd";

let newRoutes: any[] = []

// 查询组信息，生成路由
export function render(oldRender: () => void) {
    getAllGraph().then(data => {
        const gs = data.groups
        if (gs) {
            // 初始化图与组信息
            const {graphs, groups} = parseGroups(gs)
            // 生成路由
            const rootRoute = {
                path: '/graph',
                children: []
            }
            const routeMap: any = {1: rootRoute}
            const groupsQueue = [1]
            while (groupsQueue.length > 0) {
                const parentId = groupsQueue.shift()!
                const parentRoute = routeMap[parentId]
                groups.filter(g => g.parentId === parentId).forEach(g => {
                    const p = parentRoute.path + '/' + g.id
                    const crtRoute = {
                        path: p,
                        name: g.desc,
                        children: graphs.filter(graph => graph.group.id === g.id).map(graph => {
                            return {
                                path: p + '/s/' + graph.id,
                                name: graph.name,
                                element: <Graph graph={graph} key={graph.id}/>,
                            }
                        })
                    }
                    parentRoute.children.push(crtRoute)
                    routeMap[g.id] = crtRoute
                    groupsQueue.push(g.id)
                })
            }
            newRoutes.push(...rootRoute.children)
        }
        oldRender()
    })
}


// @ts-ignore
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

export async function getInitialState(): Promise<{ algos: Graph.Algo[], hotse: Graph.HotSETopic[], hhi: Graph.HHILanguage[] }> {
    const algosRes = algoGetAll()
    const hotse = getHot()
    const hhi = getHHI()
    return {
        algos: (await algosRes).algos,
        hotse: (await hotse).topics,
        hhi: (await hhi).languages,
    }
}

export const request: RequestConfig = {
    timeout: 10 * 1000,
    responseInterceptors: [
        [(response) => response, (error: Error) => {
            message.error(error.message)
            return Promise.reject(error)
        }],
    ]
}


export const layout: RunTimeLayoutConfig = () => {
    return {
        logo: require('@/assets/title.png'),
    }
}