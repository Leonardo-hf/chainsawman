// @ts-ignore
import { Request, Response } from 'express';

export default {
  'GET /api/graph/algo/vr': (req: Request, res: Response) => {
    res.status(200).send({
      base: {
        status: 82,
        msg: '许百和以江总题书了命间济老音群石务然。',
        taskId: 87,
        taskStatus: 92,
        extra: {},
      },
      ranks: [
        { nodeId: 66, score: 94 },
        { nodeId: 75, score: 95 },
        { nodeId: 86, score: 61 },
        { nodeId: 85, score: 81 },
        { nodeId: 84, score: 79 },
        { nodeId: 93, score: 85 },
        { nodeId: 68, score: 62 },
        { nodeId: 64, score: 91 },
        { nodeId: 87, score: 67 },
        { nodeId: 92, score: 84 },
        { nodeId: 79, score: 91 },
        { nodeId: 95, score: 77 },
        { nodeId: 66, score: 73 },
        { nodeId: 97, score: 80 },
        { nodeId: 84, score: 73 },
        { nodeId: 79, score: 95 },
        { nodeId: 99, score: 75 },
        { nodeId: 80, score: 98 },
        { nodeId: 68, score: 100 },
      ],
      file: '真直任界专其中单基如业革子求通系信。',
    });
  },
};
