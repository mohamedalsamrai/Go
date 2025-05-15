package fundamentals

import (
	"fmt"
	"maps"
)

func Maps() {
	map1 := make(map[string]int)
	map2 := map[string]int{
		"k1": 88,
		"k2": 62,
		"k3": 33,
	}
	map1["key1"] = 2
	map1["key2"] = 4
	map1["key3"] = 6
	fmt.Println(map1)
	fmt.Println(map2)

	//  delete key
	delete(map1, "key1")
	fmt.Println(map1)

	//  clear all value
	clear(map1)
	fmt.Println(map1)

	//  check if key not include in map
	value, isInclude := map1["jibjh"]
	fmt.Println(value)
	fmt.Println(isInclude)

	//check if maps equal
	fmt.Println(maps.Equal(map1, map2))

	for k, v := range map2 {
		fmt.Println("key is:", k, "value is:", v)
	}
}
