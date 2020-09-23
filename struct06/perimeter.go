package struct05

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Triangle struct {
	Base   float64
	Height float64
}

type Circle struct {
	Radius float64
}

func (rec Rectangle) Area() float64 {
	num1 := rec.Height
	num2 := rec.Width
	return num1 * num2
}

func (obj Circle) Area() float64 {
	return math.Pi * obj.Radius * obj.Radius
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}
func Perimeter(rec Rectangle) float64 {
	num1 := rec.Height
	num2 := rec.Width
	return 2 * (num1 + num2)
}

func Area(rec Rectangle) float64 {
	num1 := rec.Height
	num2 := rec.Width
	return num1 * num2
}
