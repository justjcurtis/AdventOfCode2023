/*
Copyright © 2023 Jacson Curtis <justjcurtis@gmail.com>
*/
package solutions

import (
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/patrickmn/go-cache"
)

type SeedMap struct {
	dests   []int
	sources []int
	lengths []int
	length  int
}

func (this *SeedMap) get(s int) (int, bool) {
	for i := 0; i < this.length; i++ {
		dest := this.dests[i]
		source := this.sources[i]
		length := this.lengths[i]
		if s >= source && s < source+length {
			return dest + s - source, true
		}
	}
	return s, false
}

func ParseMap(mapString []string) SeedMap {
	seedMap := SeedMap{}
	for _, line := range mapString {
		parts := strings.Split(line, " ")
		nums := make([]int, 3)
		for i, part := range parts {
			num, _ := strconv.Atoi(part)
			nums[i] = num
		}
		seedMap.dests = append(seedMap.dests, nums[0])
		seedMap.sources = append(seedMap.sources, nums[1])
		seedMap.lengths = append(seedMap.lengths, nums[2])
	}
	seedMap.length = len(mapString)
	return seedMap
}

func ParseSeeds(seedStringArr []string) []int {
	seeds := make([]int, len(seedStringArr))
	for i, seedString := range seedStringArr {
		seed, _ := strconv.Atoi(seedString)
		seeds[i] = seed
	}
	return seeds
}

func ParseInput(input []string) ([]int, []SeedMap) {
	mapCount := 7
	seeds := ParseSeeds(strings.Split(input[0][7:], " "))
	mapArray := make([][]string, mapCount)
	offset := 0
	for i := 0; i < mapCount; i++ {
		for lineIndex := 2 + offset; lineIndex < len(input); lineIndex++ {
			offset++
			line := input[lineIndex]
			if len(line) < 1 {
				break
			}
			if line[len(line)-1] == ':' {
				continue
			}
			mapArray[i] = append(mapArray[i], line)
		}
	}
	maps := make([]SeedMap, mapCount)
	for i, mapString := range mapArray {
		maps[i] = ParseMap(mapString)
	}
	return seeds, maps
}

var Day3Cache = cache.New(5*time.Minute, 10*time.Minute)

func GetLocation(seed int, maps []SeedMap) int {
	if val, found := Day3Cache.Get(strconv.Itoa(seed)); found {
		return val.(int)
	}
	currentSeed := seed
	for i := 0; i < len(maps); i++ {
		currentSeed, _ = maps[i].get(currentSeed)
	}
	Day3Cache.Set(strconv.Itoa(seed), currentSeed, cache.DefaultExpiration)
	return currentSeed
}

func SolveDay5Part1(seeds []int, maps []SeedMap) int {
	minLocation := -1
	for _, seed := range seeds {
		location := GetLocation(seed, maps)
		if minLocation == -1 || location < minLocation {
			minLocation = location
		}
	}
	return minLocation
}

type SeedRange struct {
	start  int
	length int
}

func ParseSeedRanges(seeds []int) []SeedRange {
	newSeeds := make([]SeedRange, len(seeds)/2)
	for i := 0; i < len(seeds)-1; i += 2 {
		newSeeds[i/2].start = seeds[i]
		newSeeds[i/2].length = seeds[i+1]
	}
	return newSeeds
}

func (this *SeedRange) Find(target int, maxTarget int) int {
	if maxTarget >= this.start && target < this.start+this.length {
		if target < this.start {
			return this.start - target
		}
		return 0
	}
	return -1
}

func (this *SeedMap) GetRangesForTargetRange(target int, maxTarget int) [][]int {
	ranges := make([][]int, this.length)
	minDest := -1
	maxDest := -1
	minSource := -1
	maxSource := -1
	for i := 0; i < this.length; i++ {
		dest := this.dests[i]
		if minDest == -1 || dest < minDest {
			minDest = dest
		}
		if maxDest == -1 || dest > maxDest {
			maxDest = dest
		}

		source := this.sources[i]
		if minSource == -1 || source < minSource {
			minSource = source
		}
		if maxSource == -1 || source > maxSource {
			maxSource = source
		}

		length := this.lengths[i]

		ranges[i] = []int{dest, dest + length - 1, source, source + length - 1}
	}
	if minSource-1 > target {
		ranges = append(ranges, []int{target, minSource - 1, target, minSource - 1})
	}
	if maxSource+1 < maxTarget {
		ranges = append(ranges, []int{maxSource + 1, maxTarget, maxSource + 1, maxTarget})
	}
	return ranges
}

func (this *SeedMap) GetMinMapInputRangeForTargetRange(target int, maxTarget int) [][]int {
	resultRanges := [][]int{}
	ranges := this.GetRangesForTargetRange(target, maxTarget)
	for i := 0; i < len(ranges); i++ {
		destStart, destEnd := ranges[i][0], ranges[i][1]
		sourceStart, sourceEnd := ranges[i][2], ranges[i][3]
		if target+maxTarget >= destStart && target < destEnd {
			startDiff := target - destStart
			if startDiff < 0 {
				startDiff = 0
			}
			endDiff := destEnd - maxTarget
			if endDiff < 0 {
				endDiff = 0
			}
			if sourceStart+startDiff > sourceEnd-endDiff {
				continue
			}
			resultRanges = append(resultRanges, []int{sourceStart + startDiff, sourceEnd - endDiff, destStart + startDiff})
		}
	}
	sort.Slice(resultRanges, func(i, j int) bool {
		return resultRanges[i][2] < resultRanges[j][2]
	})
	return resultRanges
}

func GetMaxTarget(seedMap SeedMap) int {
	maxTarget := -1
	for i := 0; i < seedMap.length; i++ {
		dest := seedMap.dests[i]
		length := seedMap.lengths[i]
		if dest+length > maxTarget {
			maxTarget = dest + length
		}
	}
	return maxTarget
}

func RecursePart2(seeds []SeedRange, maps []SeedMap, target int, maxTarget int, mapIndex int) int {
	if mapIndex < 0 {
		for _, seed := range seeds {
			offset := seed.Find(target, maxTarget)
			if offset != -1 {
				return target + offset
			}
		}
		return -1
	}

	seedMap := maps[mapIndex]
	ranges := seedMap.GetMinMapInputRangeForTargetRange(target, maxTarget)

	for _, rangePair := range ranges {
		sourceStart, sourceEnd := rangePair[0], rangePair[1]
		destStart := rangePair[2]
		result := RecursePart2(seeds, maps, sourceStart, sourceEnd, mapIndex-1)
		if result != -1 {
			diff := destStart - sourceStart
			return result + diff
		}
	}
	return -1
}

func SolveDay5Part2(seeds []SeedRange, maps []SeedMap) int {
	target := 0
	maxTarget := GetMaxTarget(maps[len(maps)-1])
	result := RecursePart2(seeds, maps, target, maxTarget, len(maps)-1)
	if result != -1 {
		return result
	}
	return -1
}
func Day5(input []string) []string {
	seeds, maps := ParseInput(input)
	part1 := SolveDay5Part1(seeds, maps)
	secondSeeds := ParseSeedRanges(seeds)
	part2 := SolveDay5Part2(secondSeeds, maps)
	return []string{strconv.Itoa(part1), strconv.Itoa(part2)}
}
