package main

import "fmt"

// variadic functions accepts muliple parameters 
// we can say these functions accepts unlimited parameters
// for example fmt.Println(1,2,3,4,5,6,7,8,"hello")
func operation(nums ... int) (int,int) {
	total := 0
	mul := 1
	for _, num := range nums {
		total += num
		mul *= num
	}
	return total,mul
}

func main()  {

	nums := []int{1,2,3,4,5,6,7,8,9,10}
	fmt.Println(operation(nums...))
}
