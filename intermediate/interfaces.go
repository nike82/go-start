package intermediate

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perimeter() float64
}

type rect struct {
	width, height float64
}

type rect1 struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.height * r.width
}

func (r rect1) area() float64 {
	return r.height * r.width
}

func (r rect) perimeter() float64 {
	return 2 * (r.height + r.width)
}

func (r rect1) perimeter() float64 {
	return 2 * (r.height + r.width)
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (c circle) diameter() float64 {
	return 2 * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}

func main() {

	r := rect{width: 3, height: 4}
	c := circle{radius: 5}
	//r1 := rect1{width: 3, height: 4}
	measure(r)
	measure(c)
	//measure(r1)

	myPrinter(1, "nike", 45.6, false)

}

func printType(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("unknown")
	}
}

func myPrinter(i ...any) {
	for _, v := range i {
		fmt.Println(v)
	}
}
