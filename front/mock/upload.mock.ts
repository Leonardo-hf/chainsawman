// @ts-ignore
import { Request, Response } from 'express';

export default {
  'POST /api/file/upload': (req: Request, res: Response) => {
    res.status(200).send({ id: 'aBb86386-c5bb-E8E7-02b7-1CFBBa1170e4', size: 89 });
  },
};
