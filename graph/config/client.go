package config

import (
	"chainsawman/graph/db"
	"time"
)

var NebulaClient db.NebulaClient

var MysqlClient db.MysqlClient

var RedisClient db.RedisClient

// TODO: 这种东西应该放到配置文件里
const (
	nebulaAddr     = "127.0.0.1"
	nebulaPort     = 9669
	nebulaUsername = "root"
	nebulaPasswd   = "nebula"

	MysqlAddr = "root:12345678@(localhost:3306)/graph?charset=utf8mb4&parseTime=True&loc=Local"

	//TODO 可以考虑拆分成两个redis，一个用于缓存，一个用来当队列
	redisAddr       = "localhost:6379"
	redisTopic      = "task"
	redisGroup      = "task_consumers"
	redisExpiration = time.Hour
)

func initDB() {
	NebulaClient = db.InitNebulaClient(&db.NebulaConfig{
		Addr:     nebulaAddr,
		Port:     nebulaPort,
		Username: nebulaUsername,
		Passwd:   nebulaPasswd,
	})
	MysqlClient = nil
	RedisClient = nil
}
func init() {

	//initMysql()

	//initDB()

}
