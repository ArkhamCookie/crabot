package timestamp_test

import (
	"testing"

	"crabot/timestamp"
)

func TestUnixTimeBasic(t *testing.T) {
	output, err := timestamp.UnixTime(2025, 12, 12, 12, 12, 00, "utc")

	if err != nil {
		t.Fatal("error thrown:", err)
	}

	if output != 1765541532 {
		t.Fatalf("unix time did not match, wanted: 1765541532 got: %v", output)
	}
}
