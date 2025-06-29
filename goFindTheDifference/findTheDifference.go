package main

import "fmt"

func main() {
	c := findTheDifference("mynameis", "mamyrisne")
	fmt.Println(c)
}

func findTheDifference(s string, t string) byte {
	mymap := make(map[rune]int)
	for _, c := range s {
		mymap[c]++
	}
	for _, c := range t { //looking for the odd character out, s and t should be the same but for one char
		mymap[c]--

		if mymap[c] < 0 {
			return byte(c)
		}
	}
	return byte('c')
}
