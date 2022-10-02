/*
Требуется найти в бинарном векторе самую длинную последовательность
единиц и вывести её длину.
*/

package main

import (
	"fmt"
)

func fillInSlice(array []int) []int {

	var slice []int
	var N int

	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		var n int

		fmt.Scan(&n)

		slice = append(slice, n)
	}
	return slice
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var slice1 []int

	slice1 = fillInSlice(slice1)

	counter := 0
	best := 0

	for _, el := range slice1 {

		if el == 1 {
			counter++
		} else if el == 0 {
			best = max(counter, best)
			counter = 0
		}
		best = max(counter, best)
	}
	fmt.Println(best)
}