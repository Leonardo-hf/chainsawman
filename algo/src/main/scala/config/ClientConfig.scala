package config

import com.typesafe.config.Config
import dao.{MinioClientImpl, MysqlClient, MysqlClientImpl, OSSClient, SparkClient, SparkClientImpl}


object ClientConfig {

  var mysqlClient: MysqlClient = _

  var sparkClient: SparkClient = _

  var ossClient: OSSClient= _


  def Init(config: Config): Unit = {
//    mysqlClient = MysqlClientImpl.Init()
    sparkClient = SparkClientImpl.Init(config.getConfig("spark"))
    ossClient = MinioClientImpl.Init(config.getConfig("minio"))
  }
}
