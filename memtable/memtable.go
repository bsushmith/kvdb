package memtable

import "fmt"

type Memtable struct {
	root *Node
}

type orderType int

const (
	inorder orderType = iota
	preorder
	postOrder
)

func (mt *Memtable) Insert(key int) {
	mt.root = mt.root.Insert(key)
}

func (mt *Memtable) Search(key int) *Node {
	return mt.root.Search(key)
}

func (mt *Memtable) Delete(key int) {
	mt.root.Delete(key)
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
	Key   int
	Left  *Node
	Right *Node
}

func (n *Node) Insert(key int) *Node {
	panic("implement me")
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

func (n *Node) Delete(key int) {
	panic("implement me")
}

// printInorder prints the tree in inorder
// inorder: left, root, right
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

// printPreorder prints the tree in preorder
// preorder: root, left, right
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

// printPostorder prints the tree in postorder
// postorder: left, right, root
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
