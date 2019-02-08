package avltree

import (
	"awesomeProject/src/avltree"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestNodeCreation(t *testing.T) {
	val := "val"
	node := avltree.Node{Value: val}
	assert.Equal(t, node.Value, val)
	assert.Nil(t, node.Left)
	assert.Nil(t, node.Right)
	assert.Equal(t, node.BalanceFactor, 0)
}

func TestNodeInsertion(t *testing.T) {
	arr := [6]int{3, 7, 2, 8, 4, 6}
	node := avltree.Node{Value: "5"}
	for _, val := range arr {
		node.Insert(strconv.Itoa(val))
	}
	assert.Equal(t, node.Value, "5")
	assert.Equal(t, node.Left.Left.Value, "2")
	assert.Equal(t, node.Right.Right.Value, "8")
}

func TestNodeFind(t *testing.T) {
	positive := [6]int{3, 7, 2, 8, 4, 6}
	negative := [6]int{1, 72, 20, 9, 44, -6}
	node := &avltree.Node{Value: "5"}
	for _, val := range positive {
		node.Insert(strconv.Itoa(val))
	}

	for _, val := range positive {
		result := node.Find(strconv.Itoa(val))
		assert.Equal(t, true, result, "positive", val)
	}

	for _, val := range negative {
		result := node.Find(strconv.Itoa(val))
		assert.Equal(t, false, result, "negative", val)
	}
}

func TestNodeRotateLeft(t *testing.T) {
	arr := [3]int{4, 3, 2}
	node := &avltree.Node{Value: "5"}
	for _, val := range arr {
		node.Insert(strconv.Itoa(val))
	}
	assert.Equal(t, node.Left.Left.Value, "2")
	assert.Nil(t, node.Right)
}

func TestNodeRotateRight(t *testing.T) {
	arr := [3]int{3, 4, 5}
	node := &avltree.Node{Value: "2"}
	for _, val := range arr {
		node.Insert(strconv.Itoa(val))
	}
	assert.Equal(t, node.Value, "2")
	assert.Equal(t, node.Right.Right.Value, "5")
	assert.Nil(t, node.Left)
}

func TestNodeBigRotateLeft(t *testing.T) {
	arr := [3]int{4, 2, 3}
	node := &avltree.Node{Value: "1"}
	for _, val := range arr {
		node.Insert(strconv.Itoa(val))
	}
	assert.Equal(t, node.Value, "1")
	assert.Equal(t, node.Right.Right.Value, "4")
	assert.Equal(t, node.Right.Value, "3")
	assert.Nil(t, node.Left)
}

func TestNodeBigRotateRight(t *testing.T) {
	arr := [3]int{2, 4, 3}
	node := &avltree.Node{Value: "5"}
	for _, val := range arr {
		node.Insert(strconv.Itoa(val))
	}
	assert.Equal(t, node.Value, "5")
	assert.Equal(t, node.Left.Value, "3")
	assert.Equal(t, node.Left.Left.Value, "2")
	assert.Nil(t, node.Right)
}
