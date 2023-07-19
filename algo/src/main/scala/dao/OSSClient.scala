package dao

trait OSSClient {
  def upload(name: String, content: String): (String, Option[Exception])
}
