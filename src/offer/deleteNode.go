package offer

func deleteNode(head *ListNode, val int) *ListNode {

	res := &ListNode{}
	res.Next = head
	p := res
	for cur := head; nil != cur.Next; cur = cur.Next {
		if val == p.Next.Val {
			p.Next = cur.Next
			break
		}
		p = cur
	}
	return res.Next
}
