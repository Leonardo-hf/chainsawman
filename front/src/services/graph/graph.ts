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

/** Depth算法 GET /api/graph/algo/depth */
export async function algoDepth(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: Graph.algoDepthParams,
  options?: { [key: string]: any },
) {
  return request<Graph.AlgoRankReply>('/api/graph/algo/depth', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** Ecology算法 GET /api/graph/algo/ecology */
export async function algoEcology(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: Graph.algoEcologyParams,
  options?: { [key: string]: any },
) {
  return request<Graph.AlgoRankReply>('/api/graph/algo/ecology', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** Integration算法 GET /api/graph/algo/integration */
export async function algoIntegration(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: Graph.algoIntegrationParams,
  options?: { [key: string]: any },
) {
  return request<Graph.AlgoRankReply>('/api/graph/algo/integration', {
    method: 'GET',
    params: {
      ...params,
    },
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

/** Quantity算法 GET /api/graph/algo/quantity */
export async function algoQuantity(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: Graph.algoQuantityParams,
  options?: { [key: string]: any },
) {
  return request<Graph.AlgoRankReply>('/api/graph/algo/quantity', {
    method: 'GET',
    params: {
      // iter has a default value: 100
      iter: '100',
      ...params,
    },
    ...(options || {}),
  });
}

/** 新建图 POST /api/graph/create */
export async function createGraph(
  body: Graph.CreateGraphRequest,
  options?: { [key: string]: any },
) {
  return request<Graph.GraphInfoReply>('/api/graph/create', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 获取图的详细边节点信息 GET /api/graph/detail */
export async function getGraph(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: Graph.getGraphParams,
  options?: { [key: string]: any },
) {
  return request<Graph.GetGraphDetailReply>('/api/graph/detail', {
    method: 'GET',
    params: {
      // max has a default value: 2000
      max: '2000',
      ...params,
    },
    ...(options || {}),
  });
}

/** 删除图 POST /api/graph/drop */
export async function dropGraph(body: Graph.DropGraphRequest, options?: { [key: string]: any }) {
  return request<Graph.BaseReply>('/api/graph/drop', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 获得文件下载链接 GET /api/graph/file/get */
export async function fileGetPresigned(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: Graph.fileGetPresignedParams,
  options?: { [key: string]: any },
) {
  return request<Graph.PresignedReply>('/api/graph/file/get', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** 获得文件上传链接 GET /api/graph/file/put */
export async function filePutPresigned(options?: { [key: string]: any }) {
  return request<Graph.PresignedReply>('/api/graph/file/put', {
    method: 'GET',
    ...(options || {}),
  });
}

/** 获取图列表 GET /api/graph/getAll */
export async function getAllGraph(options?: { [key: string]: any }) {
  return request<Graph.GetAllGraphReply>('/api/graph/getAll', {
    method: 'GET',
    ...(options || {}),
  });
}

/** 新建策略组 POST /api/graph/group/create */
export async function createGroup(
  body: Graph.CreateGroupRequest,
  options?: { [key: string]: any },
) {
  return request<Graph.GroupInfoReply>('/api/graph/group/create', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 删除策略组 POST /api/graph/group/drop */
export async function dropGroup(body: Graph.DropGroupRequest, options?: { [key: string]: any }) {
  return request<Graph.BaseReply>('/api/graph/group/drop', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 根据名称查询图信息 GET /api/graph/info */
export async function getGraphInfo(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: Graph.getGraphInfoParams,
  options?: { [key: string]: any },
) {
  return request<Graph.GraphInfoReply>('/api/graph/info', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** 获得全部节点 GET /api/graph/node/getAll */
export async function getNodes(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: Graph.getNodesParams,
  options?: { [key: string]: any },
) {
  return request<Graph.GetNodesReply>('/api/graph/node/getAll', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** 获得匹配节点 GET /api/graph/node/getMatch */
export async function getMatchNodes(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: Graph.getMatchNodesParams,
  options?: { [key: string]: any },
) {
  return request<Graph.GetMatchNodesReply>('/api/graph/node/getMatch', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** 获取节点信息及邻居节点 GET /api/graph/node/nbr */
export async function getNeighbors(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: Graph.getNeighborsParams,
  options?: { [key: string]: any },
) {
  return request<Graph.GetGraphDetailReply>('/api/graph/node/nbr', {
    method: 'GET',
    params: {
      // max has a default value: 2000
      max: '2000',
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

/** 获取图计算任务 GET /api/graph/task/task/getAll */
export async function getGraphTasks(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: Graph.getGraphTasksParams,
  options?: { [key: string]: any },
) {
  return request<Graph.GetTasksReply>('/api/graph/task/task/getAll', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** 更新图 POST /api/graph/update */
export async function updateGraph(
  body: Graph.UpdateGraphRequest,
  options?: { [key: string]: any },
) {
  return request<Graph.GraphInfoReply>('/api/graph/update', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}
