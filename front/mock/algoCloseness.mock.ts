// @ts-ignore
import { Request, Response } from 'express';

export default {
  'GET /api/graph/algo/closeness': (req: Request, res: Response) => {
    res.status(200).send({
      base: {
        status: 97,
        msg: '利相级容资一以之性白影强变件。',
        taskId: '18bDC6c8-cC2B-4eb1-CBE9-CBe8d3A65Ae3',
        taskStatus: 88,
        extra: {},
      },
      ranks: [
        { nodeId: 83, score: 93 },
        { nodeId: 72, score: 68 },
        { nodeId: 74, score: 66 },
        { nodeId: 90, score: 84 },
        { nodeId: 78, score: 72 },
        { nodeId: 76, score: 69 },
      ],
      file: '好算划府想会己与王年备习种共龙从论。',
    });
  },
};
