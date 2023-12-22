package main

import (
	"fmt"
	"io"
	"os"
	"strconv"

	//"strconv"
	"strings"
	//"unicode"
)


const max_green int = 13
const max_red int = 12
const max_blue int = 14
const max_count int = 14

type Subset struct {
	color string
	count int
}

type Set struct {
	subsets []Subset
}

type Game struct {
	game_id int
	sets []Set
}


func isSubsetPossible(subset Subset) bool {
	if subset.count > 0 && subset.count <= max_count{
		switch i := subset.color; i {
		case "red":
			if subset.count > max_red {
				return false
			}
		case "blue":
			if subset.count > max_blue {
				return false
			}
		case "green":
			if subset.count > max_green {
				return false
			}
		}
	} else if subset.count > max_count {
		return false
	}

	return true
}

func isSetPossible(set Set) bool {
	subsets := set.subsets
	for _, subset := range subsets {
		if !isSubsetPossible(subset) {
			return false
		}
	}

	return true
}

func isGamePossible(game Game) (bool, int) {
	sets := game.sets
	for _, set := range sets {
		if !isSetPossible(set) {
			return false, 0
		}
	}

	return true, game.game_id
}

func parseGameID(line string) (int, error) {
	end := -1
	start := 5
	runes := []rune(line)
	for i, r := range runes {
		if r == ':' {
			end = i
		}
	}

	game_id_string := line[start:end]
	game_id, err := strconv.Atoi(game_id_string)
	if err != nil {
		return 0, fmt.Errorf("error converting game_id_string: %s into an integer", game_id_string)
	}

	return game_id, nil
}

func parseGameSets(line string) []Set {
	sets := []Subset{}
	colon := 0
	runes := []rune(line)
	for i, rune := range runes {
		if rune == ':' {
			colon = i
		}
		games_no_id := line[colon + 1:len(line)-1]
		
	}

	return sets
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
		game_id, err := parseGameID(line) 
		if err != nil {
			fmt.Printf("error parsing game ID: %v", err)
		}
		fmt.Println(game_id)
	}

}