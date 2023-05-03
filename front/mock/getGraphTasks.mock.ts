// @ts-ignore
import {Request, Response} from 'express';

export default {
    'GET /api/graph/task/getAll': (req: Request, res: Response) => {
        res.status(200).send({
            base: {
                status: 81,
                msg: '并所记路织江器入身本等没标众求。',
                taskId: 66,
                taskStatus: 80,
                extra: {},
            },
            tasks: [
                {
                    idf: 0,
                    desc: 'degree',
                    createTime: 93,
                    updateTime: 72,
                    status: 1,
                    req: {},
                    res: {
                        ranks: [
                            {
                                nodeID: 64,
                                score: 93
                            },
                            {
                                nodeID: 74,
                                score: 91
                            },
                            {
                                nodeID: 75,
                                score: 88
                            },
                        ],
                        file: 'betweenness108364e3-3282-4e2e-8704-6d0e203739c5.csv'
                    },
                },
                {
                    idf: 1,
                    desc: 'average clustering coefficient',
                    createTime: 61,
                    updateTime: 71,
                    status: 1,
                    req: {},
                    res: {
                        score: 100
                    },
                },
                {
                    idf: 2,
                    desc: 'louvain',
                    createTime: 92,
                    updateTime: 74,
                    status: 0,
                    req: {},
                    res: {},
                },
            ],
        });
    },
};
