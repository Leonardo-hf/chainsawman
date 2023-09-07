// @ts-ignore
import { Request, Response } from 'express';

export default {
  'POST /api/graph/drop': (req: Request, res: Response) => {
    res.status(200).send({
      status: 65,
      msg: '划切规等面都如引位的器公机。',
      taskId: 'F5ce869F-56b2-bf4D-5692-45CF5c66D587',
      taskStatus: 82,
      extra: {},
    });
  },
};
