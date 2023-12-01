package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data, err := getData("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(data, "\n")
	globalSum, err := sum(lines)
	fmt.Printf("Total: %d\n", globalSum)
}

func sum(lines []string) (int, error) {
	var sumGlobal int
	for _, line := range lines {
		l := regexp.MustCompile(`\d`).FindAllString(line, -1)
		line := l[0] + l[len(l)-1]
		//fmt.Printf("Line: %s\n", line)
		lineIntValue, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		sumGlobal += lineIntValue
	}
	return sumGlobal, nil
}

func getData(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return string(data), nil
}
