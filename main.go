package main

import(
	"fmt"
	"os"
	"io"
	"strings"
	"unicode"
	"math"
)

func combineRunes(a, b rune) int {
    // Convert runes to integers (assuming they represent digits)
    intA := int(a - '0')
    intB := int(b - '0')

    // Count the number of digits in intB
    digits := 1
    if intB > 9 {
        digits = int(math.Log10(float64(intB))) + 1
    }

    // Combine the integers
    return intA*int(math.Pow(10, float64(digits))) + intB
}

func main() {
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
		characters := []rune(line)
		start := -1
		end := -1

		for i, char := range characters {
			if unicode.IsDigit(char) {
				if start == -1 {
					start = i
				}
				end = i
			}
		}

		combined_digits := combineRunes(characters[start], characters[end])
		total += combined_digits

		if start != -1 {
			fmt.Printf("Character at start index (%d): %c\n", start, characters[start])
            fmt.Printf("Character at end index (%d): %c\n", end, characters[end])
			fmt.Printf("Joined first and last digit: %d\n", combined_digits)
			fmt.Printf("Current total: %d\n", total)
		}
	}

}