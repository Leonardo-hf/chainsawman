// @ts-ignore
import { Request, Response } from 'express';

export default {
  'GET /api/graph/file/get': (req: Request, res: Response) => {
    res.status(200).send({
      url: 'https://procomponents.ant.design/',
      filename: '斯头热战第斯九共电说所分委有天型。',
    });
  },
};
