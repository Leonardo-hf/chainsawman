import {defineConfig} from '@umijs/max';

export default defineConfig({
    antd: {},
    access: {},
    model: {},
    initialState: {},
    request: {},
    dva: {},
    layout: {
        title: 'chainsawman',
    },
    plugins: ['@umijs/max-plugin-openapi'],
    openAPI: {
        // /umi/plugin/openapi
        requestLibPath: "import { request } from '@umijs/max';",
        schemaPath: '/home/approdite/IdeaProjects/chainsawman-frontend/graph.json',
        mock: false,
    },
    routes: [
        {
            path: '/',
            redirect: '/home',
        },
        // {
        //     name: '首页',
        //     path: '/home',
        //     component: './Home',
        // },
        // {
        //     name: '权限演示',
        //     path: '/access',
        //     component: './Access',
        // },
        // {
        //     name: ' CRUD 示例',
        //     path: '/table',
        //     component: './Table',
        // },
        {
            name: 'graph',
            path: '/graph',
            routes: [
                {
                    name: 'python',
                    path: '/graph/test',
                    component: './Graph'
                }
            ]
        },
    ],
    npmClient: 'npm',
});

