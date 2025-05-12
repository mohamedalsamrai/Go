package lessons

import "fmt"

func OperatorsLesson() {
	// Logical operators
	fmt.Println("Logical operators:")
	fmt.Println("NOT (!true):", !true)
	fmt.Println("OR (false || true):", false || true)
	fmt.Println("AND (true && false):", true && false)

	a := 5 // 0101
	b := 3 // 0011
	fmt.Println("\nBitwise operators")
	fmt.Println("a =", a)
	fmt.Println("b =", b)

	fmt.Println("a & b =", a&b)   // 0101 & 0011 = 0001 => 1  
	fmt.Println("a | b =", a|b)   // 0101 | 0011 = 0111 => 7
	fmt.Println("a ^ b =", a^b)   // 0101 ^ 0011 = 0110 => 6
	fmt.Println("^a =", ^a)       // ^0101 = 1010 => -6 (2 complement)
	fmt.Println("a << 1 =", a<<1) // 0101 << 1 = 1010 => 10
	fmt.Println("a >> 1 =", a>>1) // 0101 >> 1 = 0010 => 2

	// Comparison operators
	fmt.Println("\nComparison operators")
	fmt.Println("a == b:", a == b) // false
	fmt.Println("a != b:", a != b) // true
	fmt.Println("a > b:", a > b)   // true
	fmt.Println("a < b:", a < b)   // false
	fmt.Println("a >= b:", a >= b) // true
	fmt.Println("a <= b:", a <= b) // false
}
