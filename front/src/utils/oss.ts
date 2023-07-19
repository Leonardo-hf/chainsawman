import {fileGetPresigned, filePutPresigned} from '@/services/graph/graph'
import {request} from "@umijs/max";
import {RcFile} from 'antd/es/upload';

export const upload = async (file: RcFile) => {
    const res = await filePutPresigned()
    const filename = res.filename
    return await request(res.url, {
        headers: {
            'Content-Type': file.type,
        },
        data: file,
        method: 'put'
    }).then(_ => filename)
}

export const getPreviewURL = async (fileId: string) => {
    const res = await fileGetPresigned({filename: fileId})
    return res.url
}