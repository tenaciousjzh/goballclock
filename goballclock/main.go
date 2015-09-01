package main

import (
	"fmt"
	"github.com/x/goballclock/clock"
	"github.com/x/goballclock/validator"
	"log"
	"os"
)

func init() {
	//Change the device for logging to stdout
	log.SetOutput(os.Stdout) //sets it from default stderr to stdout
}

type BallClockArgs struct {
	NumBalls int
	Duration int
}

func main() {
	args := parseArgs()
	bc, err := clock.NewBallClock(args.NumBalls, args.Duration)
	if err != nil {
		log.Printf("Unable to create the Ball Clock. Error : %s", err.Error())
	}
	result := bc.RunClock()
	fmt.Println(result)
}

func parseArgs() BallClockArgs {
	ballResult := validator.ValidateBallInput(os.Args[1])
	numBalls := 0
	if ballResult.IsValid {
		numBalls = ballResult.Value
	}

	//go log.Printf("numBalls = %d\n", numBalls)

	duration := 0
	if len(os.Args) > 2 {
		durationResult := validator.ValidateDuration(os.Args[2])
		if durationResult.IsValid {
			duration = durationResult.Value
		}
	}
	//go log.Printf("duration = %d\n", duration)
	return BallClockArgs{NumBalls: numBalls, Duration: duration}
}
