package com.grewalas.notesgraph
package files

import java.io.File
import scala.annotation.tailrec
import scala.io.Source
import scala.util.Using

case class FileManager(path: String) {
  private val MdExtension = ".md"
  def getFiles: Map[String, String] = walkTree(Seq(new File(path)), Map.empty)

  @tailrec
  private def walkTree(dirs: Seq[File], acc: Map[String, String]): Map[String, String] =
    dirs match {
      case Nil                                                                           => acc
      case dir :: tail if dir.isDirectory && !dir.isHidden                               =>
        walkTree(tail ++ Option(dir.listFiles()).getOrElse(Array.empty[File]), acc)
      case file :: tail if file.isFile && !file.isHidden && file.getName.endsWith(".md") =>
        val content = Using.resource(Source.fromFile(file))(_.getLines.mkString("\n"))
        walkTree(tail, acc + (file.getName.stripSuffix(MdExtension) -> content))
      case _ :: tail                                                                     =>
        walkTree(tail, acc)
    }
}
