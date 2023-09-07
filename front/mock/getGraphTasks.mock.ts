// @ts-ignore
import { Request, Response } from 'express';

export default {
  'GET /api/graph/task/task/getAll': (req: Request, res: Response) => {
    res.status(200).send({
      base: {
        status: 62,
        msg: '书了传专次也立你度身世真与十从单性生。',
        taskId: 'D6C8D85c-b3bE-f29D-3fb5-44D2f42DCa5A',
        taskStatus: 90,
        extra: {},
      },
      tasks: [
        {
          id: '27c9E94A-Bdb6-bB91-70AF-bD533a4cDca2',
          idf: '条立月用算条相目边县花象里原治可。',
          createTime: 73,
          updateTime: 75,
          status: 0,
          req: '{}',
          res: '{}',
        },
        {
          id: '156ee99C-b187-2F04-e4B2-De34BA1F4c8A',
          idf: '太为小亲做阶存会做来张给斗五。',
          createTime: 70,
          updateTime: 85,
          status: 1,
          req: '{}',
          res: '{}',
        },
        {
          id: 'f75Bdddd-77eC-Ef29-ce85-DaFafF6CC6B5',
          idf: '气在上无间更片方期运外开步历做转今元。',
          createTime: 71,
          updateTime: 99,
          status: 0,
          req: '{}',
          res: '{}',
        },
        {
          id: '3b2959bf-deBE-08BB-8833-Db8Fff24c84b',
          idf: '因得机实系且业快府构商每热根军容低离。',
          createTime: 75,
          updateTime: 80,
          status: 1,
          req: '{}',
          res: '{}',
        },
      ],
    });
  },
};
