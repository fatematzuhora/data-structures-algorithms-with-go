// using recursive method
package main

import "fmt"

func binarySearch(datalist []int, target int) (result int, searchCount int) {
	mid := len(datalist) / 2

	switch {
	case len(datalist) == 0:
		result = -1 // not found
	case datalist[mid] > target:
		result, searchCount = binarySearch(datalist[:mid], target)
	case datalist[mid] < target:
		result, searchCount = binarySearch(datalist[mid+1:], target)
		if result >= 0 { // if anything but the -1 "not found" result
			result += mid + 1
		}
	default: // a[mid] == search
		result = mid // found
	}

	searchCount++
	return
}

func main() {
	numbers := []int{1, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59}
	target := 37

	result, searchCount := binarySearch(numbers, target)
	fmt.Printf("The number %d is found in position %d after %d steps.", target, result, searchCount)
}
