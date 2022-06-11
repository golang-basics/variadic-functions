package main

func main() {
	f(slice)
	// cannot use this because:
	// []T is different from ...T
	// f(variadic)
}

func f(cb func([]int)) {
	cb([]int{1, 2, 3})
}

func slice([]int) {
}

func variadic(...int) {
}
