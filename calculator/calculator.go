package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

// delete aditional settings in settings extension go extension

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	operation := scanner.Text()
	values := strings.Split(operation, "+")
	fmt.Println(values)
	fmt.Println(operation)
	fmt.Println(values)

	operator1, err1 := strconv.Atoi(values[0]) // _ for catching errors 
	if err1 != nil {// nil<=>null
		fmt.Println(err1)
	}

	operator2, _ := strconv.Atoi(values[1])

	fmt.Println(operator1 + operator2)

}
