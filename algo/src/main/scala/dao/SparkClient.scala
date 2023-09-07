package dao

import org.apache.spark.sql.DataFrame
import service.{LouvainConfig, PRConfig, VoteConfig}

trait SparkClient {
  def degree(graphID: Long, edgesTag: Seq[String]): (DataFrame, Option[Exception])

  def pagerank(graphID: Long, edgesTag: Seq[String], cfg: PRConfig): (DataFrame, Option[Exception])

  def betweenness(graphID: Long, edgesTag: Seq[String]): (DataFrame, Option[Exception])

  def closeness(graphID: Long, edgesTag: Seq[String]): (DataFrame, Option[Exception])

  def voterank(graphID: Long, edgesTag: Seq[String], cfg: VoteConfig): (DataFrame, Option[Exception])

  def clusteringCoefficient(graphID: Long, edgesTag: Seq[String]): (Double, Option[Exception])

  def louvain(graphID: Long, edgesTag: Seq[String], cfg: LouvainConfig): (DataFrame, Option[Exception])

  def depth(graphID: Long, edgeTags: Seq[String]): (DataFrame, Option[Exception])

  def ecology(graphID: Long, edgeTags: Seq[String]): (DataFrame, Option[Exception])
}
