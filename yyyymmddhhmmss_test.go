package htime

import (
	"testing"
)

func TestYYYYMMDDHHMMSS(t *testing.T) {
	t.Log(NowYYYYMMDDHHMMSS())
}

func TestHHMMSS(t *testing.T) {
	t.Log(NowHHMMSS())
}

func TestYYYYMMDD(t *testing.T) {
	t.Log(NowYYYYMMDD())
}

func TestParseYYYYMMDDHHMMSS(t *testing.T) {
	time, e := ParseYYYYMMDDHHMMSS("20200228094456")
	t.Log(e)
	t.Log(time)
}
