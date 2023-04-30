// @ts-ignore
import { Request, Response } from 'express';

export default {
  'GET /api/graph/algo/louvain': (req: Request, res: Response) => {
    res.status(200).send({
      base: {
        status: 68,
        msg: '片表复群因把现油查性叫从决术。',
        taskId: 88,
        taskStatus: 66,
        extra: {},
      },
      ranks: [
        { nodeId: 95, score: 87 },
        { nodeId: 65, score: 69 },
        { nodeId: 98, score: 78 },
        { nodeId: 73, score: 75 },
        { nodeId: 64, score: 93 },
        { nodeId: 73, score: 91 },
        { nodeId: 75, score: 90 },
        { nodeId: 64, score: 82 },
        { nodeId: 74, score: 75 },
        { nodeId: 62, score: 97 },
        { nodeId: 78, score: 95 },
        { nodeId: 79, score: 98 },
        { nodeId: 78, score: 99 },
      ],
      file: '单查主电道目族实作布为声。',
    });
  },
};
