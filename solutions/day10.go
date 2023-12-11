/*
Copyright © 2023 Jacson Curtis <justjcurtis@gmail.com>
*/
package solutions

type day10Node struct {
	x          int
	y          int
	isStart    bool
	neighbours []*day10Node
}

type day10Parsed struct {
	startX int
	startY int
	pipes  [][]day10Node
}

var charMap = map[rune][][]int{
	'|': {{0, 1}, {0, -1}},
	'-': {{1, 0}, {-1, 0}},
	'L': {{0, 1}, {1, 0}},
	'J': {{0, 1}, {-1, 0}},
	'7': {{0, -1}, {1, 0}},
	'F': {{0, -1}, {-1, 0}},
	'S': {{0, 1}, {0, -1}, {1, 0}, {-1, 0}},
}

func ParseDay10Input(input []string) day10Parsed {
	startX := -1
	startY := -1
	neigbourMap := [][][][]int{}
	pipes := [][]day10Node{}
	for y, line := range input {
		row := [][][]int{}
		pipeRow := []day10Node{}
		for x, char := range line {
			if char == '.' {
				row = append(row, [][]int{})
			}
			isStart := char == 'S'
			if isStart {
				startX = x
				startY = y
			}
			row = append(row, charMap[char])
			pipeRow = append(pipeRow, day10Node{x, y, isStart, []*day10Node{}})
		}
		neigbourMap = append(neigbourMap, row)
	}
	for y, row := range pipes {
		for x, pipe := range row {
			nearby := neigbourMap[y][x]
			for _, dir := range nearby {
				nx := x + dir[0]
				ny := y + dir[1]
				if nx < 0 || ny < 0 || nx >= len(pipes) || ny >= len(pipes[nx]) {
					continue
				}
				pipe.neighbours = append(pipe.neighbours, &pipes[ny][nx])
			}
		}
	}
	return day10Parsed{startX, startY, pipes}
}

func aStart(start *day10Node, end *day10Node) int {
	visited := map[*day10Node]bool{}
	gScore := map[*day10Node]int{}
	fScore := map[*day10Node]int{}
	queue := []*day10Node{start}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if visited[current] {
			continue
		}
		visited[current] = true
		if current == end {
			if len(visited) > 1 {
				return 0
			}
		}
		for _, neighbour := range current.neighbours {
			queue = append(queue, neighbour)
		}
	}
	return -1
}

func Day10(input []string) []string {
	ParseDay10Input(input)
	return []string{"", ""}
}
