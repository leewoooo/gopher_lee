package valiable

import (
	"math"
	"testing"
	"unsafe"
)

func TestOverFlow(t *testing.T) {
	// cannot use 255 + 1 (untyped int constant 256) as uint8 value in variable declaration (overflows)
	// var a uint8 = 255 + 1

	var b uint8 = math.MaxUint8
	b = b + 1
	t.Log(b)

	var c int8 = math.MaxInt8
	c = c + 1
	t.Log(c)
}

func TestUnderFlow(t *testing.T) {
	// var a uint8 = 0 - 1

	var b uint8 = 0
	b = b - 1
	t.Log(b)
}

func TestUnSafe(t *testing.T) {
	var a int8 = 123
	var b int16 = 32000

	t.Log(unsafe.Sizeof(a)) //1
	t.Log(unsafe.Sizeof(b)) //2
}
