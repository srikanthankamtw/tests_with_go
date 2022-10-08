package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("say Hello with name", func(t *testing.T) {
		actual := Hello("Satya", "")
		expected := "Hello, Satya"
		assertCorrectMessage(t, actual, expected)
	})

	t.Run("say Hello, world when empty string supplied", func(t *testing.T) {
		actual := Hello("", "")
		expected := "Hello, world"
		assertCorrectMessage(t, actual, expected)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "French")
		want := "Bonjour, Elodie"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, actual string, expected string) {
	t.Helper()
	if expected != actual {
		t.Errorf("Expceted %q, Actual %q", expected, actual)
	}
}
