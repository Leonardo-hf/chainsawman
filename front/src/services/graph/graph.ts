// @ts-ignore
/* eslint-disable */
import { request } from '@umijs/max';

/** 创建算法 POST /api/graph/algo/create */
export async function algoCreate(body: Graph.CreateAlgoRequest, options?: { [key: string]: any }) {
  return request<Graph.BaseReply>('/api/graph/algo/create', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 删除算法 POST /api/graph/algo/drop */
export async function algoDrop(body: Graph.AlgoIDRequest, options?: { [key: string]: any }) {
  return request<Graph.BaseReply>('/api/graph/algo/drop', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 执行算法 POST /api/graph/algo/exec */
export async function algoExec(body: Graph.ExecAlgoRequest, options?: { [key: string]: any }) {
  return request<Graph.GetAlgoTaskReply>('/api/graph/algo/exec', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 查询算法 GET /api/graph/algo/getAll */
export async function algoGetAll(options?: { [key: string]: any }) {
  return request<Graph.GetAlgoReply>('/api/graph/algo/getAll', {
    method: 'GET',
    ...(options || {}),
  });
}

/** 查询算法文档 GET /api/graph/algo/getDoc */
export async function algoGetDoc(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: Graph.algoGetDocParams,
  options?: { [key: string]: any },
) {
  return request<Graph.GetAlgoDocReply>('/api/graph/algo/getDoc', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** 删除任务 POST /api/graph/algo/task/drop */
export async function dropAlgoTask(
  body: Graph.DropAlgoTaskRequest,
  options?: { [key: string]: any },
) {
  return request<Graph.BaseReply>('/api/graph/algo/task/drop', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 获取单个图计算任务结果 GET /api/graph/algo/task/get */
export async function getAlgoTaskByID(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: Graph.getAlgoTaskByIDParams,
  options?: { [key: string]: any },
) {
  return request<Graph.GetAlgoTaskReply>('/api/graph/algo/task/get', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** 获取图计算任务 GET /api/graph/algo/task/getAll */
export async function getAlgoTask(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: Graph.getAlgoTaskParams,
  options?: { [key: string]: any },
) {
  return request<Graph.GetAlgoTasksReply>('/api/graph/algo/task/getAll', {
    method: 'GET',
    params: {
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

/** 获得文件下载链接 GET /api/graph/file/get/algo */
export async function fileAlgoGetPresigned(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: Graph.fileAlgoGetPresignedParams,
  options?: { [key: string]: any },
) {
  return request<Graph.PresignedReply>('/api/graph/file/get/algo', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** 获得算法库文件上传链接 GET /api/graph/file/put/lib */
export async function fileLibPutPresigned(options?: { [key: string]: any }) {
  return request<Graph.PresignedReply>('/api/graph/file/put/lib', {
    method: 'GET',
    ...(options || {}),
  });
}

/** 获得图源文件上传链接 GET /api/graph/file/put/source */
export async function fileSourcePutPresigned(options?: { [key: string]: any }) {
  return request<Graph.PresignedReply>('/api/graph/file/put/source', {
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

/** 获得匹配节点 GET /api/graph/node/getMatchByTag */
export async function getMatchNodesByTag(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: Graph.getMatchNodesByTagParams,
  options?: { [key: string]: any },
) {
  return request<Graph.GetMatchNodesByTagReply>('/api/graph/node/getMatchByTag', {
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

/** 获得软件类型HHI指数 GET /api/graph/se/hhi */
export async function getHHI(options?: { [key: string]: any }) {
  return request<Graph.GetHHIReply>('/api/graph/se/hhi', {
    method: 'GET',
    ...(options || {}),
  });
}

/** 获得热门软件概要 GET /api/graph/se/hot */
export async function getHot(options?: { [key: string]: any }) {
  return request<Graph.GetHotSEReply>('/api/graph/se/hot', {
    method: 'GET',
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
