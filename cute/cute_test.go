package cute

import "testing"

func TestHelloWithName(t *testing.T) {
	got := Hello("Mina")
	want := "Hi, Mina. Welcome to the tiny cat cafe."

	if got != want {
		t.Fatalf("Hello() = %q, want %q", got, want)
	}
}

func TestHelloWithoutName(t *testing.T) {
	got := Hello("")
	want := "Hi, friend. Welcome to the tiny cat cafe."

	if got != want {
		t.Fatalf("Hello() = %q, want %q", got, want)
	}
}
////test