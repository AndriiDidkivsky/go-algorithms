package btree_maze

import (
	"math/rand"
)

type Node struct {
	BottomWall, RightWall bool
}

func GenerateMaze(width, height int) [][]Node {
	maze := initMaze(width, height)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if y != 0 {
				//if right
				if rand.Intn(2) == 0 {
					if x != width {
						maze[y][x].RightWall = false
					} else {
						maze[y-1][x].BottomWall = false
					}
				} else {
					maze[y-1][x].BottomWall = false
				}

			} else if x != width {
				maze[y][x].RightWall = false
			}
		}
	}
	return maze
}

func initMaze(width, height int) [][]Node {
	var maze [][]Node = make([][]Node, width)
	for i := 0; i < width; i++ {
		maze[i] = make([]Node, height)
		for j := 0; j < height; j++ {
			maze[i][j].BottomWall = true
			maze[i][j].RightWall = true
		}
	}

	return maze
}
