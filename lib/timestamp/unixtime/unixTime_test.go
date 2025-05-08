package unixtime_test

import (
	"testing"

	"github.com/ArkhamCookie/crabot/lib/timestamp/unixtime"
)

func TestUnixTimeBasic(t *testing.T) {
	output, err := unixtime.UnixTime(2025, 12, 12, 12, 12, 00, "UTC")

	if err != nil {
		t.Fatal("error thrown:", err)
	}

	if output != 1765541532 {
		t.Fatalf("unix time did not match, wanted: 1765541532 got: %v", output)
	}
}
