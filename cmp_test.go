package iwant

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsZero(t *testing.T) {
	// string
	assert.True(t, IsZero(""))
	// int
	assert.True(t, IsZero(0))
	// int64
	assert.True(t, IsZero(int64(0)))
}
