package avltree

import (
	"awesomeProject/src/utils"
	"strconv"
)

type Node struct {
	Value         string
	Left          *Node
	Right         *Node
	BalanceFactor int
}

func (n *Node) Find(s string) bool {
	if n == nil {
		return false
	}

	switch {
	case s == n.Value:
		return true
	case s < n.Value:
		return n.Left.Find(s)
	default: // case s > n.Value
		return n.Right.Find(s)
	}
}

func (n *Node) Insert(value string) bool {
	switch {
	case value == n.Value:
		return false
	case value < n.Value:
		if n.Left == nil {
			n.Left = &Node{Value: value}
			n.BalanceFactor = utils.Ternary(n.Right == nil, -1, 0)
			return true
		} else {
			if n.Left.Insert(value) {
				if utils.Abs(n.Left.BalanceFactor) > 1 {
					n.balance(n.Left)
				} else {
					n.BalanceFactor--
				}
				return true
			}
		}
	case value > n.Value:
		if n.Right == nil {
			n.Right = &Node{Value: value}
			n.BalanceFactor = utils.Ternary(n.Left == nil, 1, 0)
			return true
		} else {
			if n.Right.Insert(value) {
				if utils.Abs(n.Right.BalanceFactor) > 1 {
					n.balance(n.Right)
				} else {
					n.BalanceFactor++
				}
				return true
			}
		}
	}

	return false
}

func (n *Node) String() string {
	if n == nil {
		return ""
	}
	return "(" + n.Value + ":" + strconv.Itoa(n.BalanceFactor) + n.Left.String() + n.Right.String() + ")"
}

func (n *Node) rotateLeft(parent *Node) {
	b := parent.Right
	parent.Right = b.Left
	b.Left = parent
	if parent == n.Left {
		n.Left = b
	} else {
		n.Right = b
	}
	parent.BalanceFactor = 0
	b.BalanceFactor = 0
}

func (n *Node) rotateRight(parent *Node) {
	b := parent.Left
	parent.Left = b.Right
	b.Right = parent
	if parent == n.Left {
		n.Left = b
	} else {
		n.Right = b
	}
	parent.BalanceFactor = 0
	b.BalanceFactor = 0
}

func (n *Node) bigRotateLeft(parent *Node) {
	parent.Right.Left.BalanceFactor = 1
	parent.rotateRight(parent.Right)
	parent.Right.BalanceFactor = 1
	n.rotateLeft(parent)
}

func (n *Node) bigRotateRight(parent *Node) {
	parent.Left.Right.BalanceFactor = -1
	parent.rotateLeft(parent.Left)
	parent.Left.BalanceFactor = -1
	n.rotateRight(parent)
}

func (n *Node) balance(parent *Node) {
	switch {
	case parent.BalanceFactor == -2 && utils.Abs(parent.Left.BalanceFactor) == 1:
		if parent.Left.BalanceFactor == 1 {
			n.bigRotateRight(parent)
		} else {
			n.rotateRight(parent)
		}
	case parent.BalanceFactor == 2 && utils.Abs(parent.Right.BalanceFactor) == 1:
		if parent.Right.BalanceFactor == -1 {
			n.bigRotateLeft(parent)
		} else {
			n.rotateLeft(parent)
		}
	}
}
