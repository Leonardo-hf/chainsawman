declare namespace API {
  type BaseReply = {
    status: number;
    msg: string;
    extra: Record<string, any>;
  };

  type DropRequest = {
    graph: string;
  };

  type Edge = {
    source: string;
    target: string;
  };

  type getGraphParams = {
    graph: string;
  };

  type Graph = {
    name: string;
    desc: string;
    nodes: number;
    edges: number;
  };

  type Node = {
    name: string;
    desc: string;
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

  type SearchRequest = {
    graph: string;
  };
}
