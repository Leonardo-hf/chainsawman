// Code generated by goctl. DO NOT EDIT.
package types

type BaseReply struct {
	Status     int64             `json:"status"`
	Msg        string            `json:"msg"`
	TaskID     string            `json:"taskId"`
	TaskStatus int64             `json:"taskStatus"`
	Extra      map[string]string `json:"extra"`
}

type Graph struct {
	Id       int64  `json:"id"`
	Status   int64  `json:"status"`
	GroupID  int64  `json:"groupId"`
	Name     string `json:"name"`
	Desc     string `json:"desc"`
	NumNode  int64  `json:"numNode"`
	NumEdge  int64  `json:"numEdge"`
	CreatAt  int64  `json:"creatAt"`
	UpdateAt int64  `json:"updateAt"`
}

type Structure struct {
	Id            int64   `json:"id"`
	Name          string  `json:"name"`
	Desc          string  `json:"desc"`
	EdgeDirection bool    `json:"edgeDirection"`
	Display       string  `json:"display"`
	Attrs         []*Attr `json:"attrs,optional"`
}

type Attr struct {
	Name    string `json:"name"`
	Desc    string `json:"desc"`
	Primary bool   `json:"primary"`
	Type    int64  `json:"type"`
}

type Pair struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Node struct {
	Id    int64   `json:"id"`
	Deg   int64   `json:"deg"`
	Attrs []*Pair `json:"attrs"`
}

type MatchNode struct {
	Id          int64  `json:"id"`
	PrimaryAttr string `json:"primaryAttr"`
}

type NodePack struct {
	Tag   string  `json:"tag"`
	Nodes []*Node `json:"nodes"`
}

type MatchNodePacks struct {
	Tag   string       `json:"tag"`
	Match []*MatchNode `json:"match"`
}

type Edge struct {
	Source int64   `json:"source"`
	Target int64   `json:"target"`
	Attrs  []*Pair `json:"attrs"`
}

type EdgePack struct {
	Tag   string  `json:"tag"`
	Edges []*Edge `json:"edges"`
}

type Group struct {
	Id           int64        `json:"id"`
	Name         string       `json:"name"`
	Desc         string       `json:"desc"`
	NodeTypeList []*Structure `json:"nodeTypeList"`
	EdgeTypeList []*Structure `json:"edgeTypeList"`
	Graphs       []*Graph     `json:"graphs"`
}

type GetAllGraphReply struct {
	Base   *BaseReply `json:"base"`
	Groups []*Group   `json:"groups"`
}

type GetGraphDetailRequest struct {
	TaskID  string `form:"taskId,optional"`
	GraphID int64  `form:"graphId"`
	Top     int64  `form:"top"`
	Max     int64  `form:"max,default=2000"`
}

type GetGraphDetailReply struct {
	Base      *BaseReply  `json:"base"`
	NodePacks []*NodePack `json:"nodePacks"`
	EdgePacks []*EdgePack `json:"edgePacks"`
}

type GetNeighborsRequest struct {
	TaskID    string `form:"taskId,optional"`
	GraphID   int64  `form:"graphId"`
	NodeID    int64  `form:"nodeId"`
	Direction string `form:"direction"`
	Max       int64  `form:"max,default=2000"`
}

type DropGraphRequest struct {
	GraphID int64 `json:"graphId"`
}

type CreateGraphRequest struct {
	TaskID  string `json:"taskId,optional"`
	GraphID int64  `json:"graphId,optional"`
	Graph   string `json:"graph"`
	Desc    string `json:"desc,optional"`
	GroupID int64  `json:"groupId"`
}

type UpdateGraphRequest struct {
	TaskID       string  `json:"taskId,optional"`
	GraphID      int64   `json:"graphId"`
	NodeFileList []*Pair `json:"nodeFileList"`
	EdgeFileList []*Pair `json:"edgeFileList"`
}

type GraphInfoReply struct {
	Base  *BaseReply `json:"base"`
	Graph *Graph     `json:"graph"`
}

type GetGraphInfoRequest struct {
	Name string `form:"name"`
}

type GetNodesRequest struct {
	TaskID  string `form:"taskId,optional"`
	GraphID int64  `form:"graphId"`
}

type GetMatchNodesRequest struct {
	GraphID  int64  `form:"graphId"`
	Keywords string `form:"keywords"`
}

type GetMatchNodesReply struct {
	Base           *BaseReply        `json:"base"`
	MatchNodePacks []*MatchNodePacks `json:"matchNodePacks"`
}

type GetNodesReply struct {
	Base      *BaseReply  `json:"base"`
	NodePacks []*NodePack `json:"nodePacks"`
}

type CreateGroupRequest struct {
	Name         string       `json:"name"`
	Desc         string       `json:"desc"`
	NodeTypeList []*Structure `json:"nodeTypeList"`
	EdgeTypeList []*Structure `json:"edgeTypeList"`
}

type GroupInfoReply struct {
	Base  *BaseReply `json:"base"`
	Group *Group     `json:"group"`
}

type DropGroupRequest struct {
	GroupID int64 `json:"groupId"`
}

type PresignedRequest struct {
	Filename string `form:"filename"`
}

type PresignedReply struct {
	Url      string `json:"url"`
	Filename string `json:"filename"`
}

type Task struct {
	Id         string `json:"id"`
	Idf        string `json:"idf"`
	CreateTime int64  `json:"createTime"`
	UpdateTime int64  `json:"updateTime"`
	Status     int64  `json:"status"`
	Req        string `json:"req"`
	Res        string `json:"res"`
}

type GetTasksRequest struct {
	GraphID int64 `form:"graphId"`
}

type GetTasksReply struct {
	Base  *BaseReply `json:"base"`
	Tasks []*Task    `json:"tasks"`
}

type DropTaskRequest struct {
	TaskID string `json:"taskId,optional"`
}

type Rank struct {
	Tag   string  `json:"tag"`
	Node  *Node   `json:"node"`
	Score float64 `json:"score"`
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

type AlgoRequest struct {
	TaskID  string `form:"taskId,optional"`
	GraphID int64  `form:"graphId"`
}

type AlgoDegreeRequest struct {
	TaskID  string `form:"taskId,optional"`
	GraphID int64  `form:"graphId"`
}

type AlgoPageRankRequest struct {
	TaskID  string  `form:"taskId,optional"`
	GraphID int64   `form:"graphId"`
	Iter    int64   `form:"iter,default=3"`
	Prob    float64 `form:"prob,default=0.85"`
}

type AlgoVoteRankRequest struct {
	TaskID  string `form:"taskId,optional"`
	GraphID int64  `form:"graphId"`
	Iter    int64  `form:"iter,default=100"`
}

type AlgoLouvainRequest struct {
	TaskID       string  `form:"taskId,optional"`
	GraphID      int64   `form:"graphId"`
	MaxIter      int64   `form:"maxIter,default=10"`
	InternalIter int64   `form:"internalIter,default=5"`
	Tol          float64 `form:"tol,default=0.5"`
}
