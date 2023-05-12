package common

const (
	GraphCreate    = "graph:create"
	GraphGet       = "graph:get"
	GraphNeighbors = "graph:neighbors"

	AlgoDegree      = "algo:degree"
	AlgoPagerank    = "algo:pagerank"
	AlgoBetweenness = "algo:between"
	AlgoCloseness   = "algo:close"
	AlgoVoterank    = "algo:voterank"
	AlgoLouvain     = "algo:louvain"
	AlgoAvgCC       = "algo:avgCC"
)

const (
	PHigh   = "high"
	PMedium = "medium"
	PLow    = "low"
)

const (
	TaskMqEd  = "task_ed"
	TaskMqEd2 = "task_ed2"
)

type TaskAttr struct {
	Queue      string
	Persistent bool
}

func TaskIdf(idf string) *TaskAttr {
	return taskAttrMap[idf]
}

var taskAttrMap = map[string]*TaskAttr{
	GraphCreate: {
		Queue:      PMedium,
		Persistent: false,
	},
	GraphGet: {
		Queue: PHigh, Persistent: false,
	},
	GraphNeighbors: {
		Queue:      PHigh,
		Persistent: false,
	},
	AlgoDegree: {
		Queue:      PLow,
		Persistent: true,
	},
	AlgoPagerank: {
		Queue:      PLow,
		Persistent: true,
	},
	AlgoBetweenness: {
		Queue:      PLow,
		Persistent: true,
	},
	AlgoCloseness: {
		Queue:      PLow,
		Persistent: true,
	},
	AlgoVoterank: {
		Queue:      PLow,
		Persistent: true,
	},
	AlgoLouvain: {
		Queue:      PLow,
		Persistent: true,
	},
	AlgoAvgCC: {
		Queue:      PLow,
		Persistent: true,
	},
}
