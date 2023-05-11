package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 8 // short initialization

	fmt.Println(x)

	val, err := strconv.ParseInt("NaN", 0, 64) // int parser. (value, base, size)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(val)
	}

	// Slices <=> arrays y maps
	m := make(map[int]string)

	m[0] = "11"

	s := []int{1, 2, 3}

	s = append(s, 15) // adding new elements

	for index, value := range s { // range allows iterate over s
		fmt.Println(index)
		fmt.Println(value)
	}

}
