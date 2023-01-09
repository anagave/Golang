package CalculatorFunctions

func ADDITION(x int, y int) int {
	return x + y
}

func SUBSTRACTION(x int, y int) int {
	return x - y
}

func MULTIPLICATION(x int, y int) int {
	return x * y
}

func DIVISION(x int, y int) (int, int) {
	q := x / y
	r := x % y
	return q, r
}
