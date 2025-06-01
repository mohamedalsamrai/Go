package lessons
import (
	"fmt"	
)
func Panic() {	
	// Panic is used to stop the normal execution of a function and start panicking.
	// It is often used for unrecoverable errors or unexpected situations.
	
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
			// You can handle the panic here, log it, or take other actions
			fmt.Println("Panic handled gracefully, continuing execution.")
		}
	}()

	fmt.Println("Starting the function")
	panic("Something went wrong!") // This will cause a panic
	fmt.Println("This line will not be executed") // This line will not be executed
}