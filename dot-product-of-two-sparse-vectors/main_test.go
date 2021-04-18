package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSparseVector_dotProduct(t *testing.T) {
	//  [0,1,0,0,0,0,0,3] * [0,2,0,0,0,0,0,0] = 2
	//v1 := Constructor([]int{0,1,0,0,0,0,0,3})
	//v2 := Constructor([]int{0,2,0,0,0,0,0,0})
	//assert.Equal(t, 2, v1.dotProduct(v2))


	//	[0,0,0,0,0,0,0,0] * [1,2,3,4,5,7,8] = 0
	//v1 = Constructor([]int{0,0,0,0,0,0,0,0})
	//v2 = Constructor([]int{1,2,3,4,5,7,8})
	//assert.Equal(t, 0, v1.dotProduct(v2))

	//	[1,0,0,0,0,0,0,1] * [1,2,3,4,5,7,8] = 9
	v1 := Constructor([]int{1,0,0,0,0,0,0,1})
	v2 := Constructor([]int{1,2,3,4,5,6,7,8})
	assert.Equal(t, 9, v1.dotProduct(v2))

}