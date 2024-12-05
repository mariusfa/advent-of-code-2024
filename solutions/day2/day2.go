package day2

// testdata
// 7 6 4 2 1
// 1 2 7 8 9
// 9 7 6 2 1
// 8 6 4 4 1
// 1 3 2 4 5
// 1 3 6 7 9

func Solve() {
	// TODO: Implement
}

func getSafeReports(reports [][]int) int {
	safeCount := 0
	for _, report := range reports {
		pastValue := 0
		isReportIncreasing := false

		if report[0] == report[1] {
			println("unsafe")
			continue
		}

		if report[0] < report[1] {
			println("increasing")
			isReportIncreasing = true
		} else {
			println("decresing")
			isReportIncreasing = false
		}

		isSafe := true
		for index, value := range report {
			println("value", value)
			println("pastValue", pastValue)
			if index == 0 {
				pastValue = value
				continue
			}
			if isReportIncreasing {
				if value-pastValue < 1 && value-pastValue > 3 {
					isSafe = false
					break
				}
			} else {
				println("value-pastValue", pastValue-value)
				if pastValue-value < 1 && pastValue-value > 3 {
					isSafe = false
					break
				}
			}
			pastValue = value
		}
		if isSafe {
			println("safe")
			safeCount++
		}
	}

	return safeCount
}
