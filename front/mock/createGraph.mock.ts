// @ts-ignore
import { Request, Response } from 'express';

export default {
  'POST /api/graph/create': (req: Request, res: Response) => {
    res.status(200).send({
      base: {
        status: 98,
        msg: '关和热好构程品改方水研将族斯出权。',
        taskId: 82,
        taskStatus: 98,
        extra: {},
      },
      graph: { name: '郑平', desc: '十世非们公见活金厂都算公连它满员。', nodes: 60, edges: 93 },
    });
  },
};
