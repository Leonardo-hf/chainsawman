package applerodite.config

import applerodite.dao.{GraphClient, MinioClientImpl, NebulaClientImpl, OSSClient}
import com.facebook.thrift.protocol.TCompactProtocol
import com.typesafe.config.ConfigFactory
import org.apache.spark.SparkConf
import org.apache.spark.sql.SparkSession


object ClientConfig {

  var ossClient: OSSClient = _

  var spark: SparkSession = _

  var graphClient: GraphClient = _

  Init()

  def Init(): Unit = {
    var conf = ConfigFactory.parseString("akka.http.server.preview.enable-http2 = on")
      .withFallback(ConfigFactory.defaultApplication())
    if (System.getenv("CHS_ENV") == "pre") {
      conf = ConfigFactory.load("application-pre.conf")
    }
        ossClient = MinioClientImpl.Init(conf.getConfig("minio"))
        val sparkConf = new SparkConf()
          .set("spark.serializer", "org.apache.spark.serializer.KryoSerializer")
          .registerKryoClasses(Array[Class[_]](classOf[TCompactProtocol]))
        spark = SparkSession
          .builder()
          .master(conf.getConfig("spark").getString("url"))
          .config(sparkConf)
          .getOrCreate()
    graphClient = NebulaClientImpl.Init(conf.getConfig("nebula"))
  }
}
