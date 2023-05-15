package errs

import (
	"fmt"
)

// Check implementations
var _ error = &ErrYear{}
var _ error = &ErrMonth{}
var _ error = &ErrDay{}
var _ error = &ErrHour{}
var _ error = &ErrMinute{}
var _ error = &ErrSecond{}
var _ error = &ErrMilliSecond{}
var _ error = &ErrMicroSecond{}

// ErrYear if year is invalid
type ErrYear struct {
	i int
}

func NewErrYear(m int) ErrYear {
	return ErrYear{i: m}
}
func (e ErrYear) Error() string {
	return fmt.Sprintf(`invalid year: %d`, e.i)
}

const MonthMax = 12

// ErrMonth if month is invalid
type ErrMonth struct {
	i int
}

func NewErrMonth(m int) ErrMonth {
	return ErrMonth{i: m}
}
func (e ErrMonth) Error() string {
	return fmt.Sprintf(`invalid month: %d`, e.i)
}

const DayMax = 31

// ErrDay if day is invalid
type ErrDay struct {
	i int
}

func NewErrDay(d int) ErrDay {
	return ErrDay{i: d}
}
func (e ErrDay) Error() string {
	return fmt.Sprintf(`invalid day: %d`, e.i)
}

const HourMax = 23

// ErrHour if hour is invalid
type ErrHour struct {
	i int
}

func NewErrHour(h int) ErrHour {
	return ErrHour{i: h}
}

func (e ErrHour) Error() string {
	return fmt.Sprintf(`invalid hour: %d`, e.i)
}

const MinuteMax = 59

// ErrMinute if minute is invalid
type ErrMinute struct {
	i int
}

func NewErrMinute(m int) ErrMinute {
	return ErrMinute{i: m}
}

func (e ErrMinute) Error() string {
	return fmt.Sprintf(`invalid minute: %d`, e.i)
}

const SecondMax = 59

// ErrSecond if minute is invalid
type ErrSecond struct {
	i int
}

func NewErrSecond(i int) ErrSecond {
	return ErrSecond{i: i}
}

func (e ErrSecond) Error() string {
	return fmt.Sprintf(`invalid second: %d`, e.i)
}

const MilliSecondMax = 999

// ErrMilliSecond if minute is invalid
type ErrMilliSecond struct {
	i int
}

func NewErrMilliSecond(ms int) ErrMilliSecond {
	return ErrMilliSecond{i: ms}
}

func (e ErrMilliSecond) Error() string {
	return fmt.Sprintf(`invalid millisecond: %d`, e.i)
}

const MicroSecondMax = 999

// ErrMicroSecond if minute is invalid
type ErrMicroSecond struct {
	i int
}

func NewErrMicroSecond(ms int) ErrMicroSecond {
	return ErrMicroSecond{i: ms}
}

func (e ErrMicroSecond) Error() string {
	return fmt.Sprintf(`invalid Microsecond: %d`, e.i)
}
