package main

import "fmt"

func linearSearch(datalist []int, key int) bool {
	for _, item := range datalist {
		if item == key {
			return true
		}
	}
	return false
}

func main() {
	numbers := []int{10, 13, 8, 7, 2, 9, 0, 6}

	fmt.Println(linearSearch(numbers, 9))
}
