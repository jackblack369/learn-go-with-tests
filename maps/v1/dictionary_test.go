package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	got := dictionary.Search("test")
	want := "this is just a test"

	assertStrings(t, got, want)
}

func TestAsianCountry_Search(t *testing.T) {
	asianCountry := AsianCountry{1: "Indonesia", 2: "Malaysia", 3: "Singapore", 4: "Thailand", 5: "Vietnam", 6: "Philippines"}

	got := asianCountry.Search(2)
	want := "Indonesia"

	assertStrings(t, got, want)

}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}
