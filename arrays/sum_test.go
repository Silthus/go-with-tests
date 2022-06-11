package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSum(t *testing.T) {
	assert.Equal(t, 10, Sum([]int{2, 4, 4}))
}
