package dao

import com.facebook.thrift.protocol.TCompactProtocol
import com.vesoft.nebula.algorithm.config.PRConfig
import com.vesoft.nebula.algorithm.lib.PageRankAlgo
import com.vesoft.nebula.connector.connector.NebulaDataFrameReader
import com.vesoft.nebula.connector.{NebulaConnectionConfig, ReadNebulaConfig}
import model.RankPO
import org.apache.spark.SparkConf
import org.apache.spark.sql.{DataFrame, SparkSession}

object SparkClientImpl extends SparkClient {

  var spark: SparkSession = _

  var nebulaCfg: NebulaConnectionConfig = _

  def Init(): SparkClient = {
    val sparkConf = new SparkConf()
      .set("spark.serializer", "org.apache.spark.serializer.KryoSerializer")
      .registerKryoClasses(Array[Class[_]](classOf[TCompactProtocol]))
    spark = SparkSession
      .builder()
      .master("local")
      .config(sparkConf)
      .getOrCreate()
    nebulaCfg =
      NebulaConnectionConfig
        .builder()
        .withMetaAddress("127.0.0.1:9669")
        .withTimeout(6000)
        .withConenctionRetry(2)
        .build()
    this
  }

  def readDf(graph: Long): DataFrame = {
    val nebulaReadEdgeConfig: ReadNebulaConfig = ReadNebulaConfig
      .builder()
      .withSpace(graph.toString)
      .withLabel("sedges")
      .withNoColumn(true)
      //      .withLimit(2000)
      //      .withPartitionNum(100)
      .build()
    spark.read.nebula(nebulaCfg, nebulaReadEdgeConfig).loadEdgesToDF()
  }

  override def degree(graph: Long): (RankPO, Option[Exception]) = {
    val pageRankConfig = PRConfig(3, 0.85)
    val df = PageRankAlgo.apply(spark, readDf(graph), pageRankConfig, hasWeight = false)
    //    df.map(row=>row)
    df.show()
    (RankPO(ranks = List.empty), Option.empty)
  }


}
