package main

import (
	"fmt"
	"strconv"
)

func add(i int, j int) int {
	return i + j
}

func sub(i int, j int) int {
	return i - j
}

func mult(i int, j int) int {
	return i * j
}

func main() {
	var funcMap = map[string]func(int, int) int{
		"+": add,
		"-": sub,
		"*": mult}

	expressions := [][]string{[]string{"2", "8", "3"},
		[]string{"2", "o", "3", "4"},
		[]string{"five", "+", "3"},
	}

	for _, expression := range expressions {
		if len(expression) != 3 {
			fmt.Println("Incorrect length")
			continue
		}
		p1, err := strconv.Atoi(expression[0])
		if err != nil {
			fmt.Println(err)
			continue
		}
		p2, err := strconv.Atoi(expression[2])

		if err != nil {
			fmt.Println(err)
			continue
		}
		op1 := expression[1]
		opMap, ok := funcMap[op1]
		if !ok {
			fmt.Println("incorrect operator")
			continue
		}

		fmt.Println(p1, op1, p2)
		result := opMap(p1, p2)
		fmt.Println(result)
	}

}
