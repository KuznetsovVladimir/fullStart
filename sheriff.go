package main

import (
	"fmt"
	"os"
)

func minValue(m map[string]int) int {
	min := 200000
	for _, value := range m {
		if value < min {
			min = value
		}
	}
	return min
}

func main() {
	var line string

	fmt.Fscan(os.Stdin, &line)

	dictionary := map[string]int{
		"s": 0,
		"h": 0,
		"e": 0,
		"r": 0,
		"i": 0,
		"f": 0,
	}

	for _, char := range line {
		_, ok := dictionary[string(char)]
		if ok {
			dictionary[string(char)]++
		}
	}

	result := minValue(dictionary)
	fCount := dictionary[`f`]
	if result <= fCount/2 {
		fmt.Println(result)
	} else {
		if fCount >= 2 {
			fmt.Println(fCount / 2)
		} else {
			fmt.Println(0)
		}
	}
}
