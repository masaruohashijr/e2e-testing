package services

import (
	"strings"
)

func JaroWinklerDistance(s1, s2 string) float64 {

	s1Matches := make([]bool, len(s1))
	s2Matches := make([]bool, len(s2))

	var matchingCharacters = 0.0
	var transpositions = 0.0

	if len(s1) == 0 || len(s2) == 0 {
		return 0
	}

	if len(s1) == 0 && len(s2) == 0 {
		return 1
	}

	if strings.EqualFold(s1, s2) {
		return 1
	}

	matchDistance := len(s1)
	if len(s2) > matchDistance {
		matchDistance = len(s2)
	}
	matchDistance = matchDistance/2 - 1

	for i := range s1 {
		low := i - matchDistance
		if low < 0 {
			low = 0
		}
		high := i + matchDistance + 1
		if high > len(s2) {
			high = len(s2)
		}
		for j := low; j < high; j++ {
			if s2Matches[j] {
				continue
			}
			if s1[i] != s2[j] {
				continue
			}
			s1Matches[i] = true
			s2Matches[j] = true
			matchingCharacters++
			break
		}
	}

	if matchingCharacters == 0 {
		return 0
	}

	k := 0
	for i := range s1 {
		if !s1Matches[i] {
			continue
		}
		for !s2Matches[k] {
			k++
		}
		if s1[i] != s2[k] {
			transpositions++
		}
		k++
	}
	transpositions /= 2

	weight := (matchingCharacters/float64(len(s1)) + matchingCharacters/float64(len(s2)) + (matchingCharacters-transpositions)/matchingCharacters) / 3

	l := 0
	p := 0.1

	s1 = strings.ToLower(s1)
	s2 = strings.ToLower(s2)

	if weight > 0.7 {
		for (l < 4) && s1[l] == s2[l] {
			l++
		}
		weight = weight + float64(l)*p*(1-weight)
	}

	return weight
}

func CompareTwoSentences(s1, s2 string, percent float64) bool {
	//logging.Debug.Println("CompareTwoSentences")
	d := JaroWinklerDistance(s1, s2)
	if d*100 > percent {
		return true
	}
	return false
}
