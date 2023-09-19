declare namespace Graph {
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

  type algoDepthParams = {
    taskId?: string;
    graphId: number;
  };

  type algoEcologyParams = {
    taskId?: string;
    graphId: number;
  };

  type algoIntegrationParams = {
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

  type algoQuantityParams = {
    taskId?: string;
    graphId: number;
    iter?: number;
  };

  type AlgoRankReply = {
    base: BaseReply;
    ranks: Rank[];
    file: string;
  };

  type AlgoRequest = {
    taskId?: string;
    graphId: number;
  };

  type AlgoVoteRankRequest = {
    taskId?: string;
    graphId: number;
    iter: number;
  };

  type Attr = {
    name: string;
    desc: string;
    primary: boolean;
    type: number;
  };

  type BaseReply = {
    status: number;
    msg: string;
    taskId: string;
    taskStatus: number;
    extra: Record<string, any>;
  };

  type CreateGraphRequest = {
    taskId?: string;
    graphId?: number;
    graph: string;
    desc?: string;
    groupId: number;
  };

  type CreateGroupRequest = {
    name: string;
    desc: string;
    nodeTypeList: Structure[];
    edgeTypeList: Structure[];
  };

  type DropGraphRequest = {
    graphId: number;
  };

  type DropGroupRequest = {
    groupId: number;
  };

  type DropTaskRequest = {
    taskId?: string;
  };

  type Edge = {
    source: number;
    target: number;
    attrs: Pair[];
  };

  type EdgePack = {
    tag: string;
    edges: Edge[];
  };

  type fileGetPresignedParams = {
    filename: string;
  };

  type GetAllGraphReply = {
    base: BaseReply;
    groups: Group[];
  };

  type GetGraphDetailReply = {
    base: BaseReply;
    nodePacks: NodePack[];
    edgePacks: EdgePack[];
  };

  type GetGraphDetailRequest = {
    taskId?: string;
    graphId: number;
    top: number;
    max: number;
  };

  type getGraphInfoParams = {
    name: string;
  };

  type GetGraphInfoRequest = {
    name: string;
  };

  type getGraphParams = {
    taskId?: string;
    graphId: number;
    top: number;
    max?: number;
  };

  type getGraphTasksParams = {
    graphId: number;
  };

  type getMatchNodesParams = {
    graphId: number;
    keywords: string;
  };

  type GetMatchNodesReply = {
    base: BaseReply;
    matchNodePacks: MatchNodePacks[];
  };

  type GetMatchNodesRequest = {
    graphId: number;
    keywords: string;
  };

  type getNeighborsParams = {
    taskId?: string;
    graphId: number;
    nodeId: number;
    direction: string;
    max?: number;
  };

  type GetNeighborsRequest = {
    taskId?: string;
    graphId: number;
    nodeId: number;
    direction: string;
    max: number;
  };

  type getNodesParams = {
    taskId?: string;
    graphId: number;
  };

  type GetNodesReply = {
    base: BaseReply;
    nodePacks: NodePack[];
  };

  type GetNodesRequest = {
    taskId?: string;
    graphId: number;
  };

  type GetTasksReply = {
    base: BaseReply;
    tasks: Task[];
  };

  type GetTasksRequest = {
    graphId: number;
  };

  type Graph = {
    id: number;
    status: number;
    groupId: number;
    name: string;
    desc: string;
    numNode: number;
    numEdge: number;
    creatAt: number;
    updateAt: number;
  };

  type GraphInfoReply = {
    base: BaseReply;
    graph: Graph;
  };

  type Group = {
    id: number;
    name: string;
    desc: string;
    nodeTypeList: Structure[];
    edgeTypeList: Structure[];
    graphs: Graph[];
  };

  type GroupInfoReply = {
    base: BaseReply;
    group: Group;
  };

  type MatchNode = {
    id: number;
    primaryAttr: string;
  };

  type MatchNodePacks = {
    tag: string;
    match: MatchNode[];
  };

  type Node = {
    id: number;
    deg: number;
    attrs: Pair[];
  };

  type NodePack = {
    tag: string;
    nodes: Node[];
  };

  type Pair = {
    key: string;
    value: string;
  };

  type PresignedReply = {
    url: string;
    filename: string;
  };

  type PresignedRequest = {
    filename: string;
  };

  type Rank = {
    tag: string;
    node: Node;
    score: number;
  };

  type Structure = {
    id: number;
    name: string;
    desc: string;
    edgeDirection: boolean;
    display: string;
    attrs?: Attr[];
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

  type UpdateGraphRequest = {
    taskId?: string;
    graphId: number;
    nodeFileList: Pair[];
    edgeFileList: Pair[];
  };
}
