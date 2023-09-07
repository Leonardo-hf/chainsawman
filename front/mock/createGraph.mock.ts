// @ts-ignore
import { Request, Response } from 'express';

export default {
  'POST /api/graph/create': (req: Request, res: Response) => {
    res.status(200).send({
      base: {
        status: 84,
        msg: '动利而引过东选用义基基半识解三五车。',
        taskId: '86BAe65e-89C5-b1e3-6D8B-fC3d4CcBb3E2',
        taskStatus: 93,
        extra: {},
      },
      graph: {
        id: 87,
        status: 73,
        groupId: 76,
        name: '卢杰',
        desc: '通物技包率查次向除色其那将节持马数。',
        numNode: 67,
        numEdge: 73,
        creatAt: 77,
        updateAt: 61,
      },
    });
  },
};
