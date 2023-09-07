package db

import (
	"chainsawman/common"
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
		return fmt.Errorf("[NEBULA] nGQL error: %v, stats: %v", res.GetErrorMsg(), drop)
	}
	return err
}

func (n *NebulaClientImpl) MatchNodes(graph int64, keywords string, group *model.Group) (map[string][]*MatchNode, error) {
	session, err := n.getSession()
	if err != nil {
		return nil, err
	}
	defer func() { session.Release() }()
	matchNodePackMap := make(map[string][]*MatchNode)
	for _, nt := range group.Nodes {
		primary := ""
		// 检查节点类型的主属性
		for _, attr := range nt.NodeAttrs {
			if common.Int642Bool(attr.Primary) {
				primary = attr.Name
				break
			}
		}
		// 如果没有主属性，则返回
		if primary == "" {
			continue
		}
		tag := nt.Name
		primary = fmt.Sprintf("%v.%v", tag, primary)
		stat := fmt.Sprintf("USE G%v;"+
			"MATCH (v:%v) "+
			"WHERE v.%v STARTS WITH \"%v\" "+
			"RETURN v.%v AS p, id(v) as id "+
			"LIMIT %v;",
			graph, common.BaseTag, primary, keywords, primary, common.MaxMatchCandidates)
		res, err := session.Execute(stat)
		if err != nil {
			return nil, err
		}
		if !res.IsSucceed() {
			return nil, fmt.Errorf("[NEBULA] nGQL error: %v, stats: %v", res.GetErrorMsg(), stat)
		}
		matchNodePacks := make([]*MatchNode, res.GetRowSize())
		for i := 0; i < res.GetRowSize(); i++ {
			r, _ := res.GetRowValuesByIndex(i)
			matchNodePacks[i] = &MatchNode{
				Id:          common.ParseInt(r, "id"),
				PrimaryAttr: common.Parse(r, "p"),
			}
		}
		matchNodePackMap[tag] = matchNodePacks
	}
	return matchNodePackMap, nil
}

// TODO: 连接池会被用尽，这个写法是错误的
func (n *NebulaClientImpl) getSession() (*nebula.Session, error) {
	return n.Pool.GetSession(n.Username, n.Password)
}
