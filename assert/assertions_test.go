package assert_test

import (
	"go-with-tests/assert"
	"testing"
)

func TestIsEqualTo(t *testing.T) {
	assert.That(t, "test").IsEqualTo("test")
	assert.That(t, 2).IsEqualTo(2)
	assert.That(t, []string{"1", "2", "3"}).IsEqualTo([]string{"1", "2", "3"})
	assert.That(t, map[string]interface{}{
		"test": "test",
		"deep": map[string]string{
			"equal": "true",
		},
	}).IsEqualTo(map[string]interface{}{
		"test": "test",
		"deep": map[string]string{
			"equal": "true",
		},
	})
}

func TestContains(t *testing.T) {
	assert.ThatString(t, "foo bar is cool").
		Contains("cool").
		IsEqualTo("foo bar is cool")
}

type TestStruct struct {
	prop1 string
	prop2 *TestStruct
}
