/*
Copyright © 2023 Jacson Curtis <justjcurtis@gmail.com>
*/
package solutions

import (
	"AdventOfCode2023/utils"
	"sort"
	"strconv"
	"unicode"
)

var camelCards = map[string]int{
	"A": 14,
	"K": 13,
	"Q": 12,
	"J": 11,
	"T": 10,
}

type CamelHand struct {
	cards              []int
	bid                int
	handType           int
	handTypeWithJokers int
}

func GetHandType(cards []int) int {
	handMap := make(map[int]int)
	for _, card := range cards {
		handMap[card]++
	}
	if len(handMap) == 1 {
		return 6
	}
	if len(handMap) == 2 {
		for _, count := range handMap {
			if count == 4 || count == 1 {
				return 5
			}
		}
		return 4
	}
	if len(handMap) == 3 {
		for _, count := range handMap {
			if count == 3 {
				return 3
			}
		}
		return 2
	}
	if len(handMap) == 4 {
		return 1
	}
	return 0
}

func GetHandTypeWithJokers(handType int, wildCount int) int {
	result := handType
	switch handType {
	case 0:
		if wildCount > 0 {
			result = 1
		}
		break
	case 2:
		if wildCount == 2 {
			result = 5
		} else if wildCount == 1 {
			result = 4
		}
		break
	default:
		if wildCount > 0 {
			result = handType + 2
		}
		break
	}
	if result > 6 {
		result = 6
	}
	return result

}

func ParseHand(line string) CamelHand {
	cards := make([]int, 5)
	wildCount := 0
	for i := 0; i < 5; i++ {
		char := line[i]
		if char == 'J' {
			wildCount++
		}
		if unicode.IsDigit(rune(char)) {
			num, _ := strconv.Atoi(string(char))
			cards[i] = num
			continue
		}
		cards[i] = camelCards[string(char)]
	}
	bid, _ := strconv.Atoi(line[6:])

	handType := GetHandType(cards)
	handTypeWithJokers := GetHandTypeWithJokers(handType, wildCount)
	return CamelHand{cards, bid, handType, handTypeWithJokers}
}

func ParseDay7(lines []string) []CamelHand {
	fn := func(j int) []CamelHand {
		return []CamelHand{ParseHand(lines[j])}
	}
	return utils.Parallelise(utils.ArrAcc[CamelHand], fn, len(lines))
}

func SolveDay7Part1(hands []CamelHand) int {
	total := 0
	sort.Slice(hands, func(i, j int) bool {
		if hands[i].handType == hands[j].handType {
			for c := 0; c < 5; c++ {
				if hands[i].cards[c] != hands[j].cards[c] {
					return hands[i].cards[c] < hands[j].cards[c]
				}
			}
		}
		return hands[i].handType < hands[j].handType
	})
	for i, hand := range hands {
		total += hand.bid * (i + 1)
	}
	return total
}

func SolveDay7Part2(hands []CamelHand) int {
	total := 0
	sort.Slice(hands, func(i, j int) bool {
		if hands[i].handTypeWithJokers == hands[j].handTypeWithJokers {
			for c := 0; c < 5; c++ {
				a := hands[i].cards[c]
				b := hands[j].cards[c]
				if a == 11 {
					a = 1
				}
				if b == 11 {
					b = 1
				}
				if a != b {
					return a < b
				}
			}
		}
		return hands[i].handTypeWithJokers < hands[j].handTypeWithJokers
	})
	for i, hand := range hands {
		total += hand.bid * (i + 1)
	}
	return total
}

func Day7(input []string) []string {
	hands := ParseDay7(input)
	part1 := SolveDay7Part1(hands)
	part2 := SolveDay7Part2(hands)
	return []string{strconv.Itoa(part1), strconv.Itoa(part2)}
}
