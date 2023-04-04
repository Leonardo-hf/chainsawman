package config

import (
	"chainsawman/graph/db"
	"fmt"

	nebula "github.com/vesoft-inc/nebula-go/v3"
)

var NebulaClient db.NebulaClient

var MysqlClient db.MysqlClient

const (
	address  = "127.0.0.1"
	port     = 9669
	username = "root"
	password = "nebula"

	Dsn = "root:12345678@(localhost:3306)/graph?charset=utf8mb4&parseTime=True&loc=Local"
)

var log = nebula.DefaultLogger{}

func initNebula() {
	hostAddress := nebula.HostAddress{Host: address, Port: port}
	hostList := []nebula.HostAddress{hostAddress}
	testPoolConfig := nebula.GetDefaultConf()
	pool, err := nebula.NewConnectionPool(hostList, testPoolConfig, log)
	if err != nil {
		msg := fmt.Sprintf("Fail to initialize the connection pool, host: %s, port: %d, %s", address, port, err.Error())
		log.Fatal(msg)
		panic(msg)
	}
	NebulaClient = &db.NebulaClientImpl{
		Pool:     pool,
		Username: username,
		Password: password,
	}
}

func initMysql() {
	MysqlClient = db.InitClientImpl(Dsn)
}

func init() {
	//initNebula()
	//initMysql()
}
