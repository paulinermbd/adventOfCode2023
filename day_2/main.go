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
	mapIndex, err := prepareData(data)
	if err != nil {
		panic(err)
	}
	gameTotal, err := getGameSet(mapIndex)
	if err != nil {
		return
	}
	fmt.Printf("gameTotal: %d\n", gameTotal)
}

func prepareData(data string) (map[int]string, error) {
	mapIndexSet := make(map[int]string)
	sets := strings.Split(data, "\n")

	for i, set := range sets {
		prefix := fmt.Sprintf("Game %d:", i+1)
		str, _ := strings.CutPrefix(set, prefix)
		mapIndexSet[i+1] = str
		i++
	}

	return mapIndexSet, nil
}

func getGameSet(mapIndex map[int]string) (int, error) {
	sumGame := 0
	for i, set := range mapIndex {
		isValid := true
		gameSet := strings.Split(set, ";")
		fmt.Printf("gameSet %d: %+v\n", i, gameSet)

		for _, game := range gameSet {
			subBlue := regexp.MustCompile(`(\d+) blue`).FindStringSubmatch(game)
			if len(subBlue) > 0 {
				fmt.Printf("sub: %+v\n", subBlue[1])
				valueInt, _ := strconv.Atoi(subBlue[1])
				if valueInt > 14 {
					isValid = false
					break
				}
			}

			subGreen := regexp.MustCompile(`(\d+) green`).FindStringSubmatch(game)
			if len(subGreen) > 0 {
				fmt.Printf("sub: %+v\n", subGreen[1])
				valueInt, _ := strconv.Atoi(subGreen[1])
				if valueInt > 13 {
					isValid = false
					break
				}
			}

			subRed := regexp.MustCompile(`(\d+) red`).FindStringSubmatch(game)
			if len(subRed) > 0 {
				fmt.Printf("sub: %+v\n", subRed[1])
				valueInt, _ := strconv.Atoi(subRed[1])
				if valueInt > 12 {
					isValid = false
					break
				}
			}
		}

		if isValid {
			sumGame += i
		}
		i++
	}

	return sumGame, nil
}

func getData(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return string(data), nil
}
