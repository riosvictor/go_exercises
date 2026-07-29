package main

import (
	"fmt"
	"time"
)

func main() {
	// If statement
	if true {
		fmt.Println("This is true")
	}

	// If else statement
	if false {
		fmt.Println("This is true")
	} else {
		fmt.Println("This is false")
	}

	// Switch statement
	switch 1 {
	case 1:
		fmt.Println("This is one")
	case 2:
		fmt.Println("This is two")
	default:
		fmt.Println("This is default")
	}

	// For loop
	fmt.Print("For loop: ")
	for i := 0; i <= 10; i++ {
		fmt.Print(i)
	}

	// While loop
	i := 0
	fmt.Print("\nWhile loop: ")
	for i < 10 {
		fmt.Print(i)
		i++
	}

	// Even numbers
	fmt.Print("\nEven numbers: ")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Print(i)
		}
	}

	//
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("\nIt's the weekend")
	default:
		fmt.Println("\nIt's a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
