package common

const (
	GraphCreate    = "graph:create"
	GraphUpdate    = "graph:update"
	GraphGet       = "graph:get"
	GraphNodes     = "graph:nodes"
	GraphNeighbors = "graph:neighbors"
	CronPython     = "cron:python"
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

type PreparedGraphInfo struct {
	GroupID int64
	GraphID int64
	Name    string
}

const (
	SoftwareDepsGroup = 3
	PythonDepsGraph   = 1
)

var PreparedGraph = map[int64]*PreparedGraphInfo{
	PythonDepsGraph: {
		GraphID: PythonDepsGraph,
		GroupID: SoftwareDepsGroup,
		Name:    "Python软件依赖图谱",
	},
}

type TaskAttr struct {
	Queue      string
	Cron       string
	FixedGraph *PreparedGraphInfo
}

func TaskIdf(idf string) *TaskAttr {
	return taskAttrMap[idf]
}

var taskAttrMap = map[string]*TaskAttr{
	GraphCreate: {
		Queue: PLow,
	},
	GraphUpdate: {
		Queue: PLow,
	},
	GraphNodes: {
		Queue: PHigh,
	},
	GraphGet: {
		Queue: PHigh,
	},
	GraphNeighbors: {
		Queue: PHigh,
	},
	CronPython: {
		Queue:      PMedium,
		FixedGraph: PreparedGraph[PythonDepsGraph],
		Cron:       "0 * * * *",
	},
}
