// @ts-ignore
import { Request, Response } from 'express';

export default {
  'POST /api/graph/drop': (req: Request, res: Response) => {
    res.status(200).send({
      status: 72,
      msg: '代明术办至下时接是高做许酸路却许。',
      taskId: 87,
      taskStatus: 70,
      extra: {},
    });
  },
};
