declare namespace Graph {
  type Algo = {
    id?: number;
    name: string;
    detail: string;
    groupId: number;
    tag: string;
    params?: AlgoParam[];
  };

  type AlgoParam = {
    key: string;
    keyDesc: string;
    type: number;
    initValue?: string;
    max?: string;
    min?: string;
  };

  type AlgoReply = {
    base: BaseReply;
    algoId: number;
    file: string;
  };

  type AlgoTask = {
    id: number;
    createTime: number;
    updateTime: number;
    status: number;
    graphId: number;
    req: string;
    algoId: number;
    output: string;
  };

  type Attr = {
    name: string;
    desc: string;
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

  type DropAlgoTaskRequest = {
    id: number;
  };

  type DropGraphRequest = {
    graphId: number;
  };

  type DropGroupRequest = {
    groupId: number;
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

  type getAlgoTaskByIDParams = {
    id: number;
  };

  type getAlgoTaskParams = {
    graphId?: number;
  };

  type GetAlgoTaskReply = {
    base: BaseReply;
    task: AlgoTask;
  };

  type GetAlgoTaskRequest = {
    id: number;
  };

  type GetAlgoTasksReply = {
    base: BaseReply;
    tasks: AlgoTask[];
  };

  type GetAlgoTasksRequest = {
    graphId?: number;
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

  type GetHHIReply = {
    languages: HHILanguage[];
  };

  type GetHotSEReply = {
    topics: HotSETopic[];
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

  type HHI = {
    name: string;
    score: number;
  };

  type HHILanguage = {
    hhIs: HHI[];
    language: string;
    updateTime: number;
  };

  type HotSE = {
    artifact: string;
    version: string;
    homePage: string;
    score: number;
  };

  type HotSETopic = {
    software: HotSE[];
    language: string;
    topic: string;
    updateTime: number;
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
    edgeDirection?: boolean;
    display: string;
    primary?: string;
    attrs?: Attr[];
  };

  type UpdateGraphRequest = {
    taskId?: string;
    graphId: number;
    nodeFileList: Pair[];
    edgeFileList: Pair[];
  };
}
