package main

import (
	"fmt"
	"github.com/Rakanixu/factorial"
	"strconv"
)

func main() {
	var input int64
	var result int64

	fmt.Println("------------Calculate factorial------------")
	fmt.Print("Enter integer:")
	fmt.Scanln(&input)
	result = factorial.CalculateFactorial(input)
	fmt.Println("Result: " + strconv.Itoa(result))
	fmt.Println("-------------------------------------------")
}
