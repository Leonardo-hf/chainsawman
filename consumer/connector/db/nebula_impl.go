package db

import (
	"chainsawman/consumer/connector/msg"
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
		panic(fmt.Sprintf("Fail to initialize the connection pool, host: %s, port: %d, %s", cfg.Addr, cfg.Port, err.Error()))
	}
	return &NebulaClientImpl{
		Pool:     pool,
		Username: cfg.Username,
		Password: cfg.Passwd,
	}
}

func (n *NebulaClientImpl) InsertNode(graphID int64, node *msg.NodeBody) (int, error) {
	session, err := n.getSession()
	defer func() { session.Release() }()
	if err != nil {
		return 0, err
	}
	insert := fmt.Sprintf("USE G%v;"+
		"INSERT VERTEX snode(name, intro, deg) VALUES \"%v\":(\"%v\", \"%v\", %v);",
		graphID, node.ID, node.Name, node.Desc, 0)
	res, err := session.Execute(insert)
	if err != nil {
		return 0, err
	}
	if !res.IsSucceed() {
		return 0, fmt.Errorf("[NEBULA] nGQL error: %v", res.GetErrorMsg())
	}
	return res.GetColSize(), nil
}

func (n *NebulaClientImpl) UpdateNode(graphID int64, node *msg.NodeBody) (int, error) {
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

func (n *NebulaClientImpl) addDeg(graphID int64, nodeID int64) (int, error) {
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

func (n *NebulaClientImpl) InsertEdge(graphID int64, edge *msg.EdgeBody) (int, error) {
	session, err := n.getSession()
	defer func() { session.Release() }()
	if err != nil {
		return 0, err
	}
	insert := fmt.Sprintf("USE G%v;"+
		"INSERT EDGE sedge() VALUES \"%v\"->\"%v\";", graphID, edge.Source, edge.Target)
	res, err := session.Execute(insert)
	if err != nil {
		return 0, err
	}
	if !res.IsSucceed() {
		return 0, fmt.Errorf("[NEBULA] nGQL error: %v", res.GetErrorMsg())
	}
	_, err = n.addDeg(graphID, edge.Source)
	if err != nil {
		return 0, err
	}
	_, err = n.addDeg(graphID, edge.Target)
	if err != nil {
		return 0, err
	}
	return res.GetColSize(), nil
}

func (n *NebulaClientImpl) DeleteEdge(graphID int64, edge *msg.EdgeBody) (int, error) {
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
