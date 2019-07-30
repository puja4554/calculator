package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	var num1, num2 int

	fmt.Println("Which operation you want to perform : +/-/*/%")

	reader := bufio.NewReader(os.Stdin)
	sign, _ := reader.ReadString('\n')

	switch sign {
	case "+":
		fmt.Println("Enter the numbers")

		if _, err := fmt.Scanf("%d %d", &num1, &num2); err != nil {
			log.Print("  Scan for num1 & num2 failed, due to ", err)
			return
		}

		fmt.Println(num1 + num2)

	case "-":
		fmt.Println("Enter the numbers")

		if _, err := fmt.Scanf("%d %d", &num1, &num2); err != nil {
			log.Print("  Scan for num1 & num2 failed, due to ", err)
			return
		}

		fmt.Println(num1 - num2)
	case "*":
		fmt.Println("Enter the numbers")

		if _, err := fmt.Scanf("%d %d", &num1, &num2); err != nil {
			log.Print("  Scan for num1 & num2 failed, due to ", err)
			return
		}

		fmt.Println(num1 * num2)
	case "%":
		fmt.Println("Enter the numbers")

		if _, err := fmt.Scanf("%d %d", &num1, &num2); err != nil {
			log.Print("  Scan for num1 & num2 failed, due to ", err)
			return
		}

		fmt.Println(num1 % num2)

	}

}
