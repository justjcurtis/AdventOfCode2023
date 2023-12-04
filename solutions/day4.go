/*
Copyright © 2023 Jacson Curtis <justjcurtis@gmail.com>
*/
package solutions

import (
	"AdventOfCode2023/utils"
	"strconv"
	"strings"
)

func GetWinningCount(card string) int {
	numberString := strings.Split(card, ":")[1]
	winners := GetNumbers(strings.Split(numberString, "|")[0])
	numbers := GetNumbers(strings.Split(numberString, "|")[1])
	count := 0
	for _, number := range numbers {
		for _, winner := range winners {
			if number == winner {
				count++
				break
			}
		}
	}
	return count
}

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

func GetCardValue(card string) int {
	numberString := strings.Split(card, ":")[1]
	winners := GetNumbers(strings.Split(numberString, "|")[0])
	numbers := GetNumbers(strings.Split(numberString, "|")[1])
	value := 0
	for _, number := range numbers {
		for _, winner := range winners {
			if number == winner {
				if value == 0 {
					value++
				} else {
					value *= 2
				}
				break
			}
		}
	}
	return value
}

func GetCardValues(cards []string) int {
	fn := func(j int) int {
		return GetCardValue(cards[j])
	}
	return utils.Parallelise(utils.IntAcc, fn, len(cards))
}

func GetNextCardForProcessing(cardMap map[int][]int) (int, int) {
	for index := 0; index < len(cardMap); index++ {
		card := cardMap[index]
		if card[0] > card[1] {
			return index, card[0] - card[1]
		}
	}
	return -1, 0
}

func GetTotalCards(cards []string) int {
	cardCount := len(cards)
	totalCards := len(cards)
	cardMap := map[int][]int{}
	for i := range cards {
		cardMap[i] = []int{1, 0}
	}
	for cardCount > 0 {
		i, n := GetNextCardForProcessing(cardMap)
		if i == -1 {
			break
		}
		winCount := GetWinningCount(cards[i])
		cardMap[i] = []int{0, cardMap[i][1] + n}
		cardCount -= n
		for j := 1; j <= winCount; j++ {
			z := i + j
			if _, ok := cardMap[z]; !ok {
				cardMap[z] = []int{0, 0}
			}
			cardMap[z] = []int{
				cardMap[z][0] + n,
				cardMap[z][1],
			}
			cardCount += n
			totalCards += n
		}
	}
	return totalCards
}

func Day4(cards []string) []string {
	part1 := GetCardValues(cards)
	part2 := GetTotalCards(cards)
	return []string{strconv.Itoa(part1), strconv.Itoa(part2)}
}
