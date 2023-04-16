// @ts-ignore
import { Request, Response } from 'express';

export default {
  'GET /api/graph/get': (req: Request, res: Response) => {
    res.status(200).send({
      base: {
        status: 89,
        msg: '边没真点养达九府级共都族并主可据。',
        taskId: 95,
        taskStatus: 1,
        extra: {},
      },
      nodes: [
        { name: '马娜', desc: '界层给提色何过做高金报位从。', deg: 78 },
        { name: '姚磊', desc: '已究日构片型府展己斯好务算节标必面。', deg: 85 },
        { name: '郭强', desc: '三拉等气清天作则红前连加也先往样。', deg: 92 },
        { name: '李娟', desc: '们真四究王示住什众即些件百色教角。', deg: 97 },
        { name: '锺敏', desc: '育经上信家相年林能想温圆快从于院。', deg: 70 },
        { name: '萧强', desc: '五意信率族政近民从在家和书加。', deg: 60 },
        { name: '方芳', desc: '四十文水正严把正运派料用战。', deg: 78 },
        { name: '毛洋', desc: '办龙那音她族走派业音示阶正使活空选交。', deg: 89 },
        { name: '余明', desc: '千学金何特件于边样越须酸比增构。', deg: 70 },
        { name: '徐丽', desc: '支调业同内周打力温速结建图力院。', deg: 65 },
      ],
      edges: [
        { source: '马娜', target: '姚磊' },
        { source: '马娜', target: '郭强' },
        {
          source: '马娜',
          target: '李娟',
        },
        { source: '李娟', target: '锺敏' },
        { source: '李娟', target: '萧强' },
        { source: '萧强', target: '郭强' },
        { source: '郭强', target: '方芳' },
        {
          source: '方芳',
          target: '毛洋',
        },
        { source: '余明', target: '毛洋' },
        {
          source: '余明',
          target: '徐丽',
        },
        { source: '徐丽', target: '毛洋' },
      ],
    });
  },
};
