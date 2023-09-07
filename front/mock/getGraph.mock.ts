// @ts-ignore
import {Request, Response} from 'express';

export default {
    'GET /api/graph/detail': (req: Request, res: Response) => {
        res.status(200).send({
            base: {
                status: 94,
                msg: '金权管角常山养际到及拉他如属机声量。',
                taskId: '74bc5E2f-B141-4ee6-80dC-cE9eDcEeADA5',
                taskStatus: 1,
                extra: {},
            },
            nodePacks: [
                {
                    tag: '新建节点1',
                    nodes: [
                        {
                            id: 1,
                            deg: 1,
                            attrs: [
                                {key: 'aaa', value: '节点1'},
                                {key: 'bbb', value: 1},
                            ],
                        },
                        {
                            id: 2,
                            deg: 1,
                            attrs: [
                                {key: 'aaa', value: '节点2'},
                                {key: 'bbb', value: 2},
                            ],
                        },
                        {
                            id: 3,
                            deg: 1,
                            attrs: [
                                {key: 'aaa', value: '节点3'},
                                {key: 'bbb', value: 3},
                            ],
                        },
                        {
                            id: 4,
                            deg: 1,
                            attrs: [
                                {key: 'aaa', value: '节点4'},
                                {key: 'bbb', value: 4},
                            ],
                        },
                        {
                            id: 5,
                            deg: 1,
                            attrs: [
                                {key: 'aaa', value: '节点5'},
                                {key: 'bbb', value: 5},
                            ],
                        },
                    ],
                },
                {
                    tag: '新建节点2',
                    nodes: [
                        {
                            id: 6,
                            deg: 3,
                            attrs: [
                                {key: 'ccc', value: '节点A'},
                            ],
                        },
                        {
                            id: 7,
                            deg: 2,
                            attrs: [
                                {key: 'ccc', value: '节点B'},
                            ],
                        },
                    ],
                },
            ],
            edgePacks: [
                {
                    tag: '新建边',
                    edges: [
                        {
                            source: 1,
                            target: 6,
                            attrs: [
                                {key: '描述', value: 'depend'},
                                {key: '权重', value: 1},
                            ],
                        },
                        {
                            source: 2,
                            target: 6,
                            attrs: [
                                {key: '描述', value: 'depend'},
                                {key: '权重', value: 1},
                            ],
                        },
                        {
                            source: 3,
                            target: 6,
                            attrs: [
                                {key: '描述', value: 'depend'},
                                {key: '权重', value: 1},
                            ],
                        },
                        {
                            source: 4,
                            target: 7,
                            attrs: [
                                {key: '描述', value: 'depend'},
                                {key: '权重', value: 1},
                            ],
                        },
                        {
                            source: 5,
                            target: 7,
                            attrs: [
                                {key: '描述', value: 'depend'},
                                {key: '权重', value: 1},
                            ],

                        }
                    ]
                }
            ]
        })
    }
}
