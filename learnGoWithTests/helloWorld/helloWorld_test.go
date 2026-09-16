package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Test spanish greeting", func(t *testing.T) {
		got := Hello("Amigo", "spanish")
		want := "Hola Amigo"
		assertCorrectMessage(t, got, want)
	})
	t.Run("Test french greeting", func(t *testing.T) {
		got := Hello("Paul", "french")
		want := "Bonjour Paul"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Test empty spanish greeting", func(t *testing.T) {
		got := Hello("", "spanish")
		want := "Hola mundo"
		assertCorrectMessage(t, got, want)
	})
	t.Run("Test empty french greeting", func(t *testing.T) {
		got := Hello("", "french")
		want := "Bonjour le monde"
		assertCorrectMessage(t, got, want)
	})

}

func assertCorrectMessage(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
