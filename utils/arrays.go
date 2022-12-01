package utils

import "strconv"

// sum array of numbers
func Sum(numbers []int) int {
	var sum int
	for _, num := range numbers {
		sum += num
	}
	return sum
}

// run function for each array of int in array and return array of results
func ReduceSum(input [][]int) []int {
	var output []int
	for _, line := range input {
		output = append(output, Sum(line))
	}
	return output
}

func MatrixInt(input [][]string) [][]int {
	var output [][]int
	for _, line := range input {
		output = append(output, ArrayInt(line))
	}
	return output
}

func ArrayInt(input []string) []int {
	var output []int
	for _, line := range input {
		num, _ := strconv.Atoi(line)
		output = append(output, num)
	}
	return output
}
