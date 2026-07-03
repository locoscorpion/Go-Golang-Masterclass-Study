package main

import "fmt"

func main() {
	var name string
	var num1, num2, sum int
	name = "no name"
	fmt.Printf("\x1b[1m\x1b[32m----------------------Greeter and calculator-------------------------------\x1b[0m\n") //using ANSI colors for print format
	fmt.Printf("Welcome to my first program, please introduce your name: ")
	fmt.Scanln(&name)
	fmt.Printf("\nNow, introduce the first number\n")
	fmt.Scanln(&num1)
	fmt.Printf("Now, introduce the second number\n")
	fmt.Scanln(&num2)
	sum = num1 + num2
	fmt.Printf("Thanks %v, The sum of the two numbers is: %d\n", name, sum)
}
