package main

import (
	"github.com/wyzfzu/golang_study/advance"
	"github.com/wyzfzu/golang_study/basics"
)

func TestBasics() {
	basics.TestSingleNumber()
	basics.TestIsPalindrome()
	basics.TestIsValidStr()
	basics.TestLongestCommonPrefix()
	basics.TestPlusOne()
	basics.TestRemoveDuplicates()
	basics.TestMergeIntervals()
	basics.TestTwoSum()
}

func TestAdcance() {
	advance.TestPointer()
	advance.TestGoRutine()
	advance.TestOOP()
	advance.TestChan()
	advance.TestLock()
}

func main() {
	TestBasics()
	TestAdcance()
}
