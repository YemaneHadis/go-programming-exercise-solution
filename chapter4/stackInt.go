package main

import "fmt"

type Stack struct {
	stackSlice []int
}

func (stack *Stack) push(number int) {
	stack.stackSlice = append(stack.stackSlice, number)

}
func (stack *Stack) poll() int {
	size := len(stack.stackSlice)
	top := stack.stackSlice[size-1]
	stack.stackSlice = stack.stackSlice[:size-1]
	return top
}

func main() {
	stack1 := Stack{
		stackSlice: []int{1, 2, 3},
	}
	fmt.Println("Before Output")
	fmt.Println(stack1)

	stack1.push(12)
	fmt.Println("After Output")
	fmt.Println(stack1)

	stack1.poll()
	fmt.Println(stack1)

	stack1.push(13)
	fmt.Println("After Output")
	fmt.Println(stack1)

}
