package main

import (
	"bufio"
	"fmt"
	"os"
	// "regexp"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

const (
	MAX_RED = 12
	MAX_BLUE = 13
	MAX_GREEN = 14
)

func main() {
	
	f, err := os.Open("./input.txt")
	check(err)
	scanner := bufio.NewScanner(f)

	scanner.Split(bufio.ScanLines)

	impossibleGameSum := 0
	minGamesPowerSum :=0

	for scanner.Scan() {
		gameId, result := parseLine(scanner.Text())

		// fmt.Printf("Game id: %d. Results: %v\n", gameId, result)
		
		if (isImpossible(result)) {
			impossibleGameSum += gameId
		}

		minGame := getMinGame(result)
		minGamesPowerSum += minGame.Power()
	}

	fmt.Printf("The sum of ids of impossible games is: %d\n", impossibleGameSum)
	fmt.Printf("The sum of the minimum games powers is: %d\n", minGamesPowerSum)
}

func isImpossible(games []Game) bool {
	for _, game := range games {
		if (game.Blue > MAX_BLUE || game.Red > MAX_RED || game.Green > MAX_GREEN) {
			return true
		}
	}
	return false
}

type Game struct {
	Blue	int
	Red 	int
	Green	 int
}

func (g *Game) Power() int {
	return g.Red * g.Blue * g.Green 
}

func getMinGame(games []Game) Game {
	result := games[0]

	for _, game := range games {
		if game.Red > result.Red {
			result.Red = game.Red
		}
		if game.Blue > result.Blue {
			result.Blue = game.Blue
		}
		if game.Green > result.Green {
			result.Green = game.Green
		}
	}

	return result
}

func parseLine(line string) (gameId int, games []Game) {
	gameInfo := strings.Split(line, ":")
	id := strings.Split(gameInfo[0], " ")[1]
	// fmt.Printf("gameInfo: %v\n", gameInfo[1])
	
	rounds := strings.Split(gameInfo[1], ";")
	
	var result []Game
	for _, val:= range rounds {
		var game Game
		round := strings.Split(val, ",")
		for _, roundVal := range round {
			taken:= strings.Split(strings.TrimSpace(roundVal), " ")
			// fmt.Printf("taken: %v\n", taken)

			numTaken, err := strconv.Atoi(taken[0])
			check(err)
			switch(taken[1]) {
			case "red":
				game.Red = numTaken
				break
			case "blue":
				game.Blue = numTaken
				break
			case "green":
				game.Green = numTaken
				break
			default:
				break
			}
		}

		result = append(result, game)
	}
	

	gameId, err:= strconv.Atoi(id)
	check(err)

	return gameId, result

}