package com.grewalas.notesgraph
package graph

import files.{ FileManager, LinkFinder }

case class Graph(nodes: Map[String, Node], edges: Map[Node, Set[Node]], edgesSet: Set[Edge])

object Graph {
  def apply(path: String): Graph = {
    val fileManager = FileManager(path)
    val files = fileManager.getFiles
    val linkFinder = LinkFinder(files)

    val nodes = files.map { case (fileName, fileContents) =>
      fileName -> Node(fileName, fileContents)
    }

    val edgesSet = (linkFinder.getExternalLinks ++ linkFinder.getInternalLinks).flatMap {
      case (from, to) =>
        nodes.get(from).flatMap(fromNode => nodes.get(to).map(toNode => Edge(fromNode, toNode)))
    }.toSet
    val edges = edgesSet.groupBy(_.source).map(e => e._1 -> e._2.map(_.target))

    Graph(nodes, edges, edgesSet)
  }
}
