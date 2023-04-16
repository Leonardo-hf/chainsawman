import {getAllGraph} from "@/services/graph/graph";
import Graph from "@/pages/Graph";
import {RequestConfig} from "@@/plugin-request/request";

let graphs: any

// export function render(oldRender) {
//     getAllGraph().then(data => {
//         const routes: { name: string, path: string, element: JSX.Element }[] = []
//         if (data) {
//             data.graphs.forEach((graph: { name: string; }) => routes.push({
//                 path: '/graph/' + graph.name,
//                 element: <Graph/>,
//                 name: graph.name,
//             }))
//         }
//         graphs = routes
//         oldRender()
//     })
// }


// export function patchClientRoutes({routes}) {
//     // TODO: 这个是直接根据路由的序号找的，扩展性差
//     let menu = routes[0].children[2]
//     // menu["routes"] = []
//     menu["children"] = []
//     graphs.forEach(graph => {
//         // menu.routes.push(graph)
//         menu.children.push(graph)
//     })
//     console.log(routes)
// }

export async function getInitialState(): Promise<{ name: string }> {
    return {name: '@umijs/m322x'};
}

export const request: RequestConfig = {
    timeout: 1000,
};

export const layout = () => {
    return {
        logo: 'https://i.328888.xyz/2023/03/22/YMzqZ.png',
        menu: {
            locale: false,
        },
    };
};