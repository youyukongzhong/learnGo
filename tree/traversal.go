package tree

// 中序遍历
// 先访问左子树，再访问当前节点，最后访问右子树。
func (node *Node) Traverse() {
	if node == nil {
		return
	}
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}
