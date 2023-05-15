package base10

import (
	"testing"
	"time"
)

func TestTimeUintTime(t *testing.T) {
	want := time.Date(
		2022, time.June, 3,
		23, 12, 59,
		120000, // Note accuracy
		time.UTC,
	)

	got, err := TimeToUint64(want)
	if err != nil {
		t.Fail()
	}

	gotTime, err := Uint64ToTime(got, time.UTC)
	if err != nil {
		t.Fail()
	}

	if gotTime != want {
		t.Errorf(`got %v want %v`, gotTime, want)
		t.Fail()
	}
}

func TestUintTimeUint(t *testing.T) {
	// 2022-12-06 20:41:40 466 ms 810 microseconds
	want := uint64(2022120620414046681)

	gotTime, err := Uint64ToTime(want, time.UTC)
	if err != nil {
		t.Fail()
	}

	got, err := TimeToUint64(gotTime)
	if err != nil {
		t.Fail()
	}

	if got != want {
		t.Errorf(` got: %d`, got)
		t.Errorf(`want: %d`, want)
		t.Fail()
	}
}
