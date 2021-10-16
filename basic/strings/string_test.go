package strings

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"testing"
	"unsafe"
)

func TestCircuitString(t *testing.T) {
	// given
	gString := "Hello 고랭"

	//when
	// #1
	for i := 0; i < len(gString); i++ {
		fmt.Printf("값: %d 문자:%c\n", gString[i], gString[i])
	}

	// #2
	arr := []rune(gString)
	for i := 0; i < len(arr); i++ {
		fmt.Printf("값: %d 문자:%c\n", arr[i], arr[i])
	}

	// #3
	for _, v := range gString {
		fmt.Printf("값: %d 문자:%c\n", v, v)
	}
}

func TestStringPointer(t *testing.T) {
	//givne
	str1 := "Hello"

	str2 := "Hello"

	log.Println((*reflect.StringHeader)(unsafe.Pointer(&str1)))
	log.Println((*reflect.StringHeader)(unsafe.Pointer(&str2)))

	//result
	// 2021/10/15 20:42:24 &{4337888730 5}
	// 2021/10/15 20:42:24 &{4337888730 5}
}

func TestStringSize(t *testing.T) {
	str1 := "Hello String"
	str2 := ""

	log.Println(unsafe.Sizeof(str1))
	log.Println(unsafe.Sizeof(str2))
}

func TestStringImmutable(t *testing.T) {
	str1 := "Hello String"
	log.Println((*reflect.StringHeader)(unsafe.Pointer(&str1)))

	// 부분 수정은 불가능 하다
	// str1[2] = 'e'  error -> cannot assign to str1[2] (value of type byte)compiler (UnassignableOperand)

	str1 = "Hello Golang"
	log.Println((*reflect.StringHeader)(unsafe.Pointer(&str1)))

	// output (서로 다른 주소값을 나타낸다.) -> 즉 string을 새로 만든다.
	// 2021/10/16 21:53:05 &{4300555103 12}
	// 2021/10/16 21:53:05 &{4300555091 12}
}

func TestStringCompile(t *testing.T) {
	str1 := "12"
	log.Println((*reflect.StringHeader)(unsafe.Pointer(&str1)))

	str2 := "1" + "2"
	log.Println((*reflect.StringHeader)(unsafe.Pointer(&str2)))
}

func TestStringRunTime(t *testing.T) {
	str1 := "12"
	log.Println((*reflect.StringHeader)(unsafe.Pointer(&str1)))

	str2 := strconv.Itoa(12)
	log.Println((*reflect.StringHeader)(unsafe.Pointer(&str2)))
}
