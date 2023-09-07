// @ts-ignore
import { Request, Response } from 'express';

export default {
  'GET /api/graph/algo/vr': (req: Request, res: Response) => {
    res.status(200).send({
      base: {
        status: 62,
        msg: '级据群科认国总报把元复同长王。',
        taskId: 'd7f66BEf-B96E-Aa18-bAEc-44E0C2ED2096',
        taskStatus: 90,
        extra: {},
      },
      ranks: [
        { nodeId: 70, score: 82 },
        { nodeId: 76, score: 94 },
        { nodeId: 93, score: 93 },
        { nodeId: 78, score: 70 },
        { nodeId: 97, score: 69 },
        { nodeId: 70, score: 97 },
        { nodeId: 86, score: 70 },
        { nodeId: 85, score: 91 },
        { nodeId: 91, score: 64 },
        { nodeId: 90, score: 90 },
      ],
      file: '际带油现圆划千称石器感问两题节派七百。',
    });
  },
};
