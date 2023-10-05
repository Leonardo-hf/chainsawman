import {
    fileAlgoGetPresigned,
    fileLibPutPresigned,
    fileSourcePutPresigned
} from '@/services/graph/graph'
import {request} from "@umijs/max";
import {RcFile} from 'antd/es/upload';

export const uploadSource = async (file: RcFile) => {
    const res = await fileSourcePutPresigned()
    const filename = res.filename
    const url = res.url
    return await request(url.substring(url.indexOf('/source')), {
        timeout: 3600 * 1000,
        headers: {
            'Content-Type': file.type,
        },
        data: file,
        method: 'put'
    }).then(_ => filename)
}

export const uploadLib = async (file: RcFile) => {
    const res = await fileLibPutPresigned()
    const filename = res.filename
    const url = res.url
    return await request(url.substring(url.indexOf('/source')), {
        timeout: 3600 * 1000,
        headers: {
            'Content-Type': file.type,
        },
        data: file,
        method: 'put'
    }).then(_ => filename)
}

export const getPreviewURL = async (fileId: string) => {
    const res = await fileAlgoGetPresigned({filename: fileId})
    const url =  res.url
    return url.substring(url.indexOf('/algo'))
}
