package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.Equal(t, "134", addStrings("11", "123"))
}
