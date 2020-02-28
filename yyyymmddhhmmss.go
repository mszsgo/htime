// 日期时间14位长度格式处理
package htime

import (
	"errors"
	"time"
)

const (
	YYYYMMDDHHMMSS = "20060102150405"
	YYYYMMDD       = "20060102"
	HHMMSS         = "150405"
)

var (
	ErrParseTime = errors.New("解析时间格式错误")
)

// 取当前日期时间

func NowYYYYMMDDHHMMSS() string {
	return time.Now().Format(YYYYMMDDHHMMSS)
}

func NowYYYYMMDD() string {
	return time.Now().Format(YYYYMMDD)
}

func NowHHMMSS() string {
	return time.Now().Format(HHMMSS)
}

// 解析日期字符串

func ParseYYYYMMDDHHMMSS(v string) (time.Time, error) {
	t, e := time.Parse(YYYYMMDDHHMMSS, v)
	if e != nil {
		return time.Time{}, ErrParseTime
	}
	return t, nil
}

func ParseYYYYMMDD(v string) (time.Time, error) {
	t, e := time.Parse(YYYYMMDD, v)
	if e != nil {
		return time.Time{}, ErrParseTime
	}
	return t, nil
}
