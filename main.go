package main

import (
	"fmt"
	"time"
)

func printA(a string) {
	fmt.Println(a)
}
func main() {
	// the main function itself is a go-routine; run a very lightweight thread

	words := []string{
		"alpha", "beta", "delta", "epsilon", "gamma", "pi", "theta", "zeta", "omega",
	}

	for i, word := range words {
		go printA(fmt.Sprintf("no. %d: %s", i, word))
	}

	time.Sleep(1 * time.Second) // using this method is not a good idea, change to wait group instead

	printA("Print a letter ''B'' to the monitor")
}
