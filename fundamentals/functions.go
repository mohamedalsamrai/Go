package fundamentals

import "fmt"

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
