package common

const (
	GraphCreate    = "graph:create"
	GraphUpdate    = "graph:update"
	GraphGet       = "graph:get"
	GraphNodes     = "graph:nodes"
	GraphNeighbors = "graph:neighbors"

	AlgoDegree      = "algo:degree"
	AlgoPagerank    = "algo:pagerank"
	AlgoBetweenness = "algo:between"
	AlgoCloseness   = "algo:close"
	AlgoLouvain     = "algo:louvain"
	AlgoAvgCC       = "algo:avgCC"
	AlgoQuantity    = "algo:quantity"
	AlgoDepth       = "algo:depth"
	AlgoIntegration = "algo:integration"
	AlgoEcology     = "algo:ecology"

	AlgoComp = "algo:comp"
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
	GraphUpdate: {
		Queue:      PMedium,
		Persistent: false,
	},
	GraphNodes: {
		Queue:      PLow,
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
	AlgoLouvain: {
		Queue:      PLow,
		Persistent: true,
	},
	AlgoAvgCC: {
		Queue:      PLow,
		Persistent: true,
	},
	AlgoComp: {
		Queue:      PLow,
		Persistent: true,
	},
	AlgoQuantity: {
		Queue:      PLow,
		Persistent: true,
	},
	AlgoDepth: {
		Queue:      PLow,
		Persistent: true,
	},
	AlgoIntegration: {
		Queue:      PLow,
		Persistent: true,
	},
	AlgoEcology: {
		Queue:      PLow,
		Persistent: true,
	},
}
