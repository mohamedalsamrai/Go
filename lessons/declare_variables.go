/*
this lesson discribs ways to declare variables
*/

package lessons

import (
	"fmt"
)

func DeclareVariablesLesson() {
	/* basic declaration using var keyword + variable name + type
	declarea variable without immediately initialising it */
	var massage1 string //value ""

	var massage2 string = "Hello World" //Declaration with initialisation

	/* Type inference Go can infer the type based on the provided value when initialising a variable */
	var massage3 = "Mohammed"

	/* Short variable declaration := operator it use to declare and initialise a variable */
	message4 := "Ahmed"

	// Multiple variable declaration
	var x1, y1, z1 int = 1, 2, 3
	var x2, y2, z2 = 10, 20, 30
	x3, y3, z3 := 100, 200, 300

	// Grouped declaration
	var (
		a int
		b float64
		c string = "Hello, grouped declaration!"
	)
    
	// Constants
	const Pi = 3.14159

	fmt.Printf("massage 1:%s\nmassage 2:%s\nmassage 3:%s\nmassage 4:%s\n", massage1, massage2, massage3, message4)
	fmt.Printf("x1:%d y1:%d z1:%d\nx2:%d y2:%d z2:%d\nx3:%d y3:%d z3:%d\n", x1, y1, z1, x2, y2, z2, x3, y3, z3)
	fmt.Printf("a:%d b:%f c:%s\n", a, b, c)
	fmt.Print("Pi:",Pi,"\n")
}
