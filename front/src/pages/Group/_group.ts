import {formatEnum2Options} from "@/utils/format";

const nodeStyleEnum = {
    color: {
        text: '色彩',
        status: 'color'
    },
    icon: {
        text: '图标',
        status: 'icon'
    }
}
const nodeStyleOptions = formatEnum2Options(nodeStyleEnum)
const edgeStyleEnum = {
    real: {
        text: '实线',
        status: 'real'
    },
    dash: {
        text: '虚线',
        status: 'dash'
    }
}
const edgeStyleOptions = formatEnum2Options(edgeStyleEnum)
const edgeDirectEnum = {
    true: {
        text: '有向',
        status: true
    },
    false: {
        text: '无向',
        status: false
    }
}
const edgeDirectOptions = formatEnum2Options(edgeDirectEnum)

const defaultGroup = [{
    name: 'normal',
    desc: '标准节点',
    type: 'node',
    attrs: [{
        name: 'name',
        desc: '名称',
        type: 0,
        primary: true
    }, {
        name: 'desc',
        desc: '描述',
        type: 0
    }]
}, {
    name: 'normal',
    desc: '标准边',
    type: 'edge',
}]
export  {
    nodeStyleEnum, nodeStyleOptions,
    edgeStyleEnum, edgeStyleOptions,
    edgeDirectEnum, edgeDirectOptions,
    defaultGroup
}