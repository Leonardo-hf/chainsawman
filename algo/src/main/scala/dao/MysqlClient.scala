package dao

trait MysqlClient {
  def createAlgo(algo: model.AlgoPO): (Int, Option[Exception])

  def queryAlgo(): (List[model.AlgoPO], Option[Exception])

  def dropAlgo(name: String): (Int, Option[Exception])
}