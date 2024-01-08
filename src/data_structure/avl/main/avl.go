package main

import "fmt"

/**
平衡二叉树
1. 本身首先是一棵二叉搜索树
2. 带有平衡条件：每个结点的左右子树的高度之差的绝对值（平衡因子）最多为1
*/

type AVLTreeNode struct {
	Data   int
	Height int
	Left   *AVLTreeNode
	Right  *AVLTreeNode
}

type AVLTree struct {
	root *AVLTreeNode
}

// 新建一个节点
func NewNode(data int) *AVLTreeNode {
	return &AVLTreeNode{
		Data:   data,
		Height: 1,
		Left:   nil,
		Right:  nil,
	}
}

// 新建一个平衡二叉树
func NewAVLTree() *AVLTree {
	return &AVLTree{root: nil}
}

// 获取节点的高度
func (n *AVLTreeNode) height() int {
	if n == nil {
		return 0
	}
	return n.Height
}

// 获取节点的平衡因子
func (n *AVLTreeNode) balanceFactor() int {
	return n.Left.height() - n.Right.height()
}

// max 获取两个数的最大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 更新节点的高度
func (n *AVLTreeNode) updateHeight() {
	n.Height = max(n.Left.height(), n.Right.height()) + 1
}

// 右旋操作
// 对节点y进行向右旋转操作，返回旋转后新的根节点x
//        y                              x
//       / \                           /   \
//      x   T4     向右旋转 (y)        z     y
//     / \       - - - - - - - ->    / \   / \
//    z   T3                       T1  T2 T3 T4
//   / \
// T1   T2
func rotateRight(y *AVLTreeNode) *AVLTreeNode {
	x := y.Left
	T3 := x.Right

	// 旋转
	x.Right = y
	y.Left = T3

	//更新节点高度
	y.updateHeight()
	x.updateHeight()

	return x
}

// 左旋操作
// 对节点y进行向左旋转操作，返回旋转后新的根节点x
//    y                             x
//  /  \                          /   \
// T1   x      向左旋转 (y)       y     z
//     / \   - - - - - - - ->   / \   / \
//   T2  z                     T1 T2 T3 T4
//      / \
//     T3 T4
func rotateLeft(y *AVLTreeNode) *AVLTreeNode {
	x := y.Right
	T2 := x.Left

	// 执行旋转
	x.Left = y
	y.Right = T2

	// 更新节点的高度
	y.updateHeight()
	x.updateHeight()

	return x
}

// 插入节点到AVL树
func insert(root *AVLTreeNode, data int) *AVLTreeNode {
	// 1、执行标准的BST插入
	if root == nil {
		return NewNode(data)
	}
	if data < root.Data {
		root.Left = insert(root.Left, data)
	} else if data > root.Data {
		root.Right = insert(root.Right, data)
	} else {
		// 相同的数据不插入
		return root
	}

	// 2、更新节点的高度
	root.updateHeight()

	// 3、获取节点的平衡因子
	balance := root.balanceFactor()

	// 4、进行平衡调整
	// 左子树不平衡，右旋
	if balance > 1 && data < root.Left.Data {
		return rotateRight(root)
	}
	// 右子树不平衡，左旋
	if balance < -1 && data > root.Right.Data {
		return rotateLeft(root)
	}
	// 左右不平衡，先左旋再右旋
	if balance > 1 && data > root.Left.Data {
		root.Left = rotateLeft(root.Left)
		return rotateRight(root)
	}
	// 右左不平衡，先右旋再左旋
	if balance < -1 && data < root.Right.Data {
		root.Right = rotateRight(root.Right)
		return rotateLeft(root)
	}
	return root
}

// 插入数据到 AVL 树
func (t *AVLTree) Insert(data int) {
	t.root = insert(t.root, data)
}

// 中序遍历 AVL 树
func inorderTraversal(root *AVLTreeNode) {
	if root != nil {
		inorderTraversal(root.Left)
		fmt.Print(root.Data, " ")
		inorderTraversal(root.Right)
	}
}

// 中序遍历 AVL 树
func (t *AVLTree) InorderTraversal() {
	inorderTraversal(t.root)
	fmt.Println()
}

func main() {
	myAVLTree := NewAVLTree()

	// 插入数据
	dataList := []int{10, 20, 30, 40, 50, 25}
	for _, data := range dataList {
		myAVLTree.Insert(data)
	}

	// 中序遍历 AVL 树
	fmt.Println("Inorder Traversal:")
	myAVLTree.InorderTraversal()
}
