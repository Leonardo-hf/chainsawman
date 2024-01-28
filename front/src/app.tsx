import {RequestConfig} from "@@/plugin-request/request";
import React from "react";
import {RunTimeLayoutConfig} from "@umijs/max";
import Graph from "./pages/Graph"

import {algoGetAll, getAllGraph, getHHI, getHot} from "./services/graph/graph";
import {parseGroups} from "@/models/global";
import {message} from "antd";
import {RootGroupID} from "@/constants";
import AlgoMenu from "@/pages/AlgoMenu";
import Algo from "@/pages/Algo";

let newGraphRoutes: any[] = []
let newAlgoRoutes: any = {}
let initAlgos: Graph.Algo[] = []

// 查询组信息和算法信息，生成路由
export function render(oldRender: () => void) {
    const getGraph = getAllGraph().then(data => {
        const gs = data.groups
        if (gs) {
            // 初始化图与组信息
            let {graphs, groups} = parseGroups(gs)
            graphs = graphs.filter(a => a.status == 1)
            // 生成路由
            const rootRoute = {
                path: '/graph',
                children: []
            }
            const routeMap: any = {1: rootRoute}
            const groupsQueue = [RootGroupID]
            while (groupsQueue.length > 0) {
                const parentId = groupsQueue.shift()!
                const parentRoute = routeMap[parentId]
                groups.filter(g => g.parentId === parentId).forEach(g => {
                    const p = parentRoute.path + '/' + g.id
                    const crtRoute = {
                        path: p,
                        name: <span style={{color: '#708090'}}>{g.desc}</span>,
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
            newGraphRoutes.push(...rootRoute.children)
        }
        // oldRender()
    })
    const getAlgo = algoGetAll().then(res => {
        initAlgos = res.algos
        const category = ['软件影响力', '软件风险', '其他']
        // 数据结构为: {tag: {tag, algos}[]}[]，即分类下包含子分类
        // 分类为 '软件影响力', '软件风险', '其他'，其他算法中包含除 '软件影响力', '软件风险' 的算法
        newAlgoRoutes = Object.entries<{ tag: string, algos: Graph.Algo[] }[]>(initAlgos.reduce((group: any, algo) => {
            let {tag} = algo
            if (tag != '软件风险' && tag != '软件影响力') {
                const item = group['其他']?.filter((it: { tag: string, algos: Graph.Algo[] }) => it.tag == tag)
                if (item && item.length) {
                    item[0].algos = [...item[0].algos, algo]
                } else {
                    group['其他'] = [...group['其他'] ?? [], {tag: tag, algos: [algo]}]
                }
                return group
            }
            group[tag] = group[tag] ?? [{tag: tag, algos: []}]
            group[tag][0].algos.push(algo)
            return group
        }, {}))
            // 排序
            .sort((a, b) => category.indexOf(a[0]) - category.indexOf(b[0]))
            // 渲染为路由
            .map((e, id: number) => {
                return {
                    path: `/algo/${id}`,
                    name: e[0],
                    children: [
                        {
                            path: `/algo/${id}/index`,
                            name: '概述',
                            element: <AlgoMenu category={e[1]}/>
                        },
                        // 收集每个分类下的节点
                        ...e[1].map(as => {
                            // 收集每个子分类下的节点
                            return as.algos.map(a => {
                                return {
                                    path: `/algo/${id}/${a.id}`,
                                    name: a.name,
                                    element: <Algo algo={a}/>
                                }
                            })
                        }).reduce((a, b) => [...a, ...b])
                    ]
                }
            })
    })
    Promise.all([getGraph, getAlgo]).then(_ => oldRender())
}


// @ts-ignore
export function patchClientRoutes({routes}) {
    // TODO: 这个是直接根据路由的序号找的，扩展性差
    routes[0].children[2].children = newGraphRoutes
    routes[0].children[3].children = newAlgoRoutes
}

export async function getInitialState(): Promise<{
    algos: Graph.Algo[], hotse: Graph.HotSETopic[], hhi: Graph.HHILanguage[],
}> {
    const hotse = getHot()
    const hhi = getHHI()
    return {
        algos: initAlgos,
        hotse: (await hotse).topics,
        hhi: (await hhi).languages,
    }
}

export const request: RequestConfig = {
    timeout: 10 * 1000,
    errorConfig: {
        errorHandler: (err) => {
            message.error(err.message)
        }
    },
    responseInterceptors: [
        (res) => {
            //@ts-ignore
            const base = res.data.base
            if (base && base.status > 2000) {
                throw new Error(base.msg)
            }
            return res
        }
    ]
}


export const layout: RunTimeLayoutConfig = () => {
    return {
        logo: require('@/assets/title.png'),
    }
}