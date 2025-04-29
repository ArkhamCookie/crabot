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

	output := timestamp.ShortDate()

	if !want.MatchString(output) {
		t.Fatalf("`timestamp.ShortDate()` = %q want match for %#q\n", output, want)
	}
}

func TestLongDate(t *testing.T) {
	now := currentTime()

	want := regexp.MustCompile("<t:" + now + ":D>")

	output := timestamp.LongDate()

	if !want.MatchString(output) {
		t.Fatalf("`timestamp.LongDate()` = %q want match for %#q\n", output, want)
	}
}

func TestShortTime(t *testing.T) {
	now := currentTime()

	want := regexp.MustCompile("<t:" + now + ":t>")

	output := timestamp.ShortTime()

	if !want.MatchString(output) {
		t.Fatalf("`timestamp.ShortTime()` = %q want match for %#q\n", output, want)
	}
}

func TestLongTime(t *testing.T) {
	now := currentTime()

	want := regexp.MustCompile("<t:" + now + ":T>")

	output := timestamp.LongTime()

	if !want.MatchString(output) {
		t.Fatalf("`timestamp.LongTime()` = %q want match for %#q\n", output, want)
	}
}

func TestShortFull(t *testing.T) {
	now := currentTime()

	want := regexp.MustCompile("<t:" + now + ":f>")

	output := timestamp.ShortFull()

	if !want.MatchString(output) {
		t.Fatalf("`timestamp.ShortFull()` = %q want match for %#q\n", output, want)
	}
}

func TestLongFull(t *testing.T) {
	now := currentTime()

	want := regexp.MustCompile("<t:" + now + ":F>")

	output := timestamp.LongFull()

	if !want.MatchString(output) {
		t.Fatalf("`timestamp.LongFull()` = %q want match for %#q\n", output, want)
	}
}

func TestRelativeTime(t *testing.T) {
	now := currentTime()

	want := regexp.MustCompile("<t:" + now + ":R>")

	output := timestamp.RelativeTime()

	if !want.MatchString(output) {
		t.Fatalf("`timestamp.RelativeTime()` = %q want match for %#q\n", output, want)
	}
}
