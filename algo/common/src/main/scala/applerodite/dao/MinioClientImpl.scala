package applerodite.dao

import io.minio.{BucketExistsArgs, MakeBucketArgs, MinioClient, PutObjectArgs}

import java.io.ByteArrayInputStream


object MinioClientImpl extends OSSClient {

  case class MinioConfig(endpoint: String, user: String, passwd: String, bucket: String)

  var minioClient: MinioClient = _

  var bucket: String = _

  def Init(config: MinioConfig): OSSClient = {
    minioClient =
      MinioClient.builder()
        .endpoint(config.endpoint)
        .credentials(config.user, config.passwd)
        .build()
    bucket = config.bucket
    if (!minioClient.bucketExists(BucketExistsArgs.builder().bucket(bucket).build())) {
      minioClient.makeBucket(MakeBucketArgs.builder().bucket(bucket).build())
    }
    this
  }

  override def upload(name: String, content: String): (String, Option[Exception]) = {
    val stream = new ByteArrayInputStream(content.getBytes())
    minioClient.putObject(PutObjectArgs.builder().bucket(bucket).`object`(name).stream(
      stream, stream.available(), -1
    ).build())
    stream.close()
    (name, Option.empty)
  }
}
