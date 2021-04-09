// To execute Go code, please declare a func main() in a package "main"

package main

/*
   1. Paraphase The Problem
       Given an m*n array of arrays that represent islands (1) and water,
       count the Islands.

   2. Test Cases
       1 island
         ["1","1","1","1","0"],
         ["1","1","0","1","0"],
         ["1","1","0","0","0"],
         ["0","0","0","0","0"]

       4 islands
        1 0 1
        0 0 0
        1 0 1

   3. Approach
       Go through the array of arrays and visit every coordinate. If you hit and island, then you neat to "walk" the land. Visit every spot that that is adjacent looking for 1s. Alwauys keep track of places you've visisted.

   4. Psuedo Code

     loop through first dimension
           loop through second dimension
             if we've visted this piece already continue
             if this is not an island mark it as visited, contune
               if this is an island create a slice with all the adjcacent pieces
               we can avoid the ones we obviously have already visisted
               mark the piece as visisted
               while this slice is not empty
                 check the co-ordinate is visited, if it is continue
                 if the co-rodinate is not an island, markit visisted, cotinue
                 if the co-ordinate is an island, then mark it visisted, put all
                 the adjacent pieces onto the slice


*/

func numIslands(grid [][]byte) int {
	// create a new matrix to track visited so we don't mutate the slice externally

	if len(grid) == 0 {
		return 0
	}
	dimx := len(grid)
	dimy := len(grid[0])

	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[i]))
	}

	islands := 0
	// loop through first dimension
	for x, xrow := range grid {
		// loop through second dimension
		for y, val := range xrow {
			// if we've visted this piece already continue
			if visited[x][y] {
				continue
			}

			// this spot has now been visted
			visited[x][y] = true

			// if this is not an island mark it as visited, continue
			if val != '1' {
				continue
			}

			// this is part of an island, create a slice with all the adjcacent pieces
			adj := calcAdjacent(dimx, dimy, x, y)

			// while this slice is not empty
			var coord []int
			for len(adj) > 0 {
				coord, adj = adj[0], adj[1:]
				x, y := coord[0], coord[1]

				// check the coordinate is visited, if it is continue
				if visited[x][y] {
					continue
				}

				// mark visited
				visited[x][y] = true

				// if the coordinate is not an island, mark it visited, continue
				if grid[x][y] != '1' {
					continue
				}

				// put all the adjacent coordinates in the the slice to make sure we check
				adj = append(adj, calcAdjacent(dimx, dimy, x, y)...)
			}
			islands++
		}
	}
	return islands
}

func calcAdjacent(dimx int, dimy int, x int, y int) [][]int {

	var adj [][]int

	// just bounds check up down left and right and  put them into the result
	if x-1 > -1 {
		adj = append(adj, []int{x - 1, y})
	}
	if x+1 < dimx {
		adj = append(adj, []int{x + 1, y})

	}
	if y-1 > -1 {
		adj = append(adj, []int{x, y - 1})
	}
	if y+1 < dimy {
		adj = append(adj, []int{x, y + 1})

	}
	return adj
}
