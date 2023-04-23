package model

type Node struct {
	ID   int64
	Name string
	Desc string
	Deg  int64
}

type Edge struct {
	Source int64
	Target int64
}
