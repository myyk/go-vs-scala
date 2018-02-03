package rant1

object Subsets {
  def subsets[T](data: Set[T], size: Int): Iterator[Set[T]] ={
     data.subsets(size)
  }
}
