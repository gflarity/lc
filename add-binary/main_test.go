package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddBinary(t *testing.T) {
	assert.Equal(t, "10", addBinary("01", "1"))
	assert.Equal(t, "100", addBinary("01", "11"))
	assert.Equal(t, "11", addBinary("", "11"))
	assert.Equal(t, "1000000000000000000000000000000000000000000000000000000000000000000", addBinary("100000000000000000000000000000000000000000000000000000000000000000", "100000000000000000000000000000000000000000000000000000000000000000"))
}
