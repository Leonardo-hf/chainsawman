import {ProColumns, ProTable} from '@ant-design/pro-components';
import {Row} from 'antd';
import React from 'react';

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
            dataIndex: 'rank',
            renderText: text => text.toFixed(3)
        },
    ]
    return <ProTable
        rowKey='node'
        dataSource={rows}
        columns={columns}
        search={false}
        options={false}
        toolBarRender={() => [
            <a key='out' target='_blank' href={'/api/file/fetch/' + file}>导出数据</a>
        ]}
        pagination={false}
    />
};

export default RankTable;
