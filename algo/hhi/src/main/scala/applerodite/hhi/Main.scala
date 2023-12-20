package applerodite.hhi

import applerodite.config.CommonService
import org.apache.spark.graphx.{Edge, Graph, VertexRDD}
import org.apache.spark.sql.Row

object Main extends Template {
  override def exec(svc: CommonService, param: Param): Seq[Row] = {
    val spark = svc.getSparkSession

    val library = GraphView.Node.LIBRARY.NAME
    val topic = GraphView.Node.LIBRARY.ATTR_TOPIC
    val belong2 = GraphView.Edge.BELONG2.NAME
    val depend = GraphView.Edge.DEPEND.NAME
    val release = GraphView.Node.RELEASE.NAME
    val sql = f"match (s:$library)<-[b1:$belong2]-(r1:$release)-[d:$depend]->(r2:$release)-[b2:$belong2]->(t:$library) " +
      f"return distinct id(s) as source, s.$library.$topic as s_topic, id(t) as target, t.$library.$topic as t_topic;"
    val ctrRes = svc.getGraphClient.gql(param.graphID, sql)
    if (ctrRes.isEmpty) {
      return Seq.empty
    }
    val res = ctrRes.get
    val edges = spark.sparkContext.parallelize((0 until res.rowsSize()).map(
      i => {
        Edge.apply(res.rowValues(i).get("source").asLong(), res.rowValues(i).get("target").asLong(), None)
      }
    ))
    val vertices: VertexRDD[String] = VertexRDD.apply(spark.sparkContext.parallelize((0 until res.rowsSize()).map(
      i => {
        List.apply(
          (res.rowValues(i).get("source").asLong(), res.rowValues(i).get("s_topic").asString()),
          (res.rowValues(i).get("target").asLong(), res.rowValues(i).get("t_topic").asString())
        )
      }
    ).toList.flatten).distinct())
    val graph = Graph.apply(vertices, edges)
    graph.outerJoinVertices(graph.degrees) {
      (_, t, d) => (t, d.getOrElse(0))
    }.vertices.map(v => v._2).groupBy(vd => vd._1).map(vds => {
      println(vds._1, vds._2)
      val r = ResultRow.apply(`class` = vds._1, score = vds._2.map(vd => math.pow(vd._2, 2)).sum / math.pow(vds._2.map(vd => vd._2).sum, 2))
      println(r)
      r
    }
    ).sortBy(r => r.`score`, ascending = false).map(r => r.toRow(`score` = f"${100 * r.`score`}%.2f" + "%")).collect()
  }
}
