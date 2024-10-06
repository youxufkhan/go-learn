package main

import (
	"errors"
	"fmt"
)

func main() {
	var printValue string = "Hello World"
	printMe(printValue)

	var numerator int = 11
	var denominator int = 2
	var result int = intDivision(numerator, denominator)
	fmt.Println(result)

	var result2, remainder, err = intDivisionWithRemainder(numerator, denominator)
	if err != nil {
		fmt.Printf(err.Error())
	}
	fmt.Printf("The result of integer division is %v with remainder %v", result2, remainder)

	denominator = 0
	var result3, remainder2, err2 = intDivisionWithRemainder(numerator, denominator)

	if err2 != nil {
		fmt.Printf(err2.Error())
	} else if remainder2 == 0 {
		fmt.Printf("The result of integer division is %v", result3)

	} else {
		fmt.Printf("The result of integer division is %v with remainder %v", result3, remainder2)
	}

	//below switch statement is exactly the same as above if else statements
	switch {
	case err2 != nil:
		fmt.Printf(err2.Error())
	case remainder == 0:
		fmt.Printf("The result of integer division is %v", result3)
	default:
		fmt.Printf("The result of integer division is %v with remainder %v", result3, remainder2)
	}

	//conditional switch statement
	switch remainder {
	case 0:
		fmt.Println("The division was exact")
	case 1, 2:
		fmt.Println("The division was close")
	default:
		fmt.Println("The division was not close")
	}
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) int {
	var result int = numerator / denominator
	return result
}

func intDivisionWithRemainder(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("Cannot divide by Zero")
		return 0, 0, err
	}
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err
}
