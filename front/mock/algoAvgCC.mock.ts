// @ts-ignore
import { Request, Response } from 'express';

export default {
  'GET /api/graph/algo/avgCC': (req: Request, res: Response) => {
    res.status(200).send({
      base: {
        status: 86,
        msg: '活清如东相算事革维很转机广每。',
        taskId: 75,
        taskStatus: 81,
        extra: {},
      },
      score: 60,
    });
  },
};
