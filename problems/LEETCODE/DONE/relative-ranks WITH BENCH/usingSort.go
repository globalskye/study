package main

import (
	"sort"
	"strconv"
)

func findRelativeRanksUsingSort(score []int) []string {
	type ScoreWithIndex struct {
		Score int
		Index int
	}
	scores := make([]ScoreWithIndex, len(score))
	for i, s := range score {
		scores[i] = ScoreWithIndex{s, i}
	}
	sort.Slice(scores, func(i, j int) bool {
		return scores[i].Score > scores[j].Score
	})

	res := make([]string, len(score))
	for i, s := range scores {
		switch i {
		case 0:
			res[s.Index] = "Gold Medal"
		case 1:
			res[s.Index] = "Silver Medal"
		case 2:
			res[s.Index] = "Bronze Medal"
		default:
			res[s.Index] = strconv.Itoa(i + 1)
		}
	}

	return res
}
