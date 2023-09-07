// @ts-ignore
import { Request, Response } from 'express';

export default {
  'GET /api/graph/algo/pr': (req: Request, res: Response) => {
    res.status(200).send({
      base: {
        status: 73,
        msg: '话说石角但情道队空儿龙起历维温类油。',
        taskId: 'e9Db7f62-E9f0-3e50-5dAa-0EB67E914Db7',
        taskStatus: 62,
        extra: {},
      },
      ranks: [
        { nodeId: 63, score: 63 },
        { nodeId: 65, score: 84 },
      ],
      file: '数度引日张张设严光反除整次工必。',
    });
  },
};
