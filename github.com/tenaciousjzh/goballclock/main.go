package main

import (
	"fmt"
	"github.com/davecheney/profile"
	"github.com/tenaciousjzh/goballclock/clock"
	"github.com/tenaciousjzh/goballclock/validator"
	"log"
	"os"
	"time"
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
	defer profile.Start(profile.CPUProfile).Stop()
	start := time.Now()
	args := parseArgs()
	bc, err := clock.NewBallClock(args.NumBalls, args.Duration)
	if err != nil {
		log.Printf("Unable to create the Ball Clock. Error : %s", err.Error())
	}
	result := bc.RunClock()
	fmt.Println(result)
	end := time.Now()
	log.Printf("Elapsed time : %s", end.Sub(start).String())
}

func parseArgs() BallClockArgs {
	if validator.IsMissingArgs(os.Args) {
		log.Println("You must supply at least one argument of how many balls to add to the clock to proceed.")
		os.Exit(1)
	}

	ballResult := validator.ValidateBallInput(os.Args[1])
	numBalls := 0
	if ballResult.IsValid {
		numBalls = ballResult.Value
	}

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
