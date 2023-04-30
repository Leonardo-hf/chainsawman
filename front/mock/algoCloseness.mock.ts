// @ts-ignore
import { Request, Response } from 'express';

export default {
  'GET /api/graph/algo/closeness': (req: Request, res: Response) => {
    res.status(200).send({
      base: {
        status: 81,
        msg: '如记江那年使更风山共山领月县派或。',
        taskId: 97,
        taskStatus: 98,
        extra: {},
      },
      ranks: [
        { nodeId: 99, score: 75 },
        { nodeId: 73, score: 88 },
        { nodeId: 64, score: 91 },
        { nodeId: 96, score: 61 },
      ],
      file: '现者她正格理型意去图青解身际道。',
    });
  },
};
