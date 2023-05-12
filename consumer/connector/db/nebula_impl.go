package db

import (
	"chainsawman/consumer/connector/model"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"

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

func (n *NebulaClientImpl) GetOutNeighbors(graph int64, nodeID int64) ([]int64, error) {
	session, err := n.getSession()
	defer func() { session.Release() }()
	if err != nil {
		return nil, err
	}
	query := fmt.Sprintf("USE G%v;"+
		"GET SUBGRAPH WITH PROP 1 STEPS FROM \"%v\" OUT sedge "+
		"YIELD VERTICES AS nodes;", graph, nodeID)
	res, err := session.Execute(query)
	if !res.IsSucceed() {
		return nil, fmt.Errorf("[NEBULA] nGQL error: %v", res.GetErrorMsg())
	}
	nodes := make([]int64, 0)
	for i := 0; i < res.GetRowSize(); i++ {
		record, _ := res.GetRowValuesByIndex(i)
		snodes, _ := record.GetValueByColName("nodes")
		snodesList, _ := snodes.AsList()
		for _, snodeWrapper := range snodesList {
			snode, _ := snodeWrapper.AsNode()
			id, _ := snode.GetID().AsString()
			idInt, _ := strconv.ParseInt(id, 10, 64)
			nodes = append(nodes, idInt)
		}
	}
	return nodes, nil
}

func InitNebulaClient(cfg *NebulaConfig) NebulaClient {
	hostAddress := nebula.HostAddress{Host: cfg.Addr, Port: cfg.Port}
	hostList := []nebula.HostAddress{hostAddress}
	testPoolConfig := nebula.GetDefaultConf()
	pool, err := nebula.NewConnectionPool(hostList, testPoolConfig, nebula.DefaultLogger{})
	if err != nil {
		panic(fmt.Sprintf("Fail to initialize the connection pool, host: %s, port: %d, %s", cfg.Addr, cfg.Port, err.Error()))
	}
	logx.Info("[Connector] nebula init.")
	return &NebulaClientImpl{
		Pool:     pool,
		Username: cfg.Username,
		Password: cfg.Passwd,
	}
}

func (n *NebulaClientImpl) MultiInsertNodes(graph int64, nodes []*model.NebulaNode) (int, error) {
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
			return 0, fmt.Errorf("[NEBULA] nGQL error: %v", res.GetErrorMsg())
		}
	}
	return len(nodes), nil
}

func (n *NebulaClientImpl) UpdateNode(graphID int64, node *model.NebulaNode) (int, error) {
	session, err := n.getSession()
	defer func() { session.Release() }()
	if err != nil {
		return 0, err
	}
	insert := fmt.Sprintf("USE G%v;"+
		"UPDATE VERTEX ON snode \"%v\" SET name = \"%v\", desc = \"%v\";",
		graphID, node.ID, node.Name, node.Desc)
	res, err := session.Execute(insert)
	if err != nil {
		return 0, err
	}
	if !res.IsSucceed() {
		return 0, fmt.Errorf("[NEBULA] nGQL error: %v", res.GetErrorMsg())
	}
	return res.GetColSize(), nil
}

func (n *NebulaClientImpl) AddDeg(graphID int64, nodeID int64) (int, error) {
	session, err := n.getSession()
	defer func() { session.Release() }()
	if err != nil {
		return 0, err
	}
	insert := fmt.Sprintf("USE G%v;"+
		"UPDATE VERTEX ON snode \"%v\" SET deg = deg + 1",
		graphID, nodeID)
	res, err := session.Execute(insert)
	if err != nil {
		return 0, err
	}
	if !res.IsSucceed() {
		return 0, fmt.Errorf("[NEBULA] nGQL error: %v", res.GetErrorMsg())
	}
	return res.GetColSize(), nil
}

// DeleteNode TODO: 删除节点导致边删除，会导致度数错误
func (n *NebulaClientImpl) DeleteNode(graphID int64, nodeID int64) (int, error) {
	session, err := n.getSession()
	defer func() { session.Release() }()
	if err != nil {
		return 0, err
	}
	insert := fmt.Sprintf("USE G%v;"+
		"DELETE VERTEX \"%v\" WITH EDGE;",
		graphID, nodeID)
	res, err := session.Execute(insert)
	if err != nil {
		return 0, err
	}
	if !res.IsSucceed() {
		return 0, fmt.Errorf("[NEBULA] nGQL error: %v", res.GetErrorMsg())
	}
	return res.GetColSize(), nil
}

func (n *NebulaClientImpl) InsertEdge(graphID int64, edge *model.NebulaEdge) (int, error) {
	return n.MultiInsertEdges(graphID, []*model.NebulaEdge{edge})
}

func (n *NebulaClientImpl) MultiInsertEdges(graph int64, edges []*model.NebulaEdge) (int, error) {
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
			return 0, fmt.Errorf("[NEBULA] nGQL error: %v", res.GetErrorMsg())
		}
	}
	return len(edges), nil
}

func (n *NebulaClientImpl) DeleteEdge(graphID int64, edge *model.NebulaEdge) (int, error) {
	session, err := n.getSession()
	defer func() { session.Release() }()
	if err != nil {
		return 0, err
	}
	insert := fmt.Sprintf("USE G%v;"+
		"DELETE EDGE sedge \"%v\" -> \"%v\";",
		graphID, edge.Source, edge.Target)
	res, err := session.Execute(insert)
	if err != nil {
		return 0, err
	}
	if !res.IsSucceed() {
		return 0, fmt.Errorf("[NEBULA] nGQL error: %v", res.GetErrorMsg())
	}
	return res.GetColSize(), nil
}

func (n *NebulaClientImpl) getSession() (*nebula.Session, error) {
	return n.Pool.GetSession(n.Username, n.Password)
}
