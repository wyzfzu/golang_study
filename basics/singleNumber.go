package basics

import "fmt"

func FindNumber(nums []int) int {
	var numLen = len(nums)
	var numMap map[int]int = make(map[int]int, numLen)
	for _, v := range nums {
		numMap[v]++
	}

	for k, v := range numMap {
		if v == 1 {
			return k
		}
	}

	return -1
}

func FindNumber2(nums []int) int {
	var oneTimeNum = nums[0]
	for i := 1; i < len(nums); i++ {
		oneTimeNum ^= nums[i]
	}

	return oneTimeNum
}

func TestSingleNumber() {
	numsArr := [][]int{
		{2, 2, 1},
		{4, 1, 2, 1, 2},
	}
	for _, nums := range numsArr {
		num := FindNumber(nums)
		fmt.Println("find single number use map, nums=", nums, "single number=", num)
		num = FindNumber2(nums)
		fmt.Println("find single number use xor, nums=", nums, "single number=", num)
	}
}
