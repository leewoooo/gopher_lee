package timezero

import (
	"testing"
	"time"
)

func TestTimeZeroValue(t *testing.T) {
	//given
	testTime := time.Time{}

	//when
	isZero := testTime.IsZero()

	//then
	if !isZero {
		t.Fatal("testTime should be zero")
	}

	t.Log(testTime)
	t.Log(testTime.Second())
	t.Log(testTime.Nanosecond())
}
