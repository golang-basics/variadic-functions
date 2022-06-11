package main

import "fmt"

func main() {
	variadic("1")
	variadic("2", []int{}...)
	variadic("3", []int(nil)...)
	variadic("4", nil...)
}

func variadic(label string, nums ...int) {
	if nums == nil {
		fmt.Println(label, "nil")
	}
	// will always result in at least
	// an empty slice (not a nil slice)
	fmt.Println(nums)
}
