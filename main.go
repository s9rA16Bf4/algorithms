package main

import (
	"algorithm/bubblesort"
	"fmt"
)

func main() {
	testData1 := []string{"a", "c", "b"}

	fmt.Println(bubblesort.Bubblesort(testData1))
}
