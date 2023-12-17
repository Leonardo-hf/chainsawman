package applerodite.diversity

import applerodite.config.CommonService
import com.vesoft.nebula.client.graph.data.ResultSet
import org.apache.spark.sql.Row

object Main extends Template {

  override def exec(svc: CommonService, param: Param): Seq[Row] = {
    val graphService = svc.getGraphClient
    val developer = GraphView.Node.DEVELOPER.NAME
    val contribute = GraphView.Edge.CONTRIBUTE.NAME
    val library = GraphView.Node.LIBRARY.NAME
    val artifact = GraphView.Node.LIBRARY.ATTR_ARTIFACT
    val maintain = GraphView.Edge.MAINTAIN.NAME
    val organization = GraphView.Node.ORGANIZATION.NAME
    val host = GraphView.Edge.HOST.NAME

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
    param.`libraries`.map(s => {
      var score1, score2 = 1.0
      val ctrRes = graphService.gql(param.graphID, s"match (u:$developer)-[c:$contribute]->(p:$library{$artifact:$yh$s$yh}) return u, c;")
      if (ctrRes.isDefined) {
        score1 = getScore1(ctrRes.get)
      }
      val mtnRes = graphService.gql(param.graphID, s"match (u:$developer)-[m:$maintain]->(o:$organization)-[h:$host]->(p:$library{$artifact:$yh$s$yh}) return u;")
      if (mtnRes.isDefined) {
        score2 = getScore2(mtnRes.get)
      }
      ResultRow.apply(`artifact` = s, s1 = score1 * 100, s2 = score2 * 100, score = (1 - score1) * (1 - score2) * 100)
    }).sortBy(r => -r.`score`).map(r => r.toRow(s1 = f"${r.`s1`}%.2f" + "%", s2 = f"${r.`s2`}%.2f" + "%", `score` = f"${r.`score`}%.2f" + "%"))
  }
}
