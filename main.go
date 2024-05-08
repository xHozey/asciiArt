package main

import (
	"fmt"
	"os"
	"strings"

	ft "asciiart/features"
)

func main() {
	// Define the maximum allowed line length
	const MaxInputLength = 27

	// Specify the ASCII art banner file to use
	const bannerFile = "standard.txt"
	// const bannerFile = "shadow.txt"
	// const bannerFile = "thinkertoy.txt"

	if err := ft.CheckArguments(os.Args[1:]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	input := os.Args[1]

	// Check input lines do not exceed max length
	isWithinMaxLength, indicesOfExceedingLines := ft.CheckMaxInputLength(input, MaxInputLength)
	if !isWithinMaxLength {
		fmt.Printf("Error: Lines exceed the %d-character limit at line: %s.\n", MaxInputLength, indicesOfExceedingLines)
		os.Exit(1)
	}

	if !ft.CheckValidInput(input) {
		fmt.Println("Error: The input contains characters without corresponding ASCII art representation!")
		os.Exit(1)
	}

	data, err := os.ReadFile(bannerFile)
	if err != nil {
		fmt.Print(err)
		return
	}

	splittedInput := strings.Split(input, "\\n")
	hasNonEmptyLines := ft.CheckNewLine(splittedInput)

	content := strings.Split(string(data[1:]), "\n\n")
	characterMatrix := ft.ConvertToCharacterMatrix(content)

	ft.DrawASCIIArt(characterMatrix, splittedInput, hasNonEmptyLines)
}
