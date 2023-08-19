package main

import (
	"fmt"
	"strconv"
	"strings"

	"slices"
)

func main() {

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	idx, ok := slices.BinarySearch(nums, 5)
	fmt.Println(idx, ok)

	data := []int{1, 10, 11, 2} // sorted lexicographically
	cmp := func(a int, b string) int {
		return strings.Compare(strconv.Itoa(a), b)
	}
	pos, found := slices.BinarySearchFunc(data, "2", cmp)
	fmt.Println(pos, found)
}
