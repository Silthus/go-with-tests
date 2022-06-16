package assert

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func That[V](t *testing.T, object V) (assert *ObjectAsserter[V]) {
	return &ObjectAsserter[V]{
		t:      t,
		object: object,
	}
}

func ThatString(t *testing.T, str string) (stringAssert *StringAsserter) {
	return &StringAsserter{
		t:   t,
		str: str,
	}
}

type Asserter interface {
	IsEqualTo(expected interface{}) Asserter
}

type ObjectAsserter[V interface{}] struct {
	Asserter
	t      *testing.T
	object V
}

func (a *ObjectAsserter[V]) IsEqualTo(expected interface{}) *ObjectAsserter[V] {
	assert.Equal(a.t, expected, a.object)
	return a
}

type E any

func (a *ObjectAsserter[V]) Extracting(f func(a *ObjectAsserter[V]) any) *ObjectAsserter[any] {
	return That(a.t, f(a))
}

type StringAsserter struct {
	Asserter
	t   *testing.T
	str string
}

func (a *StringAsserter) IsEqualTo(equal interface{}) *StringAsserter {
	assert.Equal(a.t, equal, a.str)
	return a
}

func (a *StringAsserter) Contains(contains string) *StringAsserter {
	assert.Contains(a.t, a.str, contains)
	return a
}
