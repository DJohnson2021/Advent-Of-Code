package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"
)

//var result []int
//var currentWord strings.Builder

func findFirstValue(input string, numMap map[string]rune) rune {
	result := '0'

	for i, char := range input {

		//fmt.Println( unicode.IsLetter(char))
		if unicode.IsLetter(char) {

			endIndex := i + 1
			startIndex := 0
			for word := range numMap {
				new_slice := ""
				startIndex = endIndex - len(word)
				if startIndex < 0 {
					new_slice = input[0:endIndex]
				} else {
					new_slice = input[startIndex:endIndex]
				}
				//fmt.Println(new_slice)
				if value, exists := numMap[new_slice]; exists {
					result = value
					return result
				}
			}

		} else if unicode.IsDigit(char) {
			result = char
			return result
		}
	}

	return result
}

func findLastValue(input string, numMap map[string]rune) rune {
	result := '0'
	runes := []rune(input)

	for i := len(runes) - 1; i >= 0; i-- {

		//fmt.Println( unicode.IsLetter(char))
		if unicode.IsLetter(runes[i]) {

			endIndex := i + 1
			startIndex := 0
			for word := range numMap {
				new_slice := ""
				startIndex = endIndex - len(word)
				if startIndex < 0 {
					new_slice = input[0:endIndex]
				} else {
					new_slice = input[startIndex:endIndex]
				}
				//fmt.Println(new_slice)
				if value, exists := numMap[new_slice]; exists {
					result = value
					return result
				}
			}

		} else if unicode.IsDigit(runes[i]) {
			result = runes[i]
			return result
		}
	}

	return result
}

func main() {
	// Create a map with string keys and int values
	numMap := make(map[string]rune)

	numMap["one"] = '1'
	numMap["two"] = '2'
	numMap["three"] = '3'
	numMap["four"] = '4'
	numMap["five"] = '5'
	numMap["six"] = '6'
	numMap["seven"] = '7'
	numMap["eight"] = '8'
	numMap["nine"] = '9'
	numMap["zero"] = '0'

	file, err := os.Open("Day-1-puzzle.txt")
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}

	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file: ", err)
	}

	lines := strings.Split(string(content), "\n")

	total := 0

	for _, line := range lines {
		fmt.Println("Line: ", line)

		start := findFirstValue(line, numMap)
		end := findLastValue(line, numMap)
		//fmt.Println(start)

		if start > 0 && end > 0 {
			//nums_length := len(numbers)
			//start := numbers[0]
			//end := numbers[nums_length - 1]

			left_digit := string(start)
			right_digit := string(end)
			combined_digits := left_digit + right_digit
			line_number, err := strconv.Atoi(combined_digits)
			if err != nil {
				fmt.Printf("Error converting combined digits %s into an integer", combined_digits)
			}
			total += line_number

			fmt.Printf("First Digit %d: \n", start)
			fmt.Printf("Last Digit %d: \n", end)
			fmt.Printf("Joined first and last digit: %s\n", combined_digits)
		} else {
			fmt.Println("No numbers found in this line.")
		}

		fmt.Printf("Current total: %d\n", total)
	}

}

/*
func subStringToIntSlice(input string, numMap map[string]int) []int {
	result := []int{}
	currentWord := strings.Builder{}
	for _, char := range input {
		if unicode.IsLetter(char) {

			currentWord.WriteRune(char)
			word := currentWord.String()


			if value, exists := numMap[word]; exists {
				result = append(result, value)
				currentWord.Reset()
			}

		} else if unicode.IsDigit(char) {
			result = append(result, int(char - '0'))
			currentWord.Reset()

		} else if currentWord.Len() > 0 {
            exception := currentWord.String()
            fmt.Println("Current work if len greater than 0: ", exception)
			currentWord.Reset()
		}
	}

	return result
}
*/

/*
func subStringToIntSlice(input string, numMap map[string]int) []int {
    result := []int{}
    currentWord := strings.Builder{}

    for i := 0; i < len(input); i++ {
        char := rune(input[i])
        if unicode.IsLetter(char) {
            currentWord.WriteRune(char)
            for j := i + 1; j <= len(input); j++ {
                word := currentWord.String()
                if value, exists := numMap[word]; exists {
                    result = append(result, value)
                    i = j - 1 // Move the outer loop's index past the current word
                    currentWord.Reset()
                    break
                }
                if j < len(input) && unicode.IsLetter(rune(input[j])) {
                    currentWord.WriteRune(rune(input[j]))
                } else {
                    break
                }
            }
            currentWord.Reset()
        } else if unicode.IsDigit(char) {
            result = append(result, int(char-'0'))
        }
    }

    return result
}
*/
