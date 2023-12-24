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

/*
const max_green int = 13
const max_red int = 12
const max_blue int = 14
const max_count int = 14
*/
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

func parseSubset(subsetStr string) (Subset, error) {
	parts := strings.Fields(subsetStr)
	if len(parts) != 2 {
		return Subset{}, fmt.Errorf("invalid format\n")
	}

	count, err := strconv.Atoi(parts[0])
	if err != nil {
		return Subset{}, fmt.Errorf("invalid count")
	}

	return Subset{color: parts[1], count: count}, nil
}

func parseSets(gameStr string) ([]Set, error) {
    sets := []Set{}

    // Find the position of the colon and slice the string from that point
    colonPos := strings.Index(gameStr, ":")
    if colonPos == -1 {
        return nil, fmt.Errorf("invalid game string format")
    }

    // Split the string into sets at each semicolon
    setStrs := strings.Split(gameStr[colonPos+1:], ";")

    for _, setStr := range setStrs {
        currentSet := Set{}

        // Split each set into subsets
        subsetStrs := strings.Split(setStr, ",")
        for _, subsetStr := range subsetStrs {
            subsetStr = strings.TrimSpace(subsetStr)
            if subsetStr != "" {
                subset, err := parseSubset(subsetStr)
                if err != nil {
                    return nil, fmt.Errorf("error parsing '%s': %v", subsetStr, err)
                }
                currentSet.subsets = append(currentSet.subsets, subset)
            }
        }

        sets = append(sets, currentSet)
    }

    return sets, nil
}

func parseGame(gameStr string) Game {
	game_id, err := parseGameID(gameStr)
	if err != nil {
		fmt.Println("Error parsing game id: ", err)
	}
	sets , err := parseSets(gameStr) 
	if err != nil {
		fmt.Println("Error parsing game sets: ", err)
	}

	game := Game{game_id: game_id, sets: sets}
	return game
}

/*
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

func isGamePossible(game Game) bool {
	sets := game.sets
	for _, set := range sets {
		if !isSetPossible(set) {
			return false
		}
	}

	return true
}
*/

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


func getMaxColorCounts(game Game) (int, int, int) {
	max_blue := 0
	max_green := 0
	max_red := 0

	sets := game.sets
	for _, set := range sets {
		subsets := set.subsets
		for _, subset := range subsets {
			switch color := subset.color; color {
			case "blue":
				if subset.count > max_blue {
					max_blue = subset.count
				}
			case "green":
				if subset.count > max_green {
					max_green = subset.count
				}
			case "red":
				if subset.count > max_red {
					max_red = subset.count
				}
			}
		}
	}

	return max_blue, max_green, max_red
}

func main() {
	file, err := os.Open("Day-2-puzzle.txt")
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
	//games := []Game{}

	//total_games_possible := 0
	total_set_power := 0

	for _, line  := range lines {
		fmt.Println(line)
		game := parseGame(line)
		fmt.Printf("Game %d: %+v\n", game.game_id, game.sets)
		//games = append(games, game)

		max_blue, max_green, max_red := getMaxColorCounts(game)
		fmt.Printf("Max Blue: %d, Max Green: %d, Max Red: %d\n", max_blue, max_green, max_red)

		set_power := max_blue * max_green * max_red
		fmt.Printf("Set Power: %d\n", set_power)

		total_set_power += set_power
		fmt.Printf("Total Set Power: %d\n", total_set_power)

		/*
		if isGamePossible(game) {
			total_games_possible += game.game_id
		}

		fmt.Println(total_games_possible)

		
		game_id, err := parseGameID(line) 
		if err != nil {
			fmt.Printf("error parsing game ID: %v", err)
		}
		fmt.Println(game_id)
		sets, err := parseSets(line) 
		if err != nil {
			fmt.Printf("error parsing subsets: %v", err)
		}
		for i, set := range sets {
			fmt.Printf("Set %d: %+v\n", i+1, set)
		}
		*/
	}

}