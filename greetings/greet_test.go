package greetings

import (
	"testing"
)

func TestHelloEmpty(t *testing.T) {
	msg, error := Greet("")
	if msg != "" || error == nil {
		t.Fatalf(`Hello("") = %q, %v , want "", error`, msg, error)
	}
}
