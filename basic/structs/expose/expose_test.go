package expose

import (
	"log"
	"testing"

	"gopher_lee/basic/structs"
)

func TestExpose(t *testing.T) {
	// package가 다르면 expose가 되지 않은 field는 초기화가 불가능하다.
	// package명이 같지만 다른 디렉토리에 있어도 초기화는 불가능 하다.
	user := structs.User{
		Name: "lee",
	}

	log.Println("user", user)
}
