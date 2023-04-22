import {defineConfig} from '@umijs/max';



export default defineConfig({
    antd: {
    },
    // access: {},
    model: {},
    initialState: {},
    request: {
        dataField: 'data'
    },

    proxy: {
        '/api': {
            'target': 'http://127.0.0.1:8888/',
            'changeOrigin': true,
            // 'pathRewrite': {'^/api': ''},
        },
    },
    // mock: false,
    dva: {},
    layout: {
        title: 'chainsawman',
    },
    plugins: ['@umijs/max-plugin-openapi'],

    openAPI: [{
        requestLibPath: "import { request } from '@umijs/max';",
        schemaPath: [__dirname, 'graph.json'].join('/'),
        mock: true,
        namespace: 'Graph',
        projectName: 'graph'
    }, {
        requestLibPath: "import { request } from '@umijs/max';",
        schemaPath: [__dirname, 'file.json'].join('/'),
        mock: true,
        namespace: 'File',
        projectName: 'file'
    }],
    routes: [
        {
            path: '/',
            redirect: '/home',
        },
        {
            name: '首页',
            path: '/home',
            component: './Home',
        },

        {
            name: 'graph',
            path: '/graph',
            routes: [
                {
                    name: 'test',
                    path: '/graph/test',
                    component: './Graph'
                }
            ]
        },
    ],
    npmClient: 'npm',
});

