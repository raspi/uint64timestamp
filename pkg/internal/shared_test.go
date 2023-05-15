package internal

import "testing"

func TestGetDigit(t *testing.T) {
	from := uint64(654321)
	want := uint8(4)
	got := GetDigit(from, 3)

	if got != want {
		t.Fail()
	}

}
