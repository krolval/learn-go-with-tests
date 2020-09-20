package main

func Sum(slice []int) int {
	var result int
	for _, v := range slice {
		result += v
	}
	return result
}

func SumAll(args ...[]int) []int {
	sums := make([]int, len(args))
	for k, v := range args {
		sums[k] = Sum(v) //--> append function which takes a slice and a new value
		//sums = append(sums, Sum(v))
	}
	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
