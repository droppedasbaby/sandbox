package com.grewalas.notesgraph

import files.{ FileManager, LinkFinder }
import graph.Graph

@main def run(args: String*): Unit = {
  val path = "../notes/"

  val files = FileManager(path).getFiles
  val linkFinder = LinkFinder(files)

  val graph = Graph(path)
  println("--------------------graph.nodes--------------------")
  graph.nodes.foreach(n =>
    println(n._1 + " -> " + n._2.Content.substring(0, 32).replace("\n", " "))
  )
  println("--------------------graph.edges--------------------")
  graph.edges.foreach(e => println(e._1.name + " -> " + e._2.map(_.name)))
  println("--------------------graph.edgesSet--------------------")
  graph.edgesSet.toSeq
    .sortBy(_._1.name)
    .foreach(e => println(e.source.name + " -> " + e.target.name))
}
