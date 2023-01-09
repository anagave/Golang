// Assignment 5:

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("// 1. Using the for loop print 1 to 100 numbers")
	for i := 1; i <= 100; i++ {
		fmt.Println(i)
	}
	fmt.Println("**************************************************************************************************")

	fmt.Println("// 2. Create a for loop using this syntax\n//      for condition { }\n// print out the odd numbers in 1 to 50.")
	for i := 1; i <= 50; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}
	fmt.Println("**************************************************************************************************")

	fmt.Println("// 3. Create a for loop using this syntax\n//      for condition { }\n// print out the even numbers in 1 to 50.")
	for i := 1; i <= 50; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	fmt.Println("**************************************************************************************************")

	fmt.Println("// 4. Print out the quotient for each number between 50 and 105 when it is divided by 6.")
	for i := 50; i <= 105; i++ {
		fmt.Println(float64(i) / float64(6))
	}
	fmt.Println("**************************************************************************************************")

	fmt.Println("// 5. Read the user input.\n// If the input is equal to Golang tutorial print welcome else print end")
	fmt.Print("Enter your input for 5.")
	inputReader := bufio.NewReader(os.Stdin)
	input, _ := inputReader.ReadString('\n')
	input = strings.TrimSpace(input)
	fmt.Println(input)
	if input == "Golang tutorial" {
		fmt.Println("Welcome")
	} else {
		fmt.Println("End")
	}
	fmt.Println("**************************************************************************************************")

	fmt.Println("// 6. Write a program to print the numbers from 1 to 80.\n// But, for multiples of two print Golang instead of the number and for the multiples of four print tutorial.\n// For numbers which are multiples of both two and four print Golang tutorial.")
	for i := 1; i <= 80; i++ {
		if i%2 == 0 {
			if i%4 == 0 {
				fmt.Println("Golang tutorial")
			} else {
				fmt.Println("Golang")
			}
		} else if i%4 == 0 {
			fmt.Println("tutorial")
		} else {
			fmt.Println(i)
		}
	}
	fmt.Println("**************************************************************************************************")
}
