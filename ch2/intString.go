package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Print provide an integer.")
		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	input := strconv.Itoa(n) // Best option IMO
	fmt.Printf("strconv.Itoa() %s of type %T\n", input, input)

	input = strconv.FormatInt(int64(n), 10)
	fmt.Printf("strconv.FormatInt() %s of type %T\n", input, input)

	input = string(n)
	fmt.Printf("string() %s of type %T\n", input, input)
}
