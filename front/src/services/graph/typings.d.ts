declare namespace Graph {
    type BaseReply = {
        status: number;
        msg: string;
        taskId: number;
        taskStatus: number;
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
        taskId: number;
        graph: string;
        min: number;
    };

    type getNeighborsParams = {
        taskId: number;
        graph: string;
        node: string;
        distance: number;
        min: number;
    };

    type Graph = {
        id: number
        name: string;
        desc: string;
        nodes: number;
        edges: number;
        status: number
    };

    type Node = {
        name: string;
        desc: string;
        deg: number;
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
        taskId: number;
        graph: string;
        node: string;
        distance: number;
        min: number;
    };

    type SearchRequest = {
        taskId: number;
        graph: string;
        min: number;
    };

    type UploadRequest = {
        taskId: number;
        graph: string;
        desc: string
        nodeId: string;
        edgeId: string;
    };
}
