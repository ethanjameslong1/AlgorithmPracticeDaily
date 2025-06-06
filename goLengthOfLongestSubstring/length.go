package main

import (
	"fmt"
)

func main() {
	test := "pwwkew"
	fmt.Println(test)
	fmt.Println(lengthOfLongestSubstring(test))
	/*
		I think I go thorugh the strng and look for a substring of length i from i = 0 to max
		for each iteration if i find a substrng equal to i, i increment i and loop it, when I fail to find one i reutrn i - 1
	*/
}

func lengthOfLongestSubstring(s string) int {
	for i := 1; i < 50000; i++ {
		if !findSubstring(s, i) {
			return i - 1
		}
	}
	return 0
}

func findSubstring(s string, l int) bool {

	for i := range s {
		if i > len(s)-l {
			return false
		}
		mm := make(map[byte]bool)
		found := false
		for j := i; j <= i-1+l; j++ {

			if mm[s[j]] {
				found = true
			} else {
				mm[s[j]] = true
			}

		}
		if !found {
			fmt.Println(i)
			return true
		}

	}
	return false
}
