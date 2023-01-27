package main

import (
	"fmt"
	"math/rand"
)

func quickSort(in []int) []int {
	if len(in) < 2 {
		return in
	}
	pivot := rand.Int() % len(in)
	left := 0
	right := len(in) - 1
	in[pivot], in[right] = in[right], in[pivot]

	for i := range in {
		if in[i] < in[right] {
			in[left], in[i] = in[i], in[left]
			left++
		}
	}

	in[left], in[right] = in[right], in[left]
	quickSort(in[:left])
	quickSort(in[left+1:])
	return in

}
func main() {
	arr := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 23, 25, 26, 28, 29, 32, 1, 4, 6, 3, 11, 9, 10, 12, 13, 16, 122, 23, 456, 23, 10, 203, 202}
	var slice []int = arr[2:27]
	fmt.Print(quickSort(slice))
}
