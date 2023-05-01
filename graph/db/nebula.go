package db

type NebulaClient interface {
	DropGraph(graph int64) error

	CreateGraph(graph int64) error
}
