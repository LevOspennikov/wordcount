package counter

import (
	"awesomeProject/src/counter"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestDiffWords(t *testing.T) {
	arr := make([]string, 0)
	arr = append(arr, "a", "b", "c", "c ", "a", "bv")
	arr = append(arr, " a   ", "  b ", " c", "c ", "a", "bv")
	res := counter.DiffWordsCounter(arr, strings.TrimSpace)
	assert.Equal(t, res, 4)
	arr = append(arr, " a   ", "  b ", " c", "c ", "a", "bv")
	res = counter.DiffWordsCounter(arr, strings.TrimSpace)
	assert.Equal(t, res, 4)
}


func TestWords(t *testing.T) {
	arr := make([]string, 0)
	arr = append(arr, "a", "b", "c", "c ", "a", "b")
	res := counter.WordsCounter(arr, func(s string) string {
		if s == "b" {
			return ""
		}
		return s
	})
	assert.Equal(t, res, 4)
	res = counter.WordsCounter(arr)
	assert.Equal(t, res, 6)
}

func TestCountWords(t *testing.T) {
	arr := "one   two  free four "
	res := counter.CountWords(arr)
	assert.Equal(t, res, 4)
	arr = "one"
	res = counter.CountWords(arr)
	assert.Equal(t, res, 1)
	arr = ""
	res = counter.CountWords(arr)
	assert.Equal(t, res, 0)
}

func TestCountDiffWords(t *testing.T) {
	arr := "one   two  FREE four free OnE \n !one!"
	res := counter.CountDiffWords(arr)
	assert.Equal(t, res, 4)
	arr = "one"
	res = counter.CountDiffWords(arr)
	assert.Equal(t, res, 1)
	arr = ""
	res = counter.CountDiffWords(arr)
	assert.Equal(t, res, 0)
}



