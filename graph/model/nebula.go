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
