package db

import (
	nebula2 "chainsawman/graph/model"
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
	//_, _ = session.Execute("ADD HOSTS 127.0.0.1:9779;")
	createSchema := fmt.Sprintf(
		"CREATE SPACE IF NOT EXISTS %v(vid_type = FIXED_STRING(30));"+
			"USE %v;"+
			"CREATE TAG IF NOT EXISTS snode(intro string);"+
			"CREATE TAG INDEX IF NOT EXISTS snode_tag_index on snode();"+
			"CREATE EDGE IF NOT EXISTS sedge();"+
			"CREATE EDGE INDEX IF NOT EXISTS sedge_tag_index on sedge();",
		graph, graph)
	res, err := session.Execute(createSchema)
	if !res.IsSucceed() {
		return fmt.Errorf("[NEBULA] nGQL error: %v", res.GetErrorMsg())
	}
	return err
}

func (n *NebulaClientImpl) InsertNode(graph string, node *nebula2.Node) (int, error) {
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
	if !res.IsSucceed() {
		return 0, fmt.Errorf("[NEBULA] nGQL error: %v", res.GetErrorMsg())
	}
	return res.GetColSize(), nil
}

func (n *NebulaClientImpl) MultiInsertNodes(graph string, nodes []*nebula2.Node) (int, error) {
	session, err := n.getSession()
	if err != nil {
		return 0, err
	}
	for i, node := range nodes {
		insertSchema := fmt.Sprintf("USE %v;"+
			"INSERT VERTEX snode(intro) VALUES \"%v\":(\"%v\");", graph, node.Name, node.Desc)
		res, err := session.Execute(insertSchema)
		if err != nil {
			return i, err
		}
		if !res.IsSucceed() {
			return 0, fmt.Errorf("[NEBULA] nGQL error: %v", res.GetErrorMsg())
		}
	}
	return len(nodes), nil
}

func (n *NebulaClientImpl) InsertEdge(graph string, edge *nebula2.Edge) (int, error) {
	session, err := n.getSession()
	if err != nil {
		return 0, err
	}
	insertSchema := fmt.Sprintf("USE %v;"+
		"INSERT EDGE sedge() VALUES \"%v\"->\"%v\";", graph, edge.Source, edge.Target)
	res, err := session.Execute(insertSchema)
	if err != nil {
		return 0, err
	}
	if !res.IsSucceed() {
		return 0, fmt.Errorf("[NEBULA] nGQL error: %v", res.GetErrorMsg())
	}
	return res.GetColSize(), nil
}

func (n *NebulaClientImpl) MultiInsertEdges(graph string, edges []*nebula2.Edge) (int, error) {
	session, err := n.getSession()
	if err != nil {
		return 0, err
	}
	for i, edge := range edges {
		insertSchema := fmt.Sprintf("USE %v;"+
			"INSERT EDGE sedge() VALUES \"%v\"->\"%v\":();", graph, edge.Source, edge.Target)
		res, err := session.Execute(insertSchema)
		if err != nil {
			return i, err
		}
		if !res.IsSucceed() {
			return 0, fmt.Errorf("[NEBULA] nGQL error: %v", res.GetErrorMsg())
		}
	}
	return len(edges), nil
}

func count(graph string, target string, session *nebula.Session) (int64, error) {
	cntSchema := fmt.Sprintf("USE %v;"+
		"LOOKUP ON %v| YIELD COUNT(*) AS cnt;",
		graph, target)
	res, err := session.Execute(cntSchema)
	if err != nil {
		return 0, err
	}
	if !res.IsSucceed() {
		return 0, fmt.Errorf("[NEBULA] nGQL error: %v", res.GetErrorMsg())
	}
	record, err := res.GetRowValuesByIndex(0)
	if err != nil {
		return 0, err
	}
	cnt, _ := record.GetValueByColName("cnt")
	cntInt64, _ := cnt.AsInt()
	return cntInt64, nil
}

func countNodes(graph string, session *nebula.Session) (int64, error) {
	return count(graph, "snode YIELD id(vertex)", session)
}

func countEdges(graph string, session *nebula.Session) (int64, error) {
	return count(graph, "sedge YIELD edge AS e", session)
}

func (n *NebulaClientImpl) GetGraphs() ([]*nebula2.Graph, error) {
	session, err := n.getSession()
	if err != nil {
		return nil, err
	}
	getSchema := "SHOW SPACES;"
	res, err := session.Execute(getSchema)
	if err != nil {
		return nil, err
	}
	if !res.IsSucceed() {
		return nil, fmt.Errorf("[NEBULA] nGQL error: %v", res.GetErrorMsg())
	}
	var graphs []*nebula2.Graph
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
		graphs = append(graphs, &nebula2.Graph{
			Name:  nameStr,
			Nodes: nodeCnt,
			Edges: edgeCnt,
		})
	}
	return graphs, nil
}

func (n *NebulaClientImpl) GetNodes(graph string) ([]*nebula2.Node, error) {
	session, err := n.getSession()
	if err != nil {
		return nil, err
	}
	getSchema := fmt.Sprintf("USE %v;"+
		"LOOKUP ON snode YIELD id(vertex) AS name, properties(vertex).intro AS intro;",
		graph)
	res, err := session.Execute(getSchema)
	if err != nil {
		return nil, err
	}
	if !res.IsSucceed() {
		return nil, fmt.Errorf("[NEBULA] nGQL error: %v", res.GetErrorMsg())
	}
	var nodes []*nebula2.Node
	for i := 0; i < res.GetRowSize(); i++ {
		record, _ := res.GetRowValuesByIndex(i)
		name, _ := record.GetValueByColName("name")
		nameStr, _ := name.AsString()
		desc, _ := record.GetValueByColName("intro")
		descStr, _ := desc.AsString()
		nodes = append(nodes, &nebula2.Node{
			Name: nameStr,
			Desc: descStr,
		})
	}
	return nodes, nil
}

func (n *NebulaClientImpl) GetEdges(graph string) ([]*nebula2.Edge, error) {
	session, err := n.getSession()
	if err != nil {
		return nil, err
	}
	getSchema := fmt.Sprintf("USE %v;"+
		"LOOKUP ON sedge YIELD src(edge) AS src, dst(edge) AS dst;",
		graph)
	res, err := session.Execute(getSchema)
	if err != nil {
		return nil, err
	}
	if !res.IsSucceed() {
		return nil, fmt.Errorf("[NEBULA] nGQL error: %v", res.GetErrorMsg())
	}
	var edges []*nebula2.Edge
	for i := 0; i < res.GetRowSize(); i++ {
		record, _ := res.GetRowValuesByIndex(i)
		source, _ := record.GetValueByColName("src")
		sourceStr, _ := source.AsString()
		target, _ := record.GetValueByColName("dst")
		targetStr, _ := target.AsString()
		edges = append(edges, &nebula2.Edge{
			Source: sourceStr,
			Target: targetStr,
		})
	}
	return edges, nil
}

func (n *NebulaClientImpl) DropGraph(graph string) error {
	session, err := n.getSession()
	if err != nil {
		return err
	}
	dropSchema := fmt.Sprintf("DROP SPACE IF EXISTS %v;",
		graph)
	res, err := session.Execute(dropSchema)
	if !res.IsSucceed() {
		return fmt.Errorf("[NEBULA] nGQL error: %v", res.GetErrorMsg())
	}
	return err
}

func (n *NebulaClientImpl) getSession() (*nebula.Session, error) {
	return n.Pool.GetSession(n.Username, n.Password)
}
