package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x      int
	y      int
	Radius int
}

func crateStruct() {
	var c = Circle{
		x:      100,
		y:      100,
		Radius: 50,
	}
	fmt.Println(c)
	fmt.Printf("%+v\n", c)
	var c1 Circle = Circle{100, 100, 50}
	fmt.Printf("%+v\n", c1)
	var c2 *Circle = &Circle{100, 100, 50}
	fmt.Printf("%+v\n", c2)
	var c3 *Circle = new(Circle)
	fmt.Printf("%+v\n", c3)
	expandByValue(&c1) // 指针传递
	fmt.Printf("%+v\n", c1)

}
func expandByValue(c *Circle) {
	c.Radius *= 2
}

func (c Circle) Area() float64 {
	//Go 没有 self 和 this 这样的关键字来指代当前的对象，它是用户自己定义的变量名称，通常我们都使用单个字母来表示。
	return math.Pi * float64(c.Radius) * float64(c.Radius)
}
func (c Circle) Circumference() float64 {
	return 2 * math.Pi * float64(c.Radius)
}

func funcStruct() {
	var c = Circle{Radius: 50}
	fmt.Println(c.Area(), c.Circumference())
}

type innerS struct {
	in1 int
	in2 int
}
type outerS struct {
	// 在一个结构体中对于每一种数据类型只能有一个匿名字段
	b      int
	c      float32
	int    //anonymous field
	innerS // anonymous field
}

// 匿名字段
func anonymousFieldsStruct() {
	outer := new(outerS)
	outer.b = 6
	outer.c = 7.5
	outer.int = 60
	outer.in1 = 5
	outer.in2 = 10
	fmt.Printf("outer.b is: %d\n", outer.b)
	fmt.Printf("outer.c is: %f\n", outer.c)
	fmt.Printf("outer.int is: %d\n", outer.int)
	fmt.Printf("outer.in1 is: %d\n", outer.in1)
	fmt.Printf("outer.in2 is: %d\n", outer.in2)
	outer2 := outerS{6, 7.5, 60, innerS{5, 10}}
	fmt.Println("outer2 is:", outer2)

}

func embedStruct() {
	// 嵌套结构体
	type A struct {
		ax, ay int
	}
	type B struct {
		A
		bx, by float32
	}
	b := B{A{1, 2}, 3.0, 4.0}
	fmt.Println(b.ax, b.ay, b.bx, b.by)
	fmt.Println(b.A)

}
func main() {
	//crateStruct()
	//funcStruct()
	//anonymousFieldsStruct()
	embedStruct()
}
