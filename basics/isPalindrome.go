package basics

import (
	"fmt"
	"strconv"
)

func isPalindrome(num int) bool {
	if num < 0 {
		return false
	}
	n := num

	target := 0

	for n > 0 {
		digit := n % 10
		target = target*10 + digit
		n /= 10
	}

	return target == num
}

func isPalindrome2(num int) bool {
	if num < 0 {
		return false
	}

	n := []rune(strconv.Itoa(num))
	j := len(n) - 1

	for i := 0; i < j; i++ {
		if n[i] != n[j] {
			return false
		}
		j--
	}

	return true
}

func TestIsPalindrome() {
	nums := []int{-2, 4, 123, 121, 1234321}
	for _, v := range nums {
		fmt.Println("isPalindrome use num mod, num=", v, ", result=", isPalindrome(v))
		fmt.Println("isPalindrome use string, num=", v, ", result=", isPalindrome2(v))
	}
}
