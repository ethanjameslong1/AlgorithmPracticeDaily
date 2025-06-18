package main

import "fmt"

func main() {
	fmt.Println("How many ways are there to climb 5 stairs? Assuming i'm able to step 1 or 2 stairs at a time?")

	fmt.Printf("There are %d ways to climb 5 stairs given that information", climbStairs(5))

}

func climbStairs(n int) int {
	x, y := 0, 1
	for i := 1; i <= n; i++ {
		x, y = y, x+y
	}
	return y
}
