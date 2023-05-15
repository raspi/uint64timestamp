package base2

import (
	"github.com/raspi/uint64timestamp/pkg/errs"
	"github.com/raspi/uint64timestamp/pkg/internal"
	"math/bits"
	"time"
)

// Computer parsable timestamp as uint64
//
// y 9999 = 14 bits or more
// m 12 = 4 bits
// d 31 = 5 bits
//
// h 23 = 5 bits
// m 59 = 6 bits
// s 59 = 6 bits
//
//	ms 999 = 10 bits
//
// micro 999 = 10 bits
// nano (hundreds)  9 = 4 bits (0 -- 900)

// uint64 bit positions
const (
	MaskYear           = uint64(0b11111111_11111100_00000000_00000000_00000000_00000000_00000000_00000000)
	MaskMonth          = uint64(0b00000000_00000011_11000000_00000000_00000000_00000000_00000000_00000000)
	MaskDay            = uint64(0b00000000_00000000_00111110_00000000_00000000_00000000_00000000_00000000)
	MaskHour           = uint64(0b00000000_00000000_00000001_11110000_00000000_00000000_00000000_00000000)
	MaskMinute         = uint64(0b00000000_00000000_00000000_00001111_11000000_00000000_00000000_00000000)
	MaskSecond         = uint64(0b00000000_00000000_00000000_00000000_00111111_00000000_00000000_00000000)
	MaskMilliSec       = uint64(0b00000000_00000000_00000000_00000000_00000000_11111111_11000000_00000000)
	MaskMicroSec       = uint64(0b00000000_00000000_00000000_00000000_00000000_00000000_00111111_11110000)
	MaskNanoSecHundred = uint64(0b00000000_00000000_00000000_00000000_00000000_00000000_00000000_00001111)

	MaskYMD     = MaskYear | MaskMonth | MaskDay
	MaskHMS     = MaskHour | MaskMinute | MaskSecond
	MaskMsMicro = MaskMilliSec | MaskMicroSec
	MaskAll     = MaskYMD | MaskHMS | MaskMsMicro
)

func TimeToUint64(t time.Time) (uint64, error) {
	t = t.Truncate(time.Nanosecond)

	// No negative years
	if t.Year() < 0 {
		return 0, errs.NewErrYear(t.Year())
	}

	y := uint64(t.Year()) << bits.TrailingZeros64(MaskYear)
	m := uint64(t.Month()) << bits.TrailingZeros64(MaskMonth)
	d := uint64(t.Day()) << bits.TrailingZeros64(MaskDay)

	h := uint64(t.Hour()) << bits.TrailingZeros64(MaskHour)
	mi := uint64(t.Minute()) << bits.TrailingZeros64(MaskMinute)
	s := uint64(t.Second()) << bits.TrailingZeros64(MaskSecond)

	ns := time.Nanosecond * time.Duration(t.Nanosecond())

	ms := uint64(ns.Milliseconds() << bits.TrailingZeros64(MaskMilliSec))
	ns %= time.Millisecond

	micro := uint64(ns.Microseconds() << bits.TrailingZeros64(MaskMicroSec))
	ns %= time.Microsecond

	// Hundreds of nanoseconds (100 -- 900)
	nsHundred := uint64(internal.GetDigit(uint64(ns), 2))

	return y | m | d | h | mi | s | ms | micro | nsHundred, nil
}

func Uint64ToTime(t uint64, l *time.Location) (time.Time, error) {
	t = MaskAll & t

	y := MaskYear & t >> bits.TrailingZeros64(MaskYear)

	if bits.OnesCount64(y) > bits.OnesCount64(MaskYear) {
		return time.Time{}, errs.NewErrYear(int(y))
	}

	m := MaskMonth & t >> bits.TrailingZeros64(MaskMonth)
	if m > errs.MonthMax {
		return time.Time{}, errs.NewErrMonth(int(m))
	}

	d := MaskDay & t >> bits.TrailingZeros64(MaskDay)
	if d > errs.DayMax {
		return time.Time{}, errs.NewErrDay(int(d))
	}

	h := MaskHour & t >> bits.TrailingZeros64(MaskHour)
	if h > errs.HourMax {
		return time.Time{}, errs.NewErrHour(int(h))
	}

	mi := MaskMinute & t >> bits.TrailingZeros64(MaskMinute)
	if mi > errs.MinuteMax {
		return time.Time{}, errs.NewErrMinute(int(mi))
	}

	s := MaskSecond & t >> bits.TrailingZeros64(MaskSecond)
	if s > errs.SecondMax {
		return time.Time{}, errs.NewErrSecond(int(s))
	}

	ms := MaskMilliSec & t >> bits.TrailingZeros64(MaskMilliSec)
	if ms > errs.MilliSecondMax {
		return time.Time{}, errs.NewErrMilliSecond(int(ms))
	}

	micro := MaskMicroSec & t >> bits.TrailingZeros64(MaskMicroSec)
	if micro > errs.MicroSecondMax {
		return time.Time{}, errs.NewErrMicroSecond(int(micro))
	}

	var ns time.Duration
	ns += time.Millisecond * time.Duration(ms)
	ns += time.Microsecond * time.Duration(micro)
	ns += time.Nanosecond * time.Duration((MaskNanoSecHundred&t)*100)

	return time.Date(
		int(y), time.Month(m), int(d),
		int(h), int(mi), int(s),
		int(ns),
		l,
	), nil
}
