package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	// initial/obvious test cases
	assert.ElementsMatch(t, []string{}, fizzBuzz(0))
	assert.ElementsMatch(t, []string{"1"}, fizzBuzz(1))
	assert.ElementsMatch(t, []string{}, fizzBuzz(-1))
	assert.ElementsMatch(t, []string{"1", "2", "Fizz"}, fizzBuzz(3))
	assert.ElementsMatch(t, []string{"1", "2", "Fizz", "4", "Buzz"}, fizzBuzz(5))
	assert.ElementsMatch(t, []string{"1", "2", "Fizz", "4", "Buzz",  "Fizz",  "7",  "8", "Fizz",  "Buzz",  "11", "Fizz", "13", "14",  "FizzBuzz"}, fizzBuzz(15))
	// 3. test case second pass
	// 4. implementation
}
