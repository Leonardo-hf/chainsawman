// @ts-ignore
import { Request, Response } from 'express';

export default {
  'POST /api/file/upload': (req: Request, res: Response) => {
    res.status(200).send({ id: '2f0c21A6-A7Bc-aFc9-fC55-81436E90c7EF', size: 64 });
  },
};
