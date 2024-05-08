# ASCII Art Project

## Description

The ASCII Art project is a Go program that converts a given input string into a graphic representation using ASCII characters. It supports handling inputs with numbers, letters, spaces, special characters, and newline characters (`\n`). The program uses different banner files to format the input string into ASCII art, showcasing your ability to manipulate data and use the Go file system API.

## Project Objectives

- Learn how to use Go's file system (fs) API.
- Practice data manipulation in Go.
- Implement ASCII art rendering for different characters using provided banner files.

## Banner Files

The project uses banner files to provide the graphical template representations of ASCII art characters. Each character is represented with a height of 8 lines, and the characters are separated by a newline (`\n`). The banners provided include:

- `shadow`: A specific graphical template representation using ASCII.
- `standard`: A specific graphical template representation using ASCII.
- `thinkertoy`: A specific graphical template representation using ASCII.

## Usage

To run the program, use the following command:

```
go run . "<input_string>"
```

For example:

```
go run . "Hello\nThere" | cat -e
```

The input string can include numbers, letters, spaces, special characters, and newline characters. The program will output the ASCII art representation of the input string. 

### Example Outputs

```
go run . "Hello There" | cat -e

 _    _          _   _               _______   _                           $
| |  | |        | | | |             |__   __| | |                          $
| |__| |   ___  | | | |   ___          | |    | |__     ___   _ __    ___  $
|  __  |  / _ \ | | | |  / _ \         | |    |  _ \   / _ \ | '__|  / _ \ $
| |  | | |  __/ | | | | | (_) |        | |    | | | | |  __/ | |    |  __/ $
|_|  |_|  \___| |_| |_|  \___/         |_|    |_| |_|  \___| |_|     \___| $
                                                                           $
                                                                           $
```

## Project Structure

The project consists of the following files:

- `main.go`: The main program file that handles input and ASCII art conversion.
- `banner files`: The banner files contain ASCII art representations of characters.
- `test files`: Test files for unit testing to ensure the program functions correctly.

## Good Practices

- Follow Go's standard coding style and best practices.
- Write clear and concise code comments.
- Implement comprehensive unit tests to verify the program's correctness.

## Getting Started

1. Clone the repository to your local machine.
2. Ensure you have Go installed.
3. Navigate to the project directory.
4. Run the program using the command `go run . "<input_string>"`.
5. Review the output ASCII art representation.

