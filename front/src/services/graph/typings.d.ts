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
    taskId?: string;
    graphId: number;
  };

  type algoBetweennessParams = {
    taskId?: string;
    graphId: number;
  };

  type algoClosenessParams = {
    taskId?: string;
    graphId: number;
  };

  type algoDegreeParams = {
    taskId?: string;
    graphId: number;
  };

  type AlgoDegreeRequest = {
    taskId?: string;
    graphId: number;
  };

  type algoLouvainParams = {
    taskId?: string;
    graphId: number;
    maxIter?: number;
    internalIter?: number;
    tol?: number;
  };

  type AlgoLouvainRequest = {
    taskId?: string;
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
    taskId?: string;
    graphId: number;
    iter?: number;
    prob?: number;
  };

  type AlgoPageRankRequest = {
    taskId?: string;
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
    taskId?: string;
    graphId: number;
  };

  type algoVoteRankParams = {
    taskId?: string;
    graphId: number;
    iter?: number;
  };

  type AlgoVoteRankRequest = {
    taskId?: string;
    graphId: number;
    iter: number;
  };

  type BaseReply = {
    status: number;
    msg: string;
    taskId: string;
    taskStatus: number;
    extra: Record<string, any>;
  };

  type DropRequest = {
    graphId: number;
  };

  type DropTaskRequest = {
    taskId?: string;
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
    taskId?: string;
    graphId: number;
    min: number;
  };

  type getGraphTasksParams = {
    graphId: number;
  };

  type getNeighborsParams = {
    taskId?: string;
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
    taskId?: string;
    graphId: number;
    nodeId: number;
    distance: number;
    min: number;
  };

  type SearchRequest = {
    taskId?: string;
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
    id: string;
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
    taskId?: string;
    graph: string;
    desc?: string;
    nodeId: string;
    edgeId: string;
    graphId?: number;
  };
}
