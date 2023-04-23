package db

// NebulaClient TODO: 删掉用不上的接口
type NebulaClient interface {
	DropGraph(graph int64) error
}
