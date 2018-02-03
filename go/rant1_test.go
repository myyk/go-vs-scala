package rant1

import (
	"fmt"
	"log"
	"math/big"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Run with command:
//    go test -bench .

var Result [][][]string

func BenchmarkSubsets(b *testing.B) {
	k := 2
	input := make([]string, 31)
	for i := range input {
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

func Test_GenerateSubsetsKLength(t *testing.T) {
	scenarios := []struct {
		input    []string
		k        int
		expected [][]string
	}{
		{[]string{}, -1, nil},
		{[]string{"a", "b", "c"}, -1, nil},
		{nil, 3, nil},
		{[]string{}, 0, [][]string{{}}},
		{[]string{"a", "b", "c"}, 0, [][]string{{}}},
		{[]string{"a", "b", "c"}, 4, nil},
		{[]string{"a"}, 0, [][]string{{}}},
		{[]string{"a", "b"}, 1, [][]string{{"a"}, {"b"}}},
		{[]string{"a", "b", "c"}, 2, [][]string{{"a", "b"}, {"a", "c"}, {"b", "c"}}},
		{[]string{"a", "b", "a"}, 2, [][]string{{"a", "b"}, {"a", "a"}, {"b", "a"}}},
		{[]string{"a", "b", "c", "d"}, 2, [][]string{{"a", "b"}, {"a", "c"}, {"a", "d"}, {"b", "c"}, {"b", "d"}, {"c", "d"}}},
		{[]string{"a", "b", "c", "d"}, 3, [][]string{{"a", "b", "c"}, {"a", "b", "d"}, {"a", "c", "d"}, {"b", "c", "d"}}},
		{[]string{"a", "b", "c", "d", "e"}, 3, [][]string{
			{"a", "b", "c"}, {"a", "b", "d"}, {"a", "b", "e"},
			{"a", "c", "d"}, {"a", "c", "e"}, {"a", "d", "e"},
			{"b", "c", "d"}, {"b", "c", "e"}, {"b", "d", "e"}, {"c", "d", "e"},
		}},
	}

	for _, scenario := range scenarios {
		description := fmt.Sprintf("k: %d, input: %v", scenario.k, scenario.input)

		subsets := Subsets(scenario.input, scenario.k)

		// collect results from the chan
		var actual [][]string
		for next := range subsets {
			actual = append(actual, next)
		}

		assert.Len(t, actual, len(scenario.expected), description)
		for _, next := range scenario.expected {
			assert.Contains(t, actual, next, description)
		}
	}
}
