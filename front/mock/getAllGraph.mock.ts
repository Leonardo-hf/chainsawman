// @ts-ignore
import { Request, Response } from 'express';

export default {
  'GET /api/graph/getAll': (req: Request, res: Response) => {
    res.status(200).send({
      base: {
        status: 98,
        msg: '角前照性果办铁进治物研养拉同所公路里。',
        taskId: 75,
        taskStatus: 81,
        extra: {},
      },
      graphs: [
        { name: '程平', desc: '温备机切国格只经住农际什。', nodes: 64, edges: 63 },
        { name: '江洋', desc: '易线全果里信维山极和看商得商老想。', nodes: 87, edges: 96 },
        { name: '毛霞', desc: '年观与问引小复县具什族规且支选文明都。', nodes: 73, edges: 98 },
        { name: '文芳', desc: '识行日三现数计转精之用称风接化事。', nodes: 78, edges: 92 },
        { name: '万娜', desc: '最重感信热重部长作路革任做金查清前。', nodes: 77, edges: 92 },
        { name: '邱杰', desc: '片济方边起型建问林所不一地没。', nodes: 71, edges: 87 },
        { name: '赖刚', desc: '史参交并思员运问马现效选队压置术其。', nodes: 63, edges: 95 },
        { name: '邵勇', desc: '严些志非我周是等再几知转不。', nodes: 64, edges: 79 },
        { name: '魏芳', desc: '格增支越即展方两题种再但几般手它。', nodes: 89, edges: 62 },
        { name: '许超', desc: '路易长眼无列相文部张去世下见按几实。', nodes: 66, edges: 89 },
        { name: '赖军', desc: '组做线林众斗回被度人战号较门头。', nodes: 97, edges: 70 },
        { name: '白芳', desc: '示器步间金容设美亲众七儿办各县。', nodes: 73, edges: 94 },
        { name: '潘平', desc: '约种算中信验听要年长断员万级。', nodes: 94, edges: 86 },
        { name: '曹杰', desc: '市往实克话压色位列质二么划立全争积。', nodes: 87, edges: 64 },
        { name: '锺超', desc: '广意马身第眼物养织火书学群万离率。', nodes: 83, edges: 75 },
        { name: '任芳', desc: '专今目去律行步报段政整太群。', nodes: 82, edges: 81 },
      ],
    });
  },
};
