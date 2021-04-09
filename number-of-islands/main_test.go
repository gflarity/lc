package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalcAdjacent(t *testing.T) {
	assert.ElementsMatch(t, calcAdjacent(2,2,0,0), [][]int{{1,0}, {0, 1}}, "simple case should match")
}

func TestNumIslands(t *testing.T) {
	tc1 := [][]byte{{'1','1','1','1','0'}, {'1','1','0','1','0'},{'1','1','0','0','0'},{'0','0','0','0','0'}}
	assert.Equal(t,  1, numIslands(tc1), "there should be one island")
	tc2 := [][]byte{
	{'1','1','0','0','0'},
	{'1','1','0','0','0'},
	{'0','0','1','0','0'},
	{'0','0','0','1','1'}}
	assert.Equal(t, 3, numIslands(tc2), "there should be three islands")
}
