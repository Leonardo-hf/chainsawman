// @ts-ignore
import { Request, Response } from 'express';

export default {
  'POST /api/file/upload': (req: Request, res: Response) => {
    res.status(200).send({ id: '32e9d62D-cBAE-1fa0-Bd04-DdeE2ECc3cFa', size: 80 });
  },
};
