package applerodite.config

import applerodite.dao.{GraphClient, MinioClientImpl, MysqlClient, MysqlClientImpl, MysqlConfig, NebulaClientImpl, NebulaConfig, OSSClient}
import com.facebook.thrift.protocol.TCompactProtocol
import com.typesafe.config.{Config, ConfigFactory}
import org.apache.spark.SparkConf
import org.apache.spark.sql.SparkSession

trait CommonService {
  def getOSSClient: OSSClient

  def getSparkSession: SparkSession

  def getGraphClient: GraphClient

  def getMysqlClient: MysqlClient

}

// TODO: logs
object CommonServiceImpl extends CommonService {

  var ossClient: OSSClient = _

  var spark: SparkSession = _

  var graphClient: GraphClient = _

  var mysqlClient: MysqlClient = _

  def parseMinioConfig(conf: Config): MinioClientImpl.MinioConfig = {
    val minioConf = conf.getConfig("minio")
    MinioClientImpl.MinioConfig(minioConf.getString("url"), minioConf.getString("user"),
      minioConf.getString("password"), minioConf.getString("bucket"))
  }

  def parseNebulaConfig(conf: Config): NebulaConfig = {
    val nebulaConf = conf.getConfig("nebula")
    NebulaConfig(nebulaConf.getString("metaHost"), nebulaConf.getInt("metaPort"),
      nebulaConf.getString("graphdHost"), nebulaConf.getInt("graphdPort"),
      nebulaConf.getString("user"), nebulaConf.getString("password"))
  }

  def parseMysqlConfig(conf: Config) : MysqlConfig = {
    val mysqlConf = conf.getConfig("mysql")
    MysqlConfig.apply(driver = mysqlConf.getString("driver"), url = mysqlConf.getString("url"),
      user = mysqlConf.getString("user"), password = mysqlConf.getString("password"))
  }

  Init()

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
    graphClient = new NebulaClientImpl().GetGraphClient(parseNebulaConfig(conf), spark)
    mysqlClient = MysqlClientImpl.Init(parseMysqlConfig(conf))
  }

  override def getOSSClient: OSSClient = ossClient

  override def getSparkSession: SparkSession = spark

  override def getGraphClient: GraphClient = graphClient

  override def getMysqlClient: MysqlClient = mysqlClient
}