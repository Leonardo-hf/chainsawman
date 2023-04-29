package config

import akka.actor.ActorSystem
import akka.grpc.GrpcClientSettings
import dao.{MysqlClient, MysqlClientImpl, SparkClient, SparkClientImpl}
import service.fileClient

import scala.concurrent.ExecutionContextExecutor


object ClientConfig {

  var mysqlClient: MysqlClient = _

  var sparkClient: SparkClient = _

  var fileRPC: fileClient = _

  implicit val sys: ActorSystem = ActorSystem("FileServiceClient")
  implicit val ec: ExecutionContextExecutor = sys.dispatcher


  def Init(): Unit = {
    mysqlClient = MysqlClientImpl.Init()
    sparkClient = SparkClientImpl.Init()
    val clientSettings = GrpcClientSettings.connectToServiceAt("127.0.0.1", 8080).withTls(false)
    // val clientSettings = GrpcClientSettings.fromConfig(GreeterService.name)
    fileRPC = fileClient(clientSettings)
  }
}
