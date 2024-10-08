package helloworld

import "testing"

func TestHello(t *testing.T) {
	t.Run("Hello w/o name", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying Hello in Spanish", func(t *testing.T) {
		got := Hello("Andre", "Spanish")
		want := "Hola, Andre"
		assertCorrectMessage(t, got, want)
	})
	t.Run("saying Hello in French", func(t *testing.T) {
		got := Hello("Elise", "French")
		want := "Bonjour, Elise"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Got %q want %q", got, want)
	}
}
