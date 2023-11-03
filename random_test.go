package iwant

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func BenchmarkRandString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandString(16, LowerAlpha+UpperAlpha+Number+Symbol)
	}
}

func TestRandString(t *testing.T) {
	pool := LowerAlpha + UpperAlpha
	l := 16
	s := RandString(l, pool)
	assert.Equal(t, l, len(s))
	for _, c := range s {
		assert.Contains(t, pool, string(c))
	}
}
