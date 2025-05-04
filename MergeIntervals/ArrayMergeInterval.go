package main

import "sort"

func mergeIntervals(intervals [][]int) [][]int {
	if len(intervals) == 1 {
		return intervals
	}

	visitedIndices := map[int]bool{}
	answer := [][]int{}

	for i := range intervals {
		_, visited := visitedIndices[i]
		if visited {
			continue
		}

		initInterval := intervals[i]
		for j := i + 1; j < len(intervals); j++ {
			_, visited := visitedIndices[j]
			if visited {
				continue
			}

			if initInterval[0] < intervals[j][0] {
				if initInterval[1] >= intervals[j][0] && intervals[j][1] >= initInterval[1] {
					initInterval = []int{initInterval[0], intervals[j][1]}
					visitedIndices[i], visitedIndices[j] = true, true
				} else if initInterval[1] >= intervals[j][0] && intervals[j][1] < initInterval[1] {
					initInterval = []int{initInterval[0], initInterval[1]}
					visitedIndices[i], visitedIndices[j] = true, true
				}
			} else if initInterval[0] > intervals[j][0] {
				if intervals[j][1] >= initInterval[0] && intervals[j][1] >= initInterval[1] {
					initInterval = []int{intervals[j][0], intervals[j][1]}
					visitedIndices[i], visitedIndices[j] = true, true
				} else if intervals[j][1] >= initInterval[0] && intervals[j][1] < initInterval[1] {
					initInterval = []int{intervals[j][0], initInterval[1]}
					visitedIndices[i], visitedIndices[j] = true, true
				}
			} else {
				continue
			}

			// if initInterval[1] < intervals[j][0] {
			//     continue
			// }

			// initInterval = []int{intervals[i][0], intervals[j][1]}
		}
		answer = append(answer, initInterval)
	}

	return answer
}

func mergeIntervalsOptimized(intervals [][]int) [][]int {
	if len(intervals) == 1 {
		return intervals
	}

	// Sort the array in ascending order
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	answer := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		currentInterval := intervals[i]
		if answer[len(answer)-1][1] >= currentInterval[0] {
			if answer[len(answer)-1][1] < currentInterval[1] {
				answer[len(answer)-1][1] = currentInterval[1]
			}
		} else {
			answer = append(answer, currentInterval)
		}
	}

	return answer
}
