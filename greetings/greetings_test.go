package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "X"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("X")

	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("X") = %q, %v, want match for %#q, nul`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")

	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
