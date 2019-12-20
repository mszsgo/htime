package htime

import (
	"errors"
	"fmt"
	"time"
)

func NowRFC3339() string {
	return time.Now().Format(time.RFC3339)
}

func ParseRFC3339(v string) time.Time {
	t, e := time.Parse(time.RFC3339, v)
	if e != nil {
		panic(errors.New(fmt.Sprintf("99210:%s", e.Error())))
	}
	return t
}
