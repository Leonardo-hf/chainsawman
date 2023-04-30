// @ts-ignore
import { Request, Response } from 'express';

export default {
  'GET /api/graph/algo/pr': (req: Request, res: Response) => {
    res.status(200).send({
      base: {
        status: 96,
        msg: '入便自题化照先九影件商容压程分接支按。',
        taskId: 64,
        taskStatus: 70,
        extra: {},
      },
      ranks: [
        { nodeId: 79, score: 74 },
        { nodeId: 75, score: 99 },
        { nodeId: 85, score: 97 },
        { nodeId: 95, score: 80 },
        { nodeId: 79, score: 72 },
        { nodeId: 91, score: 96 },
        { nodeId: 83, score: 98 },
      ],
      file: '较来标众快红西感参标相清下斗。',
    });
  },
};
