package basics

import "fmt"

func plusOne(digits []int) []int {
	idx := len(digits)
	r := make([]int, idx+1)
	sign := 1
	for idx > 0 {
		r[idx] = digits[idx-1] + sign
		if r[idx] > 9 {
			r[idx] -= 10
			sign = 1
		} else {
			sign = 0
		}
		idx--
	}

	if sign == 1 {
		r[0] = 1
		return r
	} else {
		return r[1:]
	}
}

func TestPlusOne() {
	arr := [][]int{
		{1, 2, 3},
		{4, 3, 2, 1},
		{9},
		{9, 9},
		{4, 9, 9, 9},
	}

	for _, v := range arr {
		fmt.Println("plus one, num=", v, ", result=", plusOne(v))
	}
}
