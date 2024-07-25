package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, world"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHello2(t *testing.T) {
	got := HelloBrook()
	want := "Hello, Brook"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
