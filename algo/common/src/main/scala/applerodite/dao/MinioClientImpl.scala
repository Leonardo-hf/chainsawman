package applerodite.dao

import com.typesafe.config.Config
import io.minio.{BucketExistsArgs, MakeBucketArgs, MinioClient, PutObjectArgs}

import java.io.ByteArrayInputStream

object MinioClientImpl extends OSSClient {
  var minioClient: MinioClient = _

  var bucket: String = _

  def Init(config: Config): OSSClient = {
    minioClient =
      MinioClient.builder()
        .endpoint(config.getString("url"))
        .credentials(config.getString("user"), config.getString("password"))
        .build()
    bucket = config.getString("bucket")
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
