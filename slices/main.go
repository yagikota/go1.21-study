package main

import (
	"fmt"

	"golang.org/x/exp/slices"
)

// BinarySearch/BinarySearchFunc
// func main() {

// 	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

// 	idx, ok := slices.BinarySearch(nums, 5)
// 	fmt.Println(idx, ok)

// 	data := []int{1, 10, 11, 2} // sorted lexicographically
// 	cmp := func(a int, b string) int {
// 		return strings.Compare(strconv.Itoa(a), b)
// 	}
// 	pos, found := slices.BinarySearchFunc(data, "2", cmp)
// 	fmt.Println(pos, found)
// }

// Contains/ContainsFunc
// func main() {
// 	numbers := []int{0, 42, -10, 8}
// 	hasNegative := slices.Contains(numbers, -10)
// 	fmt.Println("Has a negative:", hasNegative)
// }

// IndexFunc
// func main() {
// 	numbers := []int{0, 42, -10, 8}
// 	i := slices.IndexFunc(numbers, func(n int) bool {
// 		return n < 0
// 	})
// 	fmt.Println("First negative at index", i)
// }

// Clip
// func main() {
// 	s := make([]int, 0, 5)
// 	fmt.Printf("slice: %v, cap:%v , len: %v \n", s, cap(s), len(s))
// 	s = append(s, 0)
// 	fmt.Printf("slice: %v, cap:%v , len: %v \n", s, cap(s), len(s))
// 	s = slices.Clip(s)
// 	fmt.Printf("%v, %v\n", s, s)
// 	fmt.Printf("slice: %v, cap:%v , len: %v \n", s, cap(s), len(s))
// }

// https://go.dev/blog/slices-intro
// s[:len(s):len(s)] vs s[:len(s)]
func main() {
	s1 := make([]int, 0, 5)
	fmt.Printf("slice1: %v, cap:%v , len: %v \n", s1, cap(s1), len(s1))
	s1 = s1[:3]
	fmt.Printf("slice1: %v, cap:%v , len: %v \n", s1, cap(s1), len(s1))

	s2 := make([]int, 0, 5)
	fmt.Printf("slice2: %v, cap:%v , len: %v \n", s2, cap(s2), len(s2))
	// s2 = s2[:len:cap]
	s2 = s2[:3:3]
	fmt.Printf("slice2: %v, cap:%v , len: %v \n", s2, cap(s2), len(s2))
}
