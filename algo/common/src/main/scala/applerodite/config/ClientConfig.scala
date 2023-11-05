package applerodite.config

import applerodite.dao.{GraphClient, MinioClientImpl, NebulaClientImpl, OSSClient}
import com.facebook.thrift.protocol.TCompactProtocol
import com.typesafe.config.{Config, ConfigFactory}
import org.apache.spark.SparkConf
import org.apache.spark.sql.SparkSession

// TODO: logs
object ClientConfig {

  var ossClient: OSSClient = _

  var spark: SparkSession = _

  var graphClient: GraphClient = _

  Init()

  def parseMinioConfig(conf: Config): MinioClientImpl.MinioConfig = {
    val minioConf = conf.getConfig("minio")
    MinioClientImpl.MinioConfig(minioConf.getString("url"), minioConf.getString("user"),
      minioConf.getString("password"), minioConf.getString("bucket"))
  }

  def parseNebulaConfig(conf: Config): NebulaClientImpl.NebulaConfig = {
    val nebulaConf = conf.getConfig("nebula")
    NebulaClientImpl.NebulaConfig(nebulaConf.getString("metaHost"), nebulaConf.getInt("metaPort"),
      nebulaConf.getString("graphdHost"), nebulaConf.getInt("graphdPort"),
      nebulaConf.getString("user"), nebulaConf.getString("password"))
  }

  def Init(): Unit = {
    var conf = ConfigFactory.parseString("akka.http.server.preview.enable-http2 = on")
      .withFallback(ConfigFactory.defaultApplication())
    if (System.getenv("CHS_ENV") == "pre") {
      conf = ConfigFactory.load("application-pre.conf")
    }

    ossClient = MinioClientImpl.Init(parseMinioConfig(conf))

    val sparkConf = new SparkConf()
      .set("spark.serializer", "org.apache.spark.serializer.KryoSerializer")
      .registerKryoClasses(Array[Class[_]](classOf[TCompactProtocol]))
    spark = SparkSession
      .builder()
      .master(conf.getConfig("spark").getString("url"))
      .config(sparkConf)
      .getOrCreate()
    graphClient = NebulaClientImpl.Init(parseNebulaConfig(conf))
  }
}
