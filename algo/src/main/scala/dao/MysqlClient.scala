package dao

trait MysqlClient {
  def createAlgo(algo: model.AlgoPO): (Int, Exception)

  def queryAlgo(): (List[model.AlgoPO], Exception)

  def dropAlgo(name: String): (Int, Exception)
}