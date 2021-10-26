package interfaces

import (
	"fmt"
	"testing"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rect struct {
	width  float64
	height float64
}

func (r *Rect) Area() float64 {
	return r.width * r.height
}

func (r *Rect) Perimeter() float64 {
	return 2 * (r.width * r.height)
}

func TestInterfaceCreate(t *testing.T) {
	var s Shape

	fmt.Println("value of s is", s)
	fmt.Printf("type of s is %T\n", s)
}

func TestImplement(t *testing.T) {
	var shape Shape

	rect := Rect{
		width:  3.0,
		height: 4.0,
	}

	shape = &rect

	t.Logf("Arar : %f", shape.Area())
	t.Logf("Perimeter : %f", shape.Perimeter())

	fmt.Printf("Shape Type: %T, rect Type: %T \n", shape, rect)

	r, ok := shape.(*Rect)
	if !ok {
		t.Fatal("interface를 구현한 type은 Rect가 아닙니다.")
	}

	t.Log("width :", r.width)
	t.Log("height :", r.height)
}
