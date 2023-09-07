// @ts-ignore
import { Request, Response } from 'express';

export default {
  'GET /api/graph/algo/degree': (req: Request, res: Response) => {
    res.status(200).send({
      base: {
        status: 74,
        msg: '满即并次铁将带很备整制争只三二速五。',
        taskId: 'eD55ff40-3eA5-a8AE-2C91-8BB3bDfCBcDB',
        taskStatus: 89,
        extra: {},
      },
      ranks: [
        { nodeId: 66, score: 86 },
        { nodeId: 89, score: 96 },
        { nodeId: 86, score: 84 },
        { nodeId: 98, score: 92 },
        { nodeId: 100, score: 94 },
        { nodeId: 72, score: 76 },
        { nodeId: 82, score: 70 },
        { nodeId: 62, score: 71 },
        { nodeId: 91, score: 75 },
        { nodeId: 88, score: 63 },
        { nodeId: 84, score: 69 },
        { nodeId: 86, score: 88 },
        { nodeId: 61, score: 75 },
        { nodeId: 80, score: 63 },
      ],
      file: '率利选还写近积思里划许按军方整。',
    });
  },
};
