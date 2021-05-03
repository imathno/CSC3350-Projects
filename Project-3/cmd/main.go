package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

func openFile(path string) (string, bool) {
	data, err := ioutil.ReadFile(path)

	if err != nil {
		fmt.Println("File Read Error", err)
		return "", true
	}
	return string(data), false
}

func main() {
	var dat string
	var err bool
	var wg sync.WaitGroup
	wg.Add(1) // Telling the sync function that we're waiting for 1 thread to finish

	// This creates a Go subroutine which is a fancy way of saying thread.
	go func() {
		dat, err = openFile(os.Args[1])
		wg.Done() // The subroutine / thread is finished doing its task
	}()
	wg.Wait() // Syncs the function call and waits for openFile to finish and return data

	if err {
		fmt.Println("Error in opening file ", os.Args[1])
	}
	fmt.Println(dat)
}
