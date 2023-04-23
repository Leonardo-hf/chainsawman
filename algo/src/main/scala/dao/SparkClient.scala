package dao

import model.RankPO

trait SparkClient {
  def degree(graph: Long): (RankPO, Option[Exception])
}
