package util

import org.apache.spark.sql.DataFrame
import com.google.protobuf.ByteString

object CSVUtil {
  val sep: String = System.lineSeparator()

  def df2CSV(df: DataFrame): ByteString = {
    val content: StringBuilder = new StringBuilder
    content.append(df.columns.mkString(",")).append(sep)
    df.rdd.foreach(r=>content.append(r.toSeq.mkString(",")).append(sep))
    ByteString.copyFromUtf8(content.toString())
  }

}
