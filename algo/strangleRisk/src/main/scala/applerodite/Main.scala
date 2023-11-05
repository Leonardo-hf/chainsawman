package applerodite

import applerodite.config.{AlgoConstants, ClientConfig}
import applerodite.util.CSVUtil
import com.alibaba.fastjson.JSON
import com.vesoft.nebula.client.graph.data.ResultSet
import org.apache.spark.sql.Row
import org.apache.spark.sql.functions._
import org.apache.spark.sql.types.{StringType, StructField, StructType}

import scala.collection.mutable.ListBuffer

object Main {
  val SCHEMA_STRANGLE: StructType = StructType(
    List(
      StructField("library", StringType, nullable = false),
      StructField("s1", StringType, nullable = false),
      StructField("s2", StringType, nullable = false),
      StructField(AlgoConstants.SCORE_COL, StringType, nullable = false)
    ))

  def main(args: Array[String]): Unit = {
    if (args.length < 1) {
      return
    }

    val json = JSON.parseObject(args.apply(0))
    val graphID: String = json.getString("graphID")
    val target: String = json.getString("target")

    // 候选识别软件列表
    val jCandidates = json.getJSONArray("libraries")
    val candidates = ListBuffer[String]()
    for (i <- 0 until jCandidates.size()) {
      candidates.append(jCandidates.getString(i))
    }
    val spark = ClientConfig.spark

    def getScore1(res: ResultSet): Double = {
      var known = 0L
      var china = 0L
      for (i <- 0 until res.rowsSize()) {
        val country = res.rowValues(i).get("u").asNode().properties("developer").get("location").asString()
        val commits = res.rowValues(i).get("c").asRelationship().properties().get("commits").asLong()

        if (country != "None") {
          known += commits
        }
        if (country == "China") {
          china += commits
        }
      }

      if (known != 0) {
        return 1.0 * china / known
      }
      0
    }

    def getScore2(res: ResultSet): Double = {
      var known = 0L
      var china = 0L
      for (i <- 0 until res.rowsSize()) {
        val country = res.rowValues(i).get("u").asNode().properties("developer").get("location").asString()
        if (country != "None") {
          known += 1
        }
        if (country == "China") {
          china += 1
        }
      }
      if (known != 0) {
        return 1.0 * china / known
      }
      0
    }

    val yh = "\""
    val res = candidates.map(s => {
      var score1, score2 = 1.0
      val ctrRes = ClientConfig.graphClient.gql(graphID, s"match (u:developer)-[c:contribute]->(p:library{artifact:$yh$s$yh}) return u, c;")
      if (ctrRes.isDefined) {
        score1 = getScore1(ctrRes.get)
      }
      val mtnRes = ClientConfig.graphClient.gql(graphID, s"match (u:developer)-[m:maintain]->(o:organization)-[h:host]->(p:library{artifact:$yh$s$yh}) return u;")
      if (mtnRes.isDefined) {
        score2 = getScore2(mtnRes.get)
      }
      Row.apply(s, f"$score1%.4f", f"$score2%.4f", f"${(1 - score1) * (1 - score2) * 100}%.2f" + "%")
    })

    val df = spark.sqlContext.createDataFrame(spark.sparkContext.parallelize(res), SCHEMA_STRANGLE).orderBy(desc(AlgoConstants.SCORE_COL))
    ClientConfig.ossClient.upload(name = target, content = CSVUtil.df2CSV(df))
    spark.close()
  }
}
