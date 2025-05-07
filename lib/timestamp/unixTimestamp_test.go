package timestamp_test

import (
	"testing"

	"crabot/timestamp"
)

func TestUnixShortDate(t *testing.T) {
	want := "<t:1765541532:d>"
	output := timestamp.UnixShortDate(1765541532)

	if output != want {
		t.Fatalf("`timestamp.UnixShortDate(1765541532)` = %q, wanted match for %q\n", output, want)
	}
}

func TestUnixLongDate(t *testing.T) {
	want := "<t:1765541532:D>"
	output := timestamp.UnixLongDate(1765541532)

	if output != want {
		t.Fatalf("`timestamp.UnixLongDate(1765541532)` = %q, wanted match for %q\n", output, want)
	}
}

func TestUnixShortTime(t *testing.T) {
	want := "<t:1765541532:t>"
	output := timestamp.UnixShortTime(1765541532)

	if output != want {
		t.Fatalf("`timestamp.UnixShortTime(1765541532)` = %q, wanted match for %q\n", output, want)
	}
}

func TestUnixLongTime(t *testing.T) {
	want := "<t:1765541532:T>"
	output := timestamp.UnixLongTime(1765541532)

	if output != want {
		t.Fatalf("`timestamp.UnixLongTime(1765541532)` = %q, wanted match for %q\n", output, want)
	}
}

func TestUnixShortFull(t *testing.T) {
	want := "<t:1765541532:f>"
	output := timestamp.UnixShortFull(1765541532)

	if output != want {
		t.Fatalf("`timestamp.UnixShortFull(1765541532)` = %q, wanted match for %q\n", output, want)
	}
}

func TestUnixLongFull(t *testing.T) {
	want := "<t:1765541532:F>"
	output := timestamp.UnixLongFull(1765541532)

	if output != want {
		t.Fatalf("`timestamp.UnixLongFull(1765541532)` = %q, wanted match for %q\n", output, want)
	}
}

func TestUnixRelativeTime(t *testing.T) {
	want := "<t:1765541532:R>"
	output := timestamp.UnixRelativeTime(1765541532)

	if output != want {
		t.Fatalf("`timestamp.UnixRelativeTime(1765541532)` = %q, wanted match for %q\n", output, want)
	}
}
