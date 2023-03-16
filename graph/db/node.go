package db

type Node struct {
	Name string
	Desc string
}

type Edge struct {
	Source string
	Target string
}

type Graph struct {
	Name  string `json:"name"`
	Nodes int64  `json:"nodes"`
	Edges int64  `json:"edges"`
}
