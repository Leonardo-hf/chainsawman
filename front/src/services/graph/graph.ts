// @ts-ignore
/* eslint-disable */
import { request } from '@umijs/max';

/** 平均聚类系数 GET /api/graph/algo/avgCC */
export async function algoAvgCC(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: Graph.algoAvgCCParams,
  options?: { [key: string]: any },
) {
  return request<Graph.AlgoMetricReply>('/api/graph/algo/avgCC', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** Betweenness算法 GET /api/graph/algo/betweenness */
export async function algoBetweenness(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: Graph.algoBetweennessParams,
  options?: { [key: string]: any },
) {
  return request<Graph.AlgoRankReply>('/api/graph/algo/betweenness', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** Closeness算法 GET /api/graph/algo/closeness */
export async function algoCloseness(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: Graph.algoClosenessParams,
  options?: { [key: string]: any },
) {
  return request<Graph.AlgoRankReply>('/api/graph/algo/closeness', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** 度中心度算法 GET /api/graph/algo/degree */
export async function algoDegree(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: Graph.algoDegreeParams,
  options?: { [key: string]: any },
) {
  return request<Graph.AlgoRankReply>('/api/graph/algo/degree', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** 查询算法 GET /api/graph/algo/getAll */
export async function algoGetAll(options?: { [key: string]: any }) {
  return request<Graph.AlgoReply>('/api/graph/algo/getAll', {
    method: 'GET',
    ...(options || {}),
  });
}

/** Louvain聚类算法 GET /api/graph/algo/louvain */
export async function algoLouvain(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: Graph.algoLouvainParams,
  options?: { [key: string]: any },
) {
  return request<Graph.AlgoRankReply>('/api/graph/algo/louvain', {
    method: 'GET',
    params: {
      // maxIter has a default value: 10
      maxIter: '10',
      // internalIter has a default value: 5
      internalIter: '5',
      // tol has a default value: 0.5
      tol: '0.5',
      ...params,
    },
    ...(options || {}),
  });
}

/** PageRank算法 GET /api/graph/algo/pr */
export async function algoPageRank(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: Graph.algoPageRankParams,
  options?: { [key: string]: any },
) {
  return request<Graph.AlgoRankReply>('/api/graph/algo/pr', {
    method: 'GET',
    params: {
      // iter has a default value: 3
      iter: '3',
      // prob has a default value: 0.85
      prob: '0.85',
      ...params,
    },
    ...(options || {}),
  });
}

/** VoteRank算法 GET /api/graph/algo/vr */
export async function algoVoteRank(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: Graph.algoVoteRankParams,
  options?: { [key: string]: any },
) {
  return request<Graph.AlgoRankReply>('/api/graph/algo/vr', {
    method: 'GET',
    params: {
      // iter has a default value: 100
      iter: '100',
      ...params,
    },
    ...(options || {}),
  });
}

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

/** 创建空图 POST /api/graph/createEmpty */
export async function createEmptyGraph(
  body: Graph.UploadEmptyRequest,
  options?: { [key: string]: any },
) {
  return request<Graph.SearchGraphReply>('/api/graph/createEmpty', {
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

/** 删除任务 POST /api/graph/task/drop */
export async function dropTask(body: Graph.DropTaskRequest, options?: { [key: string]: any }) {
  return request<Graph.BaseReply>('/api/graph/task/drop', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 获取图计算任务 GET /api/graph/task/getAll */
export async function getGraphTasks(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: Graph.getGraphTasksParams,
  options?: { [key: string]: any },
) {
  return request<Graph.SearchTasksReply>('/api/graph/task/getAll', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}
