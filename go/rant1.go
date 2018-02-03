package rant1

// Subsets generates all subsets of the slice of strings with the length provided.
func Subsets(data []string, length int) <-chan []string {
	c := make(chan []string)
	go func(c chan []string) {
		defer close(c)
		makeSubsets(c, data, length)
	}(c)
	return c
}

func makeSubsets(c chan []string, inputs []string, k int) {
	inputLength := len(inputs)
	if k > inputLength {
		return
	}

	for num := 0; num < 1<<uint(inputLength); num++ {
		var ones int
		for i := num; i != 0; i = i >> 1 {
			ones += i & 1
		}

		if ones == k {
			next := num
			subsetIndex := 0
			setIndex := 0
			output := make([]string, k)
			for subsetIndex < k {
				if next&1 == 1 {
					x := inputs[setIndex]
					output[subsetIndex] = x
					subsetIndex++
				}

				setIndex++
				next = next >> 1
			}
			c <- output
		}
	}
}
