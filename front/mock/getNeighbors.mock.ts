// @ts-ignore
import {Request, Response} from 'express';

export default {
    'GET /api/graph/node/get': (req: Request, res: Response) => {
        res.status(200).send({
            base: {
                status: 75,
                msg: '明于低划亲主示适维有期年结阶格无机。',
                taskId: 72,
                taskStatus: 1,
                extra: {},
            },
            node: {name: '郭强', desc: '三拉等气清天作则红前连加也先往样。', deg: 92},
            nodes: [
                {name: '马娜', desc: '界层给提色何过做高金报位从。', deg: 78},
                {name: '姚磊', desc: '已究日构片型府展己斯好务算节标必面。', deg: 85},
                {name: '郭强', desc: '三拉等气清天作则红前连加也先往样。', deg: 92},
                {name: '李娟', desc: '们真四究王示住什众即些件百色教角。', deg: 97},
                {name: '萧强', desc: '五意信率族政近民从在家和书加。', deg: 60},
                {name: '方芳', desc: '四十文水正严把正运派料用战。', deg: 78},
                {name: '毛洋', desc: '办龙那音她族走派业音示阶正使活空选交。', deg: 89},
            ],
            edges: [
                {source: '马娜', target: '姚磊'},
                {source: '马娜', target: '郭强'},
                {
                    source: '马娜',
                    target: '李娟',
                },
                {source: '李娟', target: '萧强'},
                {source: '萧强', target: '郭强'},
                {source: '郭强', target: '方芳'},
                {
                    source: '方芳',
                    target: '毛洋',
                },
            ],
        });
    },
};
