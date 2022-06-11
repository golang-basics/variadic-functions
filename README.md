# Variadic function in Go (Practical use cases)

Variadic functions in Go (practical use cases)

### Motivation

- Accept an arbitrary number of arguments
- Simulate optional arguments

### Examples

Here are some examples of variadic function in the Go standard library

### Rules

- variadic argument has to be the last one

### Pros

- Clean and elegant API

### Cons

- Generic and ambiguous arguments and argument names
- No automatic conversion for `[]interface{}` aka `[]any`

### Caveats

Returning the passed slice
You can’t use a variadic param as a result type, but, you can return it as a slice.

func f(nums ...int) []int {
nums[1] = 10
return nums
}

### Resources

- [Appending and copying slices - Go Spec](https://go.dev/ref/spec#Appending_and_copying_slices)
- [Passing arguments to variadic functions - Go Spec](https://go.dev/ref/spec#Passing_arguments_to_..._parameters)
- [Types Assignability - Go Spec](https://go.dev/ref/spec#Assignability)
- [Types Identity - Go Spec](https://go.dev/ref/spec#Type_identity)
