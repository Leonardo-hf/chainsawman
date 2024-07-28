import {defineConfig} from '@umijs/max';


export default defineConfig({
    antd: {},
    history: {type: 'hash'},
    hash: true,
    model: {},
    initialState: {},
    request: {
        dataField: 'data'
    },
    proxy: {
        '/api/graph': {
            'target': 'http://127.0.0.1:8888/',
            // 'target': 'http://127.0.0.1:30130/',
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
        '/api/vul': {
            'target': 'http://129.211.191.135/',
            'changeOrigin': true,
            'pathRewrite': {'^/api/vul': '/'},
        },
        '/source': {
            'target': 'http://127.0.0.1:9000/',
            'changeOrigin': true,
        },
        '/algo': {
            'target': 'http://127.0.0.1:9000/',
            'changeOrigin': true,
        },
        '/assets': {
            'target': 'http://127.0.0.1:9000/',
            'changeOrigin': true,
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
            component: './Dashboard',
        },
        {
            name: '数据仓库',
            path: '/database',
            routes: [
                {
                    name: '基础软件资产',
                    path: '/database/software',
                },
                {
                    name: '软件依赖图谱',
                    path: '/database/graph',
                    routes: []
                },
                {
                    name: '漏洞情报库',
                    path: '/database/vul',
                    component: './Vul'
                },
            ]
        },
        {
            name: '度量评估模型',
            path: '/algo',
            routes: []
        },
        {
            name: '算法执行',
            path: '/exec',
            routes: [
                {
                    name: '高影响力软件识别',
                    path: '/exec/impact',
                    component: './Impact'
                },
                {
                    name: '卡脖子软件识别',
                    path: '/exec/strangle',
                    component: './Strangle'
                },
                {
                    name: '执行记录',
                    path: '/exec/records',
                    component: './AlgoExec'
                },
            ]
        },
        {
            name: '分析工具',
            path: '/util',
            routes: [
                {
                    name: '软件成分分析',
                    path: '/util/extractor',
                    component: './Extractor'
                },
                {
                    name: '软件静态分析',
                    path: '/util/lint',
                    component: './Lint'
                },
                {
                    name: '软件动态测试',
                    path: '/util/dynamic',
                },
                {
                    name: '软件漏洞修复',
                    path: '/util/packet',
                },
            ]
        },
        {
            name: '数据源管理',
            path: '/config',
            routes: [
                {
                    name: '网络结构配置',
                    path: '/config/group',
                    component: './Group'
                },
                {
                    name: '网络数据配置',
                    path: '/config/source',
                    component: './GTable',
                },
            ]
        },
        {
            name: '系统监控',
            path: '/monitor',
            component: './Monitor'
        },
    ],
    npmClient: 'npm',
});

