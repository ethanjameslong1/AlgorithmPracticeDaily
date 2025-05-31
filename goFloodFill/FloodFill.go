package main

import (
	"container/list"
	"fmt"
)

type Point struct {
	X int
	Y int
}

func main() {
	matrix := [][]int{
		{1, 0, 0, 1, 0, 0},
		{1, 1, 0, 1, 0, 1},
		{1, 0, 1, 1, 1, 0},
		{1, 0, 1, 1, 0, 1},
		{1, 0, 0, 1, 0, 1},
		{1, 1, 0, 1, 1, 0},
	}
	for i := range matrix {
		fmt.Println(matrix[i])
	}

	matrix = transformGraph(matrix, 1, 1)

	for i := range matrix {
		fmt.Println(matrix[i])
	}
}

// takes in a 2d graph with 0's and 1's, all pixels connected to the passed in coords will be turned to a 3.
func transformGraph(matrix [][]int, x int, y int) [][]int {
	marked := make([][]bool, len(matrix))

	for i := range marked {
		marked[i] = make([]bool, len(matrix[0]))
	}
	queue := list.New()
	queue.PushFront(Point{x, y})
	for queue.Len() > 0 {

		point := queue.Front().Value.(Point)
		queue.Remove(queue.Front())
		if marked[point.X][point.Y] == true {
			continue
		}
		for _, pt := range nextDoor(matrix, point) {
			queue.PushFront(pt)
		}
		marked[point.X][point.Y] = true
		matrix[point.X][point.Y] = 3
	}

	return matrix
}

func nextDoor(matrix [][]int, point Point) []Point {
	var Neighbors []Point
	row, col := len(matrix), len(matrix[0])
	if point.X < 0 || point.X >= row || point.Y < 0 || point.Y >= col {
		return Neighbors
	}
	curValue := matrix[point.X][point.Y]
	directions := []Point{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}
	for _, dir := range directions {
		nx, ny := point.X+dir.X, point.Y+dir.Y
		if nx < 0 || nx >= row || ny < 0 || ny >= col {
			continue
		}
		if matrix[nx][ny] == curValue {
			Neighbors = append(Neighbors, Point{nx, ny})
		}
	}
	return Neighbors
}
