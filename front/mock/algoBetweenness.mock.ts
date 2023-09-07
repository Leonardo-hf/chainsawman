// @ts-ignore
import { Request, Response } from 'express';

export default {
  'GET /api/graph/algo/betweenness': (req: Request, res: Response) => {
    res.status(200).send({
      base: {
        status: 91,
        msg: '别头它体就领南建带张效总象定都风接。',
        taskId: 'Fbf491B2-0a1d-bc3A-4Afb-B6Cd41E6c25D',
        taskStatus: 98,
        extra: {},
      },
      ranks: [
        { nodeId: 81, score: 84 },
        { nodeId: 68, score: 68 },
        { nodeId: 61, score: 81 },
        { nodeId: 92, score: 97 },
        { nodeId: 74, score: 61 },
        { nodeId: 76, score: 67 },
        { nodeId: 87, score: 75 },
        { nodeId: 75, score: 78 },
        { nodeId: 98, score: 92 },
        { nodeId: 90, score: 71 },
      ],
      file: '实专号六联持该快极军北计书且量难想。',
    });
  },
};
