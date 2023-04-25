package dao

import com.facebook.thrift.protocol.TCompactProtocol
import com.vesoft.nebula.connector.NebulaConnectionConfig
import model.{RankItem, Rank}
import org.apache.spark.SparkConf
import org.apache.spark.graphx.Graph
import org.apache.spark.graphx.lib._
import org.apache.spark.sql.SparkSession
import service.PRConfig
import util.GraphUtil

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
    spark.sparkContext.setLogLevel("WARN")
    nebulaCfg =
      NebulaConnectionConfig
        .builder()
        .withMetaAddress("127.0.0.1:9559")
        .withConenctionRetry(2)
        .withExecuteRetry(2)
        .withTimeout(6000)
        .build()
    this
  }

  override def degree(graphID: Long): (Rank, Option[Exception]) = {
    val graph: Graph[None.type, Double] = GraphUtil.loadInitGraph(graphID, hasWeight = false)
    (Rank(ranks = graph.degrees.map(r => RankItem.apply(r._1, r._2)).collect()), Option.empty)
  }

  override def pagerank(graphID: Long, cfg: PRConfig): (Rank, Option[Exception]) = {
    val graph: Graph[None.type, Double] = GraphUtil.loadInitGraph(graphID, hasWeight = false)
    val prResultRDD = PageRank.run(graph, cfg.iter.toInt, cfg.prob).vertices
    //    val schema = StructType(
    //      List(
    //        StructField(AlgoConstants.ALGO_ID_COL, LongType, nullable = false),
    //        StructField(AlgoConstants.PAGERANK_RESULT_COL, DoubleType, nullable = true)
    //      ))
    //    val algoResult = spark.sqlContext
    //      .createDataFrame(prResultRDD, schema)
    (Rank(ranks = prResultRDD.map(r => RankItem.apply(r._1, r._2)).collect()), Option.empty)
  }


}
