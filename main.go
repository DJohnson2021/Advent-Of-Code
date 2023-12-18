package main

import(
	"fmt"
	"os"
	"io"
	"strings"
	"unicode"
	"math"
)

//var result []int
//var currentWord strings.Builder

func combineIntegers(a int, b int) int {
    // Count the number of digits in b
    digits := 0
    if b > 0 {
        digits = int(math.Log10(float64(b))) + 1
    } else if b == 0 {
        digits = 1
    } else {
        // Handle negative b if necessary
    }

    // Shift a to the left by the number of digits in b and add b
    return a*int(math.Pow(10, float64(digits))) + b
}


func subStringToIntSlice(input string, numMap map[string]int) []int {
    result := []int{}
    currentWord := strings.Builder{}

    for _, char := range input {
        if unicode.IsLetter(char) {
            currentWord.WriteRune(char)
            continue
        }

        word := currentWord.String()
        if value, exists := numMap[word]; exists {
            result = append(result, value)
        }
        currentWord.Reset()

        if unicode.IsDigit(char) {
            result = append(result, int(char-'0'))
        }
    }

    // Check if the last word in the string is a number word
    word := currentWord.String()
    if value, exists := numMap[word]; exists {
        result = append(result, value)
    }

    return result
}




func main() {
	// Create a map with string keys and int values
	numMap := make(map[string]int)

	numMap["one"] = 1
	numMap["two"] = 2
	numMap["three"] = 3
	numMap["four"] = 4
	numMap["five"] = 5
	numMap["six"] = 6
	numMap["seven"] = 7
	numMap["eight"] = 8
	numMap["nine"] = 9
	numMap["zero"] = 0

	file, err := os.Open("Day-1-Part-2-sample.txt")
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

        numbers := subStringToIntSlice(line, numMap)
        fmt.Println(numbers)
        
        if len(numbers) > 0 {
            nums_length := len(numbers)
            start := numbers[0]
            end := numbers[nums_length - 1]

            combined_digits := combineIntegers(start, end)
            total += combined_digits

            fmt.Printf("First Digit %d: \n", start)
            fmt.Printf("Last Digit %d: \n", end)
            fmt.Printf("Joined first and last digit: %d\n", combined_digits)
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
			currentWord.Reset()
		}
	}

	return result
}
*/