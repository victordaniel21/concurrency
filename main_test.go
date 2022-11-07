package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_printA(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	var wg sync.WaitGroup
	wg.Add(1)

	go printA("delta", &wg)

	wg.Wait()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "delta") {
		t.Error("Expected delta, but not found")

	}
}
