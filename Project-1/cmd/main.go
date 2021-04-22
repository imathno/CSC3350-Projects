package main

import (
	"fmt"
	"time"
	"os/exec"
	"strconv"
)

func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	cmd := exec.Command("python", "../scripts/prime.py", strconv.Itoa(num))

	start_time := time.Now()
	cmd.Run()
	duration := time.Since(start_time)
	
	fmt.Println(duration.Seconds())
}