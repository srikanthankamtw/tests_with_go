package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	actual := Hello("Satya")
	expected := "Hello, Satya"
	if expected != actual {
		t.Errorf("Expceted %q, Actual %q", expected, actual)
	}
}
