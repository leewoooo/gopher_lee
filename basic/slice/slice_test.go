package slice

import (
	"log"
	"reflect"
	"testing"
	"unsafe"
)

func TestSliceInterning(t *testing.T) {
	//given
	arr1 := []int{1, 2, 3}

	arr2 := []int{1, 2, 3}

	//when
	log.Println((*reflect.SliceHeader)(unsafe.Pointer(&arr1)))

	log.Println((*reflect.SliceHeader)(unsafe.Pointer(&arr2)))
}

func TestSliceSize(t *testing.T) {
	//given
	arr1 := []int{1, 2, 3}
	arr2 := []int{1, 2, 3}

	//when
	log.Println(unsafe.Sizeof(arr1))
	log.Println(unsafe.Sizeof(arr2))
}

func changeSlice(slice []int) {
	slice[1] = 10
}

func changeArray(arr [3]int) {
	arr[1] = 10
}

func TestChange(t *testing.T) {
	slice := []int{1, 2, 3}
	arr := [3]int{1, 2, 3}

	changeSlice(slice)
	changeArray(arr)

	log.Println("slice: ", slice)
	log.Println("arr: ", arr)
}

func TestAppend(t *testing.T) {
	slice := make([]int, 3, 5)
	log.Println("origin: ", (*reflect.SliceHeader)(unsafe.Pointer(&slice)))

	slice = append(slice, 4)
	log.Println("append 1: ", (*reflect.SliceHeader)(unsafe.Pointer(&slice)))

	slice = append(slice, 5)
	log.Println("append 2: ", (*reflect.SliceHeader)(unsafe.Pointer(&slice)))

	slice = append(slice, 6)
	log.Println("append 3: ", (*reflect.SliceHeader)(unsafe.Pointer(&slice)))
}

func TestRemoveElement(t *testing.T) {
	// will remove 3
	slice := []int{1, 2, 3, 4, 5}

	slice = append(slice[:2], slice[3:]...)

	log.Println("slice: ", slice)
}

func TestNilSlice(t *testing.T) {
	var a []int

	var b []int = make([]int, 0)

	var c []int = []int{}

	log.Println("a, b, c", a, b, c)

	if a == nil {
		log.Println("a is nil")
	}

	if b == nil {
		log.Println("b is nil")
	}

	if c == nil {
		log.Println("c is nil")
	}
}
