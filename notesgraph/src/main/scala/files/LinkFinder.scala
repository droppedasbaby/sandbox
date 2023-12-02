package com.grewalas.notesgraph
package files

import scala.util.matching.Regex

case class LinkFinder(files: Map[String, String]) {
  private val internalLinkPattern: Regex = """\[\[([^\[\]]+)\]\]""".r
  private val externalLinkPattern: Regex = """\[(.+?)\]\((.+?)\)""".r

  def getInternalLinks: Seq[(String, String)] = files.toSeq.flatMap { case (fileName, content) =>
    internalLinkPattern
      .findAllMatchIn(content)
      .map(matched => (fileName, matched.group(1)))
  }

  def getExternalLinks: Seq[(String, String)] = files.toSeq.flatMap { case (fileName, content) =>
    externalLinkPattern
      .findAllMatchIn(content)
      .map(matched => (fileName, matched.group(2)))
  }
}
