// using iterative method
package main

import "fmt"

func binarySearch(datalist []int, target int) int {
	start_index := 0
	end_index := len(datalist) - 1

	for start_index <= end_index {

		mid := (start_index + end_index) / 2

		if datalist[mid] < target {
			start_index = mid + 1
		} else {
			end_index = mid - 1
		}
	}

	if start_index == len(datalist) || datalist[start_index] != target {
		return -1
	} else {
		return start_index
	}
}

func main() {
	numbers := []int{1, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59}
	target := 37

	result := binarySearch(numbers, target)
	fmt.Printf("The number %d is found in position %d.", target, result)
}
