package strings

import (
	"testing"
)

func TestHello(t *testing.T) {
	assertHello := func(t testing.TB, input, expected string) {
		t.Helper()
		eq(t, Hello(input, ""), expected)
	}
	assertI18nHello := func(t *testing.T, name string, lang string, expected string) {
		t.Helper()
		eq(t, Hello(name, lang), expected)
	}

	t.Run("Hello(...) with name", func(t *testing.T) {
		assertHello(t, "Jakob", "Hello Jakob!")
	})
	t.Run("Hello(...) with empty string prints 'Hello world!'", func(t *testing.T) {
		assertHello(t, "", "Hello world!")
	})
	t.Run("Hello(...) with german lang prints Hallo", func(t *testing.T) {
		assertI18nHello(t, "Eva", "de", "Hallo Eva!")
	})
	t.Run("Hello(...) with french lang prints Bonjour", func(t *testing.T) {
		assertI18nHello(t, "Michael", "fr", "Bonjour Michael!")
	})
}

func eq(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
