// @ts-ignore
import { Request, Response } from 'express';

export default {
  'POST /api/graph/task/drop': (req: Request, res: Response) => {
    res.status(200).send({
      status: 79,
      msg: '片保原对识高流和无界北路太先。',
      taskId: '9d5Dbc9D-07d7-5FEd-aB88-114bf4dAAaFE',
      taskStatus: 76,
      extra: {},
    });
  },
};
