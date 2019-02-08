package avltree

import (
	"awesomeProject/src/avltree"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"strconv"
	"testing"
)

func TestBtreeCreation(t *testing.T) {
	tree := avltree.Tree{}
	assert.Equal(t, tree.Size, 0)
}

func TestTreeInsertion(t *testing.T) {
	arr := [7]int{5, 3, 7, 2, 8, 4, 6}
	tree := avltree.Tree{}
	for _, val := range arr {
		tree.Put(strconv.Itoa(val))
	}
	assert.Equal(t, 7, tree.Size)
	for _, val := range arr {
		tree.Put(strconv.Itoa(val))
	}
	assert.Equal(t, 7, tree.Size)
}

func TestTreeFind(t *testing.T) {
	positive := [7]int{5, 3, 7, 2, 8, 4, 6}
	negative := [6]int{1, 72, 20, 9, 44, -6}
	tree := avltree.Tree{}
	for _, val := range positive {
		tree.Put(strconv.Itoa(val))
	}

	for _, val := range positive {
		result := tree.Has(strconv.Itoa(val))
		assert.Equal(t, true, result, "positive", val)
	}

	for _, val := range negative {
		result := tree.Has(strconv.Itoa(val))
		assert.Equal(t, false, result, "negative", val)
	}
}

func TestRandomInsertion(t *testing.T) {
	r := rand.New(rand.NewSource(31))
	tree := avltree.Tree{}
	numbers := make([]int, 0)
	size := 0

	for len(numbers) < 1000 {
		number := r.Int() % 128
		numbers = append(numbers, number)
		result := tree.Put(strconv.Itoa(number))
		if result {
			size++
		}
		hasNumber := tree.Has(strconv.Itoa(number))
		assert.Equal(t, true, hasNumber)
	}
	assert.Equal(t, size, tree.Size)
}

func TestFoldl(t *testing.T) {
	arr := [7]int{8, 7, 6, 5, 4, 3, 2}
	tree := avltree.Tree{}
	for _, val := range arr {
		tree.Put(strconv.Itoa(val))
	}
	tempString := ""
	f := func(n *avltree.Node) {
		tempString += n.Value
	}
	tree.Foldl(f)
	assert.Equal(t, "2345678", tempString)
}

func TestFoldr(t *testing.T) {
	arr := [7]int{2, 3, 4, 5, 6, 7, 8}
	tree := avltree.Tree{}
	for _, val := range arr {
		tree.Put(strconv.Itoa(val))
	}
	tempString := ""
	f := func(n *avltree.Node) {
		tempString += n.Value
	}
	tree.Foldr(f)
	assert.Equal(t, "8765432", tempString)
}

