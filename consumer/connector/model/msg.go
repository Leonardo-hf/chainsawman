package model

type Opt int

const (
	Updates Opt = iota + 1
	Creates
	Deletes
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

type EdgesBody struct {
	Edges []Edge `json:"nodes"`
}

type Edge struct {
	Source int64   `json:"source"`
	Target []int64 `json:"target"`
}
