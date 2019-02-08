package utils

import "fmt"

func ReadFile(filename string, ReadFile func(string) ([]byte, error)) (*[]byte, error) {
	b, err := ReadFile(filename)
	if err != nil {
		fmt.Print(err)
	}
	return &b, err
}