package main

import "fmt"

//func deleteNode(head *ListNode, val int) *ListNode {
//
//	res := &ListNode{}
//	res.Next = head
//	pre := res
//	cur := head
//	for cur != nil {
//		if cur.Val == val {
//			pre.Next = cur.Next
//			break
//		}
//		pre = cur
//		cur = cur.Next
//	}
//	return res.Next
//
//	//p := res
//	//for cur := head; nil != cur.Next; cur = cur.Next {
//	//	if val == p.Next.Val {
//	//		p.Next = cur.Next
//	//		break
//	//	}
//	//	p = cur
//	//}
//	//return res.Next
//}

func ReOrderArray(array []int) []int {
	// write code here

	length := len(array)
	nums := make([]int, length)

	left, right := 0, length-1

	tpLeft, tpRight := left, right

	for left < length && right >= 0 {
		if array[left]%2 == 1 {
			nums[tpLeft] = array[left]
			tpLeft++
			left++
		}

		if array[right]%2 == 0 {
			nums[tpRight] = array[right]
			tpRight--

			right--
		}
	}
	return nums

}

func main() {
	arr := []int{1, 2, 3, 4}

	array := ReOrderArray(arr)

	for _, a := range array {
		fmt.Print(a)
	}
	fmt.Println()
}

func FindKthToTail(pHead *ListNode, k int) *ListNode {
	// write code here
	if pHead == nil || k == 0 {
		return nil
	}

	stack := make([]*ListNode, 0)

	for pHead != nil {

		stack = append(stack, pHead)
		pHead = pHead.Next
	}
	if len(stack) < k {
		return nil
	}

	for i := 0; i < k-1; i++ {
		stack = stack[:len(stack)-1]
	}

	return stack[len(stack)-1]

}

func FindKthToTail1(pHead *ListNode, k int) *ListNode {
	// write code here
	fast, slow := pHead, pHead

	for i := 0; i < k; i++ {
		if fast != nil {
			fast = fast.Next
		} else {
			return nil
		}
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}
