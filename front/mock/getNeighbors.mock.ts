// @ts-ignore
import {Request, Response} from 'express';

export default {
    'GET /api/graph/node/get': (req: Request, res: Response) => {
        res.status(200).send({
            base: {
                status: 70,
                msg: '以米具习向月要到制可角理从专。',
                taskId: 82,
                taskStatus: 1,
                extra: {},
            },
            node: {id: 64, name: '熊明', desc: '除面明界劳化几学该统当三公。', deg: 93},

            nodes: [
                {id: 64, name: '熊明', desc: '除面明界劳化几学该统当三公。', deg: 93},
                {id: 75, name: '陆强', desc: '更满热教斯江认志手族利铁般直万标适。', deg: 88},
                {id: 74, name: '姜敏', desc: '基话办知长部又便千回切些基因想具果。', deg: 91},
                {id: 61, name: '刘磊', desc: '看也大形容等二放斯整外量离为据统。', deg: 74},
                {id: 78, name: '苏秀兰', desc: '此认道清个越本也同象办手书革集角然。', deg: 61},
            ],
            edges: [
                {source: 64, target: 75},
                {source: 64, target: 74},
                {source: 64, target: 61},
                {source: 78, target: 64},
            ],
        });
    },
};
