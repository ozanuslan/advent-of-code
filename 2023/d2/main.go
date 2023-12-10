package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	lines := readInput()
	p1 := part1(lines)
	p2 := part2(lines)
	println("Part 1: ", p1)
	println("Part 2: ", p2)
}

type Game struct {
	id     int
	rounds []Round
}

type Round struct {
	red   int
	green int
	blue  int
}

func part1(lines []string) int {
	games := parseGames(lines)

	availableReds := 12
	availableGreens := 13
	availableBlues := 14

	idSum := 0
	for _, game := range games {
		isPossible := true
		for _, round := range game.rounds {
			if round.red > availableReds || round.green > availableGreens || round.blue > availableBlues {
				isPossible = false
				break
			}
		}
		if isPossible {
			idSum += game.id
		}
	}
	return idSum
}

func part2(lines []string) int {
	games := parseGames(lines)

	powerSum := 0
	for _, game := range games {
		minReds := 0
		minGreens := 0
		minBlues := 0
		for _, round := range game.rounds {
			minReds = max(minReds, round.red)
			minGreens = max(minGreens, round.green)
			minBlues = max(minBlues, round.blue)
		}
		powerSum += minReds * minGreens * minBlues
	}
	return powerSum
}

func parseGames(lines []string) []Game {
	gameIdPattern := regexp.MustCompile(`Game\s(\d+):`)
	roundPattern := regexp.MustCompile(`(\d+)\s(red|green|blue)`)
	games := make([]Game, 0)
	for _, line := range lines {
		gameMatch := gameIdPattern.FindStringSubmatch(line)
		game := Game{}
		game.id, _ = strconv.Atoi(gameMatch[1])

		gameRemovedLine := gameIdPattern.ReplaceAllString(line, "")
		seperatedRounds := strings.Split(gameRemovedLine, ";")
		for _, roundGroup := range seperatedRounds {
			roundMatches := roundPattern.FindAllString(roundGroup, -1)

			red := 0
			green := 0
			blue := 0
			for _, roundMatch := range roundMatches {
				roundMatchSplit := strings.Split(roundMatch, " ")
				roundValue, _ := strconv.Atoi(roundMatchSplit[0])
				switch roundMatchSplit[1] {
				case "red":
					red = roundValue
				case "green":
					green = roundValue
				case "blue":
					blue = roundValue
				}
			}
			round := Round{red, green, blue}

			game.rounds = append(game.rounds, round)
		}

		games = append(games, game)
	}
	return games
}

func readInput() []string {
	scanner := bufio.NewScanner(os.Stdin)

	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
