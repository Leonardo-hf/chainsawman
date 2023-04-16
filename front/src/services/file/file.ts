// @ts-ignore
/* eslint-disable */
import { request } from '@umijs/max';

/** 上传文件 POST /api/file/upload */
export async function upload(options?: { [key: string]: any }) {
  return request<File.UploadReply>('/api/file/upload', {
    method: 'POST',
    ...(options || {}),
  });
}
