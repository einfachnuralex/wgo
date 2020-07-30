package main

func Sum(numbers []int) (sum int) {
	for _, summand := range numbers {
		sum += summand
	}
	return sum
}
