package config

import (
	"chainsawman/graph/db"
	"fmt"

	nebula "github.com/vesoft-inc/nebula-go/v3"
)

var NebulaClient db.NebulaClient

const (
	address  = "127.0.0.1"
	port     = 9669
	username = "root"
	password = "nebula"
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

func init() {
	initNebula()
}
