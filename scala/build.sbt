import Dependencies._

resolvers += "Sonatype OSS Snapshots" at
  "https://oss.sonatype.org/content/repositories/releases"

libraryDependencies += "com.storm-enroute" %% "scalameter" % "0.8.2"

testFrameworks += new TestFramework("org.scalameter.ScalaMeterFramework")

parallelExecution in Test := false

lazy val root = (project in file(".")).
  settings(
    inThisBuild(List(
      organization := "com.github.myyk",
      scalaVersion := "2.12.4",
      version      := "0.1.0-SNAPSHOT"
    )),
    name := "go-vs-scala",
    libraryDependencies += scalaTest % Test
  )
