package htime

import (
	"errors"
	"fmt"
	"time"
)

const (
	Z8 = "2006-01-02T15:04:05"
)

func NowZ8() string {
	return time.Now().Format(Z8)
}

func ParseZ8(v string) time.Time {
	t, e := time.Parse(Z8, v)
	if e != nil {
		panic(errors.New(fmt.Sprintf("99210:%s", e.Error())))
	}
	return t
}

func FormatZ8(t time.Time) string {
	return t.Local().Format(Z8)
}
