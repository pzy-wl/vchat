package ytime

import (
	"time"
)



func Today() time.Time {
	t := time.Now()
	return DayTime(t)
}

func DayTime(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
}

func HourTime(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), 0, 0, 0, time.Local)
}

func MinuteTime(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(),
		t.Minute(),
		0,
		0, time.Local)
}

func SecondTime(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(),
		t.Minute(),
		t.Second(),
		0, time.Local)
}
