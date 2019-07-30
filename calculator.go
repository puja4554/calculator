package main

import (
	"fmt"
	"strconv"
)

func main() {

	var inp1, inp2 interface{}

	var num1, num2 int

	var err error
	sign := ""
	fmt.Println("Which operation you want to perform : +/-/*/%")
	fmt.Scan(&sign)

	switch sign {
	case "+":
		fmt.Println("Enter the numbers ")

		fmt.Scan(&inp1, &inp2)

		if num1, err = strconv.Atoi(inp1.(string)); err != nil {
			fmt.Println("First input parameter must be integer")
			return
		}

		if num2, err = strconv.Atoi(inp2.(string)); err != nil {
			fmt.Println("Second input parameter must be integer")
			return
		}

		fmt.Println(num1 + num2)

	case "-":
		fmt.Println("Enter the numbers")

		fmt.Scan(&num1, &num2)

		fmt.Println(num1 - num2)
	case "*":
		fmt.Println("Enter the numbers")

		fmt.Scan(&num1, &num2)

		fmt.Println(num1 * num2)
	case "%":
		fmt.Println("Enter the numbers")

		fmt.Scan(&num1, &num2)

		fmt.Println(num1 % num2)

	default:
		fmt.Println("Wrong operation choice")

	}

}
