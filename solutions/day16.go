package solutions

import (
	"AdventOfCode2023/utils"
	"fmt"
	"slices"
)

type vec struct {
	x int
	y int
}

type beam struct {
	pos vec
	dir vec
}

func containsDir(s []vec, dir vec) bool {
	index := slices.IndexFunc(s, func(v vec) bool {
		return v.x == dir.x && v.y == dir.y
	})
	return index != -1
}

func countEnergized(energized [][][]vec) int {
	sum := 0
	for i := range energized {
		for j := range energized[i] {
			if len(energized[i][j]) > 0 {
				sum++
			}
		}
	}
	return sum
}

func traceBeam(input []string, b beam) [][][]vec {
	beams := []beam{b}
	energized := make([][][]vec, len(input))
	for i := range input {
		energized[i] = make([][]vec, len(input[i]))
		for j := range input[i] {
			energized[i][j] = []vec{}
		}
	}
	for {
		if len(beams) == 0 {
			break
		}
		dir := beams[0].dir
		beams[0].pos.x += dir.x
		beams[0].pos.y += dir.y
		pos := beams[0].pos

		if pos.x < 0 || pos.x >= len(input[0]) || pos.y < 0 || pos.y >= len(input) {
			beams = beams[1:]
			continue
		}

		currentChar := input[pos.y][pos.x]

		if currentChar == '.' {
			if containsDir(energized[pos.y][pos.x], dir) {
				beams = beams[1:]
				continue
			}
			energized[pos.y][pos.x] = append(energized[pos.y][pos.x], dir)
			continue
		}
		if currentChar == '|' {
			if containsDir(energized[pos.y][pos.x], dir) {
				beams = beams[1:]
				continue
			}
			energized[pos.y][pos.x] = append(energized[pos.y][pos.x], dir)
			if dir.x != 0 {
				beams[0].dir.y = 1
				beams[0].dir.x = 0
				beams = append(beams, beam{pos, vec{0, -1}})
			}
			continue
		}
		if currentChar == '-' {
			if containsDir(energized[pos.y][pos.x], dir) {
				beams = beams[1:]
				continue
			}
			energized[pos.y][pos.x] = append(energized[pos.y][pos.x], dir)
			if dir.y != 0 {
				beams[0].dir.x = 1
				beams[0].dir.y = 0
				beams = append(beams, beam{pos, vec{-1, 0}})
			}
			continue
		}
		if currentChar == '/' {
			if containsDir(energized[pos.y][pos.x], dir) {
				beams = beams[1:]
				continue
			}
			energized[pos.y][pos.x] = append(energized[pos.y][pos.x], vec{dir.x, dir.y})
			if dir.x != 0 {
				beams[0].dir.y = -1 * dir.x
				beams[0].dir.x = 0
			} else {
				beams[0].dir.x = -1 * dir.y
				beams[0].dir.y = 0
			}

			continue
		}
		if currentChar == '\\' {
			if containsDir(energized[pos.y][pos.x], dir) {
				beams = beams[1:]
				continue
			}
			energized[pos.y][pos.x] = append(energized[pos.y][pos.x], vec{dir.x, dir.y})
			if dir.x != 0 {
				beams[0].dir.y = 1 * dir.x
				beams[0].dir.x = 0
			} else {
				beams[0].dir.x = 1 * dir.y
				beams[0].dir.y = 0
			}
			continue
		}
	}
	return energized
}

func SolveDay16Part1(input []string) int {
	b := beam{vec{-1, 0}, vec{1, 0}}
	energized := traceBeam(input, b)
	return countEnergized(energized)
}

func getPossibleBeams(w int, h int) []beam {
	possibleBeams := make([]beam, (w*2)+(h*2))
	for i := 0; i < w; i++ {
		possibleBeams[i] = beam{vec{i, -1}, vec{0, 1}}
		possibleBeams[i+w] = beam{vec{i, w}, vec{0, -1}}
	}
	for i := 0; i < h; i++ {
		possibleBeams[i+(w*2)] = beam{vec{-1, i}, vec{1, 0}}
		possibleBeams[i+(w*2)+h] = beam{vec{h, i}, vec{-1, 0}}
	}
	return possibleBeams
}

func SolveDay16Part2(input []string) int {
	beams := getPossibleBeams(len(input[0]), len(input))
	fn := func(i int) int {
		energized := traceBeam(input, beams[i])
		return countEnergized(energized)
	}
	return utils.Parallelise(utils.MaxAcc, fn, len(beams))
}

func Day16(input []string) []string {
	part1 := SolveDay16Part1(input)
	part2 := SolveDay16Part2(input)
	return []string{fmt.Sprint(part1), fmt.Sprint(part2)}
}
