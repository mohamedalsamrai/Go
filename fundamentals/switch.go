package fundamentals

import "fmt"

func Switch() {

	/*
			 Switch statement
			switch var {
			case value of var it must by same type
			   to do if var value == case value

			you can put multi cases
		    case 1,2,3,4:

			}
			   whin var value == case value it will execute to do code and break

			   if all cases value != var value it will execute default to do code
			   default:
				to do code

	*/
	fmt.Println("\nSwitch statement:")
	var day int
	fmt.Println("Enter a number (1-7) for the day of the week:")
	fmt.Scanln(&day)
	switch day {
	case 1:
		fmt.Println("Sunday")
	case 2:
		fmt.Println("Monday")
	case 3:
		fmt.Println("Tuesday")
	case 4:
		fmt.Println("Wednesday")
	case 5:
		fmt.Println("Thursday")
	case 6:
		fmt.Println("Friday")

	case 7:
		fmt.Println("Saturday")
	default:
		fmt.Println("Invalid day")
	}

	switch day {
	case 1, 2, 3, 4, 5: // Weekdays
		fmt.Println("It's a weekday.")
	case 6, 7: // Weekend
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("Invalid day")
	}

	switch {
	case day >= 1 && day <= 5: // case with condition
		fmt.Println("It's a weekday.")
	case day == 6 || day == 7:
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("Invalid day")
	}

	switch day {
	case 1:
		fmt.Println("Sunday")
		fallthrough // move to second case and execute its to do code
	case 2:
		fmt.Println("Monday") // it will execute

	case 3:
		fmt.Println("Tuesday")

	case 4:
		fmt.Println("Wednesday")

	case 5:
		fmt.Println("Thursday")

	case 6:
		fmt.Println("Friday")

	case 7:
		fmt.Println("Saturday")
	default:
		fmt.Println("Invalid day")
	}

	fmt.Println("\nType Check Example:")

	checkType(42) // Example usage of checkType
	checkType("hello")
	checkType(true)
	checkType(3.14)
}

func checkType(value interface{}) {
	// type switch
	switch v := value.(type) {
	case int:
		fmt.Printf("Value is an int: %d\n", v)
		//fallthrough in type switch can not use fallthrough
	case string:
		fmt.Printf("Value is a string: %s\n", v)
	case bool:
		fmt.Printf("Value is a bool: %t\n", v)
	default:
		fmt.Printf("Unknown type: %T\n", v)
	}
}
