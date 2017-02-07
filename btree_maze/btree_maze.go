package btree_maze

import (
	"fmt"
	"math/rand"
)

type Node struct {
	bottomWall, rightWall bool
}

func GenerateMaze(width, height int) [][]Node {
	maze := initMaze(width, height)
	sy, sx := 0, 0

	for y := sy; y < height; y++ {
		for x := sx; x < width; x++ {
			if y != sy {
				//if right
				if rand.Intn(2) == 0 {
					if x != width {
						maze[y][x].rightWall = false
					} else {
						maze[y-1][x].bottomWall = false
					}
				} else {
					maze[y-1][x].bottomWall = false
				}

			} else if x != width {
				maze[y][x].rightWall = false
			}
		}
	}
	DrawMaze(maze)
	return maze
}

func initMaze(width, height int) [][]Node {
	var maze [][]Node = make([][]Node, width)
	for i := 0; i < width; i++ {
		maze[i] = make([]Node, height)
		for j := 0; j < height; j++ {
			maze[i][j].bottomWall = true
			maze[i][j].rightWall = true
		}
	}

	return maze
}
