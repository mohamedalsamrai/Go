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

	// maps
	// first way for declere map[key type ] value type
	map1 := make(map[string]int)
	fmt.Println(map1)
	// second way map name := map[key type] value type {"key1":value ......}
	map2 := map[string]int{
		"k1": 88,
		"k2": 62,
		"k3": 33,
	}
	fmt.Println(map2)
	// third way var mapName map[key type] value type this way gives nil map
	var map3 map[string]string
	fmt.Println(map3)
	// map3["key1"] = "value1" // this will give error because map3 is nil
	// map3 = make(map[string]string)
	// map3["key1"] = "value1" // now this will work


}
