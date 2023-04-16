// @ts-ignore
import { Request, Response } from 'express';

export default {
  'GET /api/graph/get': (req: Request, res: Response) => {
    res.status(200).send({
      base: {
        status: 89,
        msg: '边没真点养达九府级共都族并主可据。',
        taskId: 95,
        taskStatus: 70,
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
        { source: '选过省周志南却青结龙五土后来较常。', target: '准统直合构市小则管自总给律江。' },
        { source: '的第积深正列体导种花见己到提。', target: '历们支便连观想心多子则她活单。' },
        {
          source: '展般收即平受高斯本类不连却压走机为。',
          target: '过与华十构间它风命学间更队器。',
        },
        { source: '百经全第自党上好近西题易造面值。', target: '种中义不流得住信提质做处。' },
        { source: '着广等些部好对实我手文据都矿。', target: '南且验音他规此过清线步习直。' },
        { source: '权无安命心放等题车着经打交即。', target: '或在认县料光改合下也带天。' },
        { source: '国流单约件名又响五象光须间之华。', target: '近下义前满科里较重地制战较。' },
        {
          source: '政提展思放交报严压应为土走地设适。',
          target: '经整你如命便油线属支但能能员油命内。',
        },
        { source: '人队按出论或林报音价铁水状表科。', target: '政水头专体约级候却油七关王好族。' },
        {
          source: '采积战内律社油定长数声本产小口写。',
          target: '技革为派手性第头断条龙料院华方别。',
        },
        { source: '拉将位属色认无花话干速可适品世周。', target: '示其做看切族历已战也做划最多。' },
      ],
    });
  },
};
