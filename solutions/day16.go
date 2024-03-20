package solutions

import (
	"AdventOfCode2023/utils"
	"fmt"
)

type vec struct {
	x int
	y int
}

type beam struct {
	pos vec
	dir vec
}

func getDirIndex(dir vec) int {
	if dir.x == 0 {
		if dir.y == 1 {
			return 0
		}
		return 1
	}
	if dir.x == 1 {
		return 2
	}
	return 3
}

func containsDir(s [4]bool, dir vec) bool {
	index := getDirIndex(dir)
	return s[index]
}

func addDir(s [4]bool, dir vec) [4]bool {
	index := getDirIndex(dir)
	s[index] = true
	return s
}

func isEnergized(energized [4]bool) bool {
	return energized[0] || energized[1] || energized[2] || energized[3]
}

func traceBeam(input []string, b beam) int {
	visited := 0
	beams := []beam{b}
	energized := make([][][4]bool, len(input))
	for i := range input {
		energized[i] = make([][4]bool, len(input[i]))
		for j := range input[i] {
			energized[i][j] = [4]bool{}
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
			if !isEnergized(energized[pos.y][pos.x]) {
				visited++
			}
			energized[pos.y][pos.x] = addDir(energized[pos.y][pos.x], dir)
			continue
		}
		if currentChar == '|' {
			if containsDir(energized[pos.y][pos.x], dir) {
				beams = beams[1:]
				continue
			}
			if !isEnergized(energized[pos.y][pos.x]) {
				visited++
			}
			energized[pos.y][pos.x] = addDir(energized[pos.y][pos.x], dir)
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
			if !isEnergized(energized[pos.y][pos.x]) {
				visited++
			}
			energized[pos.y][pos.x] = addDir(energized[pos.y][pos.x], dir)
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
			if !isEnergized(energized[pos.y][pos.x]) {
				visited++
			}
			energized[pos.y][pos.x] = addDir(energized[pos.y][pos.x], vec{dir.x, dir.y})
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
			if !isEnergized(energized[pos.y][pos.x]) {
				visited++
			}
			energized[pos.y][pos.x] = addDir(energized[pos.y][pos.x], vec{dir.x, dir.y})
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
	return visited
}

func SolveDay16Part1(input []string) int {
	b := beam{vec{-1, 0}, vec{1, 0}}
	return traceBeam(input, b)
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
		return traceBeam(input, beams[i])
	}
	return utils.Parallelise(utils.MaxAcc, fn, len(beams))
}

func Day16(input []string) []string {
	part1 := SolveDay16Part1(input)
	part2 := SolveDay16Part2(input)
	return []string{fmt.Sprint(part1), fmt.Sprint(part2)}
}
