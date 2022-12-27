package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func reBinaaryTree(pre []int, vin []int) *TreeNode {
	if len(pre) == 0 {
		return nil
	}
	root := &TreeNode{pre[0], nil, nil}
	i := 0
	for i := 0; i < len(vin); i++ {
		if vin[i] == pre[0] {
			break
		}
	}
	root.Left = reBinaaryTree(pre[1:i+1], vin[:i])
	root.Right = reBinaaryTree(pre[i+1:], vin[i+1:])
	return root
}
