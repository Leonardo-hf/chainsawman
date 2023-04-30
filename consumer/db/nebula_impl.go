package db

import (
	"chainsawman/consumer/model"

	"fmt"

	nebula "github.com/vesoft-inc/nebula-go/v3"
)

type NebulaClientImpl struct {
	Pool     *nebula.ConnectionPool
	Username string
	Password string
}

type NebulaConfig struct {
	Addr     string
	Port     int
	Username string
	Passwd   string
}

func InitNebulaClient(cfg *NebulaConfig) NebulaClient {
	hostAddress := nebula.HostAddress{Host: cfg.Addr, Port: cfg.Port}
	hostList := []nebula.HostAddress{hostAddress}
	testPoolConfig := nebula.GetDefaultConf()
	pool, err := nebula.NewConnectionPool(hostList, testPoolConfig, nebula.DefaultLogger{})
	if err != nil {
		msg := fmt.Sprintf("Fail to initialize the connection pool, host: %s, port: %d, %s", cfg.Addr, cfg.Port, err.Error())
		panic(msg)
	}
	return &NebulaClientImpl{
		Pool:     pool,
		Username: cfg.Username,
		Password: cfg.Passwd,
	}
}

func (n *NebulaClientImpl) CreateGraph(graph int64) error {
	session, err := n.getSession()
	defer func() { session.Release() }()
	if err != nil {
		return err
	}
	//_, _ = session.Execute("ADD HOSTS 127.0.0.1:9779;")
	create := fmt.Sprintf(
		"CREATE SPACE IF NOT EXISTS G%v (vid_type = FIXED_STRING(30));"+
			"USE G%v;"+
			"CREATE TAG IF NOT EXISTS snode(name string, intro string, deg int);"+
			"CREATE TAG INDEX IF NOT EXISTS snode_tag_index on snode();"+
			"CREATE EDGE IF NOT EXISTS sedge();"+
			"CREATE EDGE INDEX IF NOT EXISTS sedge_tag_index on sedge();",
		graph, graph)
	res, err := session.Execute(create)
	if !res.IsSucceed() {
		return fmt.Errorf("[NEBULA] nGQL error: %v", res.GetErrorMsg())
	}
	return err
}

func (n *NebulaClientImpl) InsertNode(graph int64, node *model.Node) (int, error) {
	session, err := n.getSession()
	defer func() { session.Release() }()
	if err != nil {
		return 0, err
	}
	insert := fmt.Sprintf("USE G%v;"+
		"INSERT VERTEX snode(name, intro, deg) VALUES \"%v\":(\"%v\", \"%v\", %v);",
		graph, node.ID, node.Name, node.Desc, node.Deg)
	res, err := session.Execute(insert)
	if err != nil {
		return 0, err
	}
	if !res.IsSucceed() {
		return 0, fmt.Errorf("[NEBULA] nGQL error: %v", res.GetErrorMsg())
	}
	return res.GetColSize(), nil
}

func (n *NebulaClientImpl) MultiInsertNodes(graph int64, nodes []*model.Node) (int, error) {
	session, err := n.getSession()
	defer func() { session.Release() }()
	if err != nil {
		return 0, err
	}
	for i, node := range nodes {
		insert := fmt.Sprintf("USE G%v;"+
			"INSERT VERTEX snode(name, intro, deg) VALUES \"%v\":(\"%v\", \"%v\", %v);",
			graph, node.ID, node.Name, node.Desc, node.Deg)
		res, err := session.Execute(insert)
		if err != nil {
			return i, err
		}
		if !res.IsSucceed() {
			fmt.Println("nodes")
			return 0, fmt.Errorf("[NEBULA] nGQL error: %v", res.GetErrorMsg())
		}
	}
	return len(nodes), nil
}

func (n *NebulaClientImpl) InsertEdge(graph int64, edge *model.Edge) (int, error) {
	session, err := n.getSession()
	defer func() { session.Release() }()
	if err != nil {
		return 0, err
	}
	insert := fmt.Sprintf("USE G%v;"+
		"INSERT EDGE sedge() VALUES \"%v\"->\"%v\";", graph, edge.Source, edge.Target)
	res, err := session.Execute(insert)
	if err != nil {
		return 0, err
	}
	if !res.IsSucceed() {
		return 0, fmt.Errorf("[NEBULA] nGQL error: %v", res.GetErrorMsg())
	}
	return res.GetColSize(), nil
}

func (n *NebulaClientImpl) MultiInsertEdges(graph int64, edges []*model.Edge) (int, error) {
	session, err := n.getSession()
	defer func() { session.Release() }()
	if err != nil {
		return 0, err
	}
	for i, edge := range edges {
		insert := fmt.Sprintf("USE G%v;"+
			"INSERT EDGE sedge() VALUES \"%v\"->\"%v\":();", graph, edge.Source, edge.Target)
		res, err := session.Execute(insert)
		if err != nil {
			return i, err
		}
		if !res.IsSucceed() {
			fmt.Println("edges")
			return 0, fmt.Errorf("[NEBULA] nGQL error: %v", res.GetErrorMsg())
		}
	}
	return len(edges), nil
}

func (n *NebulaClientImpl) GetNodes(graph int64, min int64) ([]*model.Node, error) {
	session, err := n.getSession()
	defer func() { session.Release() }()
	if err != nil {
		return nil, err
	}
	query := fmt.Sprintf("USE G%v;"+
		"LOOKUP ON snode WHERE snode.deg > %v "+
		"YIELD id(vertex) AS nid, properties(vertex).name AS name, properties(vertex).intro AS intro, properties(vertex).deg AS deg;",
		graph, min)
	res, err := session.Execute(query)
	if err != nil {
		return nil, err
	}
	if !res.IsSucceed() {
		return nil, fmt.Errorf("[NEBULA] nGQL error: %v", res.GetErrorMsg())
	}
	var nodes []*model.Node
	for i := 0; i < res.GetRowSize(); i++ {
		record, _ := res.GetRowValuesByIndex(i)
		id, _ := record.GetValueByColName("nid")
		idInt, _ := id.AsInt()
		name, _ := record.GetValueByColName("name")
		nameStr, _ := name.AsString()
		desc, _ := record.GetValueByColName("intro")
		descStr, _ := desc.AsString()
		deg, _ := record.GetValueByColName("deg")
		degInt, _ := deg.AsInt()
		fmt.Println(degInt)
		nodes = append(nodes, &model.Node{
			ID:   idInt,
			Name: nameStr,
			Desc: descStr,
			Deg:  degInt,
		})
	}
	return nodes, nil
}

func (n *NebulaClientImpl) GetEdges(graph int64) ([]*model.Edge, error) {
	session, err := n.getSession()
	defer func() { session.Release() }()
	if err != nil {
		return nil, err
	}
	query := fmt.Sprintf("USE G%v;"+
		"LOOKUP ON sedge YIELD src(edge) AS src, dst(edge) AS dst;",
		graph)
	res, err := session.Execute(query)
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
		sourceInt, _ := source.AsInt()
		target, _ := record.GetValueByColName("dst")
		targetInt, _ := target.AsInt()
		edges = append(edges, &model.Edge{
			Source: sourceInt,
			Target: targetInt,
		})
	}
	return edges, nil
}

func (n *NebulaClientImpl) DropGraph(graph int64) error {
	session, err := n.getSession()
	defer func() { session.Release() }()
	if err != nil {
		return err
	}
	drop := fmt.Sprintf("DROP SPACE IF EXISTS %v;",
		graph)
	res, err := session.Execute(drop)
	if !res.IsSucceed() {
		return fmt.Errorf("[NEBULA] nGQL error: %v", res.GetErrorMsg())
	}
	return err
}

func (n *NebulaClientImpl) GetNeighbors(graph int64, nodeID int64, min int64, distance int64) ([]*model.Node, []*model.Edge, error) {
	session, err := n.getSession()
	defer func() { session.Release() }()
	if err != nil {
		return nil, nil, err
	}
	query := fmt.Sprintf("USE G%v;"+
		"GET SUBGRAPH WITH PROP %v STEPS FROM \"%v\" WHERE v.snode.deg >= %v "+
		"YIELD VERTICES AS nodes, EDGES AS relations;", graph, distance, nodeID, min)
	res, err := session.Execute(query)
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
			id, _ := snode.GetID().AsInt()
			name, _ := snodeProps["name"].AsString()
			intro, _ := snodeProps["intro"].AsString()
			deg, _ := snodeProps["deg"].AsInt()
			crt := &model.Node{
				ID:   id,
				Name: name,
				Desc: intro,
				Deg:  deg,
			}
			nodes = append(nodes, crt)
		}
		sedges, _ := record.GetValueByColName("relations")
		sedgesList, _ := sedges.AsList()
		for _, sedgeWrapper := range sedgesList {
			sedge, _ := sedgeWrapper.AsRelationship()
			src, _ := sedge.GetSrcVertexID().AsInt()
			dst, _ := sedge.GetDstVertexID().AsInt()
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
