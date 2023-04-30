// Code generated by goctl. DO NOT EDIT.
package types

type Graph struct {
	Id     int64  `json:"id"`
	Status int    `json:"status"`
	Name   string `json:"name"`
	Desc   string `json:"desc"`
	Nodes  int64  `json:"nodes"`
	Edges  int64  `json:"edges"`
}

type Node struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
	Deg  int64  `json:"deg"`
}

type Edge struct {
	Source int64 `json:"source"`
	Target int64 `json:"target"`
}

type Param struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Task struct {
	Idf        int64       `json:"idf"`
	Desc       string      `json:"desc"`
	CreateTime int64       `json:"createTime"`
	UpdateTime int64       `json:"updateTime"`
	Status     int64       `json:"status"`
	Req        interface{} `json:"req"`
	Res        interface{} `json:"res"`
}

type Rank struct {
	NodeID int64   `json:"nodeId"`
	Score  float64 `json:"score"`
}

type BaseReply struct {
	Status     int64             `json:"status"`
	Msg        string            `json:"msg"`
	TaskID     int64             `json:"taskId"`
	TaskStatus int64             `json:"taskStatus"`
	Extra      map[string]string `json:"extra"`
}

type SearchTasksReply struct {
	Base  *BaseReply `json:"base"`
	Tasks []*Task    `json:"tasks"`
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

type AlgoRankReply struct {
	Base  *BaseReply `json:"base"`
	Ranks []*Rank    `json:"ranks"`
	File  string     `json:"file"`
}

type AlgoMetricReply struct {
	Base  *BaseReply `json:"base"`
	Score float64    `json:"score"`
}

type SearchTasksRequest struct {
	GraphID int64 `form:"graphId"`
}

type SearchRequest struct {
	TaskID  int64 `form:"taskId,optional"`
	GraphID int64 `form:"graphId"`
	Min     int64 `form:"min"`
}

type SearchNodeRequest struct {
	TaskID   int64 `form:"taskId,optional"`
	GraphID  int64 `form:"graphId"`
	NodeID   int64 `form:"nodeId"`
	Distance int64 `form:"distance"`
	Min      int64 `form:"min"`
}

type DropRequest struct {
	GraphID int64 `form:"graphId"`
}

type UploadRequest struct {
	TaskID  int64  `form:"taskId,optional"`
	Graph   string `form:"graph"`
	Desc    string `form:"desc,optional"`
	NodeID  string `form:"nodeId"`
	EdgeID  string `form:"edgeId"`
	GraphID int64  `form:"graphId,optional"`
}

type AlgoRequest struct {
	TaskID  int64 `form:"taskId"`
	GraphID int64 `form:"graphId"`
}

type AlgoDegreeRequest struct {
	TaskID  int64 `form:"taskId"`
	GraphID int64 `form:"graphId"`
}

type AlgoPageRankRequest struct {
	TaskID  int64   `form:"taskId"`
	GraphID int64   `form:"graphId"`
	Iter    int64   `form:"iter,default=3"`
	Prob    float64 `form:"prob,default=0.85"`
}

type AlgoVoteRankRequest struct {
	TaskID  int64 `form:"taskId"`
	GraphID int64 `form:"graphId"`
	Iter    int64 `form:"iter,default=100"`
}

type AlgoLouvainRequest struct {
	TaskID       int64   `form:"taskId"`
	GraphID      int64   `form:"graphId"`
	MaxIter      int64   `form:"maxIter,default=10"`
	InternalIter int64   `form:"internalIter,default=5"`
	Tol          float64 `form:"tol,default=0.5"`
}
