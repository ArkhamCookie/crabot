package crabtalk_test

import (
	"regexp"
	"testing"

	"crabot/crabtalk"
)

func TestCrabtalkBasic(t *testing.T) {
	input := "foo"
	want := regexp.MustCompile("clickclickclackclick clackclackclack clackclackclack")

	output, err := crabtalk.Get(input)
	if !want.MatchString(output) || err != nil {
		t.Fatalf(`crabtalk.Get("foo") = %q, %v, want match for %#q, nil`, output, err, want)
	}
}
