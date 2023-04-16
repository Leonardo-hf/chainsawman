// @ts-ignore
/* eslint-disable */
import { request } from '@umijs/max';


/** 上传图 POST /api/graph/create */
export async function createGraph(body: Graph.UploadRequest, options?: { [key: string]: any }) {
  return request<Graph.SearchGraphReply>('/api/graph/create', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 删除图 POST /api/graph/drop */
export async function dropGraph(body: Graph.DropRequest, options?: { [key: string]: any }) {
  return request<Graph.BaseReply>('/api/graph/drop', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 获取图的详细节点信息 GET /api/graph/get */
export async function getGraph(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: Graph.getGraphParams,
  options?: { [key: string]: any },
) {
  return request<Graph.SearchGraphDetailReply>('/api/graph/get', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** 获取图列表 GET /api/graph/getAll */
export async function getAllGraph(options?: { [key: string]: any }) {

  return request<Graph.SearchAllGraphReply>('/api/graph/getAll', {
    method: 'GET',
    ...(options || {}),
  });

}

/** 获取节点信息及邻居节点 GET /api/graph/node/get */
export async function getNeighbors(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: Graph.getNeighborsParams,
  options?: { [key: string]: any },
) {
  return request<Graph.SearchNodeReply>('/api/graph/node/get', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}
