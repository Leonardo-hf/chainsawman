// @ts-ignore
/* eslint-disable */
import { request } from '@umijs/max';

/** 此处后端没有提供注释 POST /api/graph/drop */
export async function dropGraph(body: API.DropRequest, options?: { [key: string]: any }) {
  const data=request<API.BaseReply>('/api/graph/drop', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
  console.log(data)
  return data
}

/** 此处后端没有提供注释 GET /api/graph/get/${param0} */
export async function getGraph(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.getGraphParams,
  options?: { [key: string]: any },
) {
  const { graph: param0, ...queryParams } = params;
  return request<API.SearchGraphDetailReply>(`/api/graph/get/${param0}`, {
    method: 'GET',
    params: { ...queryParams },
    ...(options || {}),
  });
}

/** 此处后端没有提供注释 GET /api/graph/getAll */
export async function getAllGraph(options?: { [key: string]: any }) {
  console.log("wnm,d")
  const data=request<API.SearchAllGraphReply>('/api/graph/getAll', {
    method: 'GET',
    ...(options || {}),
  });
  console.log(data)
  return data;
}

/** 此处后端没有提供注释 POST /api/graph/upload */
export async function upload(options?: { [key: string]: any }) {
  return request<API.SearchGraphReply>('/api/graph/upload', {
    method: 'POST',
    ...(options || {}),
  });
}
