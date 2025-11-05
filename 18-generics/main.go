package main


// generics in functions
// we can pass comparable instead of writing every type 
func printSlice[T int | string](items []T) { // like this -> printSlice[T comparable]
	for _, item := range items {
		println(item)
	}
}

// generics in structs
type stack[T any] struct {
	elements []T
}

func main() {
	// generics in functions
	nums := []int{1, 2, 3, 4, 5, 6}

	names := []string{"golang", "javascript"}

	printSlice(nums)
	printSlice(names)


	// generics in structs
	myStack := stack[int]{elements: []int{10, 20, 30}}
	for _, elem := range myStack.elements {
		println(elem)
	}
}
