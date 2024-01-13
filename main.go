package main

import (
	"fmt"
	// "practice/go/dsa"
	// "practice/go/pointers"
	"practice/go/routine"
	"strings"
)

func print() {
	fmt.Println("Hello")
}

func convertString(name string) string {
	arr := strings.Split(name, "/")
	fmt.Println(arr)
	ans := "/"
	var stack []string
	for i := 0; i < len(arr); i++ {
		if arr[i] == ".." && len(stack) != 0 {
			stack = stack[:len(stack)-1]
		}
		if arr[i] != ".." && arr[i] != "" {
			stack = append(stack, arr[i])
		}
	}
	for i, directory := range stack {
		ans += directory
		if i < len(stack)-1 {
			ans += "/"
		}
	}

	print()
	// pointers.Pointer()
	// dsa.PrintMap()
	routine.TestRoutine()
	return ans
}

func main() {
	// message := convertString("/home/")

	routine.Concurrency()
	// fmt.Println(message)
}
