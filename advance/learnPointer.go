package advance

import "fmt"

func incrTen(value *int) {
	*value += 10
}

func mulTwo(arr *[]int) {
	for i := 0; i < len(*arr); i++ {
		(*arr)[i] *= 2
	}
}

func TestPointer() {
	val := 10

	incrTen(&val)

	fmt.Println("val 10 after incrTen = ", val)

	arr := []int{1, 2, 3, 4, 5}

	fmt.Print("array ", arr)

	mulTwo(&arr)

	fmt.Println(" after mulTwo = ", arr)
}
