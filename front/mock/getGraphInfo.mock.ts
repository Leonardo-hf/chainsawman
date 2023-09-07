// @ts-ignore
import { Request, Response } from 'express';

export default {
  'GET /api/graph/info': (req: Request, res: Response) => {
    res.status(200).send({
      base: {
        status: 72,
        msg: '难江义观事记济需调酸见重织五。',
        taskId: 'b8cA2c79-f9Ea-2E3D-947f-dC2549FCEBe7',
        taskStatus: 90,
        extra: {},
      },
      graph: {
        id: 66,
        status: 68,
        groupId: 79,
        name: '许艳',
        desc: '社东为们类离程长从式队能间与身总。',
        numNode: 98,
        numEdge: 62,
        creatAt: 89,
        updateAt: 79,
      },
    });
  },
};
