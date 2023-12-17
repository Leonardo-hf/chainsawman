package applerodite.degree

import applerodite.config.CommonService
import org.apache.spark.graphx.Graph
import org.apache.spark.sql.Row

object Main extends Template {
  override def exec(svc: CommonService, param: Param): Seq[Row] = {
    val graph: Graph[None.type, Double] = svc.getGraphClient.loadInitGraph(param.graphID, hasWeight = false)
    graph.degrees.map(r => ResultRow(`id` = r._1, `score` = r._2)).sortBy(r => -r.`score`).map(r => r.toRow()).collect()
  }
}
