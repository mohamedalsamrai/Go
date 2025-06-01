package lessons

import (
	"io"
	"os"
)
func Defer() {
	// Defer is used to delay the execution of a function until the surrounding function returns.
	// It is often used for cleanup tasks, such as closing files or releasing resources.

	defer func() {
		println("This will be executed at the end of the function.")
	}()

	println("This will be executed first.")

	// Example of using defer with a file operation
	file, err := os.Open("lessons/example.txt")
	if err != nil {
		println("Error opening file:", err)
		return
	}
	defer file.Close() // This will ensure the file is closed when the function exits
	data, err := io.ReadAll(file)
	if err != nil {	
		println("Error reading file:", err)
		return
	}
	println("File opened successfully data:", string(data))
}