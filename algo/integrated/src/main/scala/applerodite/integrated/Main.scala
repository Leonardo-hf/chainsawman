package applerodite.integrated

import applerodite.config.CommonService
import applerodite.{breadth, depth, mediation, stability}
import org.apache.spark.sql.Row

object Main extends Template {

  override def exec(svc: CommonService, param: Param): Seq[Row] = {
    // 获得各个算法结果，加权求和
    val p1 = breadth.Param.apply(`iter` = 2000, graphID = param.graphID, target = param.target)
    val p2 = depth.Param.apply(graphID = param.graphID, target = param.target)
    val p3 = mediation.Param.apply(graphID = param.graphID, target = param.target)
    val p4 = stability.Param.apply(graphID = param.graphID, target = param.target, `radius` = 3)
    val spark = svc.getSparkSession
    val algo: Seq[Seq[Row]] = Seq.apply(
      if (param.`weights`.head == 0) breadth.Main.exec(svc, p1) else Seq.empty,
      if (param.`weights`.apply(1) == 0) depth.Main.exec(svc, p2) else Seq.empty,
      if (param.`weights`.apply(2) == 0) mediation.Main.exec(svc, p3) else Seq.empty,
      if (param.`weights`.apply(3) == 0) stability.Main.exec(svc, p4) else Seq.empty)
    val res = algo.map(a => spark.sqlContext.createDataFrame(spark.sparkContext.parallelize(a), Constants.SCHEMA))
      .reduce((p, q) => p.join(q, Seq.apply(Constants.COL_ID, Constants.COL_ARTIFACT, Constants.COL_VERSION), joinType = "left_outer"))
    res.collect().map(
      r => ResultRow.apply(`artifact` = r.getString(1), `version` = r.getString(2),
        `score` = (3 until 7).map(i => param.`weights`.apply(i - 3) * getDouble(r, i)).sum, `id` = r.getLong(0))
    ).map(
      r => r.toRow(`score` = f"${r.`score`}%.4f")
    )
  }

  def getDouble(row: Row, i: Int): Double = {
    if (row.isNullAt(i)) {
      return 0
    }
    row.getDouble(i)
  }
}
