import {ProColumns, ProTable} from '@ant-design/pro-components';
import React, {useState} from 'react';
import Loading from "@ant-design/pro-card/es/components/Loading";
import {getPreviewURL} from "@/utils/oss";
import {readRemoteFile} from "react-papaparse";
import {Parser} from "papaparse";
import {Space} from "antd";

interface Props {
    file: string
}

const RankTable: React.FC<Props> = (props) => {
    const {file} = props;
    const [fileURL, setFileURL] = useState<string>()
    const [next, setNext] = useState<Promise<Parser>>()
    const [headers, setHeaders] = useState<ProColumns[]>([])
    const [rows, setRows] = useState<any[]>([])
    const [finish, setFinish] = useState<boolean>(false)
    if (!fileURL) {
        getPreviewURL(file).then(url => setFileURL(url))
    } else if (rows.length === 0) {
        readRemoteFile(fileURL, {
            download: true,
            header: true,
            step: (results, parser) => {
                console.log(results)
                if (!headers.length) {
                    console.log(results.meta.fields)
                    setHeaders(results.meta.fields!.map((f: string) => {
                        return {
                            title: f,
                            dataIndex: f,
                            copyable: true,
                            fixed: 'left',
                        }
                    }))
                }
                setRows([...rows, results.data])
                if (rows.length % 10 === 0) {
                    parser.pause()
                    const nextPreview = new Promise<Parser>(function (resolve) {
                        resolve(parser)
                    })
                    setNext(nextPreview)
                }
            },
            complete: () => {
                setFinish(true)
            }
        })
    }
    console.log(rows)
    return <ProTable
        rowKey='node'
        dataSource={rows}
        columns={headers}
        search={false}
        options={false}
        toolBarRender={() => [
            fileURL ? <Space>
                {finish ? <span>已全部加载</span> :
                    <a onClick={() => next?.then(parser => parser.resume())}>预览更多</a>
                }
                <a key='out' target='_blank' href={fileURL}>导出全部</a>
            </Space> : <Loading/>
        ]}
        pagination={false}
    />
};

export default RankTable;
