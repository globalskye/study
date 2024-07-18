package main

import (
	"strconv"
)

type ScoreWithIndex struct {
	Score int
	Index int
}

func findRelativeRanksUsingSort(score []int) []string {

	scores := make([]ScoreWithIndex, len(score))
	for i, s := range score {
		scores[i] = ScoreWithIndex{s, i}
	}

	shellSort(scores)

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
func mergeSort(items []ScoreWithIndex) []ScoreWithIndex {
	var num = len(items)

	if num == 1 {
		return items
	}

	middle := int(num / 2)
	var (
		left  = make([]ScoreWithIndex, middle)
		right = make([]ScoreWithIndex, num-middle)
	)
	for i := 0; i < num; i++ {
		if i < middle {
			left[i] = items[i]
		} else {
			right[i-middle] = items[i]
		}
	}

	return merge(mergeSort(left), mergeSort(right))
}

func merge(left, right []ScoreWithIndex) (result []ScoreWithIndex) {
	result = make([]ScoreWithIndex, len(left)+len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0].Score < right[0].Score {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	return
}
func shellSort(items []ScoreWithIndex) {
	var (
		n    = len(items)
		gaps = []int{1}
		k    = 1
	)

	for {
		gap := element(2, k) + 1
		if gap > n-1 {
			break
		}
		gaps = append([]int{gap}, gaps...)
		k++
	}

	for _, gap := range gaps {
		for i := gap; i < n; i += gap {
			j := i
			for j > 0 {
				if items[j-gap].Score > items[j].Score {
					items[j-gap], items[j] = items[j], items[j-gap]
				}
				j = j - gap
			}
		}
	}
}

func element(a, b int) int {
	e := 1
	for b > 0 {
		if b&1 != 0 {
			e *= a
		}
		b >>= 1
		a *= a
	}
	return e
}
