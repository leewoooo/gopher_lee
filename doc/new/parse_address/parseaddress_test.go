package parseaddress

import (
	"log"
	"net/mail"
	"regexp"
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

func TestEmailReg(t *testing.T) {
	emailreg := `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`

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
		ok, err := regexp.MatchString(emailreg, v)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("%s Match result: %v", v, ok)
	}
}
