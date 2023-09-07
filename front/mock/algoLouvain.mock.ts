// @ts-ignore
import { Request, Response } from 'express';

export default {
  'GET /api/graph/algo/louvain': (req: Request, res: Response) => {
    res.status(200).send({
      base: {
        status: 83,
        msg: '达处前东党成路还家细转外民。',
        taskId: 'ecaa45BD-ddbf-8d1D-5cA5-b73D44D8a29D',
        taskStatus: 95,
        extra: {},
      },
      ranks: [
        { nodeId: 76, score: 80 },
        { nodeId: 63, score: 75 },
        { nodeId: 91, score: 68 },
        { nodeId: 95, score: 62 },
        { nodeId: 75, score: 99 },
        { nodeId: 66, score: 87 },
        { nodeId: 88, score: 73 },
        { nodeId: 81, score: 65 },
        { nodeId: 92, score: 66 },
        { nodeId: 66, score: 83 },
        { nodeId: 70, score: 98 },
      ],
      file: '米育自持成东五太红明风条。',
    });
  },
};
