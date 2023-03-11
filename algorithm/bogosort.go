package algorithm

import (
	"math/rand"
	"time"
)

func Bogosort[V int | string | rune](unsorted_list []V) ([]V, time.Duration) {
	sorted_list := unsorted_list
	start := time.Now()

	for !checkIfSortIsCorrect(sorted_list) {
		start = time.Now()

		// Randomize the list
		for i := range sorted_list {
			j := rand.Intn(i + 1)
			sorted_list[i], sorted_list[j] = sorted_list[j], sorted_list[i]
		}
	}

	return sorted_list, time.Since(start)
}

func checkIfSortIsCorrect[V int | string | rune](unsorted_list []V) bool {

	for x := 0; x < len(unsorted_list); x++ {
		for y := x; y < len(unsorted_list); y++ {
			if unsorted_list[x] > unsorted_list[y] {
				return false
			}
		}
	}

	return true
}
