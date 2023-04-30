// @ts-ignore
import { Request, Response } from 'express';

export default {
  'GET /api/graph/get': (req: Request, res: Response) => {
    res.status(200).send({
      base: {
        status: 80,
        msg: '别统整矿及写铁思信题什动。',
        taskId: 91,
        taskStatus: 1,
        extra: {},
      },
      nodes: [
        { id: 64, name: '熊明', desc: '除面明界劳化几学该统当三公。', deg: 93 },
        { id: 75, name: '陆强', desc: '更满热教斯江认志手族利铁般直万标适。', deg: 88 },
        { id: 78, name: '苏秀兰', desc: '此认道清个越本也同象办手书革集角然。', deg: 61 },
        { id: 93, name: '戴秀英', desc: '技意志结上社间交回增许别产面断过。', deg: 78 },
        { id: 74, name: '姜敏', desc: '基话办知长部又便千回切些基因想具果。', deg: 91 },
        { id: 61, name: '刘磊', desc: '看也大形容等二放斯整外量离为据统。', deg: 74 },
      ],
      edges: [
        { source: 64, target: 75 },
        { source: 64, target: 74 },
        { source: 64, target: 61 },
        { source: 61, target: 74 },
        { source: 93, target: 61 },
        { source: 78, target: 64 },
        { source: 61, target: 75 },
      ],
    });
  },
};
