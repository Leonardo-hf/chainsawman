// @ts-ignore
import {Request, Response} from 'express';

export default {
    'GET /api/graph/getAll': (req: Request, res: Response) => {
        res.status(200).send({
            groups: [
                {
                    id: 1,
                    name: '默认分组',
                    desc: '间区精决事无革心期型划通由第又然。',
                    nodeTypeList: [
                        {
                            id: 1,
                            name: '默认节点',
                            desc: '大体题对能等热直每再那则前。',
                            display: 'color',
                            attrs: [
                                {name: '名称', desc: '亲委记铁叫队各感候王专难须化。', primary: true, type: 0},
                                {name: '描述', desc: '别制起我专由布须越及党红光又。', primary: false, type: 0},
                            ]
                        }
                    ],
                    edgeTypeList: [
                        {
                            id: 1,
                            name: '默认边',
                            desc: '再其生史想电县效天局五少展家。',
                            edgeDirection: 0,
                            display: 'real',
                            attrs: []
                        }
                    ],
                    graphs: [
                        {
                            id: 100,
                            status: 1,
                            groupId: 1,
                            name: '易强',
                            desc: '由各变导论做系厂周体群院土信共包务局。',
                            numNode: 92,
                            numEdge: 88,
                            creatAt: 66,
                            updateAt: 89,
                        },
                        {
                            id: 101,
                            status: 1,
                            groupId: 1,
                            name: '许芳',
                            desc: '正合受者六革存政重斗造五将建圆。',
                            numNode: 87,
                            numEdge: 67,
                            creatAt: 77,
                            updateAt: 91,
                        },
                        {
                            id: 102,
                            status: 1,
                            groupId: 1,
                            name: '沈娟',
                            desc: '资养放色己只圆日斗和各造军值取本会县。',
                            numNode: 84,
                            numEdge: 91,
                            creatAt: 68,
                            updateAt: 61,
                        },
                    ]
                },
                {
                    id: 2,
                    name: '新建分组',
                    desc: '律成完性京志证应毛县消持认花开。',
                    nodeTypeList: [
                        {
                            id: 2,
                            name: '新建节点1',
                            desc: '至千强回能料展此达满大平思。',
                            display: 'color',
                            attrs: [
                                {name: 'aaa', desc: '做造走除等例术照教己代消快却。', primary: true, type: 0},
                                {name: 'bbb', desc: '住种平行质水江省火应派平共备造。', primary: false, type: 1},
                            ]
                        },
                        {
                            id: 3,
                            name: '新建节点2',
                            desc: '易张马出写铁统必适说海根此部些规。',
                            display: 'icon',
                            attrs: [
                                {name: 'ccc', desc: '去性部对得须白交历较件证程。', primary: true, type: 0},
                            ]
                        }
                    ],
                    edgeTypeList: [
                        {
                            id: 1,
                            name: '新建边',
                            desc: '再其生史想电县效天局五少展家。',
                            edgeDirection: 0,
                            display: 'dash',
                            attrs: [
                                {name: '描述', desc: '条调高却老张元家立格备原外里第而看。', primary: true, type: 0},
                                {name: '权重', desc: '人构始多更那很叫个向音林而应候。', primary: false, type: 1},
                            ]
                        }
                    ],
                    graphs: [
                        {
                            id: 104,
                            status: 1,
                            groupId: 2,
                            name: '易强',
                            desc: '由各变导论做系厂周体群院土信共包务局。',
                            numNode: 92,
                            numEdge: 88,
                            creatAt: 66,
                            updateAt: 89,
                        },
                    ]
                },
            ],
        })
    }
}