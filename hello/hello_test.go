package main

import "testing"

func TestHelloWorld(t *testing.T) {
	expected := "Hello, Jesús"
	got := HelloWorld("Jesús")

	if expected != got {
		t.Errorf("got '%s' expected '%s'", got, expected)
	}
}
