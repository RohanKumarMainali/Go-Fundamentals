package dsa

import (
	"fmt"
	"slices"
)

// initialize array
// var arr [5] int
// intialize array in one line
// arr := [4]int {
// 1,2,3,4
// }
//
// if size is not defined it is slice
// arr := []int {
// 1,2,3,4
// }
// use make to create slice
// arr := make([]int,{
// 1,2,3,4
// })

func Array() {
	numbers := [3]int{1, 23, 3}
	for x := 0; x < len(numbers); x++ {
		fmt.Println(numbers[x])
	}
}

func TestArray() {
	var numbers [5]int

	for x := 0; x < len(numbers); x++ {
		numbers[x] = x + 1
	}
	fmt.Println(numbers)
}

func SortedSlice() {
	slice := []int{}
	slice = append(slice, 33, 2, 5)
	slices.Sort(slice)
	fmt.Println(slice)
	fmt.Println(slices.BinarySearch(slice, 2))
}
