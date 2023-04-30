// @ts-ignore
import { Request, Response } from 'express';

export default {
  'GET /api/graph/algo/degree': (req: Request, res: Response) => {
    res.status(200).send({
      base: {
        status: 61,
        msg: '量则车总开场生住她可出定华。',
        taskId: 62,
        taskStatus: 63,
        extra: {},
      },
      ranks: [
        { nodeId: 87, score: 86 },
        { nodeId: 64, score: 78 },
        { nodeId: 84, score: 97 },
        { nodeId: 88, score: 97 },
        { nodeId: 69, score: 84 },
        { nodeId: 100, score: 62 },
        { nodeId: 94, score: 88 },
      ],
      file: '候消采指电再八不又局听快整率计或京。',
    });
  },
};
