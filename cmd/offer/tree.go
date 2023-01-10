package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param pRoot1 TreeNode类
 * @param pRoot2 TreeNode类
 * @return bool布尔型
 */
func HasSubtree(pRoot1 *TreeNode, pRoot2 *TreeNode) bool {
	// write code here

	if pRoot2 == nil {
		return false
	}
	if pRoot1 == nil && pRoot2 != nil {
		return false
	}
	if pRoot1 == nil || pRoot2 == nil {
		return true
	}
	flag1 := recursion(pRoot1, pRoot2)
	flag2 := HasSubtree(pRoot1.Left, pRoot2)
	flag3 := HasSubtree(pRoot1.Right, pRoot2)
	return flag1 || flag2 || flag3
}

func recursion(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil && root2 != nil {
		return false
	}
	if root1 == nil || root2 == nil {
		return true
	}
	if root1.Val != root2.Val {
		return false
	}
	return recursion(root1.Left, root2.Left) && recursion(root1.Right, root2.Right)
}
