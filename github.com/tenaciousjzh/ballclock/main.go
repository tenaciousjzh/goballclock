package main

import (
	"fmt"
	"log"
	"os"
)

func init() {
	//Change the device for logging to stdout
	log.SetOutput(os.Stdout) //sets it from default stderr to stdout
}

func main() {
	numBalls := os.Args[1]
	duration := os.Args[2]

	if numBalls != nil {
		fmt.Printf("numBalls = %s", numBalls)
	}
	if duration != nil {
		fmt.Printf("duration + %s", duration)
	}
}
