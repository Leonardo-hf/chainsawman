import {ProColumns, ProTable} from '@ant-design/pro-components';
import {Button, Layout, Row, Typography} from 'antd';
import React from 'react';
import styles from './Guide.less';

type Row = {
    node: string,
    score: number
}

interface Props {
    rows: Row[],
    file: string
}

const RankTable: React.FC<Props> = (props) => {
    const {rows, file} = props;
    const columns: ProColumns[] = [
        {
            title: 'node',
            dataIndex: 'node',
        },
        {
            title: 'rank',
            dataIndex: 'rank'
        },
    ]
    return <ProTable
        rowKey='node'
        dataSource={rows}
        columns={columns}
        search={false}
        options={false}
        toolBarRender={() => [
            <a key="out" href={'http://127.0.0.1:8890/api/file/get/' + file}>导出数据</a>
        ]}
        pagination={false}
    />
};

export default RankTable;
