package applerodite

import applerodite.config.ClientConfig
import com.alibaba.fastjson.JSON
import com.vesoft.nebula.client.graph.data.ResultSet
import org.apache.spark.sql.types.{StringType, StructField, StructType}
import org.apache.spark.sql.{DataFrame, Row}

import scala.collection.mutable.ListBuffer


trait Strangle {

  def main(args: Array[String]): DataFrame

  val SCHEMA_STRANGLE: StructType = StructType(
    List(
      StructField(AlgoConstants.ARTIFACT_COL, StringType, nullable = false),
      StructField(AlgoConstants.SCORE_COL, StringType, nullable = false)
    ))
}


object Country extends Strangle {

  def main(args: Array[String]): DataFrame = {
    if (args.length < 1) {
      return null
    }

    val json = JSON.parseObject(args.apply(0))
    val graphID: String = json.getString("graphID")
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
        if (country.nonEmpty && country != "None") {
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

    spark.sqlContext.createDataFrame(spark.sparkContext.parallelize(res), SCHEMA_STRANGLE)

  }
}
