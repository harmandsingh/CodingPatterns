package main

func insertInterval(existingIntervals [][]int, newInterval []int) [][]int {
	answer := [][]int{}

	counter := 0

	for counter < len(existingIntervals) && existingIntervals[counter][0] < newInterval[0] {
		answer = append(answer, existingIntervals[counter])
		counter++
	}

	if len(answer) == 0 || answer[len(answer)-1][1] < newInterval[0] {
		answer = append(answer, newInterval)
	} else {
		if answer[len(answer)-1][1] < newInterval[1] {
			answer[len(answer)-1][1] = newInterval[1]
		}
	}

	for counter < len(existingIntervals) {
		if answer[len(answer)-1][1] >= existingIntervals[counter][0] {
			if existingIntervals[counter][1] > answer[len(answer)-1][1] {
				answer[len(answer)-1][1] = existingIntervals[counter][1]
			}
		} else {
			answer = append(answer, existingIntervals[counter])
		}
		counter++
	}

	return answer
}
