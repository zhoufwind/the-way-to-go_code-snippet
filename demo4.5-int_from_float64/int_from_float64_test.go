package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntFromFloat64_MaxInt32(t *testing.T) {
	ast := assert.New(t)
	var a float64 = 2147483647.0
	b := IntFromFloat64(a)
	ast.Equal(b, 2147483647, "2147483647 is within the int32 range(-2147483648 to 2147483647)")
}

func TestIntFromFloat64_MaxInt32_fraction_less(t *testing.T) {
	ast := assert.New(t)
	var a float64 = 2147483646.6
	b := IntFromFloat64(a)
	ast.Equal(b, 2147483647, "2147483647 is within the int32 range(-2147483648 to 2147483647)")
}

// How to test panics?
// Ref: https://stackoverflow.com/questions/31595791/how-to-test-panics
func TestIntFromFloat64_MaxInt32_fraction_greater(t *testing.T) {
	var a float64 = 2147483647.1
	assert.Panics(t, func() { IntFromFloat64(a) }, "Any value which greater than 2147483647 is out of the int32 range(-2147483648 to 2147483647)")
}

func TestIntFromFloat64_MinInt32(t *testing.T) {
	ast := assert.New(t)
	var a float64 = -2147483648.0
	b := IntFromFloat64(a)
	ast.Equal(b, -2147483648, "-2147483648 is within the int32 range(-2147483648 to 2147483647)")
}

func TestIntFromFloat64_MinInt32_fraction_less(t *testing.T) {
	var a float64 = -2147483648.1
	assert.Panics(t, func() { IntFromFloat64(a) }, "Any value which less than -2147483648.0 is out of the int32 range(-2147483648 to 2147483647)")
}

func TestIntFromFloat64_MinInt32_fraction_greater(t *testing.T) {
	ast := assert.New(t)
	var a float64 = -2147483647.9
	b := IntFromFloat64(a)
	ast.Equal(b, -2147483647, "-2147483647 is within the int32 range(-2147483648 to 2147483647)")
}
