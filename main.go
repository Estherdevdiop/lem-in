package main

import (
	"fmt"
	"lem-in/lemin"
	"os"
	"time"
)

var ()

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: go run main.go audits/file.txt")
		os.Exit(0)
	}
	startTime := time.Now()

	file := os.Args[1]
	lemin.ReadAndValidate(file)
	lemin.ValidRoms(lemin.Rooms)
	lemin.InitRoom()
	lemin.LetsMove()
	
	endTime := time.Now()
	executionTime := endTime.Sub(startTime)
	fmt.Println("Execution Time:", executionTime)

}
