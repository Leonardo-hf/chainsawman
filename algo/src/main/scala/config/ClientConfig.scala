package config

import dao.{MysqlClient, MysqlClientImpl, SparkClient}


object ClientConfig {

  var mysqlClient: MysqlClient = _

  var sparkClient: SparkClient = _

  def init(): Unit = {
    mysqlClient = MysqlClientImpl.Init()
  }
}
