package counter

import (
	"awesomeProject/src/avltree"
)

func DiffWordsCounter(arr []string, f ...func(string) string) int {
	tree := avltree.Tree{}

	for _, val := range arr {
		tempStr := val
		for _, transformer := range f {
			tempStr = transformer(tempStr)
		}
		if tempStr != "" {
			tree.Put(tempStr)
		}
	}
	return tree.Size
}