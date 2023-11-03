package iwant

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMap(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	fn := func(i int) int {
		return i * i
	}
	result := Map(data, fn)
	for i, d := range data {
		assert.Equal(t, d*d, result[i])
	}
}

func TestReduce(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	fn := func(r int, i int) int {
		return r + i
	}
	result := Reduce(data, fn, 0)
	assert.Equal(t, 15, result)
}

func TestFilter(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 10, 98}
	fn := func(i int) bool {
		return i%2 == 0
	}
	result := Filter(data, fn)
	assert.Equal(t, []int{2, 4, 6, 10, 98}, result)
}
