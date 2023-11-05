declare namespace Graph {
  type Algo = {
    id?: number;
    name: string;
    desc: string;
    groupId: number;
    isCustom: boolean;
    type: number;
    params?: AlgoParam[];
  };

  type AlgoParam = {
    key: string;
    keyDesc: string;
    type: number;
    initValue?: number;
    max?: number;
    min?: number;
  };

  type AlgoReply = {
    base: BaseReply;
    file: string;
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

  type CreateAlgoRequest = {
    algo: Algo;
    entryPoint: string;
    jar: string;
  };

  type CreateGraphRequest = {
    taskId?: string;
    graphId?: number;
    graph: string;
    groupId: number;
  };

  type CreateGroupRequest = {
    name: string;
    desc: string;
    parentId: number;
    nodeTypeList: Structure[];
    edgeTypeList: Structure[];
  };

  type DropAlgoRequest = {
    algoId: number;
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

  type ExecAlgoRequest = {
    taskId?: string;
    graphId: number;
    algoId: number;
    params?: Param[];
  };

  type fileAlgoGetPresignedParams = {
    filename: string;
  };

  type GetAlgoReply = {
    base: BaseReply;
    algos: Algo[];
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

  type getMatchNodesByTagParams = {
    graphId: number;
    keywords: string;
    nodeId: number;
  };

  type GetMatchNodesByTagReply = {
    base: BaseReply;
    matchNodes: MatchNode[];
  };

  type GetMatchNodesByTagRequest = {
    graphId: number;
    keywords: string;
    nodeId: number;
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

  type GetNodesByTagReply = {
    base: BaseReply;
    nodes: Node[];
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
    parentId: number;
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

  type Param = {
    key: string;
    type: number;
    value?: string;
    listValue?: string[];
    algoValue?: string;
  };

  type PresignedReply = {
    url: string;
    filename: string;
  };

  type PresignedRequest = {
    filename: string;
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
