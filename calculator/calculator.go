package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type calc struct{} // for create stricutures

func (calc) operate(value string, operator string) int { // receiver function, call gets this method for his structure- Like method of a class

	cleanInput := strings.Split(value, operator)
	operator1 := parseTonInt(cleanInput[0])
	operator2 := parseTonInt(cleanInput[1])

	switch operator {
	
	case "+":
		fmt.Println(operator1 + operator2)
		return operator1 + operator2
	case "-":
		fmt.Println(operator1 - operator2)
		return operator1 - operator2
	case "*":
		fmt.Println(operator1 * operator2)
		return operator1 * operator2
	case "/":
		fmt.Println(operator1 / operator2)
		return operator1 / operator2
	default:
		fmt.Println("Operator not supported")
		return 0
	}
}

func parseTonInt(enter string) int {
	operator1, err1 := strconv.Atoi(enter) // _ for catching errors 
	if err1 != nil {// nil<=>null
		fmt.Println(err1)
	}

	return operator1
}

func readInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	operation := scanner.Text()

	return operation
}

// delete aditional settings in settings extension go extension

func main() {
	
	input := readInput()
	operator:= readInput()

	calculator := calc{}
	calculator.operate(input, operator)

}
