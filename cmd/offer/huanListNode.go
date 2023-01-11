package main

import (
	"fmt"
	set "github.com/deckarep/golang-set/v2"
)

func main() {
	required := set.NewSet[string]()
	required.Add("cooking")
	required.Add("english")
	required.Add("math")
	required.Add("biology")
	fmt.Print(required)
}

func EntryNodeOfLoop(pHead *ListNode) *ListNode {
	if pHead == nil {
		return nil
	}
	slow, fast := pHead, pHead

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if slow == fast {
			break
		}
	}
	if fast == nil || fast.Next == nil {
		return nil
	}
	fast = pHead

	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return fast

}

func ReverseList(pHead *ListNode) *ListNode {
	// write code here
	return pHead
}

func Merge(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	// write code here
	head := new(ListNode)
	cur := head
	for pHead1 != nil && pHead2 != nil {
		if pHead1.Val < pHead2.Val {
			cur.Next = pHead1
			pHead1 = pHead1.Next
		} else {
			cur.Next = pHead2
			pHead2 = pHead2.Next
		}
		cur = cur.Next
	}
	if pHead1 == nil {
		cur.Next = pHead2
	} else {
		cur.Next = pHead1
	}
	return head.Next
}
