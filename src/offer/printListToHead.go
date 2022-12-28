package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func printListFromTailToHead(head *ListNode) (ans []int) {
	var stack []int
	for head != nil {
		stack = append(stack, head.Val)
		head = head.Next
	}
	for i := len(stack) - 1; i >= 0; i-- {
		ans = append(ans, stack[i])
	}
	return ans
}
