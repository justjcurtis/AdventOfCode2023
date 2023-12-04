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

	"github.com/patrickmn/go-cache"
)

func GetNumbers(numbersString string) []int {
	numbers := []int{}
	for _, number := range strings.Split(numbersString, " ") {
		if len(number) == 0 {
			continue
		}
		number, _ := strconv.Atoi(number)
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

var day4Cache = cache.New(cache.NoExpiration, cache.NoExpiration)

func GetWinCount(line string) int {
	if winCount, found := day4Cache.Get(line); found {
		return winCount.(int)
	}

	winners, numbers := ParseCard(line)
	winCount := 0
	for _, number := range numbers {
		if slices.Contains(winners, number) {
			winCount++
		}
	}
	day4Cache.Set(line, winCount, cache.NoExpiration)
	return winCount
}

func GetCardValue(line string) int {
	return int(math.Max(math.Pow(2, float64(GetWinCount(line)-1)), float64(0)))
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
	day4Cache.Flush()
	return []string{strconv.Itoa(part1), strconv.Itoa(part2)}
}
