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

// Subsets2 generates all subsets of the slice of strings with the length provided.
func Subsets2(data []string, length int) <-chan []string {
	c := make(chan []string)
	go func(c chan []string) {
		defer close(c)
		makeSubsets2(c, data, []string{}, length)
	}(c)
	return c
}

func makeSubsets2(c chan []string, inputs, outputBase []string, k int) {
	if k < 0 {
		return
	}

	inputLength := len(inputs)
	if k > inputLength {
		return
	}

	if k == 0 {
		c <- outputBase
		return
	}

	for i, next := range inputs {
		makeSubsets2(c, inputs[i+1:], append(outputBase, next), k-1)
	}
}
