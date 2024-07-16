package basic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func AddOne(x int) int {
	return x + 1
}

func TestAddOne(t *testing.T) {
	assert.Equal(t, 2, AddOne(1), "they should be equal")
}
