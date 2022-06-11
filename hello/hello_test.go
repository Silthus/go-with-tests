package hello_test

import (
	"hello"
	"testing"
)

func TestHello(t *testing.T) {
	got := hello.Hello("Jakob")
	want := "Hello Jakob!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
