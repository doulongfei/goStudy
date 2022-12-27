package offer

func deleteNode(head *ListNode, val int) *ListNode {

	res := &ListNode{}
	res.Next = head
	pre := res
	cur := head
	for cur != nil {
		if cur.Val == val {
			pre.Next = cur.Next
			break
		}
		pre = cur
		cur = cur.Next
	}
	return res.Next

	//p := res
	//for cur := head; nil != cur.Next; cur = cur.Next {
	//	if val == p.Next.Val {
	//		p.Next = cur.Next
	//		break
	//	}
	//	p = cur
	//}
	//return res.Next
}
