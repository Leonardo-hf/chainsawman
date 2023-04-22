package model

case class RankItem(node: String,
                    score: Double)

case class RankPO(ranks: Seq[RankItem])

