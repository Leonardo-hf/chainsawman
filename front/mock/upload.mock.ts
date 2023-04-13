// @ts-ignore
import { Request, Response } from 'express';

export default {
  'POST /api/file/upload': (req: Request, res: Response) => {
    res.status(200).send({ id: 'C66fEF88-cb9f-7cE5-b874-4e3F5bBFB7A6', size: 66 });
  },
};
