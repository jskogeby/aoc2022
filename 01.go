package main

import (
	"aoc2022/utils"
	"fmt"
	"sort"
)

func Day01() {
	input := utils.ReadInput("01")
	intMatrix := utils.MatrixInt(utils.DivideInput(input))
	sums := utils.ReduceSum(intMatrix)
	sort.Sort(sort.Reverse(sort.IntSlice(sums)))

	part1 := sums[0]
	fmt.Println("Part 1:", part1)
	part2 := utils.Sum(sums[0:2])
	fmt.Println("Part 2:", part2)
}
