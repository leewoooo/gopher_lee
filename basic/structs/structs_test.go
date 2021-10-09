package structs

import (
	"log"
	"testing"
	"unsafe"
)

func TestEmbeded(t *testing.T) {
	tenant := Tenant{
		User: User{
			Name: "test",
			pID:  123,
		},
		TenantName: "foo",
	}

	log.Println("tenant: ", tenant)
}

// for memory test
type People struct {
	Age   int32
	Score int32
}

func TestMemory(t *testing.T) {
	p := People{
		Age:   1,
		Score: 2,
	}

	// 예상하기로는 12 하지만 16
	log.Println(unsafe.Sizeof(p))
}
