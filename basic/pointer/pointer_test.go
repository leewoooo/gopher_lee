package pointer

import (
	"log"
	"testing"
)

func TestPointer(t *testing.T) {
	var a int = 8
	var b *int = &a

	c := *b // 8

	var d *int // nil

	t.Logf("%d", c)

	if d == nil {
		t.Log("선언 후 초기화 되지 않은 pointer")
	}
}

func IncreaseWithOutPtr(a int) {
	a++
}

func Increase(a *int) {
	*a++
}

func TestWithOutPointer(t *testing.T) {
	var a int = 1

	IncreaseWithOutPtr(a)

	if a != 1 {
		log.Fatal("a should be 1")
	}

}

func TestWithPointer(t *testing.T) {
	var a int = 1

	Increase(&a)

	if a != 2 {
		log.Fatal("a should be 2")
	}
}

type User struct {
	Name string
	Age  int
}

func NewUser(name string, age int) *User {
	u := User{
		Name: name,
		Age:  age,
	}

	return &u
}

// go test -gcflags -m   ./basic/pointer (excute root)
func TestHeapFlag(t *testing.T) {
	user := NewUser("foo", 10)

	t.Log(user)
}
