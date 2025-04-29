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

func TestShortDate(t *testing.T) {
	now := currentTime()

	want := regexp.MustCompile("<t:" + now + ":d>")

	output := timestamp.CurrentShortDate()

	if !want.MatchString(output) {
		t.Fatalf("`timestamp.CurrentShortDate()` = %q want match for %#q\n", output, want)
	}
}

func TestLongDate(t *testing.T) {
	now := currentTime()

	want := regexp.MustCompile("<t:" + now + ":D>")

	output := timestamp.CurrentLongDate()

	if !want.MatchString(output) {
		t.Fatalf("`timestamp.CurrentLongDate()` = %q want match for %#q\n", output, want)
	}
}

func TestShortTime(t *testing.T) {
	now := currentTime()

	want := regexp.MustCompile("<t:" + now + ":t>")

	output := timestamp.CurrentShortTime()

	if !want.MatchString(output) {
		t.Fatalf("`timestamp.CurrentShortTime()` = %q want match for %#q\n", output, want)
	}
}

func TestLongTime(t *testing.T) {
	now := currentTime()

	want := regexp.MustCompile("<t:" + now + ":T>")

	output := timestamp.CurrentLongTime()

	if !want.MatchString(output) {
		t.Fatalf("`timestamp.CurrentLongTime()` = %q want match for %#q\n", output, want)
	}
}

func TestShortFull(t *testing.T) {
	now := currentTime()

	want := regexp.MustCompile("<t:" + now + ":f>")

	output := timestamp.CurrentShortFull()

	if !want.MatchString(output) {
		t.Fatalf("`timestamp.CurrentShortFull()` = %q want match for %#q\n", output, want)
	}
}

func TestLongFull(t *testing.T) {
	now := currentTime()

	want := regexp.MustCompile("<t:" + now + ":F>")

	output := timestamp.CurrentLongFull()

	if !want.MatchString(output) {
		t.Fatalf("`timestamp.CurrentLongFull()` = %q want match for %#q\n", output, want)
	}
}

func TestRelativeTime(t *testing.T) {
	now := currentTime()

	want := regexp.MustCompile("<t:" + now + ":R>")

	output := timestamp.CurrentRelativeTime()

	if !want.MatchString(output) {
		t.Fatalf("`timestamp.CurrentRelativeTime()` = %q want match for %#q\n", output, want)
	}
}
