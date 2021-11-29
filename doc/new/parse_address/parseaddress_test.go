package parseaddress

import (
	"log"
	"net/mail"
	"testing"
)

func TestParseAddress(t *testing.T) {
	given := []string{
		"Alice <alice@example.com>",
		"<alice@example.com>",
		"alice@example.com",
		"alice@example",
		"bad-example",
		"",
		"@",
	}

	for _, v := range given {
		addressStruct, err := mail.ParseAddress(v)
		if err != nil {
			log.Println(err)
			continue
		}

		log.Printf("name: %s, address: %s", addressStruct.Name, addressStruct.Address)
	}
}
