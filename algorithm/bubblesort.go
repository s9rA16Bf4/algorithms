package algorithm

import "time"

func Bubblesort[V int | string | rune](unsorted_list []V) ([]V, time.Duration) {
	sorted_list := unsorted_list
	start := time.Now()

	for x := 0; x < len(unsorted_list); x++ {

		for y := x; y < len(unsorted_list); y++ {

			if unsorted_list[x] > unsorted_list[y] {
				temp := unsorted_list[x]
				unsorted_list[x] = unsorted_list[y]
				unsorted_list[y] = temp
			}
		}

	}

	return sorted_list, time.Since(start)
}
