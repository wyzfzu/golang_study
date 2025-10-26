package basics

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	dataMap := make(map[int]int, len(nums))
	for i, v := range nums {
		idx, ok := dataMap[target-v]
		if ok {
			return []int{idx, i}
		}
		dataMap[v] = i
	}

	return []int{0, 0}
}

func TestTwoSum() {
	data := [][]int{
		{2, 7, 11, 15},
		{3, 2, 4},
		{3, 3},
	}

	target := []int{9, 6, 6}

	for i, v := range data {
		fmt.Println("two sum, data=", v, ", target=", target[i], ", result=", twoSum(v, target[i]))
	}
}
