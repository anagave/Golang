// Assignment 3:
// You have to create a calculator
// Requirements
// 1) There should be at least 1 package
// 2) You need to work on exporting the functionalities outside the package
// 3) The app should perform the following
//   A) add
//   B) sum
//   C) mul
//   D) div - should
// Return the quotient and the remainder both

// Remember you will be judged based on the correctness, naming convention and packaging your app

package main

import (
	"Assignment3/CalculatorFunctions"
	"fmt"
)

func main() {
	var x int
	var y int
	var option int
	fmt.Print("Enter the first value : ")
	fmt.Scanf("%d", &x)
	fmt.Print("Enter the second value : ")
	fmt.Scanf("%d", &y)
	fmt.Print("Choose an option: 1. ADDITION 2. SUBSTRACTION 3. MULTIPLICATION 4. DIVISION : ")
	fmt.Scanf("%d", &option)

	switch option {
	case 1:
		fmt.Println(CalculatorFunctions.ADDITION(x, y))
	case 2:
		fmt.Println(CalculatorFunctions.SUBSTRACTION(x, y))
	case 3:
		fmt.Println(CalculatorFunctions.MULTIPLICATION(x, y))
	case 4:
		fmt.Println(CalculatorFunctions.DIVISION(x, y))
	default:
		fmt.Println("Choose valid option")
	}
}
