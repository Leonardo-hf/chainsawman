package model

type NebulaNode struct {
	ID   int64
	Name string
	Desc string
	Deg  int64
}

type NebulaEdge struct {
	Source int64
	Target int64
}
