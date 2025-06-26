package main

func main() {
	println(reverse(123))
}

func reverse(x int) int {
	reversed := 0
	for x != 0 {
		val1 := x % 10
		x = x / 10
		if reversed > ((1<<31)-1)/10 || reversed < -1*((1<<31)-1)/10 {
			return 0
		}
		reversed = reversed*10 + val1
	}
	return reversed
}
