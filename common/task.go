package common

const (
	GraphCreate    = "graph:create"
	GraphUpdate    = "graph:update"
	GraphGet       = "graph:get"
	GraphNodes     = "graph:nodes"
	GraphNeighbors = "graph:neighbors"
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
	Queue string
}

func TaskIdf(idf string) *TaskAttr {
	return taskAttrMap[idf]
}

var taskAttrMap = map[string]*TaskAttr{
	GraphCreate: {
		Queue: PMedium,
	},
	GraphUpdate: {
		Queue: PMedium,
	},
	GraphNodes: {
		Queue: PLow,
	},
	GraphGet: {
		Queue: PHigh,
	},
	GraphNeighbors: {
		Queue: PHigh,
	},
}
