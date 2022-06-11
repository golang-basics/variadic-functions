package main

import "fmt"

func main() {
	intNumbers := []int{1, 2, 3}
	anyNumbers := []any{1, 2, 3}
	fmt.Println(1, 2, 3)
	fmt.Println(anyNumbers...)
	// does not work because slice types do not match
	// the ...intNumbers will attempt to convert it to []interface{} and will fail
	// fmt.Println(intNumbers...)

	// converting each element of the intNumbers slice to any (interface{})
	// will make sure the unpacked resulting slice's type is of the same type
	var iNumbers = make([]any, len(intNumbers))
	for i, num := range intNumbers {
		iNumbers[i] = num
	}
	fmt.Println(iNumbers...)
}
