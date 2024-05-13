package memtable

import "fmt"

// Memtable is a simple in-memory implementation
// of a balanced binary search tree (BST) - AVL tree
type Memtable struct {
	root *Node
}

type orderType int

const (
	inorder orderType = iota
	preorder
	postOrder
)

func (mt *Memtable) Insert(node *Node) {
	mt.root = mt.root.Insert(node)
}

func (mt *Memtable) Search(key int) *Node {
	return mt.root.Search(key)
}

func (mt *Memtable) Delete(key int) {
	mt.root.Delete(key)
}

func leftRotate(n *Node) *Node {
	y := n.Right
	x := y.Left
	y.Left = n
	n.Right = x

	n.updateHeight()
	y.updateHeight()
	return y
}

func rightRotate(n *Node) *Node {
	y := n.Left
	x := y.Right
	y.Right = n
	n.Left = x
	n.updateHeight()
	y.updateHeight()
	return y
}

// Print prints the tree in the specified order
// if no order is specified, it defaults to inorder
// this is specifically for testing/debugging purposes
func (mt *Memtable) Print(ot orderType) {
	switch ot {
	case preorder:
		mt.root.printPreorder()
	case inorder:
		mt.root.printInorder()
	case postOrder:
		mt.root.printPostorder()

	default:
		// default to inorder
		mt.root.printInorder()
	}
}

type Node struct {
	Key    int
	Value  []byte
	Left   *Node
	Right  *Node
	Height int
}

func NewNode(key int, value []byte) *Node {
	return &Node{
		Key:   key,
		Value: value,
	}
}

func height(n *Node) int {
	if n == nil {
		return 0
	}
	return n.Height
}

func getBalanceFactor(n *Node) int {
	if n == nil {
		return 0
	}
	return height(n.Left) - height(n.Right)
}

func minValueNode(n *Node) *Node {
	current := n
	for current.Left != nil {
		current = current.Left
	}
	return current
}

func (n *Node) Insert(m *Node) *Node {
	if n == nil {
		return m
	}
	if m.Key < n.Key {
		n.Left = n.Left.Insert(m)
	} else if m.Key > n.Key {
		n.Right = n.Right.Insert(m)
	} else {
		n.Value = m.Value
	}
	n.updateHeight()

	return n.rebalance()
}

func (n *Node) rebalance() *Node {
	balanceFactor := getBalanceFactor(n)
	// LL
	if balanceFactor > 1 && getBalanceFactor(n.Left) >= 0 {
		return rightRotate(n)
	}

	// LR
	if balanceFactor > 1 && getBalanceFactor(n.Left) < 0 {
		n.Left = leftRotate(n.Left)
		return rightRotate(n)
	}

	// RR
	if balanceFactor < -1 && getBalanceFactor(n.Right) <= 0 {
		return leftRotate(n)
	}

	// RL
	if balanceFactor < -1 && getBalanceFactor(n.Right) > 0 {
		n.Right = rightRotate(n.Right)
		return leftRotate(n)
	}

	return n
}

func (n *Node) Search(key int) *Node {
	if n == nil {
		return nil
	}
	if key < n.Key {
		return n.Left.Search(key)
	}
	if key > n.Key {
		return n.Right.Search(key)
	}
	return n
}

func (n *Node) Delete(key int) *Node {
	if n == nil {
		return nil
	}

	if key < n.Key {
		n.Left = n.Left.Delete(key)
	} else if key > n.Key {
		n.Right = n.Right.Delete(key)
	} else if n.Left == nil || n.Right == nil {
		var temp *Node
		if n.Left == nil {
			temp = n.Right
		} else {
			temp = n.Left
		}
		if temp == nil {
			temp = n
			n = nil
		} else {
			n = temp
		}
	} else {
		temp := minValueNode(n.Right)
		n.Key = temp.Key
		n.Right = n.Right.Delete(temp.Key)
	}

	if n == nil {
		return n
	}
	n.updateHeight()
	return n.rebalance()
}

func (n *Node) updateHeight() {
	n.Height = max(height(n.Left), height(n.Right)) + 1
}

// printInorder: left, root, right
func (n *Node) printInorder() {
	if n == nil {
		return
	}
	if n.Left != nil {
		n.Left.printInorder()
	}
	fmt.Print(n.Key, " ")
	if n.Right != nil {
		n.Right.printInorder()
	}
}

// printPreorder: root, left, right
func (n *Node) printPreorder() {
	if n == nil {
		return
	}
	fmt.Print(n.Key, " ")
	if n.Left != nil {
		n.Left.printPreorder()
	}
	if n.Right != nil {
		n.Right.printPreorder()
	}
}

// printPostorder: left, right, root
func (n *Node) printPostorder() {
	if n == nil {
		return
	}
	if n.Left != nil {
		n.Left.printPostorder()
	}
	if n.Right != nil {
		n.Right.printPostorder()
	}
	fmt.Print(n.Key, " ")
}
