package db

import (
	"chainsawman/graph/model"
	"fmt"
	nebula "github.com/vesoft-inc/nebula-go/v3"
	"log"
)

type NebulaClientImpl struct {
	Pool     *nebula.ConnectionPool
	Username string
	Password string
}

type NebulaConfig struct {
	Addr     string
	Port     int
	Log      nebula.Logger
	Username string
	Passwd   string
}

func InitNebulaClient(cfg *NebulaConfig) NebulaClient {
	hostAddress := nebula.HostAddress{Host: cfg.Addr, Port: cfg.Port}
	hostList := []nebula.HostAddress{hostAddress}
	testPoolConfig := nebula.GetDefaultConf()
	pool, err := nebula.NewConnectionPool(hostList, testPoolConfig, cfg.Log)
	if err != nil {
		msg := fmt.Sprintf("Fail to initialize the connection pool, host: %s, port: %d, %s", cfg.Addr, cfg.Port, err.Error())
		log.Fatal(msg)
	}
	return &NebulaClientImpl{
		Pool:     pool,
		Username: cfg.Username,
		Password: cfg.Passwd,
	}
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
			"CREATE TAG IF NOT EXISTS snode(intro string, deg int);"+
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

func (n *NebulaClientImpl) InsertNode(graph string, node *model.Node) (int, error) {
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

func (n *NebulaClientImpl) MultiInsertNodes(graph string, nodes []*model.Node) (int, error) {
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

func (n *NebulaClientImpl) InsertEdge(graph string, edge *model.Edge) (int, error) {
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

func (n *NebulaClientImpl) MultiInsertEdges(graph string, edges []*model.Edge) (int, error) {
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

func (n *NebulaClientImpl) GetGraphs() ([]*model.Graph, error) {
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
	var graphs []*model.Graph
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
		graphs = append(graphs, &model.Graph{
			Name:  nameStr,
			Nodes: nodeCnt,
			Edges: edgeCnt,
		})
	}
	return graphs, nil
}

func (n *NebulaClientImpl) GetNodes(graph string) ([]*model.Node, error) {
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
	var nodes []*model.Node
	for i := 0; i < res.GetRowSize(); i++ {
		record, _ := res.GetRowValuesByIndex(i)
		name, _ := record.GetValueByColName("name")
		nameStr, _ := name.AsString()
		desc, _ := record.GetValueByColName("intro")
		descStr, _ := desc.AsString()
		nodes = append(nodes, &model.Node{
			Name: nameStr,
			Desc: descStr,
		})
	}
	return nodes, nil
}

func (n *NebulaClientImpl) GetEdges(graph string) ([]*model.Edge, error) {
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
	var edges []*model.Edge
	for i := 0; i < res.GetRowSize(); i++ {
		record, _ := res.GetRowValuesByIndex(i)
		source, _ := record.GetValueByColName("src")
		sourceStr, _ := source.AsString()
		target, _ := record.GetValueByColName("dst")
		targetStr, _ := target.AsString()
		edges = append(edges, &model.Edge{
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

// GetNeighbors TODO: 这方法是给消费者用的，好像可以用UNWIND优化一下输出结果；没有应用上度数过滤，考虑在代码里判断，或者给node加上度数的属性
func (n *NebulaClientImpl) GetNeighbors(graph string, node string, min int64, distance int64) ([]*model.Node, []*model.Edge, error) {
	session, err := n.getSession()
	if err != nil {
		return nil, nil, err
	}
	goSchema := fmt.Sprintf("USE %v;"+
		"GET SUBGRAPH WITH PROP %v STEPS FROM \"%v\" YIELD VERTICES AS nodes, EDGES AS relations;", graph, distance, node)
	res, err := session.Execute(goSchema)
	if !res.IsSucceed() {
		return nil, nil, fmt.Errorf("[NEBULA] nGQL error: %v", res.GetErrorMsg())
	}
	var nodes []*model.Node
	var edges []*model.Edge
	for i := 0; i < res.GetRowSize(); i++ {
		record, _ := res.GetRowValuesByIndex(i)
		snodes, _ := record.GetValueByColName("nodes")
		snodesList, _ := snodes.AsList()
		for _, snodeWrapper := range snodesList {
			snode, _ := snodeWrapper.AsNode()
			snodeProps, _ := snode.Properties("snode")
			name, _ := snode.GetID().AsString()
			intro, _ := snodeProps["intro"].AsString()
			crt := &model.Node{
				Name: name,
				Desc: intro,
			}
			nodes = append(nodes, crt)
		}
		sedges, _ := record.GetValueByColName("relations")
		sedgesList, _ := sedges.AsList()
		for _, sedgeWrapper := range sedgesList {
			sedge, _ := sedgeWrapper.AsRelationship()
			src, _ := sedge.GetSrcVertexID().AsString()
			dst, _ := sedge.GetDstVertexID().AsString()
			crt := &model.Edge{
				Source: src,
				Target: dst,
			}
			edges = append(edges, crt)
		}
	}
	return nodes, edges, nil
}

func (n *NebulaClientImpl) getSession() (*nebula.Session, error) {
	return n.Pool.GetSession(n.Username, n.Password)
}
