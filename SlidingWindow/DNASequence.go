package main

import "fmt"

func findRepeatedDnaSequences(s string) []string {
	windowSize := 10
	if len(s) <= windowSize {
		return []string{}
	}

	start, end := 0, windowSize
	initSequence := s[start:end]

	// Array to store all valid sequences
	allValidSequences := map[string]int{}
	// Array to store valid sequences that occur more that once
	answer := []string{}

	allValidSequences[initSequence] = 1

	// If we've found one valid sequence, we can check if it occurs more than once
	for end < len(s) {
		start++
		end++

		if end == len(s) {
			initSequence = s[start:]
		} else {
			initSequence = s[start:end]
		}

		value, ok := allValidSequences[initSequence]
		if ok {
			allValidSequences[initSequence] = value + 1
		} else {
			allValidSequences[initSequence] = 1
		}
	}

	for k, v := range allValidSequences {
		if v >= 2 {
			answer = append(answer, k)
		}
	}

	return answer
}

func main() {
	string1 := "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
	answer := findRepeatedDnaSequences(string1)
	fmt.Println(answer)
}
