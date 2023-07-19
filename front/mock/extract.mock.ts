// @ts-ignore
import {Request, Response} from 'express';

export default {
    'GET /api/util/java': (req: Request, res: Response) => {
        res.status(200).send({
            base: {
                status: 2000,
                msg: '成功',
            },
            package: {
                group: 'org.apache.httpcomponents',
                artifact: 'httpcore',
                version: '4.4.9'
            },
            dependencies: [{
                artifact: 'junit',
                group: 'junit',
                version: '4.12',
                scope: 'test',
                optional: false
            }, {
                artifact: 'mockito-core',
                group: 'org.mockito',
                version: '1.10.19',
                scope: 'test',
                optional: false
            }, {
                artifact: 'commons-logging',
                group: 'commons-logging',
                version: '1.2',
                scope: 'test',
                optional: false
            }, {
                artifact: 'commons-lang3',
                group: 'org.apache.commons',
                version: '3.4',
                scope: 'test',
                optional: false
            }]
        });
    },
};
