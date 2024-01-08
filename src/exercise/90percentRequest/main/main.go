package main

import (
	"container/list"
	"fmt"
	"time"
)

/**
当我们需要找到近5分钟内90%的请求时间时，我们可以利用平衡树和双向链表的结合使用，通过动态地维护这两个数据结构来实现高效的查找。

数据结构的设计
双向链表 (requests): 用于按照请求的处理时间顺序存储请求信息。每次插入一个新的请求信息时，我们将其插入到链表的合适位置，并且保持链表有序。链表头部是最早的请求，尾部是最近的请求。

平衡树 (sortedRequest): 用于辅助排序，每个节点存储了双向链表中的一个节点，按照请求的处理时间排序。这样，我们可以通过平衡树快速找到中位数位置的节点。

插入新节点
每次插入一个新的请求信息时，我们需要做以下操作：

插入到双向链表 (requests): 将请求信息按照处理时间的顺序插入到链表的合适位置。

插入到平衡树 (sortedRequest): 将请求信息插入到平衡树中，保持平衡树的有序性。

维护近5分钟内的数据结构
同时，我们需要维护这两个数据结构中的数据，以保证它们始终包含近5分钟内的请求信息。

双向链表 (requests): 在插入节点后，检查链表的头部节点是否超过5分钟前，如果超过，则删除头部节点，保证链表中只包含近5分钟内的请求。

平衡树 (sortedRequest): 平衡树中的节点数量和双向链表中的节点数量保持一致，以便在查找90%位置时能够确定节点的具体位置。

查找90%位置的请求时间
在需要获取近5分钟内90%的请求时间时，我们利用平衡树的性质进行查找。具体步骤如下：

计算90%的位置，即链表长度的90%位置。

从平衡树的根节点开始，通过比较当前节点的左子树节点数目和当前位置，确定下一步搜索的方向。

不断地向左或向右移动，直到找到90%位置的节点。

返回该节点的请求时间作为近5分钟内90%的请求时间。

代码中的关键实现
在代码中，GetPercentileDuration 方法实现了查找90%位置的请求时间。通过平衡树的辅助，我们可以在O(logN)的时间内找到90%位置的节点，而不需要完全遍历链表。
*/

// RequestTimeMonitor 模块
type RequestTimeMonitor struct {
	requests      *list.List // 双向链表 按照请求的处理时间顺序存储请求信息
	sortedRequest *AVLTree   // 平衡二叉树 按照请求的处理时间排序
	//totalDuration time.Duration
}

// RequestInfo 存储每个请求的时间信息
type RequestInfo struct {
	timestamp time.Time     // 请求到达的时间
	duration  time.Duration // 请求的处理时长
}

// Node 平衡二叉树节点
type Node struct {
	requestInfo RequestInfo // 请求信息
	height      int         // 树的高度
	size        int         // 子树的节点数目 为了能够快速定位90%的位置
	left        *Node       // 左子树
	right       *Node       // 右子树
}

// AVLTree 平衡二叉树 存储根结点
type AVLTree struct {
	root *Node
}

// NewRequestTimeMonitor 创建一个新的请求时间监控模块
func NewRequestTimeMonitor() *RequestTimeMonitor {
	return &RequestTimeMonitor{
		requests:      list.New(),
		sortedRequest: &AVLTree{root: nil},
	}
}

// RecordRequest 记录一次请求的时间信息
func (m *RequestTimeMonitor) RecordRequest(duration time.Duration) {
	requestInfo := RequestInfo{
		timestamp: time.Now(),
		duration:  duration,
	}

	// 将请求信息插入双向链表按时间顺序存储
	m.insertSorted(requestInfo)

	// 更新平衡树
	m.sortedRequest.Insert(requestInfo)

	// 更新总的请求时间
	//m.totalDuration += duration

	// 移除超过5分钟的请求信息
	fiveMinutesAgo := time.Now().Add(-5 * time.Minute)
	for m.requests.Len() > 0 {
		oldest := m.requests.Front().Value.(RequestInfo)
		if oldest.timestamp.Before(fiveMinutesAgo) {
			m.requests.Remove(m.requests.Front())
		} else {
			break
		}
	}
}

// insertSorted 将请求信息按时间顺序插入双向链表
func (m *RequestTimeMonitor) insertSorted(requestInfo RequestInfo) {
	element := m.requests.Front()
	for element != nil {
		info := element.Value.(RequestInfo)
		if requestInfo.timestamp.After(info.timestamp) {
			element = element.Next()
		} else {
			break
		}
	}

	if element == nil {
		m.requests.PushBack(requestInfo)
	} else {
		m.requests.InsertBefore(requestInfo, element)
	}
}

// GetPercentileDuration 获取近5分钟内90%的请求时间小于多少
func (m *RequestTimeMonitor) GetPercentileDuration() time.Duration {
	ninetyPercentIndex := int(0.9 * float64(m.requests.Len()))
	if ninetyPercentIndex >= m.requests.Len() {
		ninetyPercentIndex = m.requests.Len() - 1
	}

	// 从平衡树中获取近5分钟内90%的请求时间
	var i int
	root := m.sortedRequest.root
	if root != nil {
		for root != nil {
			leftSize := m.sortedRequest.getSize(root.left)
			if i == ninetyPercentIndex {
				kthNode := m.sortedRequest.kthSmallest(root, i-leftSize)
				if kthNode != nil {
					return kthNode.requestInfo.duration
				} else {
					fmt.Println("kthNode is nil")
				}
				//return m.sortedRequest.kthSmallest(root, i-leftSize).requestInfo.duration
			} else if i < ninetyPercentIndex {
				root = root.right
				i += leftSize + 1
			} else {
				root = root.left
			}
		}
	}
	return 0
}

// kthSmallest 获取平衡树中第 k 小的节点
func (t *AVLTree) kthSmallest(root *Node, k int) *Node {
	if root == nil {
		return nil
	}

	leftSize := t.getSize(root.left)

	if k == leftSize {
		return root
	} else if k < leftSize {
		return t.kthSmallest(root.left, k)
	} else {
		// 修复当右子树为空时直接返回当前节点
		if root.right == nil {
			return root
		}
		return t.kthSmallest(root.right, k-leftSize-1)
	}
}

// getSize 获取节点的子树节点数目
//func (t *AVLTree) getSize(node *Node) int {
//	if node == nil {
//		return 0
//	}
//	return node.size
//}
// getSize 返回节点的子树大小
func (t *AVLTree) getSize(node *Node) int {
	if node == nil {
		return 0
	}
	return 1 + t.getSize(node.left) + t.getSize(node.right)
}

// updateSize 更新节点的子树节点数目
func (t *AVLTree) updateSize(node *Node) {
	if node != nil {
		node.size = 1 + t.getSize(node.left) + t.getSize(node.right)
	}
}

// Insert 插入节点到平衡树
func (t *AVLTree) Insert(requestInfo RequestInfo) {
	t.root = t.insert(t.root, requestInfo)
}

// Print 打印 AVL 树的结构
func (t *AVLTree) Print(root *Node, prefix string) {
	if root != nil {
		fmt.Printf("%s%s - %s\n", prefix, root.requestInfo.timestamp, root.requestInfo.duration)
		t.Print(root.left, prefix+"L---")
		t.Print(root.right, prefix+"R---")
	}
}

// insert 将请求信息插入平衡树中
func (t *AVLTree) insert(root *Node, request RequestInfo) *Node {
	if root == nil {
		return &Node{
			requestInfo: request,
			height:      1,
			size:        1,
		}
	}

	if request.timestamp.Before(root.requestInfo.timestamp) {
		root.left = t.insert(root.left, request)
	} else {
		root.right = t.insert(root.right, request)
	}

	// 更新节点数
	root.size = 1 + t.getSize(root.left) + t.getSize(root.right)

	// 更新节点高度
	root.height = 1 + max(t.getHeight(root.left), t.getHeight(root.right))

	// 重新计算平衡因子
	balance := t.getBalance(root)

	if balance > 1 {
		if request.timestamp.Before(root.left.requestInfo.timestamp) {
			return t.rotateRight(root)
		} else {
			root.left = t.rotateLeft(root.left)
			return t.rotateRight(root)
		}
	}

	if balance < -1 {
		if request.timestamp.After(root.right.requestInfo.timestamp) {
			return t.rotateLeft(root)
		} else {
			root.right = t.rotateRight(root.right)
			return t.rotateLeft(root)
		}
	}

	return root
}

// rotateLeft 左旋
func (t *AVLTree) rotateLeft(y *Node) *Node {
	if y == nil || y.right == nil {
		return y
	}

	x := y.right
	T2 := x.left

	x.left = y
	y.right = T2

	// 更新节点的高度
	y.height = 1 + max(t.getHeight(y.left), t.getHeight(y.right))
	x.height = 1 + max(t.getHeight(x.left), t.getHeight(x.right))

	// 更新节点的子树节点数目
	t.updateSize(y)
	t.updateSize(x)

	return x
}

// rotateRight 右旋
func (t *AVLTree) rotateRight(x *Node) *Node {
	if x == nil || x.left == nil {
		return x
	}

	y := x.left
	T2 := y.right

	y.right = x
	x.left = T2

	// 更新节点的高度
	x.height = 1 + max(t.getHeight(x.left), t.getHeight(x.right))
	y.height = 1 + max(t.getHeight(y.left), t.getHeight(y.right))

	// 更新节点的子树节点数目
	t.updateSize(x)
	t.updateSize(y)

	return y
}

// getHeight 获取节点的高度
func (t *AVLTree) getHeight(node *Node) int {
	if node == nil {
		return 0
	}
	return node.height
}

// getBalance 获取节点的平衡因子
func (t *AVLTree) getBalance(node *Node) int {
	if node == nil {
		return 0
	}
	return t.getHeight(node.left) - t.getHeight(node.right)
}

// max 返回两个数中的较大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	monitor := NewRequestTimeMonitor()

	// 模拟前五分钟的请求数据
	for i := 0; i < 300; i++ {
		duration := time.Duration(i) * time.Millisecond
		monitor.RecordRequest(duration)
	}

	fmt.Println("前五分钟模拟请求数据已记录。")

	// 模拟记录请求时间
	for i := 300; i < 1000000; i++ {
		duration := time.Duration(i) * time.Millisecond
		monitor.RecordRequest(duration)

		// 输出当前近5分钟内90%的请求时间
		percentileDuration := monitor.GetPercentileDuration()
		fmt.Printf("记录第 %d 次请求，当前近5分钟内90%%的请求时间小于：%v\n", i, percentileDuration)

		time.Sleep(10 * time.Millisecond) // 模拟请求之间的时间间隔
	}
}
