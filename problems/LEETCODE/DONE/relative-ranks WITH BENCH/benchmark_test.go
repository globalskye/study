package main

import (
	"math/rand"
	"testing"
)

func generateRandomScores(n int) []int {
	scores := make([]int, n)
	for i := range scores {
		scores[i] = rand.Intn(n)
	}
	return scores
}

func BenchmarkFindRelativeRanksUsingHeap(b *testing.B) {
	scores := generateRandomScores(1_000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findRelativeRanksUsingHeap(scores)
	}
}

func BenchmarkFindRelativeRanksUsingSort(b *testing.B) {
	scores := generateRandomScores(1_000)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		findRelativeRanksUsingSort(scores)
	}

}
