package model

type Node struct {
	Name string
	Desc string
	Deg  int64
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
