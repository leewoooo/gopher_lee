package function2

import (
	"log"
	"testing"
	"unsafe"
)

func TestDeferCallStack(t *testing.T) {
	defer log.Println("1")
	defer log.Println("2")
	defer log.Println("3")
}

type Add func(a int, b int) int

func makeAdd() Add {
	return func(a, b int) int {
		return a + b
	}
}

func TestAdd(t *testing.T) {
	funcAdd := makeAdd()

	result := funcAdd(1, 2)

	if result != 3 {
		t.Fatal("result should be 3")
	}
}

func TestClosure(t *testing.T) {
	i := 0

	log.Println(unsafe.Pointer(&i))

	AddTen := func() {
		i += 10
		log.Println(unsafe.Pointer(&i))
	}
	AddTen()

	if i != 10 {
		log.Fatal("i should be 10")
	}
}
