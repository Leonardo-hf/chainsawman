import {useRequest} from 'umi';
import {getAllGraph} from "@/services/graph/graph";
// 运行时配置

// 全局初始化数据配置，用于 Layout 用户信息和权限初始化
// 更多信息见文档：https://umijs.org/docs/api/runtime-config#getinitialstate
export async function getInitialState(): Promise<{ name: string }> {
    return {name: '@umijs/m322x'};
}

// async function getMenu() {
//     const {data, error, loading} = useRequest(getAllGraph(), {throwOnError: true});
//     const routes: { name: string; path: string; component: string }[] = []
//     let menu = {
//         name: 'graph',
//         path: '/graph',
//         routes: routes
//     }
//     if (data) {
//         console.log(data)
//         for (let r in data["graphs"]) {
//             menu.routes.push(
//                 {
//                     name: r,
//                     path: '/graph/' + r,
//                     component: './Graph'
//                 })
//         }
//     }
//     return menu
// }

export const layout = () => {
    // const menu = getMenu()
    return {
        logo: 'https://i.328888.xyz/2023/03/22/YMzqZ.png',
        menu: {
            locale: false,
        },
    };
};
