package asciiart

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// checks if the argument count is exactly one
func CheckArguments(args []string) error {
	argCount := len(args)

	if argCount == 0 {
		return errors.New("no arguments provided: you should include a string to be represented")
	}

	if argCount > 1 {
		return fmt.Errorf("too many arguments provided: expected 1, got %d", argCount)
	}

	return nil
}

// validates if the input contains only printable ASCII characters
func CheckValidInput(input string) bool {
	for _, char := range input {
		if int(char) < 32 || int(char) > 126 {
			return false
		}
	}
	return true
}

// verifies each line's length in the input string against the maximum length
func CheckMaxInputLength(input string, MaxInputLength int) (bool, string) {
	splitedInput := strings.Split(input, "\\n")
	lineIndices := ""
	for i, line := range splitedInput {
		if len(line) > MaxInputLength {
			lineIndices += strconv.Itoa(i + 1)
			if i == len(splitedInput)-1 {
				continue
			}
			lineIndices += " "
		}
	}
	if lineIndices != "" {
		return false, lineIndices
	}
	return true, lineIndices
}
