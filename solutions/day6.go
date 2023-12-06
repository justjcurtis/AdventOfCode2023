/*
Copyright © 2023 Jacson Curtis <justjcurtis@gmail.com>
*/
package solutions

import (
	"math"
	"strconv"
	"strings"
)

func GetNums(arr []string) []int {
	nums := []int{}
	for _, str := range arr {
		if len(str) == 0 {
			continue
		}
		num, _ := strconv.Atoi(str)
		nums = append(nums, num)
	}
	return nums
}

func ParseDay6Input(input []string) [][]int {
	times := GetNums(strings.Split(input[0][6:], " "))
	records := GetNums(strings.Split(input[1][9:], " "))
	races := make([][]int, len(times))
	for i := 0; i < len(times); i++ {
		races[i] = []int{times[i], records[i]}
	}
	return races
}

func GetIsOdd(num int) bool {
	return num%2 == 1
}

func GetHoldTime(raceLength int, record int) int {
	return int(math.Floor((float64(-raceLength) + (math.Sqrt(float64((raceLength * raceLength) - (4 * record))))) / -2))
}

func GetRacesBetweenRecordAndMax(raceLength int, record int) int {
	maxHoldTime := raceLength / 2
	recordHoldTime := GetHoldTime(raceLength, record)
	dist := maxHoldTime - recordHoldTime
	result := 2 * dist
	isOdd := GetIsOdd(raceLength)
	if isOdd {
		result++
	}
	return result - 1
}

func SolveDay6Part1(races [][]int) int {
	result := 1
	for _, race := range races {
		opts := GetRacesBetweenRecordAndMax(race[0], race[1])
		result *= opts
	}
	return result
}

func SolveDay6Part2(races [][]int) int {
	timeStr := ""
	recordStr := ""
	for _, race := range races {
		timeStr += strconv.Itoa(race[0])
		recordStr += strconv.Itoa(race[1])
	}
	time, _ := strconv.Atoi(timeStr)
	record, _ := strconv.Atoi(recordStr)
	return GetRacesBetweenRecordAndMax(time, record)
}

func Day6(input []string) []string {
	races := ParseDay6Input(input)
	part1 := SolveDay6Part1(races)
	part2 := SolveDay6Part2(races)
	return []string{strconv.Itoa(part1), strconv.Itoa(part2)}
}
