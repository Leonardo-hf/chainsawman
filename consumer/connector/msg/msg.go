package msg

type OptFlag int

const (
	Create OptFlag = iota + 1
	Update
	Delete
)

type EntityFlag int

const (
	Node EntityFlag = iota + 1
	Edge
)

type Msg struct {
	Opt    OptFlag
	Entity EntityFlag
	Body   string
}

type UpdateBody struct {
	GraphId int64
	Edges   map[int64][]int64
}

type EdgeBody struct {
	Source int64
	Target int64
}

type NodeBody struct {
	ID   int64
	Name string
	Desc string
}
