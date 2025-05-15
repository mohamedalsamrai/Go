package lessons

import (
	"fmt"
	"slices"
)

/*
slice is dynamic array with custom lenght and capacity
*/
func SlicesLesson() {
	slice1 := []int{1, 2, 3} //slice with capacity = lenght
	fmt.Println("lenght slice1 is:", len(slice1), "capacity slice1 is:", cap(slice1))
	slice1 = append(slice1, 4, 5, 6)
	fmt.Println("lenght slice1 is:", len(slice1), "capacity slice1 is:", cap(slice1))
	array1 := [6]int{2, 4, 6, 8, 6, 5}
	slice2 := array1[1:3] // give slice from array capacity it will by array lenght - 1
	fmt.Println("lenght slice2 is:", len(slice2), "capacity slice2 is:", cap(slice2))
	slice2 = append(slice2, 5, 8, 5, 5) // if add elements make lenght > capacity the capacity will *2
	fmt.Println("lenght slice2 is:", len(slice2), "capacity slice2 is:", cap(slice2))

	// you can creat slice by make function
	slice3 := make([]int, 7) // lenght 7 capaciry 7
	copy(slice3, slice1)
	fmt.Println("slice3 is:", slice3)
	for i, v := range slice3 {
		fmt.Println(i, v)
	}
	// check slices if equal
	if slices.Equal(slice1, slice3) {
		fmt.Println("is equal")
	} else {
		fmt.Println("is not equal")
	}
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		twoD[i] = make([]int, 3)
		for j := 0; j < 3; j++ {
			twoD[i][j] = i * j
		}

	}
	fmt.Println(twoD)
	fmt.Println("the lenght is:", len(twoD), "capacity is:", cap(twoD))
}
