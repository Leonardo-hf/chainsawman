// @ts-ignore
import { Request, Response } from 'express';

export default {
  'POST /api/file/upload': (req: Request, res: Response) => {
    res.status(200).send({ id: 'c010A3DD-7bDc-b7ed-D147-4e89aCcb60EE', size: 81 });
  },
};
