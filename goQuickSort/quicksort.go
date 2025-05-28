package main
import "fmt"
import "math/rand"


func main () {
	var arr []int = make([]int, 100)
	for i:=0;i<100;i++{
		arr[i] = rand.Intn(100) + 1
		fmt.Println(arr[i])
	}
	quickSort(arr, 0, 99)
	fmt.Println("SORTED\n\n")
	for i:=0;i<100;i++{
		fmt.Println(arr[i])
	}


}


func partition(array []int, start int, end int) int {
	pivot := array[end]
	i := start-1
	for j:=start; j<end; j++ {
		if array[j] < pivot {
			i++
			array[i], array[j] = array[j], array[i]
		}
	}
	array[i+1], array[end] = array[end], array[i+1]
	return i+1
}

func quickSort(array[]int, start int, end int) {
	if end <= start {return}
	pivot := partition(array, start, end)
	quickSort(array, start, pivot-1)
	quickSort(array, pivot+1, end)
}
