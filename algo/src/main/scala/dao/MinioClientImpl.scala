package dao

import com.typesafe.config.Config
import io.minio.{BucketExistsArgs, GetObjectArgs, MakeBucketArgs, MinioClient, PutObjectArgs, UploadObjectArgs}

import java.io.ByteArrayInputStream
import java.util.UUID

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
    val id = name + '-' + UUID.randomUUID().toString
    minioClient.putObject(PutObjectArgs.builder().bucket(bucket).`object`(id).stream(
      stream, stream.available(), -1
    ).build())
    stream.close()
    (id, Option.empty)
  }
}
