package day2

import (
	"advent/utils"
	"fmt"
	"strconv"
	"strings"
)

func Solve() {
	filename := "solutions/day2/data.txt"
	content, err := utils.ReadData(filename)
	if err != nil {
		panic(err)
	}
	reports := parseReports(content)
	safeReports := getSafeReports(reports)
	fmt.Printf("Safe reports: %d\n", safeReports)
}

func parseReports(content string) [][]int {
	reports := [][]int{}
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		report := []int{}
		values := strings.Split(line, " ")
		for _, value := range values {
			intValue, _ := strconv.Atoi(value)
			report = append(report, intValue)
		}
		reports = append(reports, report)
	}
	return reports
}

func getSafeReports(reports [][]int) int {
	safeCount := 0
	for _, report := range reports {
		pastValue := 0
		isReportIncreasing := false

		if report[0] == report[1] {
			continue
		}

		if report[0] < report[1] {
			isReportIncreasing = true
		} else {
			isReportIncreasing = false
		}

		isSafe := true
		for index, value := range report {
			if index == 0 {
				pastValue = value
				continue
			}
			if isReportIncreasing {
				if !isTwoValuesWithinValidRange(pastValue, value) {
					isSafe = false
					break
				}
			} else {
				if !isTwoValuesWithinValidRange(value, pastValue) {
					isSafe = false
					break
				}
			}
			pastValue = value
		}
		if isSafe {
			safeCount++
		}
	}

	return safeCount
}

func isTwoValuesWithinValidRange(low int, high int) bool {
	if high-low < 1 || high-low > 3 {
		return false
	}
	return true
}
