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
	EdgeDirection bool    `json:"edgeDirection,optional"`
	Display       string  `json:"display"`
	Primary       string  `json:"primary,optional"`
	Attrs         []*Attr `json:"attrs,optional"`
}

type Attr struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
	Type int64  `json:"type"`
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
	ParentID     int64        `json:"parentId"`
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

type GetMatchNodesByTagRequest struct {
	GraphID  int64  `form:"graphId"`
	Keywords string `form:"keywords"`
	NodeID   int64  `form:"nodeId"`
}

type GetMatchNodesReply struct {
	Base           *BaseReply        `json:"base"`
	MatchNodePacks []*MatchNodePacks `json:"matchNodePacks"`
}

type GetMatchNodesByTagReply struct {
	Base       *BaseReply   `json:"base"`
	MatchNodes []*MatchNode `json:"matchNodes"`
}

type GetNodesReply struct {
	Base      *BaseReply  `json:"base"`
	NodePacks []*NodePack `json:"nodePacks"`
}

type GetNodesByTagReply struct {
	Base  *BaseReply `json:"base"`
	Nodes []*Node    `json:"nodes"`
}

type CreateGroupRequest struct {
	Name         string       `json:"name"`
	Desc         string       `json:"desc"`
	ParentID     int64        `json:"parentId"`
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

type AlgoTask struct {
	Id         int64  `json:"id"`
	CreateTime int64  `json:"createTime"`
	UpdateTime int64  `json:"updateTime"`
	Status     int64  `json:"status"`
	GraphID    int64  `json:"graphId"`
	Req        string `json:"req"`
	AlgoID     int64  `json:"algoId"`
	Output     string `json:"output"`
}

type GetAlgoTasksRequest struct {
	GraphID int64 `form:"graphId,optional"`
}

type GetAlgoTaskRequest struct {
	ID int64 `form:"id"`
}

type GetAlgoTaskReply struct {
	Base *BaseReply `json:"base"`
	Task *AlgoTask  `json:"task"`
}

type GetAlgoTasksReply struct {
	Base  *BaseReply  `json:"base"`
	Tasks []*AlgoTask `json:"tasks"`
}

type DropAlgoTaskRequest struct {
	ID int64 `json:"id"`
}

type Algo struct {
	Id      int64        `json:"id,optional"`
	Name    string       `json:"name"`
	Define  string       `json:"define"`
	Detail  string       `json:"detail"`
	GroupId int64        `json:"groupId"`
	Tag     string       `json:"tag"`
	TagID   int64        `json:"tagId"`
	Params  []*AlgoParam `json:"params,optional"`
}

type AlgoParam struct {
	Key       string `json:"key"`
	KeyDesc   string `json:"keyDesc"`
	Type      int64  `json:"type"`
	InitValue string `json:"initValue,optional"`
	Max       string `json:"max,optional"`
	Min       string `json:"min,optional"`
}

type AlgoReply struct {
	Base   *BaseReply `json:"base"`
	AlgoID int64      `json:"algoId"`
	File   string     `json:"file"`
}

type GetAlgoReply struct {
	Base  *BaseReply `json:"base"`
	Algos []*Algo    `json:"algos"`
}

type GetAlgoDocReply struct {
	Base *BaseReply `json:"base"`
	Doc  string     `json:"doc"`
}

type AlgoIDRequest struct {
	AlgoID int64 `form:"algoId"`
}

type CreateAlgoRequest struct {
	Algo       *Algo  `json:"algo"`
	EntryPoint string `json:"entryPoint"`
	Jar        string `json:"jar"`
}

type Param struct {
	Key       string   `json:"key"`
	Type      int64    `json:"type"`
	Value     string   `json:"value,optional"`
	ListValue []string `json:"listValue,optional"`
}

type ExecAlgoRequest struct {
	GraphID int64    `json:"graphId"`
	AlgoID  int64    `json:"algoId"`
	Params  []*Param `json:"params,optional"`
}

type HotSE struct {
	Artifact string  `json:"artifact"`
	Version  string  `json:"version"`
	HomePage string  `json:"homePage"`
	Score    float64 `json:"score"`
}

type HotSETopic struct {
	Software   []*HotSE `json:"software"`
	Language   string   `json:"language"`
	Topic      string   `json:"topic"`
	UpdateTime int64    `json:"updateTime"`
}

type GetHotSEReply struct {
	Topics []*HotSETopic `json:"topics"`
}

type HHI struct {
	Name  string  `json:"name"`
	Score float64 `json:"score"`
}

type HHILanguage struct {
	HHIs       []*HHI `json:"hhIs"`
	Language   string `json:"language"`
	UpdateTime int64  `json:"updateTime"`
}

type GetHHIReply struct {
	Languages []*HHILanguage `json:"languages"`
}
