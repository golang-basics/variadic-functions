package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(numbers)

	variadic(numbers...)
	fmt.Println(numbers)
}

func variadic(args ...int) {
	for i := 0; i < len(args); i++ {
		args[i] = (i + 1) * 10
	}
}
