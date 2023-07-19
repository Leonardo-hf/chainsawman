import {ProColumns, ProTable} from '@ant-design/pro-components';
import {Row} from 'antd';
import React, {useState} from 'react';
import Loading from "@ant-design/pro-card/es/components/Loading";
import {getPreviewURL} from "@/utils/oss";

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
    const [fileURL, setFileURL] = useState<string>('')
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
    if (fileURL === '') {
        getPreviewURL(file).then(url => setFileURL(url))
    }
    return <ProTable
        rowKey='node'
        dataSource={rows}
        columns={columns}
        search={false}
        options={false}
        toolBarRender={() => [
            fileURL === '' ? <Loading/> :
                <a key='out' target='_blank' href={fileURL}>导出数据</a>
        ]}
        pagination={false}
    />
};

export default RankTable;
