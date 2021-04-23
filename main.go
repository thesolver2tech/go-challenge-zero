package main

import "fmt"

func main() {
	arr := []int{12, 45, 67, 10, 13, 57, 89, 90}
	result := findMin(arr)
	fmt.Println("Minimum =", result)
}

func findMin(arr []int) int {
	result := 0
	for _, v := range arr {
		if v < result {
			result = v
		}
	}
	return result
}
