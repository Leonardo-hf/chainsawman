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

func (n *NebulaClientImpl) DropGraph(graph string) error {
	session, err := n.getSession()
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

// TODO: 连接池会被用尽，这个写法是错误的
func (n *NebulaClientImpl) getSession() (*nebula.Session, error) {
	return n.Pool.GetSession(n.Username, n.Password)
}
