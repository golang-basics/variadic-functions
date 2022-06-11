package main

import "fmt"

func main() {
	say1("1", "2", "3", "\n")
	say1(append([]string{"1", "2", "3"}, append([]string{}, "4", "5", "6", "\n")...)...)

	say2("1", "2", "3", "\n")
	args := []interface{}{"4", "5", "6"}
	say2(args...)
}

func say1(args ...string) {
	for _, a := range args {
		fmt.Print(a)
	}
}

// individual arguments can be of any type
// destructured slice arguments have to be of type interface{}
func say2(args ...interface{}) {
	fmt.Print(args...)
}
