// @ts-ignore
import { Request, Response } from 'express';

export default {
  'POST /api/file/upload': (req: Request, res: Response) => {
    res.status(200).send({ id: 'B43ec793-D77e-b0d1-d2C0-49BF6fe86c9f', size: 76 });
  },
};
