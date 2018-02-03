package rant1

import org.scalameter.api._
import scala.util.Random
import Subsets._

object SubsetBenchmark extends Bench.ForkedTime {
  def subsets[T](data: Set[T], size: Int): Iterator[Set[T]] ={
     data.subsets(size)
  }

  val input = for {
    _ <- 0 until 1000
  } yield {
    Random.nextString(6)
  }

  val inputs = for {
    size <- Gen.range("size")(500, 1500, 500)
  } yield {
    for {
      _ <- (0 until size).toSet[Int]
    } yield {
      Random.nextString(6)
    }
  }

  measure method "subsets" in {
    using(inputs) curve("Range") in {
      subsets(_, 2).toArray
    }
  }
}
