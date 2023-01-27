package main

import "fmt"

func fibonnaciSeries(number int) {
	x := 0
	y := 1
	z := 0
	for i := 1; i <= number; i++ {
		fmt.Printf(" %d  ", x)
		z = x + y
		x = y
		y = z
	}
	fmt.Println()
}

func main() {
	fibonnaciSeries(10)
	fibonnaciSeries(15)
	fibonnaciSeries(30)
}
