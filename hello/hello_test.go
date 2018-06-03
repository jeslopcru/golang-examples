package main

import "testing"

func TestHelloWorld(t *testing.T) {
	expected := "Hello, world"
	got := HelloWorld()

	if expected != got {
		t.Errorf("got '%s' expected '%s'", got, expected)
	}
}
