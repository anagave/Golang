// Assignment 6
// Create your own type “person” which will have an underlying type of “struct” so that it can store the following data:
// 	● first name
// 	● last name
// 	● favorite Country
// Create two VALUES of TYPE person.
// Print out the values, ranging over the elements in the slice which stores the favorite Countries.
// o/p Ex:
// First person struct {John Snow [USA Austrilia Italy]}
// Person First Name John
// Looping over the favourite Country
// 0 USA
// 1 Austrilia
// 2 Italy
// Second person struct {Tyrion Lannister [Austria Turkey Dubai]}
// Person First Name Tyrion
// Lannister
// Looping over the favorite Country
// 0 Austria
// 1 Turkey
// 2 Dubai
// 2. Take the code from the previous exercise, then store the values of type person in a map with the key of last name.
// Access each value in the map. Print out the values, ranging over the slice.
// 3. Create a new type: vehicle.
//       ○ The underlying type is a struct.
//       ○ The fields:
//                 ■ doors
//                 ■ color
// ● Create two new types: truck & sedan.
//       ○ The underlying type of each of these new types is a struct.
//       ○ Embed the “vehicle” type in both truck & sedan.
//       ○ Give truck the field “fourWheel” which will be set to bool.
//       ○ Give sedan the field “luxury” which will be set to bool. solution
// ● Using the vehicle, truck, and sedan structs:
//      ○ using a composite literal, create a value of type truck and assign values to the fields;
//      ○ using a composite literal, create a value of type sedan and assign values to the fields.
// ● Print out each of these values.
// ● Print out a single field from each of these values.
// 4.
// ● create a type SQUARE
// ● create a type CIRCLE
// ● attach a method to each that calculates AREA and returns it
//           ○ circle area= π r 2
//           ○ square area = L * W
// ● create a type SHAPE that defines an interface as anything that has the AREA method
// ● create a func INFO which takes type shape and then prints the area
// ● create a value of type square
// ● create a value of type circle
// ● use func info to print the area of square
// ● use func info to print the area of circle
// 5. write a program to count the occurrence of each word in a sentence and return a map containing word and number of occurrences of it in the given sentence

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type Person struct {
	fName    string
	lName    string
	fCountry []string
}

type Vehicle struct {
	doors int
	color string
}
type Truck struct {
	fourWheel bool
	Vehicle
}
type Sedan struct {
	luxury bool
	Vehicle
}

type Square struct {
	length float32
	width  float32
}
type Circle struct {
	radius float32
}
type Shape interface {
	Area() float32
}

func printPerson(personP []Person) {
	for i, j := range personP {
		if i == 0 {
			fmt.Println("First person struct ", j)
		} else {
			fmt.Println("Second  person struct ", j)
		}
		fmt.Println("Person First Name", j.fName)
		fmt.Println("Looping over the favourite Country")
		for i, j := range j.fCountry {
			fmt.Println(i, j)
		}
	}
}
func printPersonMap(personP []Person, personMap map[string]Person) {
	for _, a := range personP {
		personMap[a.lName] = a
		fmt.Println(personMap[a.lName])
	}

}

func info(in Shape) {
	fmt.Println(in)
	fmt.Println("Area: ", in.Area())
}

func (s Square) Area() float32 {
	return s.length * s.width
}
func (c Circle) Area() float32 {
	return math.Pi * (c.radius * c.radius)
}

func main() {
	personP := []Person{
		{
			fName: "Achyutha",
			lName: "Santhoshi",
			fCountry: []string{
				"USA",
				"UK",
				"INDIA",
			},
		},
		{
			fName: "Divyanjali",
			lName: "Rasakonda",
			fCountry: []string{
				"INDIA",
				"AUS",
				"UAE",
			},
		},
	}
	printPerson(personP)

	personMap := map[string]Person{}
	printPersonMap(personP, personMap)

	fordTruck := Truck{
		fourWheel: true,
		Vehicle: Vehicle{
			doors: 4,
			color: "Black",
		},
	}
	CamrySedan := Sedan{
		luxury: true,
		Vehicle: Vehicle{
			doors: 2,
			color: "White",
		},
	}
	fmt.Println(fordTruck)
	fmt.Println(fordTruck.fourWheel)
	fmt.Println(fordTruck.doors)
	fmt.Println(fordTruck.color)

	fmt.Println(CamrySedan)
	fmt.Println(CamrySedan.luxury)
	fmt.Println(CamrySedan.doors)
	fmt.Println(CamrySedan.color)

	info(Circle{radius: 3.6})
	info(Square{length: 2.1, width: 3.2})

	fmt.Println("Enter any sentence")
	inputReader := bufio.NewReader(os.Stdin)
	inputSentence, _ := inputReader.ReadString('\n')
	inputSentence = strings.TrimSpace(inputSentence)
	fmt.Println(inputSentence)

	var S []string
	S = strings.Split(inputSentence, " ")
	freqMap := make(map[string]int)
	for i := 0; i < len(S); i++ {
		if freqMap[S[i]] == 0 {
			freqMap[S[i]] = 1
		} else {
			v := freqMap[S[i]]
			freqMap[S[i]] = v + 1
		}
	}
	for key, value := range freqMap {
		fmt.Println(key, value)
	}

}
