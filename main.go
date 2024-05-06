package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Invalid args")
		return
	}
	data, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	file := strings.Split(string(data[1:]), "\n\n")
	input := os.Args[1]

	text := strings.Split(input, "\\n")
	fileres := map[rune][]string{}
	for i, val := range file {
		fileres[rune(i+32)] = strings.Split(val, "\n")
	}

	for _, val := range text {
		if val == "" {
			continue
		}
		for i := 0; i < 8; i++ {
			for _, val2 := range val {
				fmt.Print(fileres[val2][i])
			}
			fmt.Println()
		}
	}
}
