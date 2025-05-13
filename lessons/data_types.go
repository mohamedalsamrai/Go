package lessons

import "fmt"

var year = 2000 // it's globle variable
func DataTypesLesson() {
	age := 20                    // intger
	name := "Mohammed Al-Samrai" // string
	type Employee struct {
		firstName string
		lastName  string
		age       int
	} // struct it is like calss

	fmt.Println(name, age)

	// array
	// first way
	// var arrayName [size] type
	var numbers [5]int8
	fmt.Println(numbers)
	numbers[0] = 5
	fmt.Println("after add 5 to index 0")
	fmt.Println(numbers)

	// second way
	// arrayName := [size] type  {values}
	// if you don't need put values just put empty {}
	fruits := [3]string{}
	fmt.Println(fruits)
	fruits[0] = "apple"
	fmt.Println(fruits)

	// Matrix Array
	// var matrixName [row count][row size] type
	matrix := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("Matrix Array:", matrix)

}
