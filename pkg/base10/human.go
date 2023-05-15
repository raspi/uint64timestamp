package base10

import (
	"github.com/raspi/uint64timestamp/pkg/errs"
	"github.com/raspi/uint64timestamp/pkg/internal"
	"time"
)

// Human parsable base 10 timestamp as uint64
//
// Format:
//    20230502012359123
//    yyyyMMddHHMMSSmmm
// 18446744073709551615 max uint64, 8 bytes
//  9223372036854775808 first number requiring 64 bits
//
// 2023050202532318163 = 2023-05-02 02:53:23.18163   Timezone info is not added
//
// Maximum year is ~18446
// It's actually 61 bits until we hit year ~9223 :)

// Positions base10
const (
	Year  = uint64(1_000_000_000_000_000)
	Month = uint64(10_000_000_000_000)
	Day   = uint64(100_000_000_000)

	Hour   = uint64(1_000_000_000)
	Minute = uint64(10_000_000)
	Second = uint64(100_000)

	Millisecond         = uint64(100)
	MicrosecondHundreth = uint64(10)
	MicrosecondTenth    = uint64(1)
)

// TimeToUint64 packs timestamp into uint64
func TimeToUint64(t time.Time) (o uint64, err error) {
	t = t.Truncate(time.Microsecond)

	if t.Year() < 0 {
		return 0, errs.NewErrYear(t.Year())
	}

	o += uint64(t.Year()) * Year
	o += uint64(t.Month()) * Month
	o += uint64(t.Day()) * Day

	o += uint64(t.Hour()) * Hour
	o += uint64(t.Minute()) * Minute
	o += uint64(t.Second()) * Second

	ns := time.Nanosecond * time.Duration(t.Nanosecond())

	o += uint64(ns.Milliseconds()) * Millisecond
	ns %= time.Millisecond

	// We need some truncating because we're out of bits

	ns += time.Millisecond * 1
	ns /= 1000

	// Hundreds
	micro := MicrosecondHundreth * uint64(internal.GetDigit(uint64(ns), 2))
	// Tenths
	micro += uint64(internal.GetDigit(uint64(ns), 1))

	o += MicrosecondTenth * micro

	return o, nil
}

// Uint64ToTime converts uint64 to time.Time
// Takes timezone as it's not in the timestamp itself
// Remember possible binary endianness conversion first :)
func Uint64ToTime(t uint64, l *time.Location) (time.Time, error) {
	// Year
	y := t / Year
	t %= Year

	m := t / Month
	t %= Month

	if m > errs.MonthMax {
		return time.Time{}, errs.NewErrMonth(int(m))
	}

	d := t / Day
	t %= Day

	if d > errs.DayMax {
		return time.Time{}, errs.NewErrDay(int(d))
	}

	// Hour
	h := t / Hour
	t %= Hour

	if h > errs.HourMax {
		return time.Time{}, errs.NewErrHour(int(h))
	}

	mi := t / Minute
	t %= Minute

	if mi > errs.MinuteMax {
		return time.Time{}, errs.NewErrMinute(int(mi))
	}

	s := t / Second
	t %= Second

	if s > errs.SecondMax {
		return time.Time{}, errs.NewErrSecond(int(s))
	}

	ms := t / Millisecond
	t %= Millisecond

	nanos := time.Millisecond * time.Duration(ms)
	nanos += time.Microsecond * time.Duration(t*10)

	return time.Date(
		int(y), time.Month(m), int(d),
		int(h), int(mi), int(s),
		int(nanos),
		l,
	), nil
}
