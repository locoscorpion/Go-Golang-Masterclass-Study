package main

import (
	"fmt"
	"strings"
)

/*
Simple CLI calculator

	Objective: parse numeric input and perform arithmetic.
	What to do: support add/subtract/multiply/divide on two numbers.
	Time: 30–60 min. Concepts: types, error handling, basic parsing.
*/

// Define ANSI color constants
const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Bold   = "\033[1m"
)

// error handling
const (
	division       = "Division"
	divisionErrMsg = "Input by zero is not allowed"
)

type MathError struct {
	Operation string
	InputA    int
	InputB    int
	Message   string
}

func (e MathError) Error() string {
	var inputs []string
	if e.Operation == division {
		inputs = append(inputs, fmt.Sprintf("a=%d", e.InputA))
		inputs = append(inputs, fmt.Sprintf("b=%d", e.InputB))
	}
	return fmt.Sprintf("Math error in %s (%s): %s", e.Operation, strings.Join(inputs, ","), e.Message)
}

//Aritmetic operations

func sum(numbers ...float64) float64 {
	defer fmt.Println("Sum has finished.")
	total := 0.0
	for _, n := range numbers {
		total = total + n
	}
	return total
}

func substract(numbers ...float64) float64 {
	defer fmt.Println("Substraction has finished.")
	total := 0.0
	for _, n := range numbers {
		total = total + n
	}
	return total
}

func multiply(numbers ...float64) float64 {
	defer fmt.Println("Multiplication has finished.")
	total := 1.0
	for _, n := range numbers {
		total = total * n
	}
	return total
}

func SafeDivision(a, b int) (int, error) {
	if b == 0 {
		return 0, &MathError{
			Operation: division,
			InputA:    a,
			InputB:    b,
			Message:   divisionErrMsg,
		}
	}
	return a / b, nil
}

//CLI menu and choices

func menuDisplay() {
	fmt.Println(Bold + "-------------------CHOOSE AN OPERATION-------------------")
	fmt.Println("1. Sum a series of numbers")
	fmt.Println("2. Substract a series of numbers")
	fmt.Println("3. Multiply a series of numbers")
	fmt.Println("4. Divide two numbers")
	fmt.Println("---------------------------------------------------------" + Reset)
}

// Main
func main() {
	var optMenu string

	exitsuccess := false

	for !exitsuccess {
		menuDisplay()
		fmt.Printf("Input an option: ")
		fmt.Scanf("%s", &optMenu)
		switch optMenu {
		case "1":
			sum()
			exitsuccess = true

		case "2":
			substract()
			exitsuccess = true

		case "3":
			multiply()
			exitsuccess = true

		case "4":
			SafeDivision()
			exitsuccess = true

		default:
			fmt.Println(Red + "Invalid option try again." + Reset)
			optMenu = ""
			exitsuccess = false
		}

	}
}
