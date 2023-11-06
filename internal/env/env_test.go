package env_test

import (
	"testing"

	"internal/env"
)

var (
	err error
	// input string
	output string
	want string
)

func TestGetEnvValue(t *testing.T) {
	inputValue := "GET_TEST"
	inputFile := "test.env"
	want = "load-test-successful"

	output, err = env.GetEnvValue(inputValue, inputFile)
	if err != nil {
		t.Fatal("Error:", err)
	}
	if want != output {
		t.Fatalf(
			`env.GetEnvValue(%q, %q) = %q, want match for %#q`,
			inputValue, inputFile, output, want,
		)
	}
}
