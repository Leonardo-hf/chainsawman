// @ts-ignore
import { Request, Response } from 'express';

export default {
  'GET /api/graph/file/put': (req: Request, res: Response) => {
    res.status(200).send({
      url: 'https://github.com/umijs/dumi',
      filename: '建阶达不需金信条深列重表器象土界子。',
    });
  },
};
