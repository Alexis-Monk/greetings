package greetings

import (
	"regexp"
	"testing"
)

// Coincidencia exacta en los nombres
func TestHelloName(t *testing.T) {
	name := "Alex"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Alex")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Alex") = %q, %v, quiere coincidencia para %#q, nil`, msg, err, want)
	}
}

// Devolver un error
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, quiere "", error`, msg, err)
	}
}
