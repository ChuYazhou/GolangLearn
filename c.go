package main

import "fmt"

type Shape interface {
	Area() float64
}

// Rectangle 定义一个结构体
type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func main() {
	s := Rectangle{Width: 10.0, Height: 20.0}
	c := Circle{Radius: 10.0}
	fmt.Println("Area of Rectangle is:", s.Area())
	fmt.Printf("Area of Rectangle is:%.2f\n", s.Area())
	fmt.Println("Area of Circle is :", c.Area())
	fmt.Printf("%+v\n", s)
	fmt.Printf("%+v\n", c)
}
