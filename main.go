package main

import "fmt"

func printA(a string) {
	fmt.Println(a)
}
func main() {
	// the main function itself is a go-routine; run a very lightweight thread

	go printA("Print a letter ''A'' to the monitor") // add "go" will tell the compiler to execute in its own go-routine

	printA("Print a letter ''B'' to the monitor")
}
