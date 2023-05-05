import {ProColumns, ProTable} from '@ant-design/pro-components';
import {Layout, Row, Typography} from 'antd';
import React from 'react';
import styles from './Guide.less';

interface Props {
    score: number
}

const MetricTable: React.FC<Props> = (props) => {
    const {score} = props;
    const columns: ProColumns[] = [
        {
            dataIndex: 'key',
        },
        {
            dataIndex: 'value'
        },
    ]
    return <ProTable
        dataSource={[
            {
                key: 'score',
                value: score
            }
        ]}
        columns={columns}
        search={false}
        showHeader={false}
        options={false}
        pagination={false}
    />
};

export default MetricTable;
