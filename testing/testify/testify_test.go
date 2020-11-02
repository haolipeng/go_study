package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Calculate(x int) (result int) {
	result = x + 2
	return result
}

func TestCalculate(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		input    int
		expected int
	}{
		{2, 4},
		{-1, 1},
		{0, 2},
		{-5, -3},
		{99999, 100001},
	}

	for _, test := range tests {
		assert.Equal(Calculate(test.input), test.expected)
	}
}

func TestSomething(t *testing.T) {
	//断言相等
	assert.Equal(t, 123, 123, "they shoule be equal")

	//断言不相等
	assert.NotEqual(t, 123, 456, "they should not be equal")
}
