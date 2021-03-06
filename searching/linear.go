package main

import "fmt"

func linearSearch(datalist []int, target int) bool {
	for _, item := range datalist {
		if item == target {
			return true
		}
	}
	return false
}

func main() {
	numbers := []int{10, 14, 19, 26, 27, 31, 33, 35, 42, 44}

	fmt.Println(linearSearch(numbers, 33))
}
