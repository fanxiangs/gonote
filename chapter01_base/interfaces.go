package main

import "fmt"

type Shaper interface {
	Area() float32
}
type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func crateInterfaces() {
	sq1 := new(Square)
	sq1.side = 5
	var areaIntf Shaper
	areaIntf = sq1
	fmt.Printf("The square has area: %f\n", areaIntf.Area())
}

type Fruitable interface {
	eat()
}

type Fruit struct {
	Name      string // 属性变量
	Fruitable        // 匿名内嵌接口变量
}

type Apple struct {
	Name string
}
type Banana struct{}

func (a Apple) eat() {
	fmt.Printf("Eating Apple\n")
}
func (b Banana) eat() {
	fmt.Printf("Eating Banana\n")
}

func (f Fruit) want() {
	fmt.Printf("I like it\n")
	f.eat()
}

func polymorphism() {
	var f1 = Fruit{"Apple", Apple{}}
	f1.eat()
	var f2 = Fruit{Name: "Banana", Fruitable: Banana{}}
	f2.eat()
	//f2.want()

}

func assertTypeInterfaces() {
	sq1 := new(Square)
	sq1.side = 5
	var areaIntf Shaper
	areaIntf = sq1
	if t, ok := areaIntf.(*Square); ok {
		fmt.Printf("The type of areaIntf is: %T\n", t)
	}
}
func main() {
	//crateInterfaces()
	//polymorphism()
	assertTypeInterfaces()
}
