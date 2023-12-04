/*
Copyright © 2023 Jacson Curtis <justjcurtis@gmail.com>
*/
package solutions

import (
	"AdventOfCode2023/utils"
	"math"
	"slices"
	"strconv"
	"strings"
)

func GetNumbers(numbersString string) []int {
	numbers := []int{}
	for _, number := range strings.Split(numbersString, " ") {
		number, err := strconv.Atoi(number)
		if err != nil {
			continue
		}
		numbers = append(numbers, number)
	}
	return numbers
}

func ParseCard(line string) ([]int, []int) {
	numberString := strings.Split(strings.Split(line, ":")[1], "|")
	winners := GetNumbers(numberString[0])
	numbers := GetNumbers(numberString[1])
	return winners, numbers
}

func GetWinCount(line string) int {
	winners, numbers := ParseCard(line)
	winCount := 0
	for _, number := range numbers {
		if slices.Contains(winners, number) {
			winCount++
		}
	}
	return winCount
}

func GetCardValue(line string) int {
	winCount := GetWinCount(line)
	return int(math.Max(math.Pow(2, float64(winCount-1)), float64(0)))
}

func GetCardValues(input []string) int {
	fn := func(j int) int {
		return GetCardValue(input[j])
	}
	return utils.Parallelise(utils.IntAcc, fn, len(input))
}

func GetTotalCards(input []string, c chan<- int) {
	totalCards := len(input)
	cardCounts := []int{}
	for range input {
		cardCounts = append(cardCounts, 1)
	}
	for i := range input {
		cardsToProcess := cardCounts[i]
		winCount := GetWinCount(input[i])
		totalCards += cardsToProcess * winCount
		for j := 1; j <= winCount; j++ {
			z := i + j
			cardCounts[z] = cardCounts[z] + cardsToProcess
		}
	}
	c <- totalCards
}

func Day4(input []string) []string {
	ch := make(chan int)
	go GetTotalCards(input, ch)
	part1 := GetCardValues(input)
	part2 := <-ch
	return []string{strconv.Itoa(part1), strconv.Itoa(part2)}
}
