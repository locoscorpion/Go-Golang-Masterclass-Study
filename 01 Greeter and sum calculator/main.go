package main

import (
	"errors"
	"fmt"
)

/*
Task: Ask for the user's name and two numbers. Print a personalized greeting and perform addition/subtraction.

Go Concepts: fmt.Scanln, strconv.Atoi, error handling basics (if err != nil).
*/
func main() {
	var selectOp, name string

	var num1, num2, sum, subst int
	name = "no name"
	selectOp = ""
	fmt.Printf("\x1b[1m\x1b[32m----------------------Greeter and calculator-------------------------------\x1b[0m\n") //using ANSI colors for print format
	fmt.Printf("Welcome to my first program, please introduce your name: ")
	fmt.Scanln(&name)
	fmt.Printf("\nNow, introduce the first number\n")
	fmt.Scanln(&num1)
	fmt.Printf("Now, introduce the second number\n")
	fmt.Scanln(&num2)
	fmt.Printf("------------- CHOOSE AN OPERATION --------------------\n")
	fmt.Printf("1. Sum both numbers\n")
	fmt.Printf("2. Subtract both numbers\n")
	fmt.Printf("------------------------------------------------------\n")
	fmt.Printf("Options (input 1 - 2): ")
	fmt.Scanln(&selectOp)

	if selectOp == "1" || selectOp == "2" {
		fmt.Printf("Thanks, %s !!! Now...\n", name)
		switch selectOp {
		case "1":
			sum = num1 + num2
			fmt.Printf("The sum result of %d+%d is: %d\n", num1, num2, sum)
		case "2":
			subst = num1 - num2
			fmt.Printf("The sum result of %d-%d is: %d\n", num1, num2, subst)
		}
	} else {
		selectOpErr := errors.New("Wrong option!")
		fmt.Println(selectOpErr)
	}
}
