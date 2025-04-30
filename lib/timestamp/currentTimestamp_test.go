package timestamp_test

import (
	"regexp"
	"strconv"
	"testing"
	"time"

	"crabot/timestamp"
)

func currentTime() string {
	timeStart := time.Now().Unix()
	convertedTime := strconv.FormatInt(timeStart, 10)

	return convertedTime
}

func TestCurrentShortDate(t *testing.T) {
	now := currentTime()

	want := regexp.MustCompile("<t:" + now + ":d>")

	output := timestamp.CurrentShortDate()

	if !want.MatchString(output) {
		t.Fatalf("`timestamp.CurrentShortDate()` = %q want match for %#q\n", output, want)
	}
}

func TestCurrentLongDate(t *testing.T) {
	now := currentTime()

	want := regexp.MustCompile("<t:" + now + ":D>")

	output := timestamp.CurrentLongDate()

	if !want.MatchString(output) {
		t.Fatalf("`timestamp.CurrentLongDate()` = %q want match for %#q\n", output, want)
	}
}

func TestCurrentShortTime(t *testing.T) {
	now := currentTime()

	want := regexp.MustCompile("<t:" + now + ":t>")

	output := timestamp.CurrentShortTime()

	if !want.MatchString(output) {
		t.Fatalf("`timestamp.CurrentShortTime()` = %q want match for %#q\n", output, want)
	}
}

func TestCurrentLongTime(t *testing.T) {
	now := currentTime()

	want := regexp.MustCompile("<t:" + now + ":T>")

	output := timestamp.CurrentLongTime()

	if !want.MatchString(output) {
		t.Fatalf("`timestamp.CurrentLongTime()` = %q want match for %#q\n", output, want)
	}
}

func TestCurrentShortFull(t *testing.T) {
	now := currentTime()

	want := regexp.MustCompile("<t:" + now + ":f>")

	output := timestamp.CurrentShortFull()

	if !want.MatchString(output) {
		t.Fatalf("`timestamp.CurrentShortFull()` = %q want match for %#q\n", output, want)
	}
}

func TestCurrentLongFull(t *testing.T) {
	now := currentTime()

	want := regexp.MustCompile("<t:" + now + ":F>")

	output := timestamp.CurrentLongFull()

	if !want.MatchString(output) {
		t.Fatalf("`timestamp.CurrentLongFull()` = %q want match for %#q\n", output, want)
	}
}

func TestCurrentRelativeTime(t *testing.T) {
	now := currentTime()

	want := regexp.MustCompile("<t:" + now + ":R>")

	output := timestamp.CurrentRelativeTime()

	if !want.MatchString(output) {
		t.Fatalf("`timestamp.CurrentRelativeTime()` = %q want match for %#q\n", output, want)
	}
}
