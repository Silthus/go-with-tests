package hello_test

import (
	"hello"
	"testing"
)

func TestHello(t *testing.T) {
	assertHello := func(t testing.TB, input, expected string) {
		t.Helper()
		eq(t, hello.Hello(input), expected)
	}

	t.Run("Hello(...) with name", func(t *testing.T) {
		assertHello(t, "Jakob", "Hello Jakob!")
	})
	t.Run("Hello(...) with empty string prints 'Hello world!'", func(t *testing.T) {
		assertHello(t, "", "Hello world!")
	})
}

func eq(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
