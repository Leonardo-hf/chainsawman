// @ts-ignore
import { Request, Response } from 'express';

export default {
  'GET /api/graph/algo/betweenness': (req: Request, res: Response) => {
    res.status(200).send({
      base: {
        status: 93,
        msg: '查统何明安较志形信书东变人。',
        taskId: 85,
        taskStatus: 72,
        extra: {},
      },
      ranks: [
        { nodeId: 84, score: 98 },
        { nodeId: 94, score: 88 },
        { nodeId: 78, score: 91 },
        { nodeId: 69, score: 92 },
      ],
      file: '安等车起自照做图机四识由影制增圆。',
    });
  },
};
