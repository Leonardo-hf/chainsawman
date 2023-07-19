package config

import dao.{MinioClientImpl, MysqlClient, MysqlClientImpl, OSSClient, SparkClient, SparkClientImpl}


object ClientConfig {

  var mysqlClient: MysqlClient = _

  var sparkClient: SparkClient = _

  var ossClient: OSSClient= _


  def Init(): Unit = {
    mysqlClient = MysqlClientImpl.Init()
    sparkClient = SparkClientImpl.Init()
    ossClient = MinioClientImpl.Init()
  }
}
