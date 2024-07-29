package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}

	Countdown(buffer)

	got := buffer.String()
	want := `3
2
1
Go!`

	//	want2 := `3
	//4
	//5
	//6
	//7
	//8
	//9
	//10
	//Go!
	//`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
