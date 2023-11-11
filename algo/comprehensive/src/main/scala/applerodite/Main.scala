package applerodite

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

  def main(args: Array[String]): Unit = {
    if (args.length < 1) {
      return
    }
    val json = JSON.parseObject(args.apply(0))
    val graphID: String = json.getString("graphID")
    val target: String = json.getString("target")
    val weightArray = json.getJSONArray("weights")
    val weights = ListBuffer[Double]()
    for (i <- 0 until 4) {
      weights.append(weightArray.getDoubleValue(i))
    }

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
    val df = res.orderBy(desc(AlgoConstants.SCORE_COL))
    ClientConfig.ossClient.upload(name = target, content = CSVUtil.df2CSV(df))
    spark.close()
  }
}
