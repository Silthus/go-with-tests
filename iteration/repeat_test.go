package iteration

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepeat(t *testing.T) {
	assert.Equal(t, "aaaaa", Repeat("a", 5))
	assert.Equal(t, "abab", Repeat("ab", 2))
	assert.Panicsf(t, func() {
		Repeat("b", -1)
	}, "Repeat called with negative count.")
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 0)
	}
}
