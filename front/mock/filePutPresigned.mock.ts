// @ts-ignore
import { Request, Response } from 'express';

export default {
  'GET /api/graph/file/put': (req: Request, res: Response) => {
    res.status(200).send({
      url: 'https://www.mocky.io/v2/5cc8019d300000980a055e76',
      filename: '建阶达不需金信条深列重表器象土界子。',
    });
  },
};
