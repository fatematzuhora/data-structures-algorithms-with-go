// using built-in sort.Search()
package main

import (
	"fmt"
	"sort"
)

func main() {
	datalist := []int{1, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59}
	target := 37

	i := sort.Search(len(datalist), func(i int) bool { return datalist[i] >= target })

	if i < len(datalist) && datalist[i] == target {
		fmt.Printf("found %d at index %d", target, i)
	} else {
		fmt.Printf("%d not found", target)
	}
}
