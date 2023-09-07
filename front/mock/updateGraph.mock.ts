// @ts-ignore
import { Request, Response } from 'express';

export default {
  'POST /api/graph/update': (req: Request, res: Response) => {
    res.status(200).send({
      base: {
        status: 70,
        msg: '法严划派争走低下约化之部斯会处龙立。',
        taskId: '6d72d25e-1EF7-BB57-22Ac-2a54Ed6b9167',
        taskStatus: 95,
        extra: {},
      },
      graph: {
        id: 77,
        status: 93,
        groupId: 80,
        name: '顾勇',
        desc: '离反人体也特重花基员性以类月。',
        numNode: 86,
        numEdge: 82,
        creatAt: 64,
        updateAt: 97,
      },
    });
  },
};
