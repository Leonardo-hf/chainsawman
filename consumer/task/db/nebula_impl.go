package db

import (
	"chainsawman/common"
	"chainsawman/consumer/task/model"
	"chainsawman/consumer/task/types"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"math"
	"strconv"
	"strings"

	nebula "github.com/vesoft-inc/nebula-go/v3"
)

type NebulaClientImpl struct {
	Pool     *nebula.ConnectionPool
	Username string
	Password string
	Batch    int
}

type NebulaConfig struct {
	Addr     string
	Port     int
	Username string
	Passwd   string
	Batch    int
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
		Batch:    cfg.Batch,
	}
}

func (n *NebulaClientImpl) getName(id int64) string {
	return fmt.Sprintf("G%v", id)
}

func (n *NebulaClientImpl) CreateGraph(graph int64, group *model.Group) error {
	session, err := n.getSession()
	if err != nil {
		return err
	}
	defer func() { session.Release() }()
	//_, _ = session.Execute("ADD HOSTS 127.0.0.1:9779;")
	name := n.getName(graph)
	stat := fmt.Sprintf(
		"CREATE SPACE IF NOT EXISTS %v (vid_type = INT64);"+
			"USE %v;"+
			"CREATE TAG IF NOT EXISTS %v(%v int);"+
			"CREATE TAG INDEX IF NOT EXISTS deg_tag_index on %v();",
		name, name, common.BaseTag, common.KeyDeg, common.BaseTag)
	for _, node := range group.Nodes {
		attrNames := make([]string, len(node.NodeAttrs))
		s := "CREATE TAG IF NOT EXISTS %v(%v);"
		for i, attr := range node.NodeAttrs {
			attrNames[i] = fmt.Sprintf("`%v` %v", attr.Name, common.Type2String(attr.Type))
		}
		s = fmt.Sprintf(s, node.Name, strings.Join(attrNames, ","))
		stat += s
	}
	for _, edge := range group.Edges {
		attrNames := make([]string, len(edge.EdgeAttrs))
		s := "CREATE EDGE IF NOT EXISTS %v(%v);"
		for i, attr := range edge.EdgeAttrs {
			attrNames[i] = fmt.Sprintf("`%v` %v", attr.Name, common.Type2String(attr.Type))
		}
		s = fmt.Sprintf(s, edge.Name, strings.Join(attrNames, ","))
		stat += s
	}
	res, err := session.Execute(stat)
	if !res.IsSucceed() {
		return fmt.Errorf("[NEBULA] nGQL error: %v, stats: %v", res.GetErrorMsg(), stat)
	}
	return err
}

func (n *NebulaClientImpl) HasGraph(graph int64) (bool, error) {
	session, err := n.getSession()
	if err != nil {
		return false, nil
	}
	defer func() { session.Release() }()
	name := fmt.Sprintf("G%v", graph)
	stat := fmt.Sprintf("SHOW SPACES;")
	res, err := session.Execute(stat)
	if !res.IsSucceed() {
		return false, fmt.Errorf("[NEBULA] nGQL error: %v, stats: %v", res.GetErrorMsg(), stat)
	}
	for i := 0; i < res.GetRowSize(); i++ {
		record, _ := res.GetRowValuesByIndex(i)
		v := common.Parse(record, "Name")
		if v == name {
			return true, nil
		}
	}
	return false, nil
}

func (n *NebulaClientImpl) InsertNode(graph int64, node *model.Node, record *common.Record) (int, error) {
	return n.MultiInsertNodes(graph, node, []*common.Record{record})
}

func (n *NebulaClientImpl) MultiInsertNodes(graph int64, node *model.Node, records []*common.Record) (int, error) {
	session, err := n.getSession()
	if err != nil {
		return 0, err
	}
	defer func() { session.Release() }()
	stat := "USE G%v;INSERT VERTEX %v(%v),base(`deg`) VALUES %v;"
	size := len(node.NodeAttrs)
	names := make([]string, size)
	for i, attr := range node.NodeAttrs {
		names[i] = fmt.Sprintf("`%v`", attr.Name)
	}
	stat = fmt.Sprintf(stat, graph, node.Name, strings.Join(names, ","), "%v")
	vstat := "%v:(%v)"
	for i := 0; i < len(records); {
		values := make([]string, 0)
		for p := 0; p < n.Batch && i < len(records); p++ {
			r := records[i]
			i++
			attrs := make([]string, len(node.NodeAttrs)+1)
			id, _ := r.Get(common.KeyID)
			for q, attr := range node.NodeAttrs {
				v, err := r.Get(attr.Name)
				if err != nil {
					return 0, err
				}
				if attr.Type == 1 {
					attrs[q] = v
				} else {
					attrs[q] = fmt.Sprintf("\"%v\"", v)
				}
			}
			attrs[len(attrs)-1] = common.DefaultDeg
			values = append(values, fmt.Sprintf(vstat, id, strings.Join(attrs, ",")))
		}
		s := fmt.Sprintf(stat, strings.Join(values, ","))
		res, err := session.Execute(s)
		if err != nil {
			return i, err
		}
		if !res.IsSucceed() {
			return 0, fmt.Errorf("[NEBULA] nGQL error: %v, stats: %v", res.GetErrorMsg(), s)
		}
		logx.Infof("[NEBULA] insert %v-th nodes: %v", i, records[i-1])
	}
	return len(records), nil
}

func (n *NebulaClientImpl) MultiIncNodesDeg(graph int64, degMap map[int64]int64) (int, error) {
	session, err := n.getSession()
	if err != nil {
		return 0, err
	}
	defer func() { session.Release() }()
	stat := "USE G%v;UPDATE VERTEX ON base %v SET deg = deg + %v;"
	logx.Infof("[NEBULA] inc nodes deg start, cnt = %v", len(degMap))
	for id, d := range degMap {
		s := fmt.Sprintf(stat, graph, id, d)
		res, err := session.Execute(s)
		if err != nil {
			return 0, err
		}
		if !res.IsSucceed() {
			return 0, fmt.Errorf("[NEBULA] nGQL error: %v, stats: %v", res.GetErrorMsg(), s)
		}
	}
	logx.Infof("[NEBULA] inc nodes deg fin, cnt = %v", len(degMap))
	return len(degMap), nil
}

func (n *NebulaClientImpl) InsertEdge(graph int64, edge *model.Edge, record *common.Record) (int, error) {
	return n.MultiInsertEdges(graph, edge, []*common.Record{record})
}

func (n *NebulaClientImpl) MultiInsertEdges(graph int64, edge *model.Edge, records []*common.Record) (int, error) {
	session, err := n.getSession()
	if err != nil {
		return 0, err
	}
	defer func() { session.Release() }()
	stat := "USE G%v;INSERT EDGE %v(%v) VALUES %v;"
	size := len(edge.EdgeAttrs)
	names := make([]string, size)
	for i, attr := range edge.EdgeAttrs {
		names[i] = fmt.Sprintf("`%v`", attr.Name)
	}
	stat = fmt.Sprintf(stat, graph, edge.Name, strings.Join(names, ","), "%v")
	vstat := "%v->%v:(%v)"
	for i := 0; i < len(records); {
		values := make([]string, 0)
		for p := 0; p < n.Batch && i < len(records); p++ {
			r := records[i]
			i++
			attrs := make([]string, len(edge.EdgeAttrs))
			src, _ := r.Get(common.KeySrc)
			tgt, _ := r.Get(common.KeyTgt)
			for q, attr := range edge.EdgeAttrs {
				v, err := r.Get(attr.Name)
				if err != nil {
					return 0, err
				}
				if attr.Type == 1 {
					attrs[q] = v
				} else {
					attrs[q] = fmt.Sprintf("\"%v\"", v)
				}
			}
			values = append(values, fmt.Sprintf(vstat, src, tgt, strings.Join(attrs, ",")))
			// 无向图，插入反向边
			if !common.Int642Bool(edge.Direct) {
				values = append(values, fmt.Sprintf(vstat, tgt, src, strings.Join(attrs, ",")))
			}
		}
		s := fmt.Sprintf(stat, strings.Join(values, ","))
		res, err := session.Execute(s)
		if err != nil {
			return i, err
		}
		if !res.IsSucceed() {
			return 0, fmt.Errorf("[NEBULA] nGQL error: %v, stats: %v", res.GetErrorMsg(), s)
		}
		logx.Infof("[NEBULA] insert %v-th edge: %v", i, records[i-1])
	}
	return len(records), nil
}

// 解析 vertex
func parseVertex(res *nebula.ResultSet) (map[string][]*types.Node, error) {
	nodePackMap := make(map[string][]*types.Node)
	for i := 0; i < res.GetRowSize(); i++ {
		r, _ := res.GetRowValuesByIndex(i)
		record, _ := r.GetValueByColName("v")
		v, _ := record.AsNode()
		// 获得id
		id, _ := v.GetID().AsInt()
		node := &types.Node{
			Id:    id,
			Attrs: make([]*types.Pair, 0),
		}
		// 获得节点deg
		tag := common.BaseTag
		props, _ := v.Properties(tag)
		deg, _ := props[common.KeyDeg].AsInt()
		node.Deg = deg
		// 获得节点tag
		for _, t := range v.GetTags() {
			if t != tag {
				tag = t
				break
			}
		}
		if _, ok := nodePackMap[tag]; !ok {
			nodePackMap[tag] = make([]*types.Node, 0)
		}
		// 获得节点属性
		props, _ = v.Properties(tag)
		for k, value := range props {
			vs := value.String()
			if value.IsString() {
				vs, _ = value.AsString()
			}
			node.Attrs = append(node.Attrs, &types.Pair{
				Key:   k,
				Value: vs,
			})
		}
		nodePackMap[tag] = append(nodePackMap[tag], node)
	}
	return nodePackMap, nil
}

// GetNodesByIds 根据节点的id查询节点的详细信息，TODO: 对ids不做限量?
func (n *NebulaClientImpl) GetNodesByIds(graph int64, ids []int64) (map[string][]*types.Node, error) {
	if len(ids) == 0 {
		return make(map[string][]*types.Node), nil
	}
	session, err := n.getSession()
	if err != nil {
		return nil, err
	}
	defer func() { session.Release() }()
	stringifyIds := make([]string, len(ids))
	for i, id := range ids {
		stringifyIds[i] = fmt.Sprintf("%v", strconv.FormatInt(id, 10))
	}
	stat := fmt.Sprintf("USE G%v;FETCH PROP ON * %v YIELD vertex AS v;", graph, strings.Join(stringifyIds, ","))
	res, err := session.Execute(stat)
	if err != nil {
		return nil, err
	}
	if !res.IsSucceed() {
		return nil, fmt.Errorf("[NEBULA] nGQL error: %v, stats: %v", res.GetErrorMsg(), stat)
	}
	return parseVertex(res)
}

// GetTopNodes 获得度数前top的节点，按节点类型分类返回
func (n *NebulaClientImpl) GetTopNodes(graph int64, top int64) (map[string][]*types.Node, error) {
	session, err := n.getSession()
	if err != nil {
		return nil, err
	}
	defer func() { session.Release() }()
	stat := fmt.Sprintf("USE G%v;"+
		"MATCH (v:%v) WITH v, v.%v.%v AS deg "+
		"ORDER BY deg DESC "+
		"LIMIT %v "+
		"RETURN v;",
		graph, common.BaseTag, common.BaseTag, common.KeyDeg, top)
	res, err := session.Execute(stat)
	if err != nil {
		return nil, err
	}
	if !res.IsSucceed() {
		return nil, fmt.Errorf("[NEBULA] nGQL error: %v, stats: %v", res.GetErrorMsg(), stat)
	}
	return parseVertex(res)
}

// 解析 Relationship
func parseRelationship(res *nebula.ResultSet) (map[string][]*types.Edge, error) {
	edgePackMap := make(map[string][]*types.Edge)
	for i := 0; i < res.GetRowSize(); i++ {
		r, _ := res.GetRowValuesByIndex(i)
		record, _ := r.GetValueByColName("relations")
		relation, _ := record.AsRelationship()
		src, _ := relation.GetSrcVertexID().AsInt()
		tgt, _ := relation.GetDstVertexID().AsInt()
		edge := &types.Edge{
			Source: src,
			Target: tgt,
			Attrs:  make([]*types.Pair, 0),
		}
		// 获得边tag
		tag := relation.GetEdgeName()
		if _, ok := edgePackMap[tag]; !ok {
			edgePackMap[tag] = make([]*types.Edge, 0)
		}
		// 获得边属性
		props := relation.Properties()
		for k, v := range props {
			vs := v.String()
			if v.IsString() {
				vs, _ = v.AsString()
			}
			edge.Attrs = append(edge.Attrs, &types.Pair{
				Key:   k,
				Value: vs,
			})
		}
		edgePackMap[tag] = append(edgePackMap[tag], edge)
	}
	return edgePackMap, nil
}

// Go 以单个节点为起点展开游走，返回游走的边
func (n *NebulaClientImpl) Go(graph int64, src int64, direction string, max int64) (map[string][]*types.Edge, error) {
	return n.MultiGo(graph, []int64{src}, direction, max)
}

// MultiGo 以多个节点为起点展开游走，返回游走的边
func (n *NebulaClientImpl) MultiGo(graph int64, srcList []int64, direction string, max int64) (map[string][]*types.Edge, error) {
	session, err := n.getSession()
	if err != nil {
		return nil, err
	}
	defer func() { session.Release() }()
	names := make([]string, len(srcList))
	for i, src := range srcList {
		names[i] = fmt.Sprintf("%v", src)
	}
	start := float64(max + 1.0)
	sample := make([]string, 0)
	for start > 1 {
		start = math.Ceil(start / 2)
		sample = append(sample, strconv.FormatInt(int64(start), 10))
	}
	stat := fmt.Sprintf("USE G%v;"+
		"GO 1 TO %v STEPS FROM %v OVER * %v "+
		"YIELD DISTINCT edge AS relations "+
		"SAMPLE [%v];", graph, len(sample), strings.Join(names, ","), direction, strings.Join(sample, ","))
	res, err := session.Execute(stat)
	if err != nil {
		return nil, err
	}
	if !res.IsSucceed() {
		return nil, fmt.Errorf("[NEBULA] nGQL error: %v, stats: %v", res.GetErrorMsg(), stat)
	}
	return parseRelationship(res)
}

// GoFromTopNodes 以度数前top名的节点为起点展开游走，返回游走的边
func (n *NebulaClientImpl) GoFromTopNodes(graph int64, top int64, direction string, max int64) (map[string][]*types.Edge, error) {
	session, err := n.getSession()
	if err != nil {
		return nil, err
	}
	defer func() { session.Release() }()
	start := float64(max + 1.0)
	sample := make([]string, 0)
	for start > 1 {
		start = math.Ceil(start / 2)
		sample = append(sample, strconv.FormatInt(int64(start), 10))
	}
	stat := fmt.Sprintf("USE G%v;"+
		"LOOKUP ON base YIELD id(vertex) AS nid, properties(vertex).deg AS deg"+
		"| ORDER BY $-.deg DESC"+
		"| LIMIT %v"+
		"| GO 1 TO %v STEPS FROM $-.nid OVER * %v"+
		"  YIELD DISTINCT edge AS relations"+
		"  SAMPLE [%v];", graph, top, len(sample), direction, strings.Join(sample, ","))
	res, err := session.Execute(stat)
	if err != nil {
		return nil, err
	}
	if !res.IsSucceed() {
		return nil, fmt.Errorf("[NEBULA] nGQL error: %v, stats: %v", res.GetErrorMsg(), stat)
	}
	return parseRelationship(res)
}

// DropGraph 删除图空间
func (n *NebulaClientImpl) DropGraph(graph int64) error {
	session, err := n.getSession()
	if err != nil {
		return err
	}
	defer func() { session.Release() }()
	stat := fmt.Sprintf("DROP SPACE IF EXISTS G%v;", graph)
	res, err := session.Execute(stat)
	if !res.IsSucceed() {
		return fmt.Errorf("[NEBULA] nGQL error: %v, stats: %v", res.GetErrorMsg(), stat)
	}
	return err
}

func (n *NebulaClientImpl) getSession() (*nebula.Session, error) {
	return n.Pool.GetSession(n.Username, n.Password)
}
