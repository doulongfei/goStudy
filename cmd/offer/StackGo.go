package main

var stack1 []int
var stack2 []int

func Push(node int) {
	// write code here

	stack1 = append(stack1, node)

	if len(stack2) == 0 {
		stack2 = append(stack2, node)
		return
	}

	if node > stack2[len(stack2)-1] {
		stack2 = append(stack2, stack2[len(stack2)-1])
	} else {
		stack2 = append(stack2, node)
	}
	return

}
func Pop() {
	// write code here
}
func Top() int {
	// write code here
}
func Min() int {
	// write code here
}
