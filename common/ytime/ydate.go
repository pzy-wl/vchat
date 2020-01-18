package ytime

import (
	"database/sql/driver"
	"fmt"
	"log"
	"time"
)

const (
	//TimeZone ...
	TimeZone = "Asia/Shanghai"
	//CustomDateFmt ...
	CustomDateFmt = "2006-01-02 15:04:05"

	//DateLayout ...
	DateLayout = "2006-01-02"
)

func init() {
	SetTimeZone()
	log.Println("-----ytime init-called---------")
}

//SetTimeZone ...
func SetTimeZone() {
	lc, err := time.LoadLocation(TimeZone)
	if err == nil {
		time.Local = lc
	}
}

//Date ...
type Date struct {
	time.Time
}

//根据当天日期返回一个0时的时间值
func OfToday() Date {
	return Date{time.Now()}.ToDate()
}

//返回一个含时、分、秒的时间值
func OfNow() Date {
	return Date{Time: time.Now()}
}

//OfDatetime ...
func OfStr(in string) (Date, error) {
	out, err := time.ParseInLocation(CustomDateFmt, in, time.Local)
	return Date{out}, err
}

func OfTime(t time.Time) Date {
	return Date{t}
}

func OfInt(year, month, day int, l ...int) Date {
	hour, minute, second := 0, 0, 0
	if len(l) > 0 {
		hour = l[0]
	}
	if len(l) > 1 {
		minute = l[1]
	}
	if len(l) > 2 {
		second = l[2]
	}
	s := fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", year, month, day, hour, minute, second)
	bean, err := OfStr(s)
	if err != nil {
		return OfNow()
	}
	return bean
}

//String ...
func (p Date) ToDate() Date {
	return OfInt(p.Year(), int(p.Month()), p.Day())
}

//String ...
func (p Date) ToStr() string {
	return p.Format(CustomDateFmt)
}

func (p Date) ToStrShort() string {
	return p.Format("2006-01-02_15_04_05")
}

func (p Date) ToStrDate() string {
	return p.Format("2006-01-02")
}

func (t Date) MarshalJSON() ([]byte, error) {
	tune := t.Format(`"2006-01-02 15:04:05"`)
	return []byte(tune), nil
}

//UnmarshalJSON ...
func (p *Date) UnmarshalJSON(data []byte) error {
	local, err := time.ParseInLocation(`"`+CustomDateFmt+`"`, string(data), time.Local)
	if err != nil {
		*p = Date{Time: time.Now()}
	}
	*p = Date{Time: local}
	return nil
}

// Value insert timestamp into mysql need this function.
func (t Date) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// Scan value of time.Time
func (t *Date) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = Date{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

func (t Date) TimeShanghai() time.Time {
	return t.Time.In(time.Local)
}
