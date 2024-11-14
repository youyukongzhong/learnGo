package tree

import "fmt"

// Traverse 中序遍历
// 先访问左子树，再访问当前节点，最后访问右子树。
func (node *Node) Traverse() {
	node.TraverseFunc(func(node *Node) {
		node.Print() // 调用当前节点的 Print 方法
	})
	fmt.Println() // 在遍历结束后换行
}

func (node *Node) TraverseFunc(f func(*Node)) {
	if node == nil {
		return // 遇到空节点，直接返回
	}

	// 中序遍历：先左子树
	node.Left.TraverseFunc(f)

	// 然后执行对当前节点的操作
	f(node)

	// 最后右子树
	node.Right.TraverseFunc(f)
}
