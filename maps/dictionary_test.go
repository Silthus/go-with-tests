package maps

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var dictionary = Dictionary{
	"test": "just a test",
}

func TestSearch(t *testing.T) {
	t.Run("known word", func(t *testing.T) {
		assertDefinition(t, "test", "just a test")
	})
	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		assert.ErrorIs(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new definition", func(t *testing.T) {
		err := dictionary.Add("foo", "bar")
		assert.NoError(t, err)
		assertDefinition(t, "foo", "bar")
	})
	t.Run("existing definition throws error", func(t *testing.T) {
		err := dictionary.Add("test", "should fail")
		assert.ErrorIs(t, err, ErrWordExists)
		assertDefinition(t, "test", "just a test")
	})
}

func assertDefinition(t *testing.T, word, expectedResult string) {
	t.Helper()
	result, err := dictionary.Search(word)
	assert.NoError(t, err)
	assert.Equal(t, expectedResult, result)
}
