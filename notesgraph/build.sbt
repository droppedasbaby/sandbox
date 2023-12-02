ThisBuild / version := "0.1.0-SNAPSHOT"

ThisBuild / scalaVersion := "3.3.1"

lazy val root = (project in file("."))
  .settings(
    name := "notesgraph",
    idePackagePrefix := Some("com.grewalas.notesgraph")
  )

libraryDependencies ++= Seq(
  "com.vladsch.flexmark" % "flexmark-all" % "0.64.6"
)
