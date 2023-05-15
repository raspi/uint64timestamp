package base2

import (
	"math/bits"
	"testing"
	"time"
)

// Test only year through seconds
func TestTimeToUint64NoMicro(t *testing.T) {
	ts := time.Date(
		2023, time.June, 22,
		11, 41, 59,
		0,
		time.UTC,
	)

	gotInteger, err := TimeToUint64(ts)
	if err != nil {
		t.Fail()
	}

	got, err := Uint64ToTime(gotInteger, time.UTC)
	if err != nil {
		t.Fail()
	}

	if ts.UnixNano() != got.UnixNano() {
		t.Fail()
	}
}

// Test microseconds
func TestTimeToUint64Microseconds(t *testing.T) {
	wantNs := time.Millisecond * 123
	wantNs += time.Microsecond * 654

	ts := time.Date(
		2023, time.June, 26,
		23, 59, 59,
		int(wantNs.Nanoseconds()),
		time.UTC,
	)

	uts, err := TimeToUint64(ts)
	if err != nil {
		t.Fail()
	}

	gotts, err := Uint64ToTime(uts, time.UTC)
	if err != nil {
		t.Fail()
	}

	got := time.Nanosecond * time.Duration(gotts.Nanosecond())

	if got != wantNs {
		t.Errorf(`got %v want %v`, got, wantNs)
		t.Fail()
	}

}

// Do the numbers fit in binary format?
func TestConsts(t *testing.T) {
	if bits.OnesCount64(MaskYear) < 14 {
		t.Errorf(`year`)
		t.Fail()
	}

	if bits.OnesCount64(MaskMonth) != 4 {
		t.Errorf(`month`)
		t.Fail()
	}

	if bits.OnesCount64(MaskDay) != 5 {
		t.Errorf(`day`)
		t.Fail()
	}

	if bits.OnesCount64(MaskHour) != 5 {
		t.Errorf(`hour`)
		t.Fail()
	}

	if bits.OnesCount64(MaskMinute) != 6 {
		t.Errorf(`minute`)
		t.Fail()
	}

	if bits.OnesCount64(MaskSecond) != 6 {
		t.Errorf(`second`)
		t.Fail()
	}

	if bits.OnesCount64(MaskMilliSec) != 10 {
		t.Errorf(`millisecond`)
		t.Fail()
	}

	if bits.OnesCount64(MaskMicroSec) != 10 {
		t.Errorf(`microsecond`)
		t.Fail()
	}
}
