package applerodite

import applerodite.Country.SCHEMA_STRANGLE
import applerodite.config.AlgoConstants.{ARTIFACT_COL, NODE_ID_COL, SCHEMA_SOFTWARE, VERSION_COL}
import applerodite.config.{AlgoConstants, ClientConfig}
import applerodite.util.CSVUtil
import com.alibaba.fastjson.{JSON, JSONObject}
import org.apache.spark.sql.catalyst.encoders.RowEncoder
import org.apache.spark.sql.functions._
import org.apache.spark.sql.{Dataset, Row}

import scala.collection.mutable.ListBuffer

object Main {
  def getDouble(row: Row, i: Int): Double = {
    if (row.isNullAt(i)) {
      return 0
    }
    row.getDouble(i)
  }

  def getCandidates(graphID: String, weights: Seq[Double]): Seq[String] = {
    val limit = 2000
    val spark = ClientConfig.spark
    // 获得各个算法结果，加权平均
    val empty = spark.sqlContext.createDataFrame(spark.sparkContext.parallelize(List[Row]()), SCHEMA_SOFTWARE)
    var res: Dataset[Row] = empty
    val j = new JSONObject()
    j.put("graphID", graphID)
    j.put("iter", limit)
    val param = Array[String](j.toJSONString)
    val algo = List[Impact](Voterank, Depth, Proxy, Stable)
    for (i <- 0 until 4) {
      if (weights.apply(i) == 0) {
        res = res.join(empty, Seq(NODE_ID_COL, ARTIFACT_COL, VERSION_COL), joinType = "left_outer")
      } else {
        var df = algo.apply(i).main(param).limit(limit)
        // 变成正值
        var min = df.collect().map(r => r.getDouble(3)).min
        if (min < 0) {
          min = -min + 1
        } else {
          min = 1
        }
        val rowEncoder = RowEncoder(df.schema)
        df = df.map(r => Row.apply(r.get(0), r.get(1), r.get(2), math.log(r.getDouble(3) + min) / math.log(2)))(rowEncoder)
        val dfList = df.collect()
        val v = dfList.map(r => r.getDouble(3)).sum / dfList.length
        df = df.map(r => Row.apply(r.get(0), r.get(1), r.get(2), r.getDouble(3) * weights.apply(i) / v))(rowEncoder)
        res = res.join(df, Seq(NODE_ID_COL, ARTIFACT_COL, VERSION_COL), joinType = "left_outer")
      }
    }
    // 求和
    res = res.map(r => Row.apply(r.get(0), r.get(1), r.get(2), getDouble(r, 3) + getDouble(r, 4) + getDouble(r, 5) + getDouble(r, 6)))(RowEncoder(SCHEMA_SOFTWARE))
    res.printSchema()
    res.orderBy(desc(AlgoConstants.SCORE_COL)).select(ARTIFACT_COL).distinct().collect().map(r => r.getString(0))
  }

  def main(args: Array[String]): Unit = {
    if (args.length < 1) {
      return
    }
    val json = JSON.parseObject(args.apply(0))
    val graphID: String = json.getString("graphID")
    val target: String = json.getString("target")
    val weightArray = json.getJSONArray("impactWeights")
    val weights = ListBuffer[Double]()
    for (i <- 0 until 4) {
      weights.append(weightArray.getDoubleValue(i))
    }

    val spark = ClientConfig.spark
    val candidates = getCandidates(graphID, weights)
    println(candidates.mkString("Array(", ", ", ")"))
    val weightArray2 = json.getJSONArray("strangleWeights")
    val weights2 = ListBuffer[Double]()
    for (i <- 0 until 1) {
      weights2.append(weightArray2.getDoubleValue(i))
    }
    val j = new JSONObject()
    j.put("graphID", graphID)
    j.put("libraries", candidates)
    val param = Array[String](j.toJSONString)
    val empty = spark.sqlContext.createDataFrame(spark.sparkContext.parallelize(List[Row]()), SCHEMA_STRANGLE)
    var strangleRisks = empty
    // 获得各个算法结果
    val strangles = List[Strangle](Country)
    for (i <- 0 until 1) {
      if (weights.apply(i) == 0) {
        strangleRisks = strangleRisks.join(empty, Seq(ARTIFACT_COL), joinType = "left_outer")
      } else {
        val df = strangles.apply(i).main(param)
        strangleRisks = strangleRisks.join(df, Seq(ARTIFACT_COL), joinType = "left_outer")
      }
    }

    // 加权平均
    strangleRisks = strangleRisks.map(r => Row.apply(r.get(0), weights2.indices.map(i => getDouble(r, i + 1) * weights2.apply(i)).sum))(RowEncoder(SCHEMA_STRANGLE))
    ClientConfig.ossClient.upload(name = target, content = CSVUtil.df2CSV(strangleRisks))
    spark.stop()
  }
}
