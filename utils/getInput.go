/*
Copyright © 2023 Jacson Curtis <justjcurtis@gmail.com>
*/
package utils

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

func GetInput(day int) []string {
	cwd, _ := os.Getwd()
	f, err := os.Open(cwd + "/puzzleInput/day_" + strconv.Itoa(day) + ".txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	lines := make([]string, 0)
	rd := bufio.NewReader(f)
	for {
		line, _, err := rd.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		lines = append(lines, string(line))
	}
	return lines
}
