package basics

import "fmt"

func isValidStr(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	var stack []rune
	var top rune
	var r bool

	for _, v := range s {
		switch v {
		case '(':
			StackPush(&stack, v)
		case '[':
			StackPush(&stack, v)
		case '{':
			StackPush(&stack, v)
		case ')':
			top, r = StackTop(&stack)
			if !r || top != '(' {
				return false
			}
			StackPop(&stack)
		case ']':
			top, r = StackTop(&stack)
			if !r || top != '[' {
				return false
			}
			StackPop(&stack)
		case '}':
			top, r = StackTop(&stack)
			if !r || top != '{' {
				return false
			}
			StackPop(&stack)
		}
	}

	return len(stack) == 0
}

func StackPush(stack *[]rune, v rune) {
	*stack = append(*stack, v)
}

func StackTop(stack *[]rune) (rune, bool) {
	if len(*stack) == 0 {
		return 0, false
	}
	return (*stack)[len(*stack)-1], true
}

func StackPop(stack *[]rune) {
	if len(*stack) == 0 {
		return
	}
	*stack = (*stack)[:len(*stack)-1]
}

func TestIsValidStr() {
	strs := []string{
		"()",
		"()[]{}",
		"(]",
		"([])",
		"([)]",
	}

	for _, v := range strs {
		fmt.Println(" isValidStr, str=", v, ", result=", isValidStr(v))
	}
}
