package main

import (
	"fmt"
	"io"
	"os"
	//"strconv"
	"strings"
	//"unicode"
)


const max_green int = 13
const max_red int = 12
const max_blue int = 14
const max_count int = 14

type Set struct {
	color string
	count int
}

func isGamePossible(set Set) bool {
	if set.count > 0 && set.count <= max_count{
		switch i := set.color; i {
		case "red":
			if set.count > max_red {
				return false
			}
		case "blue":
			if set.count > max_blue {
				return false
			}
		case "green":
			if set.count > max_green {
				return false
			}
		}
	} else if set.count > max_count {
		return false
	}

	return true
}

func main() {
	file, err := os.Open("Day-2-sample.txt")
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

	for _, line  := range lines {
		fmt.Println(line)
	}

}