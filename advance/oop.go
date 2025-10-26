package advance

import (
	"fmt"
	"math"
)

/**
题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
考察点 ：接口的定义与实现、面向对象编程风格。
*/

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  int
	Height int
}

func (r *Rectangle) Area() float64 {
	return float64(r.Height) * float64(r.Width)
}

func (r *Rectangle) Perimeter() float64 {
	return (float64(r.Height) + float64(r.Width)) * 2
}

type Circle struct {
	Radius int
}

func (c *Circle) Area() float64 {
	return math.Pi * float64(c.Radius) * float64(c.Radius)
}

func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * float64(c.Radius)
}

/**
题目 ：使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
考察点 ：组合的使用、方法接收者。
*/

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeID int
}

func (e *Employee) printInfo() {
	fmt.Println("Employee Info: Name=", e.Name, ", Age=", e.Age, ", EmployeeID=", e.EmployeeID)
}

func TestOOP() {
	r := &Rectangle{3, 4}
	fmt.Println("Rect: ", r)
	fmt.Println("Rect area=", r.Area(), ", Perimeter=", r.Perimeter())

	c := &Circle{5}
	fmt.Println("Circle Raduis: ", c)
	fmt.Println("Circle area=", c.Area(), ", Perimeter=", c.Perimeter())

	e := &Employee{}
	e.Age = 25
	e.Name = "Lisi"
	e.EmployeeID = 11111

	e.printInfo()
}
