import {fileGetPresigned, filePutPresigned} from '@/services/graph/graph'
import {request} from "@umijs/max";
import {RcFile} from 'antd/es/upload';

export const upload = async (file: RcFile) => {
    const res = await filePutPresigned()
    const filename = res.filename
    const url = res.url
    return await request(url.substring(url.indexOf('/source')), {
        timeout: 10000,
        headers: {
            'Content-Type': file.type,
        },
        data: file,
        method: 'put'
    }).then(_ => filename)
}

export const getPreviewURL = async (fileId: string) => {
    const res = await fileGetPresigned({filename: fileId})
    const url =  res.url
    return url.substring(url.indexOf('/algo'))
}