package dao

trait MysqlClient {
  def createAlgo(algo: model.AlgoPO): (Long, Option[Exception])

  def multiCreateAlgoParams(param: Seq[model.AlgoParamPO]): (Int, Option[Exception])

  def queryAlgo(): (List[model.AlgoPOWithParams], Option[Exception])

  def dropAlgo(name: String): (Int, Option[Exception])
}