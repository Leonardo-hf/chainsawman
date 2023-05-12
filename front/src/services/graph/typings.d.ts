declare namespace Graph {
  type Algo = {
    id: number;
    name: string;
    desc: string;
    isCustom: boolean;
    type: number;
    params: Element[];
  };

  type algoAvgCCParams = {
    taskId?: number;
    graphId: number;
  };

  type algoBetweennessParams = {
    taskId?: number;
    graphId: number;
  };

  type algoClosenessParams = {
    taskId?: number;
    graphId: number;
  };

  type algoDegreeParams = {
    taskId?: number;
    graphId: number;
  };

  type AlgoDegreeRequest = {
    taskId?: number;
    graphId: number;
  };

  type algoLouvainParams = {
    taskId?: number;
    graphId: number;
    maxIter?: number;
    internalIter?: number;
    tol?: number;
  };

  type AlgoLouvainRequest = {
    taskId?: number;
    graphId: number;
    maxIter: number;
    internalIter: number;
    tol: number;
  };

  type AlgoMetricReply = {
    base: BaseReply;
    score: number;
  };

  type algoPageRankParams = {
    taskId?: number;
    graphId: number;
    iter?: number;
    prob?: number;
  };

  type AlgoPageRankRequest = {
    taskId?: number;
    graphId: number;
    iter: number;
    prob: number;
  };

  type AlgoRankReply = {
    base: BaseReply;
    ranks: Rank[];
    file: string;
  };

  type AlgoReply = {
    base: BaseReply;
    algos: Algo[];
  };

  type AlgoRequest = {
    taskId?: number;
    graphId: number;
  };

  type algoVoteRankParams = {
    taskId?: number;
    graphId: number;
    iter?: number;
  };

  type AlgoVoteRankRequest = {
    taskId?: number;
    graphId: number;
    iter: number;
  };

  type BaseReply = {
    status: number;
    msg: string;
    taskId: number;
    taskStatus: number;
    extra: Record<string, any>;
  };

  type DropRequest = {
    graphId: number;
  };

  type DropTaskRequest = {
    taskId?: number;
  };

  type Edge = {
    source: number;
    target: number;
  };

  type Element = {
    key: string;
    keyDesc: string;
    type: number;
  };

  type GetGraphInfoReply = {
    name: string;
    graphId: number;
  };

  type GetGraphInfoRequest = {
    name: string;
  };

  type getGraphParams = {
    taskId?: number;
    graphId: number;
    min: number;
  };

  type getGraphTasksParams = {
    graphId: number;
  };

  type getNeighborsParams = {
    taskId?: number;
    graphId: number;
    nodeId: number;
    distance: number;
    min: number;
  };

  type GetNodeReduceRequest = {
    id: number;
  };

  type Graph = {
    id: number;
    status: number;
    name: string;
    desc: string;
    nodes: number;
    edges: number;
  };

  type Node = {
    id: number;
    name: string;
    desc: string;
    deg: number;
  };

  type NodeReduce = {
    name: string;
    id: number;
  };

  type NodesInfo = {
    nodes: NodeReduce[];
  };

  type Param = {
    key: string;
    value: string;
  };

  type Rank = {
    nodeId: number;
    score: number;
  };

  type SearchAllGraphReply = {
    base: BaseReply;
    graphs: Graph[];
  };

  type SearchGraphDetailReply = {
    base: BaseReply;
    nodes: Node[];
    edges: Edge[];
  };

  type SearchGraphReply = {
    base: BaseReply;
    graph: Graph;
  };

  type SearchNodeReply = {
    base: BaseReply;
    node: Node;
    nodes: Node[];
    edges: Edge[];
  };

  type SearchNodeRequest = {
    taskId?: number;
    graphId: number;
    nodeId: number;
    distance: number;
    min: number;
  };

  type SearchRequest = {
    taskId?: number;
    graphId: number;
    min: number;
  };

  type SearchTasksReply = {
    base: BaseReply;
    tasks: Task[];
  };

  type SearchTasksRequest = {
    graphId: number;
  };

  type Task = {
    id: number;
    idf: string;
    createTime: number;
    updateTime: number;
    status: number;
    req: string;
    res: string;
  };

  type UploadEmptyRequest = {
    graph: string;
    desc?: string;
  };

  type UploadRequest = {
    taskId?: number;
    graph: string;
    desc?: string;
    nodeId: string;
    edgeId: string;
    graphId?: number;
  };
}
