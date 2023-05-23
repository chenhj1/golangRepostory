package main

import "fmt"

func main() {
	arr := []int{28, 3, 53, 2, 23, 7, 14, 10}

	quickSort(arr, 0, len(arr)-1) // 快排

	fmt.Println(arr)
}
