package db

type NebulaClient interface {
	CreateGraph(graph string) error

	InsertNode(graph string, node *Node) (int, error)

	MultiInsertNodes(graph string, nodes []*Node) (int, error)

	InsertEdge(graph string, edge *Edge) (int, error)

	MultiInsertEdges(graph string, edges []*Edge) (int, error)

	GetGraphs() ([]*Graph, error)

	GetNodes(graph string) ([]*Node, error)

	GetEdges(graph string) ([]*Edge, error)
}
