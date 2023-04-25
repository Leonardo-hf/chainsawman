package dao

import model.Rank
import service.PRConfig

trait SparkClient {
  def degree(graphID: Long): (Rank, Option[Exception])

  def pagerank(graphID: Long, cfg: PRConfig): (Rank, Option[Exception])
}
