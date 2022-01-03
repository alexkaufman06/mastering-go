package main

import (
	"fmt"
	"time"
)

func pausingPeriod() {
	for i := 0; i < 3; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf(".")
		time.Sleep(400 * time.Millisecond)
	}
}

func main() { // Entrypoint to the application
	fmt.Printf("Hello World")

	pausingPeriod()

	fmt.Printf("\nWhat is your name? ")
	var name string
	fmt.Scanln(&name)

	fmt.Printf("Hello " + name + ", it's nice to meet you!")
	time.Sleep(1 * time.Second)
	fmt.Printf("\nWell")
	pausingPeriod()
	fmt.Printf(" Goodbye World!\n")
}
