package main

import (
	"awesomeProject/src/counter"
	"awesomeProject/src/utils"
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
)

func main() {
	line := flag.Bool("l", false, "line count")
	words := flag.Bool("w", false, "words count")
	bytes := flag.Bool("b", false, "byte count")
	diffWords := flag.Bool("dw", false, "different words count")
	flag.Parse()
	fmt.Println(flag.Args())
	filename := flag.Args()[0] // first element as file name
	b, err := utils.ReadFile(filename, ioutil.ReadFile)
	if err != nil {
		fmt.Print(err)
		return
	}
	str := string(*b)
	result := ""
	switch {
	case *line:
		result += strconv.Itoa(counter.CountLines(str)) + " "
		fallthrough
	case *words:
		result += strconv.Itoa(counter.CountWords(str)) + " "
		fallthrough
	case *bytes:
		result += strconv.Itoa(counter.CountBytes(*b)) + " "
		fallthrough
	case *diffWords:
		result += strconv.Itoa(counter.CountDiffWords(str)) + " "
	}

	fmt.Println(result)
}
