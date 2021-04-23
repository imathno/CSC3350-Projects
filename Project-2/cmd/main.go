package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func main() {
	start_time := time.Now()

	_, err := ioutil.ReadFile(os.Args[1])

	if err != nil {
		fmt.Println("File Read Error", err)
	}

	duration := time.Since(start_time)

	fmt.Println(duration.Seconds())
}
