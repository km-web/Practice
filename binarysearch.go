package main

import (
	"fmt"
	"sort"
)

func binarySearch(i int, a []int) bool {

	start := 0
	end := len(a) - 1

	for start <= end {
		mid := (start + end) / 2

		if a[mid] < i {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return true
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	x := 2
	fmt.Println(binarySearch(x, a))
	i := sort.Search(len(a), func(i int) bool { return a[i] >= x })
	if i < len(a) && a[i] == x {
		fmt.Printf("found %d at index %d in %v\n", x, i, a)
	} else {
		fmt.Printf("%d not found in %v\n", x, a)
	}
}
