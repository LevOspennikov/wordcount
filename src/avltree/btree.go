package avltree

import "awesomeProject/src/utils"

type Tree struct {
	head *Node
	Size int
}

func (t *Tree) Has(s string) bool {
	if t.Size == 0 {
		return false
	}
	return t.head.Find(s)
}

func (t *Tree) Put(s string) bool {
	if t.Size == 0 {
		t.head = &Node{Value: s}
		t.Size = 1
		return true
	}
	result := t.head.Insert(s)
	if result {
		if utils.Abs(t.head.BalanceFactor) > 1 {
			link := &Node{Right: t.head, Value: ""}
			link.balance(t.head)
			t.head = link.Right
		}

		t.Size = t.Size + 1
	}
	return result
}

func (t *Tree) Foldl(f func(*Node), n ...*Node) {
	targetNode := t.getTargetNode(n...)

	if targetNode.Left != nil {
		t.Foldl(f, targetNode.Left)
	}
	f(targetNode)
	if targetNode.Right != nil {
		t.Foldl(f, targetNode.Right)
	}
}

func (t *Tree) Foldr(f func(*Node), n ...*Node) {
	targetNode := t.getTargetNode(n...)

	if targetNode.Right != nil {
		t.Foldr(f, targetNode.Right)
	}
	f(targetNode)
	if targetNode.Left != nil {
		t.Foldr(f, targetNode.Left)
	}
}

func (t *Tree) getTargetNode(n ...*Node) (targetNode *Node) {
	if n == nil {
		targetNode = t.head
	} else {
		targetNode = n[0]
	}
	return
}

func (t *Tree) String() string {
	return t.head.String()
}