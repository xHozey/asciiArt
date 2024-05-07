package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Print(err)
		return
	}
	standard := strings.Split(string(data), "\n\n")
	input := os.Args[1]
	splitInput := strings.Split(input, "\\n")
	var match bool

	for _, i := range splitInput {
		if i != "" {
			match = true
			break
		}
	}
	// "" "k" "" ""
	matrix := map[rune][]string{}
	for i, val := range standard {
		matrix[rune(32+i)] = strings.Split(val, "\n")
	}
	for i, val := range splitInput {
		if val == "" {
			if match {
				fmt.Println()
			} else if i != 0 && !match {
				fmt.Println()
			}

		} else if val != "" {
			for j := 0; j < 8; j++ {
				for _, k := range val {
					fmt.Print(matrix[k][j])
				}
				fmt.Println()
			}
		}
	}

}
