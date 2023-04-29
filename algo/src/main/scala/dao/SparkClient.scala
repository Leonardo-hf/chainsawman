package dao

import org.apache.spark.sql.DataFrame
import service.{LouvainConfig, PRConfig}

trait SparkClient {
  def degree(graphID: Long): (DataFrame, Option[Exception])

  def pagerank(graphID: Long, cfg: PRConfig): (DataFrame, Option[Exception])

  def betweenness(graphID: Long): (DataFrame, Option[Exception])

  def closeness(graphID: Long): (DataFrame, Option[Exception])

  def voterank(graphID: Long): (DataFrame, Option[Exception])

  def clusteringCoefficient(graphID: Long): (Double, Option[Exception])

  def louvain(graphID: Long, cfg: LouvainConfig): (DataFrame, Option[Exception])
}
