// @ts-ignore
import {Request, Response} from 'express';

export default {
    'GET /api/graph/node/getMatch': (req: Request, res: Response) => {
        if (req.query.keywords && req.query.keywords.length > 2) {
            console.log(req.query.keywords.length)
            res.status(200).send({
                base: {
                    status: 94,
                    msg: '金权管角常山养际到及拉他如属机声量。',
                    taskId: '74bc5E2f-B141-4ee6-80dC-cE9eDcEeADA5',
                    taskStatus: 1,
                    extra: {},
                },
                matchNodePacks: [
                    {
                        tag: '新建节点1',
                        match: [
                            {
                                id: 1,
                                primaryAttr: '节点1'
                            },
                        ]
                    },
                ]
            })
            return
        }
        res.status(200).send({
            base: {
                status: 94,
                msg: '金权管角常山养际到及拉他如属机声量。',
                taskId: '74bc5E2f-B141-4ee6-80dC-cE9eDcEeADA5',
                taskStatus: 1,
                extra: {},
            },
            matchNodePacks: [
                {
                    tag: '新建节点1',
                    match: [
                        {
                            id: 1,
                            primaryAttr: '节点1'
                        },
                        {
                            id: 2,
                            primaryAttr: '节点2'
                        },
                    ]
                },
                {
                    tag: '新建节点2',
                    match: [
                        {
                            id: 6,
                            primaryAttr: '节点6'
                        },
                    ]
                },
            ]
        })
    }
}