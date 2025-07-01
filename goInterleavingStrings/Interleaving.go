package main

import "fmt"

func main() {
	s1, s2, s3 := "aa", "bb", "abab"

	fmt.Println(isInterleave(s1, s2, s3))
}
func isInterleave(s1 string, s2 string, s3 string) bool {
	ls1, ls2, ls3 := len(s1), len(s2), len(s3)

	if ls3 != ls1+ls2 {
		return false
	}
	ansArr := make([][]bool, len(s1)+1)
	for i := range len(s1) + 1 {
		ansArr[i] = make([]bool, len(s2)+1)
		for j := range len(s2) + 1 {
			ansArr[i][j] = false
		}
	}
	vArr := make([][]bool, len(s1)+1)
	for i := range len(s1) + 1 {
		vArr[i] = make([]bool, len(s2)+1)
		for j := range len(s2) + 1 {
			vArr[i][j] = false
		}
	}
	return recHelper(0, 0, 0, s1, s2, s3, ansArr, vArr)
}

func recHelper(i int, j int, k int, s1 string, s2 string, s3 string, ansArr [][]bool, visited [][]bool) bool {
	if k == len(s3) {
		return true
	}
	if visited[i][j] {
		return ansArr[i][j]
	}
	ans := false
	if i < len(s1) && k < len(s3) && s1[i] == s3[k] {
		ans = recHelper(i+1, j, k+1, s1, s2, s3, ansArr, visited)
	}
	if !ans && j < len(s2) && k < len(s3) && s2[j] == s3[k] {
		ans = recHelper(i, j+1, k+1, s1, s2, s3, ansArr, visited)
	}

	visited[i][j] = true
	ansArr[i][j] = ans
	return ans
}
