// @ts-ignore
import { Request, Response } from 'express';

export default {
  'POST /api/graph/create': (req: Request, res: Response) => {
    res.status(200).send({
      base: {
        status: 86,
        msg: '信类参车县强数天极局真应土存过计。',
        taskId: 63,
        taskStatus: 1,
        extra: {},
      },
      graph: {
        id: 96,
        status: 96,
        name: '傅桂英',
        desc: '温动生则果照达总志林论组党写。',
        nodes: 88,
        edges: 86,
      },
    });
  },
};
