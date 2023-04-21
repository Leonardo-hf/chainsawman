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
        { name: 'a', desc: '温备机切国格只经住农际什。', nodes: 64, edges: 63 ,status:0,id:0},
        { name: 'b', desc: '易线全果里信维山极和看商得商老想。', nodes: 87, edges: 96 ,status:1,id:1},
        { name: 'c', desc: '年观与问引小复县具什族规且支选文明都。', nodes: 73, edges: 98 ,status:1,id:2},
        { name: 'd', desc: '识行日三现数计转精之用称风接化事。', nodes: 78, edges: 92 ,status:1,id:3},
        { name: 'e', desc: '最重感信热重部长作路革任做金查清前。', nodes: 77, edges: 92 ,status:1,id:4},
        { name: 'f', desc: '片济方边起型建问林所不一地没。', nodes: 71, edges: 87 ,status:1,id:5},
        { name: 'g', desc: '史参交并思员运问马现效选队压置术其。', nodes: 63, edges: 95 ,status:1,id:6},
        { name: 'h', desc: '严些志非我周是等再几知转不。', nodes: 64, edges: 79 ,status:1,id:7},
      ],
    });
  },
  'GET /api/graph/getAll1': (req: Request, res: Response) => {

    res.status(200).send({
      base: {
        status: 98,
        msg: '角前照性果办铁进治物研养拉同所公路里。',
        taskId: 75,
        taskStatus: 81,
        extra: {},
      },
      graphs: [
        { name: 'a', desc: '温备机切国格只经住农际什。', nodes: 64, edges: 63 ,status:1,id:0},
        { name: 'b', desc: '易线全果里信维山极和看商得商老想。', nodes: 87, edges: 96 ,status:1,id:1},
        { name: 'c', desc: '年观与问引小复县具什族规且支选文明都。', nodes: 73, edges: 98 ,status:1,id:2},
        { name: 'd', desc: '识行日三现数计转精之用称风接化事。', nodes: 78, edges: 92 ,status:1,id:3},
        { name: 'e', desc: '最重感信热重部长作路革任做金查清前。', nodes: 77, edges: 92 ,status:1,id:4},
        { name: 'f', desc: '片济方边起型建问林所不一地没。', nodes: 71, edges: 87 ,status:1,id:5},
        { name: 'g', desc: '史参交并思员运问马现效选队压置术其。', nodes: 63, edges: 95 ,status:1,id:6},
        { name: 'h', desc: '严些志非我周是等再几知转不。', nodes: 64, edges: 79 ,status:1,id:7},
      ],
    });
  },
};
