package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input1         int
		expectedOutput []int
	}{
		{
			[]int{2, 4},
			[]int{6},
		},
	}

	for _, test := range tests {
		assert.Equal(test.expectedOutput, add())
	}
}
