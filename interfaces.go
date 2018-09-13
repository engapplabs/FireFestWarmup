package main

import (
	"fmt"
)

type Shape interface {
	perimeter() float64
	area() float64
}

type Rect struct {
	widht, height float64
}

type Circle struct {
	radius float64
}

func (rect Rect) area() float64 {
	return rect.height * rect.widht
}

func (rect Rect) perimeter() float64 {
	return (rect.widht + rect.height) * 2
}

func (circle Circle) area() float64 {
	return circle.radius * circle.radius * 3.14
}

func (circle Circle) perimeter() float64 {
	return 2 * 3.14 * circle.radius
}

func getShapeMeasures(shape Shape) {
	fmt.Println(shape)
	fmt.Println(shape.area())
	fmt.Println(shape.perimeter())
}

func main() {
	rect := Rect{widht: 10, height: 5}
	circle := Circle{radius: 10}

	getShapeMeasures(rect)
	getShapeMeasures(circle)
}
