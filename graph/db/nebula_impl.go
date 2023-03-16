package db

import (
	"fmt"
	nebula "github.com/vesoft-inc/nebula-go/v3"
)

type NebulaClientImpl struct {
	Pool     *nebula.ConnectionPool
	Username string
	Password string
}

func (n *NebulaClientImpl) CreateGraph(graph string) error {
	session, err := n.getSession()
	if err != nil {
		return err
	}
	createSchema := fmt.Sprintf("CREATE SPACE IF NOT EXISTS %v(vid_type=FIXED_STRING(20));"+
		"USE %v;"+
		"CREATE TAG IF NOT EXISTS node(desc string DEFAULT NULL);"+
		"CREATE TAG INDEX IF NOT EXISTS node_tag_index on node();"+
		"CREATE EDGE IF NOT EXISTS edge();"+
		"CREATE EDGE INDEX IF NOT EXISTS edge_tag_index on edge();",
		graph, graph)
	_, err = session.Execute(createSchema)
	fmt.Println(n.GetGraphs())
	return err
}

func (n *NebulaClientImpl) InsertNode(graph string, node *Node) (int, error) {
	session, err := n.getSession()
	if err != nil {
		return 0, err
	}
	insertSchema := fmt.Sprintf("USE %v;"+
		"INSERT VERTEX node(desc) VALUES \"%v\":(%v);", graph, node.Name, node.Desc)
	res, err := session.Execute(insertSchema)
	if err != nil {
		return 0, err
	}
	return res.GetColSize(), nil
}

func (n *NebulaClientImpl) MultiInsertNodes(graph string, nodes []*Node) (int, error) {
	session, err := n.getSession()
	if err != nil {
		return 0, err
	}
	for i, node := range nodes {
		insertSchema := fmt.Sprintf("USE %v;"+
			"INSERT VERTEX node(desc) VALUES \"%v\":(%v);", graph, node.Name, node.Desc)
		_, err := session.Execute(insertSchema)
		if err != nil {
			return i, err
		}
	}
	return len(nodes), nil
}

func (n *NebulaClientImpl) InsertEdge(graph string, edge *Edge) (int, error) {
	session, err := n.getSession()
	if err != nil {
		return 0, err
	}
	insertSchema := fmt.Sprintf("USE %v;"+
		"INSERT EDGE edge() VALUES \"%v\"->\"%v\";", graph, edge.Source, edge.Target)
	res, err := session.Execute(insertSchema)
	if err != nil {
		return 0, err
	}
	return res.GetColSize(), nil
}

func (n *NebulaClientImpl) MultiInsertEdges(graph string, edges []*Edge) (int, error) {
	session, err := n.getSession()
	if err != nil {
		return 0, err
	}
	for i, edge := range edges {
		insertSchema := fmt.Sprintf("USE %v;"+
			"INSERT EDGE edge() VALUES \"%v\"->\"%v\";", graph, edge.Source, edge.Target)
		_, err := session.Execute(insertSchema)
		if err != nil {
			return i, err
		}
	}
	return len(edges), nil
}

func count(graph string, target string, session *nebula.Session) (int64, error) {
	cntSchema := fmt.Sprintf("USE %v;"+
		"LOOKUP ON %v | YIELD COUNT(*) AS cnt;",
		graph, target)
	res, err := session.Execute(cntSchema)
	if err != nil {
		return 0, err
	}
	record, err := res.GetRowValuesByIndex(1)
	if err != nil {
		return 0, err
	}
	cnt, _ := record.GetValueByColName("cnt")
	cntInt64, _ := cnt.AsInt()
	return cntInt64, nil
}

func countNodes(graph string, session *nebula.Session) (int64, error) {
	return count(graph, "node", session)
}

func countEdges(graph string, session *nebula.Session) (int64, error) {
	return count(graph, "edge", session)
}

func (n *NebulaClientImpl) GetGraphs() ([]*Graph, error) {
	session, err := n.getSession()
	if err != nil {
		return nil, err
	}
	getSchema := "SHOW SPACES;"
	res, err := session.Execute(getSchema)
	if err != nil {
		return nil, err
	}
	fmt.Println(res.AsStringTable())
	var graphs []*Graph
	for i := 0; i < res.GetRowSize(); i++ {
		record, _ := res.GetRowValuesByIndex(i)
		name, _ := record.GetValueByColName("Name")
		nameStr, _ := name.AsString()
		nodeCnt, err := countNodes(nameStr, session)
		if err != nil {
			return nil, err
		}
		edgeCnt, err := countEdges(nameStr, session)
		if err != nil {
			return nil, err
		}
		graphs = append(graphs, &Graph{
			Name:  nameStr,
			Nodes: nodeCnt,
			Edges: edgeCnt,
		})
	}
	return graphs, nil
}

func (n *NebulaClientImpl) GetNodes(graph string) ([]*Node, error) {
	session, err := n.getSession()
	if err != nil {
		return nil, err
	}
	getSchema := fmt.Sprintf("USE %v;"+
		"LOOKUP ON node YIELD properties(vertex).desc AS desc;",
		graph)
	res, err := session.Execute(getSchema)
	if err != nil {
		return nil, err
	}
	fmt.Println(res.AsStringTable())
	var nodes []*Node
	for i := 0; i < res.GetRowSize(); i++ {
		record, _ := res.GetRowValuesByIndex(i)
		name, _ := record.GetValueByColName("_vid")
		nameStr, _ := name.AsString()
		desc, _ := record.GetValueByColName("desc")
		descStr, _ := desc.AsString()
		nodes = append(nodes, &Node{
			Name: nameStr,
			Desc: descStr,
		})
	}
	return nodes, nil
}

func (n *NebulaClientImpl) GetEdges(graph string) ([]*Edge, error) {
	session, err := n.getSession()
	if err != nil {
		return nil, err
	}
	getSchema := fmt.Sprintf("USE %v;"+
		"LOOKUP ON edge;",
		graph)
	res, err := session.Execute(getSchema)
	if err != nil {
		return nil, err
	}
	fmt.Println(res.AsStringTable())
	var edges []*Edge
	for i := 0; i < res.GetRowSize(); i++ {
		record, _ := res.GetRowValuesByIndex(i)
		source, _ := record.GetValueByColName("_src")
		sourceStr, _ := source.AsString()
		target, _ := record.GetValueByColName("_dst")
		targetStr, _ := target.AsString()
		edges = append(edges, &Edge{
			Source: sourceStr,
			Target: targetStr,
		})
	}
	return edges, nil
}

func (n *NebulaClientImpl) getSession() (*nebula.Session, error) {
	return n.Pool.GetSession(n.Username, n.Password)
}
