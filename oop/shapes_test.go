package oop

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestPerimeter(t *testing.T) {
	assert.Equal(t, 40.0, Perimeter(Rectangle{10.0, 10.0}))
}

func TestArea(t *testing.T) {
	assertArea(t, Rectangle{10.0, 10.0}, 100.0)
	assertArea(t, Circle{10.0}, 314.1592653589793)
	assertArea(t, Triangle{12.0, 6.0}, 36.0)
}

func assertArea(t *testing.T, shape Shape, expected float64) {
	t.Helper()
	t.Run(reflect.TypeOf(shape).Name(), func(t *testing.T) {
		t.Helper()
		assert.Equal(t, expected, shape.Area())
	})
}
