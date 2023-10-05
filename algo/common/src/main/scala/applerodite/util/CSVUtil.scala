package applerodite.util

import org.apache.spark.sql.DataFrame

object CSVUtil {
  val sep: String = System.lineSeparator()

  def df2CSV(df: DataFrame): String = {
    val content: StringBuilder = new StringBuilder
    content.append(df.columns.mkString(",")).append(sep)
    df.rdd.collect().foreach(r=>content.append(r.toSeq.mkString(",")).append(sep))
    content.toString()
  }

}
