package basics

import "fmt"

func removeDuplicates(nums []int) int {
	i, j := 0, 1
	length := len(nums)
	if length == 1 {
		return 1
	}

	for i < length && j < length {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
		j++
	}

	return i + 1
}

func TestRemoveDuplicates() {
	arr := [][]int{
		{1, 1, 2},
		{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
	}

	for _, v := range arr {
		fmt.Println("removeDuplicates, before=", v, ", after=", removeDuplicates(v))
	}
}
