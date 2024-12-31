package main

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

func printAnything(val interface{}) {
	fmt.Println("Value:", val)
}

func main() {
	var s Shape

	r := Rectangle{Width: 10, Height: 5}
	c := Circle{Radius: 7}

	s = r
	fmt.Println("Rectangle Area:", s.Area())
	fmt.Println("Rectangle Perimeter:", s.Perimeter())

	s = c
	fmt.Println("Circle Area:", s.Area())
	fmt.Println("Circle Perimeter:", s.Perimeter())

	printAnything(42)
	printAnything("Hello, Go!")
	printAnything(3.14)

	var val interface{} = 42

	if i, ok := val.(int); ok {
		fmt.Println("Integer value:", i)
	} else {
		fmt.Println("Type assertion failed")
	}
}
