package main

import (
	"fmt"
	"learngo_muke/tree"
)

// myTreeNode 自定义结构体 myTreeNode 的后序遍历（postOrder Traversal）功能
type myTreeNode struct {
	*tree.Node //Embedding
}

// postOrder 后序遍历
// 后序遍历：一种树形数据结构的遍历方式。
// 后序遍历的顺序是：左子树 -> 右子树 -> 当前节点。
func (myNode *myTreeNode) postOrder() {
	//首先检查当前 myNode 是否为 nil，以及其内部的 node 是否为 nil，如果是，就直接返回，防止空指针错误
	if myNode.Node == nil || myNode.Node == nil {
		return
	}

	left := myTreeNode{myNode.Left}
	right := myTreeNode{myNode.Right}

	left.postOrder()
	right.postOrder()
	myNode.Print()
}

func main() {
	root := myTreeNode{&tree.Node{Value: 3}} // 初始化 root 节点的值为 3
	root.Left = &tree.Node{}                 // 设置 root 节点的左子节点为一个新的空 treeNode 节点
	root.Right = &tree.Node{5, nil, nil}     // 设置 root 节点的右子节点为值为 5 的新节点，左右子节点为 nil
	root.Right.Left = new(tree.Node)         // 设置 root 的右子节点的左子节点为一个新创建的 treeNode 节点，默认零值
	root.Left.Right = tree.CreateNode(2)     // 使用工厂函数 createNode 创建一个值为 2 的新节点，作为 root 的左子节点的右子节点

	// 设置 root 的右子节点的左子节点的值为 4，并调用其 Print 方法输出值
	root.Right.Left.SetValue(4)
	root.Right.Left.Print()
	fmt.Println()

	// 调用 root 的 Print 方法。由于 Print 方法不改变节点值，因此值不会影响 root
	root.Print() // 输出：3
	fmt.Println()

	// 调用 root 的 SetValue 方法，将 root 节点的值设置为 100
	// SetValue 方法接受指针，所以 root 本身的值会被更改。
	root.SetValue(100)

	// 声明一个 treeNode 指针变量 pRoot
	var pRoot *tree.Node
	// 调用 pRoot 的 SetValue 方法。pRoot 目前为 nil，因此调用 SetValue 时会检测到 node == nil，
	// 输出 "setting value to nil node. Ignored" 并返回，不会导致崩溃
	pRoot.SetValue(200)
	// 将 pRoot 设置为指向 root 节点的地址
	pRoot = root.Node

	// 调用 pRoot 的 SetValue 方法，将 root 节点的值设置为 300
	// 因为 pRoot 指向 root，所以 root 的值会更改。
	pRoot.SetValue(300)
	// 调用 pRoot 的 Print 方法打印 root 节点的值。
	// 结果输出 300。

	fmt.Println("In-order traversal:")
	root.Traverse()
	fmt.Println()

	fmt.Print("My own post-order traversal:")
	root.postOrder()
	fmt.Println()
}
