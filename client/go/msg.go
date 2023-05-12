package _go

type Opt int

const (
	Updates Opt = iota + 1
	Creates
	Deletes
	InsertEdges
	DeleteEdges
)

type Msg struct {
	Opt     Opt    `json:"opt"`
	GraphID int64  `json:"graph_id"`
	Body    string `json:"body"`
}

type NodesBody struct {
	Nodes []Node `json:"nodes"`
}

type Node struct {
	NodeID int64  `json:"node_id"`
	Name   string `json:"name"`
	Desc   string `json:"desc"`
}

type SplitEdgesBody struct {
	Edges [][2]int64 `json:"edges"`
}

type EdgesBody struct {
	Edges []Edge `json:"nodes"`
}

type Edge struct {
	Source int64   `json:"source"`
	Target []int64 `json:"target"`
}
