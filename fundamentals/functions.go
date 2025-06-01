package fundamentals

import (
	"errors"
	"fmt"
)

func Functions() {
	/*
		func <name> <parameters> <return type> {
		code
		 return <value>
		 }
	*/
	// private function start with small letter
	// public function start with capital letter

	sum := add(1, 2)
	fmt.Println("sum is:", sum)
	sub := func() int {
		return 1 - 2
	}() // anonymous function
	sub1 := func(a, b int) int {

		return a - b
	} //sub1 is function variable
	x := sub1(2, 1) // call function variable
	fmt.Println("sub is:", sub)
	fmt.Println("sub1 is:", x)
	operation := add
	fmt.Println("operation is:", operation(2, 3))
	fmt.Println("aplay is:", aplayAdd(1, 2, add))

	addFunc := returnAdd()
	fmt.Println("addFunc is:", addFunc(1, 2))
	d, r := divide(10, 3)
	fmt.Println("divide result is:", d, "remainder is:", r)
	result, err := compare(10, 10)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Comparison result:", result)
	}
	fmt.Println("Sum of variadic numbers is:", sumVariadic(1, 2, 3, 4, 5))
	fmt.Println("Add with variadic numbers is:", addVariadic(10, 1, 2, 3, 4, 5))
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Sum of slice is:", sumVariadic(numbers...)) // using ellipsis operator to pass slice as variadic parameters

}
func add(a, b int) int {
	return a + b

}

// function that take a function as parameter
func aplayAdd(x int, y int, f func(int, int) int) int {
	result := f(x, y)
	return result
}

// function that return a function
func returnAdd() func(int, int) int {
	return func(x, y int) int {
		return x + y
	}
}

// multiple return values
// func FunctionName(paerameters) (returnType1, returnType2) {
// return value1, value2
// }
func divide(a, b int) (quotient int, remainder int) {
	if b == 0 {
		return 0, 0 // return zero values if b is zero
	}
	quotient = a / b
	remainder = a % b
	return quotient, remainder // return quotient and remainder
}

func compare(a, b int) (string, error) {
	if a > b {
		return "a is greater than b", nil
	} else if a < b {
		return "b is greater than a", nil
	} else {
		return "", errors.New("unable to compare, wich is greater")
	}
}

// ... Ellipsis operator
// that lists a variable number of parameters
func sumVariadic(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func addVariadic(a int, nums ...int) int {
	sum := a
	for _, num := range nums {
		sum += num
	}
	return sum
}
