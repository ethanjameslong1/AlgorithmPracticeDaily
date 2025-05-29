package main

import (
	"fmt"
)

func main() {
	vals, weights, capacity := make([]int, 0, 10), make([]int, 0, 10), 20

	s := greedySearch(vals, weights, capacity)
	fmt.Println(s)
}

func greedySearch(value []int, weight []int, capacity int) []int {
	ratio := make([]float32, 0, len(value))
	for i := range value {
		ratio[i] = float32(value[i]) / float32(weight[i])
	}
	fmt.Println("Finding best ratio given capacity")

	return make([]int, 0, 10)
}

func roundGreedy(ratio []float32, capcity int) {

}
