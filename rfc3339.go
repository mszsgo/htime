package htime

import (
	"errors"
	"fmt"
	"time"
)

func RFC3339() string {
	return time.Now().Format(time.RFC3339)
}

func FRFC3339(t time.Time) string {
	return t.Local().Format(time.RFC3339)
}

func ParseRFC3339(v string) time.Time {
	t, e := time.Parse(time.RFC3339, v)
	if e != nil {
		panic(errors.New(fmt.Sprintf("99210:%s", e.Error())))
	}
	return t
}
