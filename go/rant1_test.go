package rant1

import (
  "math/big"
	"math/rand"
	"testing"
	"time"
  "log"
)

// Run with command:
//    go test -bench .

var Result [][][]string

func BenchmarkSubsets(b *testing.B) {
	k := 2
  input := make([]string, 31)
  for i := range input  {
    input[i] = RandStringRunes(8)
  }
  // log.Println(len(input))

	// n choose k should be output size
	outputLen := new(big.Int)
	outputLen.Binomial(int64(len(input)), int64(k))
  log.Println(outputLen)

  // a := make([][]string, int(outputLen.Int64()))

	b.ResetTimer()

  for i := 0; i < b.N; i++ {
		s := Subsets(input, k)
    j := 0
    for range s {
			// a[j] = next
      j++
		}
    // Result = append(Result, a)
  }
}

func init() {
    rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
    b := make([]rune, n)
    for i := range b {
        b[i] = letterRunes[rand.Intn(len(letterRunes))]
    }
    return string(b)
}
