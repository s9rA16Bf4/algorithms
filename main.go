package main

import (
	"fmt"

	arg "github.com/s9rA16Bf4/ArgumentParser/go/arguments"
	"github.com/s9rA16Bf4/algorithms/algorithm"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

func checkBubblesort(testData1 []string, testData2 []int, testData3 []rune) {

	fmt.Println("## Bubblesort")
	fmt.Printf("\tInput: %s\n", testData1)
	result1, time := algorithm.Bubblesort(testData1)
	fmt.Printf("\tTime %s ; Result %s\n\n", time, result1)

	fmt.Printf("\tInput: %d\n", testData2)
	result2, time := algorithm.Bubblesort(testData2)
	fmt.Printf("\tTime %s ; Result %d\n\n", time, result2)

	fmt.Printf("\tInput: %c\n", testData3)
	result3, time := algorithm.Bubblesort(testData3)
	fmt.Printf("\tTime %s ; Result %c\n\n", time, result3)
}

func checkBogosort(testData1 []string, testData2 []int, testData3 []rune) {

	fmt.Println("## Bogosort")
	fmt.Printf("\tInput: %s\n", testData1)
	result1, time := algorithm.Bogosort(testData1)
	fmt.Printf("\tTime %s ; Result %s\n\n", time, result1)

	fmt.Printf("\tInput: %d\n", testData2)
	result2, time := algorithm.Bogosort(testData2)
	fmt.Printf("\tTime %s ; Result %d\n\n", time, result2)

	fmt.Printf("\tInput: %c\n", testData3)
	result3, time := algorithm.Bogosort(testData3)
	fmt.Printf("\tTime %s ; Result %c\n\n", time, result3)
}

func main() {
	arg.Argument_add("--bubble", "-bu", false, "Sort with bubblesort")
	arg.Argument_add("--bogo", "-bo", false, "Sort with bogosort")

	parsed := arg.Argument_parse()

	if len(parsed) == 0 {
		notify.Error("No argument was provided", "main.main()")
	}

	testData1 := []string{"a", "c", "b"}
	testData2 := []int{3, 5, 1, 0, -1, 2, 9, 35, 11}
	testData3 := []rune{'a', 'c', 'b'}

	for key := range parsed {
		switch key {
		case "--bubble", "-bu":
			checkBubblesort(testData1, testData2, testData3)

		case "--bogo", "-bo":
			checkBogosort(testData1, testData2, testData3)

		default:
			notify.Error(fmt.Sprintf("Unknown command %s", key), "main.main()")
		}
	}

}
