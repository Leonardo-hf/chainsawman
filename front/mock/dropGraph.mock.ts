// @ts-ignore
import { Request, Response } from 'express';

export default {
  'POST /api/graph/drop': (req: Request, res: Response) => {
    res.status(200).send({
      status: 91,
      msg: '明改达克表存山话条长回克世月可革。',
      taskId: 70,
      taskStatus: 82,
      extra: {},
    });
  },
};
