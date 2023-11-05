package common

const (
	GraphCreate    = "graph:create"
	GraphUpdate    = "graph:update"
	GraphGet       = "graph:get"
	GraphNodes     = "graph:nodes"
	GraphNeighbors = "graph:neighbors"

	AlgoExec = "algo:exec"
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
		Queue:      PHigh,
		Persistent: false,
	},
	GraphNeighbors: {
		Queue:      PHigh,
		Persistent: false,
	},
	AlgoExec: {
		Queue:      PLow,
		Persistent: true,
	},
}
