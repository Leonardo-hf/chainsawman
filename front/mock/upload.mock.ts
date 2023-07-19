// @ts-ignore
import { Request, Response } from 'express';

export default {
  'POST /api/file/upload': (req: Request, res: Response) => {
    res.status(200).send({ id: '3Ea3fFb7-3b48-cFA7-89aB-9f217ceEE616', size: 74 });
  },
};
