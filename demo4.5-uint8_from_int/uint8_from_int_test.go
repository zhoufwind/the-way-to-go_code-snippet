package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUint8FromInt_negative(t *testing.T) {
	ast := assert.New(t)
	var a int = -1
	b, err := Uint8FromInt(a)
	fmt.Printf("%d, %v\n", b, err)
	ast.Error(err, "-1 is out of the uint8 range")
}

func TestUint8FromInt_0(t *testing.T) {
	ast := assert.New(t)
	var a int = 0
	b, err := Uint8FromInt(a)
	fmt.Printf("%d, %v\n", b, err)
	ast.Equal(b, uint8(0), "0 is within the uint8 range(0 to 255)")
}

func TestUint8FromInt_1(t *testing.T) {
	ast := assert.New(t)
	var a int = 1
	b, err := Uint8FromInt(a)
	fmt.Printf("%d, %v\n", b, err)
	ast.Equal(b, uint8(1), "0 is within the uint8 range(0 to 255)")
}

func TestUint8FromInt_255(t *testing.T) {
	ast := assert.New(t)
	var a int = 255
	b, err := Uint8FromInt(a)
	fmt.Printf("%d, %v\n", b, err)
	ast.Equal(b, uint8(255), "0 is within the uint8 range(0 to 255)")
}

func TestUint8FromInt_256(t *testing.T) {
	ast := assert.New(t)
	var a int = 256
	b, err := Uint8FromInt(a)
	fmt.Printf("%d, %v\n", b, err)
	ast.Error(err, "256 is out of the uint8 range")
}
