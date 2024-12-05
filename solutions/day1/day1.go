package day1

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Solve() {
	filename := "solutions/day1/data.txt"
	content, err := readData(filename)
	if err != nil {
		panic(err)
	}

	left, right, err := parseData(content)
	distance := calcDistance(left, right)
	fmt.Printf("Distance: %d\n", distance)
}

func calcDistance(left []int, right []int) int {
	slices.Sort(left)
	slices.Sort(right)

	distance := 0

	for i := 0; i < len(left); i++ {
		distance += absDiff(left[i], right[i])
	}

	return distance
}

func absDiff(a int, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func readData(filename string) (string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

func parseData(content string) ([]int, []int, error) {
	lines := strings.Split(content, "\n")
	var left []int
	var right []int
	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Split(line, "   ")
		leftValue, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, nil, err
		}
		left = append(left, leftValue)
		rightValue, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, nil, err
		}
		right = append(right, rightValue)
	}
	return left, right, nil
}

func printList(list []int) {
	for i := 0; i < len(list); i++ {
		println(list[i])
	}
}
