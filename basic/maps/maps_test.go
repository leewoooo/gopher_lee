package maps

import (
	"reflect"
	"testing"
)

const M = 10

func hash(d int) int {
	return d % M
}

func TestMakeHashMap(t *testing.T) {
	m := [M]string{}

	m[hash(23)] = "leewooo"
	m[hash(259)] = "foobar"

	t.Logf("%d = %s", 23, m[hash(23)])
	t.Logf("%d = %s", 259, m[hash(259)])
}

func TestMapReference(t *testing.T) {
	m := map[string]string{
		"foo": "bar",
	}

	returnM := returnMap(m)
	returnM["bar"] = "foo"

	t.Log(reflect.DeepEqual(m, returnM))

	t.Log(m)
	t.Log(returnM)
}

func returnMap(m map[string]string) map[string]string {
	return m
}
