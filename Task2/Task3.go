package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r *Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Circle struct {
	Radius float64
}

func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeId int
}

func (e *Employee) PrintInfo() {
	fmt.Printf("EmployeeId:%d#, Name:%s, Age:%d\n", e.EmployeeId, e.Name, e.Age)
}

func main() {
	var shape Shape = &Rectangle{6, 8}
	fmt.Printf("area if %.2f and perimeter is %.2f\n", shape.Area(), shape.Perimeter())

	shape = &Circle{9}
	fmt.Printf("area if %.2f and perimeter is %.2f\n", shape.Area(), shape.Perimeter())

	employee := Employee{Person{"Jack", 18}, 1}
	employee.PrintInfo()
}
