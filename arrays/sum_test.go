package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSum(t *testing.T) {
	assert.Equal(t, 10, Sum([]int{2, 4, 4}))
}

func TestSumAll(t *testing.T) {
	assert.Equal(t, []int{3, 9}, SumAll([]int{1, 2}, []int{0, 9}))
}

func TestSumAllTails(t *testing.T) {
	assert.Equal(t, []int{
		2, 5, 0, 0,
	}, SumAllTails(
		[]int{1, 2},
		[]int{1, 2, 3},
		[]int{2},
		[]int{},
	))
}
