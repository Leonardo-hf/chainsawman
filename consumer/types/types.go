// Code generated by goctl. DO NOT EDIT.
package types

type Graph struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Desc  string `json:"desc"`
	Nodes int64  `json:"nodes"`
	Edges int64  `json:"edges"`
}

type Node struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
	Deg  int64  `json:"deg"`
}

type Edge struct {
	Source string `json:"source"`
	Target string `json:"target"`
}

type BaseReply struct {
	Status     int64             `json:"status"`
	Msg        string            `json:"msg"`
	TaskID     int64             `json:"taskId"`
	TaskStatus int64             `json:"taskStatus"`
	Extra      map[string]string `json:"extra"`
}

type SearchAllGraphReply struct {
	Base   *BaseReply `json:"base"`
	Graphs []*Graph   `json:"graphs"`
}

type SearchGraphReply struct {
	Base  *BaseReply `json:"base"`
	Graph *Graph     `json:"graph"`
}

type SearchNodeReply struct {
	Base  *BaseReply `json:"base"`
	Info  *Node      `json:"node"`
	Nodes []*Node    `json:"nodes"`
	Edges []*Edge    `json:"edges"`
}

type SearchGraphDetailReply struct {
	Base  *BaseReply `json:"base"`
	Nodes []*Node    `json:"nodes"`
	Edges []*Edge    `json:"edges"`
}

type SearchRequest struct {
	TaskID int64  `form:"taskId"`
	Graph  string `form:"graph"`
	Min    int64  `form:"min"`
}

type SearchNodeRequest struct {
	TaskID   int64  `form:"taskId"`
	Graph    string `form:"graph"`
	Node     string `form:"node"`
	Distance int64  `form:"distance"`
	Min      int64  `form:"min"`
}

type DropRequest struct {
	Graph string `form:"graph"`
}

type UploadRequest struct {
	TaskID  int64  `form:"taskId"`
	GraphId int    `form:"graphId"`
	Graph   string `form:"graph"`
	Desc    string `form:"description"`
	NodeID  string `form:"nodeId"`
	EdgeID  string `form:"edgeId"`
}
