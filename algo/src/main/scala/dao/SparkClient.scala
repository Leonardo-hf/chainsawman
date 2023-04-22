package dao

import model.RankPO

trait SparkClient {
  def degree(graph: String): (RankPO, Exception)
}
