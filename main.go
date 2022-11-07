package main

import (
	"fmt"
	"sync"
)

func printA(a string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(a)
}
func main() { // the main function itself is a go-routine; run a very lightweight thread
	var wg sync.WaitGroup

	words := []string{
		"alpha", "beta", "delta", "epsilon", "gamma", "pi", "theta", "zeta", "omega",
	}

	wg.Add(9) // use wg.Add(len(words)) if words is gonna change or populate by remote API etc
	for i, word := range words {
		go printA(fmt.Sprintf("no. %d: %s", i, word), &wg)
	}
	wg.Wait()

	wg.Add(1) // wait group shouldn't be use below zero, so add 1
	printA("Print a letter ''B'' to the monitor", &wg)
}
