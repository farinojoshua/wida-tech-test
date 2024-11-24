package main

import (
	"fmt"
	"sort"
)

func findCombinations(l int, t int) [][]int {
	var result [][]int

	current := make([]int, 0)

	backtrack(l, t, 1, 0, current, &result)

	return result
}

func backtrack(l, target, start, currentSum int, current []int, result *[][]int) {
	if len(current) == l {
		if currentSum == target {
			temp := make([]int, len(current))
			copy(temp, current)
			sort.Ints(temp)
			*result = append(*result, temp)
		}
		return
	}

	if currentSum > target {
		return
	}

	for i := start; i <= 9; i++ {
		current = append(current, i)

		backtrack(l, target, i+1, currentSum+i, current, result)

		current = current[:len(current)-1]
	}
}

func main() {
	fmt.Println(findCombinations(3, 9))
	fmt.Println(findCombinations(2, 9))
}
