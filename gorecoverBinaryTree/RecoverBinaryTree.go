package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	theRoot := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 9,
			Left: &TreeNode{
				Val: 1,
			},
		},
		Right: &TreeNode{
			Val: 2,
		},
	}
	recoverTree(theRoot)

}

func recoverTree(root *TreeNode) *TreeNode {
	var prev, a, b *TreeNode
	inOrder(root, func(p *TreeNode) {
		if prev != nil && prev.Val > p.Val {
			if a == nil {
				a = prev
			}
			b = p
		}
		prev = p
	})
	swap(a, b)
	return root

}

func inOrder(root *TreeNode, doFunc func(p *TreeNode)) {
	if root == nil {
		return
	}
	inOrder(root.Left, doFunc)
	doFunc(root)
	inOrder(root.Right, doFunc)
}

func swap(a, b *TreeNode) {
	if a != nil && b != nil {
		a.Val, b.Val = b.Val, a.Val
	}
}
