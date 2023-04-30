// @ts-ignore
import { Request, Response } from 'express';

export default {
  'GET /api/graph/getAll': (req: Request, res: Response) => {
    res.status(200).send({
      base: {
        status: 61,
        msg: '高变何改并青气其处收做提设电得你走。',
        taskId: 79,
        taskStatus: 99,
        extra: {},
      },
      graphs: [
        {
          id: 98,
          status: 67,
          name: '崔丽',
          desc: '阶生收地素效意五步品老边美林。',
          nodes: 92,
          edges: 70,
        },
        {
          id: 72,
          status: 67,
          name: '范霞',
          desc: '改装完维主维立观音选火场写处思。',
          nodes: 73,
          edges: 64,
        },
      ],
    });
  },
};
