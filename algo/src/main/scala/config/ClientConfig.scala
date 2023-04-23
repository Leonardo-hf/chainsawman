package config

import dao.{MysqlClient, MysqlClientImpl, SparkClient, SparkClientImpl}


object ClientConfig {

  var mysqlClient: MysqlClient = _

  var sparkClient: SparkClient = _

  def Init(): Unit = {
    mysqlClient = MysqlClientImpl.Init()
    sparkClient = SparkClientImpl.Init()
  }
}
