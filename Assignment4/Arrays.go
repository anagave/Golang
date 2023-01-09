// Assignment 4:


package main

import (
	"fmt"
	"reflect"
)

func updateThirdElement(p *[4]string) {
	(*p)[3] = "Texas"
}

func ADDITION(x *int, y *int) *int {
	add := *x + *y
	return &add
}
func SUBSTRACTION(x *int, y *int) *int {
	sub := *x - *y
	return &sub
}
func MULTIPLICATION(x *int, y *int) *int {
	mul := *x * *y
	return &mul
}
func DIVISION(x *int, y *int) (*int, *int) {
	q := *x / *y
	r := *x % *y
	return &q, &r
}

func main() {
	fmt.Println("// 1.Create an ARRAY which holds 5 VALUES of TYPE int\n// 	● assign VALUES to each index position.\n// 	● Using format printing, print out the TYPE of the array")
	var arr [5]int
	arr[3] = 1
	arr[1] = 2
	arr[2] = 3
	arr[0] = 0
	arr[4] = 5
	fmt.Println(reflect.TypeOf(arr))
	fmt.Println("************************************************************************************************************************")

	fmt.Println("// 2.Create a SLICE of TYPE int\n// ● assign 10 VALUES\n//  ● Using format printing, print out the TYPE of the slice")
	arr2 := [10]int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	var slice []int = arr2[0:10]
	fmt.Println(reflect.TypeOf(slice))
	fmt.Println("************************************************************************************************************************")

	fmt.Println("// 3.Using the code from the previous example, use SLICING to create the following new slices which are then printed:\n// 	  ● [42 43 44 45 46]\n// 	  ● [47 48 49 50 51]\n// 	  ● [44 45 46 47 48]\n// 	  ● [43 44 45 46 47]")
	fmt.Println(arr2[:5])
	fmt.Println(arr2[5:])
	fmt.Println(arr2[2:7])
	fmt.Println(arr2[1:6])
	fmt.Println("************************************************************************************************************************")

	fmt.Println("// 4.Start with this slice\n// ○ x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}\n// ● append to that slice this value\n//○ 52\n//● print out the slice in ONE STATEMENT append to that slice these values\n// 	 ○ 53\n// 	 ○ 54\n// 	 ○ 55\n//● print out the slice\n//● append to the x slice the below slice\n//○ y := []int{56, 57, 58, 59, 60}\n//● print out the slice")
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)
	fmt.Println(x)
	x = append(x, 53, 54, 55)
	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println(x)
	fmt.Println("************************************************************************************************************************")

	fmt.Println("// 5. Write a program to pass a pointer to an array as an argument to the function\n//        a. Create an array of size 4, data type string\n//        b. Create a function with name updateThirdElement and update the value of 3rd index to Texas\n//        c. As an input to the function updateThirdElement pass pointer to an array to function updatearray")
	arrS := [4]string{"Washington", "Oklahoma", "Florida", "NewJersey"}
	fmt.Println(arrS)
	updateThirdElement(&arrS)
	fmt.Println(arrS)
	fmt.Println("************************************************************************************************************************")

	fmt.Println("6. Create calculator app using pointers")
	var x1 int
	var y1 int
	fmt.Print("Enter the first value : ")
	fmt.Scanf("%d", &x1)
	fmt.Print("Enter the second value : ")
	fmt.Scanf("%d", &y1)
	var abc int
	fmt.Print("Choose an option: 1. ADDITION 2. SUBSTRACTION 3. MULTIPLICATION 4. DIVISION : ")
	fmt.Scanf("%d", &abc)

	switch abc {
	case 1:
		add := ADDITION(&x1, &y1)
		fmt.Println(*add)
	case 2:
		sub := SUBSTRACTION(&x1, &y1)
		fmt.Println(*sub)
	case 3:
		mul := MULTIPLICATION(&x1, &y1)
		fmt.Println(*mul)
	case 4:
		q, r := DIVISION(&x1, &y1)
		fmt.Println(*q, *r)
	default:
		fmt.Println("Choose valid option")
	}
	fmt.Println("************************************************************************************************************************")
}
