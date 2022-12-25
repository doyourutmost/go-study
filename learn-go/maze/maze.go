package main

import (
	"fmt"
	"os"
)

func readMaze(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	var row, col int
	if _, err := fmt.Fscanf(file, "%d %d", &row, &col); err != nil {
		panic(err)
	}
	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			if _, err := fmt.Fscanf(file, "%d", &maze[i][j]); err != nil {
				//panic(err)
			}
		}
	}
	return maze
}

type point struct {
	i, j int
}

var dirs = [4]point{
	{-1, 0}, // 上
	{0, -1}, // 左
	{1, 0},  // 下
	{0, 1},  // 右
}

func (receiver point) add(p point) point {
	return point{receiver.i + p.i, receiver.j + p.j}
}

func (receiver point) at(grid [][]int) (int, bool) {
	if receiver.i < 0 || receiver.i >= len(grid) { // 判断上下越界
		return 0, false
	}
	if receiver.j < 0 || receiver.j >= len(grid[receiver.i]) { // 判断左右越界
		return 0, false
	}
	return grid[receiver.i][receiver.j], true
}

func bfs(maze [][]int, start, end point) [][]int {
	for _, row := range maze {
		for _, val := range row {
			fmt.Print(val)
		}
		fmt.Println()
	}
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}
	Q := []point{start}

	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:]

		if cur == end {
			break
		}
		for _, dir := range dirs {
			next := cur.add(dir)
			// 出界 或者 撞墙
			if val, ok := next.at(maze); !ok || val == 1 {
				continue
			}
			// 出界 或者 走走过的地方
			if val, ok := next.at(steps); !ok || val != 0 {
				continue
			}
			// 起点
			if next == start {
				continue
			}
			curSteps, _ := cur.at(steps)
			steps[next.i][next.j] = curSteps + 1
			Q = append(Q, next)
		}
	}
	return steps
}
func main() {
	maze := readMaze("maze/maze.in")

	for _, row := range maze {
		for _, val := range row {
			fmt.Print(val)
		}
		fmt.Println()
	}

	steps := bfs(maze, point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1})
	for _, row := range steps {
		for _, val := range row {
			fmt.Printf("%3d", val)
		}
		fmt.Println()
	}
}
