package common

import (
	"chainsawman/consumer/types"

	"github.com/zeromicro/go-zero/core/jsonx"
)

// TODO: 事实上是graph和consuemr的共用类，存在types的依赖问题
type TaskIdf int

const (
	GraphCreate = iota
	GraphGet
	GraphNeighbors

	AlgoDegree TaskIdf = iota + 2<<10
	AlgoPagerank
	AlgoBetweenness
	AlgoCloseness
	AlgoVoterank
	AlgoLouvain
	AlgoAvgCC
)

func (t TaskIdf) Visible() bool {
	return t > (2 << 10)
}

func (t TaskIdf) Desc() string {
	switch t {
	case GraphCreate:
		return "创建图"
	case GraphGet:
		return "查询图"
	case GraphNeighbors:
		return "查询节点邻居"
	case AlgoDegree:
		return "度中心度"
	case AlgoPagerank:
		return "PageRank"
	case AlgoAvgCC:
		return "平均聚类系数"
	case AlgoCloseness:
		return "紧密中心度"
	case AlgoVoterank:
		return "VoteRank"
	case AlgoLouvain:
		return "Louvain聚类"
	case AlgoBetweenness:
		return "介数中心度"
	default:
		return ""
	}
}

func (t TaskIdf) Params(s string) interface{} {
	if len(s) == 0 {
		return "{}"
	}
	switch t {
	case GraphCreate:
		p := &types.UploadRequest{}
		_ = jsonx.UnmarshalFromString(s, p)
		return p
	case GraphGet:
		p := &types.SearchRequest{}
		_ = jsonx.UnmarshalFromString(s, p)
		return p
	case GraphNeighbors:
		p := &types.SearchNodeRequest{}
		_ = jsonx.UnmarshalFromString(s, p)
		return p
	case AlgoDegree:
		p := &types.AlgoDegreeRequest{}
		_ = jsonx.UnmarshalFromString(s, p)
		return p
	case AlgoPagerank:
		p := &types.AlgoPageRankRequest{}
		_ = jsonx.UnmarshalFromString(s, p)
		return p
	case AlgoAvgCC, AlgoCloseness, AlgoBetweenness:
		p := &types.AlgoRequest{}
		_ = jsonx.UnmarshalFromString(s, p)
		return p
	case AlgoVoterank:
		p := &types.AlgoVoteRankRequest{}
		_ = jsonx.UnmarshalFromString(s, p)
		return p
	case AlgoLouvain:
		p := &types.AlgoLouvainRequest{}
		_ = jsonx.UnmarshalFromString(s, p)
		return p
	default:
		return "{}"
	}
}

func (t TaskIdf) Res(s string) interface{} {
	if len(s) == 0 {
		return "{}"
	}
	switch t {
	case GraphCreate:
		p := &types.SearchGraphReply{}
		_ = jsonx.UnmarshalFromString(s, p)
		return p
	case GraphGet:
		p := &types.SearchGraphDetailReply{}
		_ = jsonx.UnmarshalFromString(s, p)
		return p
	case GraphNeighbors:
		p := &types.SearchNodeReply{}
		_ = jsonx.UnmarshalFromString(s, p)
		return p
	case AlgoDegree, AlgoPagerank, AlgoBetweenness, AlgoCloseness, AlgoLouvain, AlgoVoterank:
		p := &types.AlgoRankReply{}
		_ = jsonx.UnmarshalFromString(s, p)
		return p
	case AlgoAvgCC:
		p := &types.AlgoMetricReply{}
		_ = jsonx.UnmarshalFromString(s, p)
		return p
	default:
		return "{}"
	}
}

func Btoi(b bool) int64 {
	if b {
		return 1
	}
	return 0
}
