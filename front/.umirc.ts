import {defineConfig} from '@umijs/max';


export default defineConfig({
    antd: {},
    // access: {},
    model: {},
    initialState: {},
    request: {
        dataField: 'data'
    },
    proxy: {
        '/api/graph': {
            'target': 'http://127.0.0.1:8888/',
            'changeOrigin': true,
            // 'pathRewrite': {'^/api': ''},
        },
        '/api/monitor': {
            'target': 'http://127.0.0.1:8890/',
            'changeOrigin': true,
            'pathRewrite': {'^/api/monitor': '/'},
        },
        '/api/util': {
            'target': 'http://127.0.0.1:8082/',
            'changeOrigin': true,
            'pathRewrite': {'^/api/util': '/'},
        },
        '/source': {
            'target': 'http://127.0.0.1:9000/',
            'changeOrigin': true,
        },
    },
    mock: false,
    dva: {},
    layout: {
        title: 'chainsawman',
    },
    plugins: ['@umijs/max-plugin-openapi'],
    openAPI: [{
        requestLibPath: "import { request } from '@umijs/max';",
        schemaPath: [__dirname, 'graph.json'].join('/'),
        mock: false,
        namespace: 'Graph',
        projectName: 'graph'
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
        {
            name: '软件成分分析',
            path: '/util/extractor',
            component: './Extractor'
        },
        {
            name: 'monitor',
            path: '/monitor',
            component: './Monitor'
        }
    ],
    npmClient: 'npm',
});

