package basics

import (
	"fmt"
	"math/rand"
	"time"
)

func GussingGame() {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	target := random.Intn(100) + 1

	fmt.Println("Welcome to gussing game ")
	fmt.Println("I have chosen a number between 1 to 100")
	fmt.Println("Can you guss what it is ?")
	var gussNum int
	for {
		fmt.Println("Enter your guss: ")
		fmt.Scanln(&gussNum)
		if gussNum == target {
			fmt.Println("Congratulations You gussed the correct number")
			break
		} else if gussNum < target {
			fmt.Println("Too low Try gussing higher number ")
		} else {
			fmt.Println("Too high Try gussing lower number ")
		}
	}
}
