package anagram

import (
	"sort"
	"strings"
)

func Detect(subject string, candidates []string) []string {
	var result []string
	for _, candidate := range candidates {
		if IsAnagram(subject, candidate) {
			result = append(result, candidate)
		}
	}
	return result
}

func IsAnagram(subject string, candidate string) bool {
	lowerSubject := strings.ToLower(subject)
	lowerCandidate := strings.ToLower(candidate)

	if lowerSubject == lowerCandidate {
		return false
	}

	if len(lowerSubject) != len(lowerCandidate) {
		return false
	}

	subjectChars := strings.Split(strings.ToLower(lowerSubject), "")
	candidateChars := strings.Split(strings.ToLower(lowerCandidate), "")
	sort.Strings(subjectChars)
	sort.Strings(candidateChars)

	return strings.Join(subjectChars, "") == strings.Join(candidateChars, "")
}
