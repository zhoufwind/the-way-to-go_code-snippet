package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSameLengthLessValue(t *testing.T) {
	assert := assert.New(t)
	a := []byte{'g', 'o', 'l', 'a', 'n', 'a'}
	b := []byte{'g', 'o', 'l', 'a', 'n', 'b'}
	assert.Equal(Compare(a, b), -1, "a should be less")
}

func TestSameLengthSameValue(t *testing.T) {
	assert := assert.New(t)
	a := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	assert.Equal(Compare(a, b), 0, "a & b should be same")
}

func TestSameLengthGreaterValue(t *testing.T) {
	assert := assert.New(t)
	a := []byte{'g', 'o', 'l', 'a', 'n', 'b'}
	b := []byte{'g', 'o', 'l', 'a', 'n', 'a'}
	assert.Equal(Compare(a, b), 1, "a should be greater")
}

func TestLessLength(t *testing.T) {
	assert := assert.New(t)
	a := []byte{'g', 'o', 'l', 'a', 'n'}
	b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	assert.Equal(Compare(a, b), -1, "a should be less")
}

func TestGreaterLength(t *testing.T) {
	assert := assert.New(t)
	a := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	b := []byte{'g', 'o', 'l', 'a', 'n'}
	assert.Equal(Compare(a, b), 1, "a should be greater")
}
