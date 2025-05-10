package lessons

import "fmt"

func FandamentalLesson() {
	//for loop

	// iteration over a range
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}
	fmt.Print("\n \n")
	// iterate over collection
	numbers := []int{1, 2, 3, 4, 5, 61}
	for index, value := range numbers {
		fmt.Printf("Index: %d Value: %d\n", index, value)
	}
	// break & continue
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue //skip
		}
		fmt.Println(i)
		if i == 5 {
			break //stop
		}
	}

	fmt.Print("\n \n")

	// for in renge like python
	for i := range 10 {
		fmt.Println(i)
	}
}
