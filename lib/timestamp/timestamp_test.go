package timestamp_test

import (
	"testing"

	"crabot/timestamp"
)

func TestShortDate(t *testing.T) {
	want := "<t:1765541532:d>"
	output, err := timestamp.ShortDate(2025, 12, 12, 12, 12, 12, "UTC")
	if err != nil {
		t.Fatal(err)
	}

	if output != want {
		t.Fatalf("`timestamp.ShortDate(2025, 12, 12, 12, 12, 12, \"UTC\")` = %q want match for %q\n", output, want)
	}
}

func TestLongDate(t *testing.T) {
	want := "<t:1765541532:D>"
	output, err := timestamp.LongDate(2025, 12, 12, 12, 12, 12, "UTC")
	if err != nil {
		t.Fatal(err)
	}

	if output != want {
		t.Fatalf("`timestamp.LongDate(2025, 12, 12, 12, 12, 12, \"UTC\")` = %q want match for %q\n", output, want)
	}
}

func TestShortTime(t *testing.T) {
	want := "<t:1765541532:t>"
	output, err := timestamp.ShortTime(2025, 12, 12, 12, 12, 12, "UTC")
	if err != nil {
		t.Fatal(err)
	}

	if output != want {
		t.Fatalf("`timestamp.ShortTime(2025, 12, 12, 12, 12, 12, \"UTC\")` = %q want match for %q\n", output, want)
	}
}

func TestLongTime(t *testing.T) {
	want := "<t:1765541532:T>"
	output, err := timestamp.LongTime(2025, 12, 12, 12, 12, 12, "UTC")
	if err != nil {
		t.Fatal(err)
	}

	if output != want {
		t.Fatalf("`timestamp.LongTime(2025, 12, 12, 12, 12, 12, \"UTC\")` = %q want match for %q\n", output, want)
	}
}

func TestShortFull(t *testing.T) {
	want := "<t:1765541532:f>"
	output, err := timestamp.ShortFull(2025, 12, 12, 12, 12, 12, "UTC")
	if err != nil {
		t.Fatal(err)
	}

	if output != want {
		t.Fatalf("`timestamp.ShortFull(2025, 12, 12, 12, 12, 12, \"UTC\")` = %q want match for %q\n", output, want)
	}
}

func TestLongFull(t *testing.T) {
	want := "<t:1765541532:F>"
	output, err := timestamp.LongFull(2025, 12, 12, 12, 12, 12, "UTC")
	if err != nil {
		t.Fatal(err)
	}

	if output != want {
		t.Fatalf("`timestamp.LongFull(2025, 12, 12, 12, 12, 12, \"UTC\")` = %q want match for %q\n", output, want)
	}
}

func TestRelativeTime(t *testing.T) {
	want := "<t:1765541532:R>"
	output, err := timestamp.RelativeTime(2025, 12, 12, 12, 12, 12, "UTC")
	if err != nil {
		t.Fatal(err)
	}

	if output != want {
		t.Fatalf("`timestamp.RelativeTime(2025, 12, 12, 12, 12, 12, \"UTC\")` = %q want match for %q\n", output, want)
	}
}

func TestUnix(t *testing.T) {
	want := "1765541532"

	output, err := timestamp.Unix(2025, 12, 12, 12, 12, 12, "UTC")
	if err != nil {
		t.Fatal(err)
	}

	if output != want {
		t.Fatalf("`timestamp.Unix(2025, 12, 12, 12, 12, 12, \"UTC\")` = %q want match for %q\n", output, want)
	}
}
