package rant1

import (
	"fmt"
	"math/big"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Run with command:
//    go test -bench .

type subsetFunc func(data []string, length int) <-chan []string

var Result [][][]string

func BenchmarkSubsets(b *testing.B) {
	benchmark(b, Subsets)
}

func BenchmarkSubsets2(b *testing.B) {
	benchmark(b, Subsets2)
}

func benchmark(b *testing.B, subsetFunc subsetFunc) {
	k := 2
	input := make([]string, 1500)
	for i := range input {
		input[i] = RandStringRunes(8)
	}

	// n choose k should be output size
	outputLen := new(big.Int)
	outputLen.Binomial(int64(len(input)), int64(k))

	a := make([][]string, int(outputLen.Int64()))

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s := subsetFunc(input, k)
		j := 0
		for next := range s {
			a[j] = next
			j++
		}
		Result = append(Result, a)
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
		for i, subsetFunc := range []subsetFunc{Subsets, Subsets2} {
			description := fmt.Sprintf("subsetFunc: %d, k: %d, input: %v", i, scenario.k, scenario.input)

			subsets := subsetFunc(scenario.input, scenario.k)

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
}
