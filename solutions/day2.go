/*
Copyright © 2023 Jacson Curtis <justjcurtis@gmail.com>
*/
package solutions

import (
	"strconv"
	"strings"
)

type day2Config struct {
	rMax int
	gMax int
	bMax int
}

func RoundIsPossible(round day2Round, config day2Config) bool {
	if round.r > config.rMax {
		return false
	}
	if round.g > config.gMax {
		return false
	}
	if round.b > config.bMax {
		return false
	}
	return true
}

func GameIsPossible(game day2Game, config day2Config) bool {
	for _, round := range game.rounds {
		if !RoundIsPossible(round, config) {
			return false
		}
	}
	return true
}

type day2Round struct {
	r int
	g int
	b int
}
type day2Game struct {
	id     int
	rounds []day2Round
}

func ParseGame(line string) day2Game {
	arr := strings.Split(line, ":")
	id, _ := strconv.Atoi(arr[0][5:])
	roundsString := strings.Split(arr[1], ";")
	numBuf := ""
	rounds := make([]day2Round, 0)
	for _, roundString := range roundsString {
		waitForComma := false
		round := day2Round{}
		for _, char := range roundString {
			if char == ',' {
				waitForComma = false
				numBuf = ""
				continue
			} else if waitForComma {
				continue
			}
			if char == ' ' {
				continue
			}
			if char == 'r' {
				round.r, _ = strconv.Atoi(numBuf)
				numBuf = ""
				waitForComma = true
				continue
			}
			if char == 'g' {
				round.g, _ = strconv.Atoi(numBuf)
				numBuf = ""
				waitForComma = true
				continue
			}
			if char == 'b' {
				round.b, _ = strconv.Atoi(numBuf)
				numBuf = ""
				waitForComma = true
				continue
			}
			numBuf += string(char)
		}
		rounds = append(rounds, round)
	}
	return day2Game{id, rounds}
}

func GetMinConfig(game day2Game) day2Config {
	config := day2Config{}
	for _, round := range game.rounds {
		if round.r > config.rMax {
			config.rMax = round.r
		}
		if round.g > config.gMax {
			config.gMax = round.g
		}
		if round.b > config.bMax {
			config.bMax = round.b
		}
	}
	return config
}

func GetConfigPower(config day2Config) int {
	return config.rMax * config.gMax * config.bMax
}

func Day2(input []string) []string {
	games := make([]day2Game, 0)
	for _, line := range input {
		games = append(games, ParseGame(line))
	}
	part1 := 0
	part2 := 0
	config1 := day2Config{rMax: 12, gMax: 13, bMax: 14}
	for _, game := range games {
		if GameIsPossible(game, config1) {
			part1 += game.id
		}
		config := GetMinConfig(game)
		part2 += GetConfigPower(config)
	}

	return []string{strconv.Itoa(part1), strconv.Itoa(part2)}
}
