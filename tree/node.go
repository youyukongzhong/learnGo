package tree

import "fmt"

// treeNode结构体定义二叉树的节点结构，包含value字段存储节点值，
// left和right字段分别指向左右子节点
type Node struct {
	Value       int
	Left, Right *Node
}

// 普通方法 Print()
// 接收者类型：(node treeNode) 指定了接收者类型 treeNode，这意味着 Print() 方法会针对 treeNode 值的副本执行，而不会修改原始值。
// 调用方式：可以通过 treeNode 实例来调用，例如 node.Print()()。
// 作用：Print() 方法只是输出节点的 value，不涉及修改。
// 注意点：这里 node 是 treeNode 值的副本，如果 Print() 方法中修改了 node 的值，原 treeNode 不受影响。
func (node Node) Print() {
	fmt.Print(node.Value, " ")
}

// 指针接收者方法 setValue
// 接收者类型：(node *treeNode) 表明接收者是指向 treeNode 的指针。
// 调用方式：setValue 可以通过 treeNode 的值或指针调用，Go 会自动处理这种转换，例如 node.setValue(5) 或 &node.setValue(5)。
// 作用：setValue 方法可以修改原 treeNode 实例的值，因为它是通过指针访问的。
// 应用场景：当方法需要修改接收者的值时，指针接收者是必选项，否则只会修改副本。
// 另外在该函数中，如果 node == nil，表示当前指针为空指针，无法赋值，因此直接返回提示。
func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("setting value to nil node.Ignored")
		return
	}
	node.Value = value
}

// 工厂函数 CreateNode
// 结构：CreateNode 是一个工厂函数，它不属于 treeNode 类型，因此没有接收者。
// 返回值：返回一个 *treeNode 类型的指针。
// 作用：封装 treeNode 的初始化逻辑，创建一个新的 treeNode 实例，并返回其地址。
// 优势：便于集中管理 treeNode 的初始化过程，如后续需为新节点附加一些初始操作，可以直接在此函数中完成。
func CreateNode(value int) *Node {
	return &Node{Value: value}
}
