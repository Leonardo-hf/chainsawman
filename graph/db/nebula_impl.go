package db

import (
	"chainsawman/graph/model"
	"fmt"
	"log"

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
		log.Fatal(msg)
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

func (n *NebulaClientImpl) DropGraph(graph int64) error {
	session, err := n.getSession()
	defer func() { session.Release() }()
	if err != nil {
		return err
	}
	drop := fmt.Sprintf("DROP SPACE IF EXISTS G%v;",
		graph)
	res, err := session.Execute(drop)
	if !res.IsSucceed() {
		return fmt.Errorf("[NEBULA] nGQL error: %v", res.GetErrorMsg())
	}
	return err
}

func (n *NebulaClientImpl) GetAllNodes(graph int64) ([]*model.Node, error) {
	session, err := n.getSession()
	defer func() { session.Release() }()
	if err != nil {
		return nil, err
	}
	query := fmt.Sprintf("USE G%v;"+
		"LOOKUP ON snode "+
		"YIELD id(vertex) AS nid, properties(vertex).name AS name, properties(vertex).intro AS intro, properties(vertex).deg AS deg;",
		graph)
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
		idInt, _ := id.AsString()
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

// TODO: 连接池会被用尽，这个写法是错误的
func (n *NebulaClientImpl) getSession() (*nebula.Session, error) {
	return n.Pool.GetSession(n.Username, n.Password)
}
