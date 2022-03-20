package main

import "testing"

func TestHello(t *testing.T) {


	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("to a person", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", spanish)
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Dan", french)
		want := "Bonjour, Dan"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Turkish", func(t *testing.T) {
		got := Hello("Mustafa", turkish)
		want := "Merhaba, Mustafa"
		assertCorrectMessage(t, got, want)
	})
}