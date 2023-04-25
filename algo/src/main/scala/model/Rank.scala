package model

case class RankItem(nodeID: Long,
                    score: Double)

case class Rank(ranks: Seq[RankItem])

