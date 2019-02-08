package counter

func WordsCounter(arr []string, f ...func(string) string) int {
	size := 0
	for _, val := range arr {
		tempStr := val
		for _, transformer := range f {
			tempStr = transformer(val)
		}
		if tempStr != "" {
			size++
		}
	}

	return size
}