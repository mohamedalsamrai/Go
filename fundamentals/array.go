package fundamentals

import "fmt"

func Array() {
	fruits := [3]string{"Apple", "Orange", "Banana"}
	// to read array
	// first way
	for i := 0; i < len(fruits); i++ {
		fmt.Println("Element at index ", i, ":", fruits[i])
	}
	// second way
	for index, value := range fruits {
		fmt.Printf("Index %d Value %s \n", index, value)

	}
	fmt.Println()
	for _, value := range fruits {
		fmt.Printf("Value %s \n", value)

	} //_ is blank identifier in go mean you don't want use value return to you from any way

	fmt.Println()

	// Comparing array
	array1 := [3]int{1, 2, 3}
	array2 := [3]int{1, 2, 3}
	fmt.Println("Array1:", array1)
	fmt.Println("Array2:", array2)
	fmt.Println("Array1 == Array2:", array1 == array2)
	array1[0] = 5
	fmt.Println("Array1 after change:", array1)
	fmt.Println("Array2:", array2)
	fmt.Println("Array1 == Array2:", array1 == array2)
}
