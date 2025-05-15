package fundamentals

import "fmt"


func Range() {
	// 1. Range with Array
	numbers := [5]int{1, 2, 3, 4, 5}
	for i, v := range numbers {
		fmt.Println("index:", i, "value:", v)
	}

	// 2. Range with Slice
	slice := []int{1, 2, 3, 4, 5}
	for i, v := range slice {
		fmt.Println("index:", i, "value:", v)
	}

	// 3. Range with Map
	map1 := map[string]int{"a": 1, "b": 2}
	for k, v := range map1 {
		fmt.Println("key:", k, "value:", v)
	}

	// 4. Range with String
	str := "hello"
	for i, v := range str {
		fmt.Println("index:", i, "value:", string(v))
	}
}