package iwant

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestPtr(t *testing.T) {
	data := rand.Int63()
	ptr := Ptr(data)
	assert.Equal(t, *ptr, data)
}
