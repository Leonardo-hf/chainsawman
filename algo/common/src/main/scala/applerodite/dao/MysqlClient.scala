package applerodite.dao

trait MysqlClient {
  def updateStatusByOutput(output: String): (Int, Option[Exception])
}
