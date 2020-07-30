package main

func Sum(numbers []int) (sum int) {
	for _, summand := range numbers {
		sum += summand
	}
	return sum
}

// SumAll sums slices in slices
func SumAll(summandSlices ...[]int) []int {
	results := make([]int, len(summandSlices))

	for index, summandSlice := range summandSlices {
		results[index] = Sum(summandSlice)
	}

	return results
}

// SumAllTails sums the tails of each slice
func SumAllTails(summandSlices ...[]int) []int {
	results := make([]int, len(summandSlices))

	for index, summandSlice := range summandSlices {
		if len(summandSlice) == 0 {
			continue
		}

		results[index] = Sum(summandSlice[1:])
	}

	return results
}
