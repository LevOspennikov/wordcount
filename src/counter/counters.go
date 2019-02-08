package counter

import "strings"

func CountWords(str string) int {
	arr := strings.Fields(str)
	return WordsCounter(arr)
}

func CountDiffWords(str string) int {
	arr := strings.Fields(str)
	return DiffWordsCounter(arr, strings.ToLower,
		strings.TrimSpace,
		TransformerWords)
}

func CountLines(str string) int {
	// windows compatibility
	arr := strings.Split(strings.Replace(str, "\r\n", "\n", -1), "\n")
	return WordsCounter(arr)
}

func CountBytes(bytes []byte) int {
	return len(bytes)
}
