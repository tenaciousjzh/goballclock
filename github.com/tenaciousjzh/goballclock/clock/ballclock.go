package clock

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/tenaciousjzh/goballclock/slice"
	"github.com/tenaciousjzh/goballclock/validator"
	"log"
)

const minTrackCapacity int = 4
const fiveMinTrackCapacity int = 11
const hourTrackCapacity = 11

type mode int

const (
	ReportDays = iota
	ReportTracks
)

type ballClock struct {
	numBalls     int
	duration     int
	ballq        []int
	orig         []int
	minTrack     []int
	fiveMinTrack []int
	hourTrack    []int
	reportMode   mode
	debug        bool
	halfDays     int
	done         bool
}

//NewBallClock is a constructor function for creating a ballClock
//object. This function returns a ballClock pointer and an error
//object if construction failed.
func NewBallClock(numBalls int, duration int) (*ballClock, error) {

	if validator.IsValidBallRange(numBalls) == false {
		return nil, errors.New(validator.InvalidBallInput)
	}
	if validator.IsValidDuration(duration) == false {
		return nil, errors.New(validator.InvalidDuration)
	}
	bc := new(ballClock)
	bc.numBalls = numBalls
	bc.duration = duration
	bc.ballq = loadBallQueue(numBalls)
	bc.orig = make([]int, numBalls)
	copy(bc.orig, bc.ballq)
	bc.minTrack = make([]int, 0)
	bc.fiveMinTrack = make([]int, 0)
	bc.hourTrack = make([]int, 0)
	bc.reportMode = ReportDays
	bc.debug = false

	if duration > 0 {
		bc.reportMode = ReportTracks
		bc.duration = duration
	}
	return bc, nil
}

func loadBallQueue(size int) []int {
	q := make([]int, size)
	for i := 0; i < len(q); i++ {
		q[i] = i + 1
	}
	return q
}

//RunClock is called once the ballClock has been successfully
//created. This function returns a message based on the processing
//performed. If one argument is provided, the number of days
//that elapse until the queue of balls returns to its original
//order will be returned.
//
//If the second argument of duration is provided, it will
//simulate the duration and print a JSON string.
func (clock *ballClock) RunClock() string {
	if clock.reportMode == ReportDays {
		return clock.simulateDays()
	} else {
		return clock.simulateDuration()
	}
}

func (clock *ballClock) simulateDays() string {
	clock.halfDays = 0
	ballVal := -1
	for clock.done == false {
		ballVal, clock.ballq = slice.Shift(clock.ballq)
		clock.updMinTrack(ballVal)
	}
	return fmt.Sprintf("%d balls cycle for %d days.", clock.numBalls, clock.halfDays/2)
}

type jsonResult struct {
	Min     []int
	FiveMin []int
	Hour    []int
	Main    []int
}

func (clock *ballClock) simulateDuration() string {
	ballVal := -1
	for i := 0; i < clock.duration; i++ {
		ballVal, clock.ballq = slice.Shift(clock.ballq)
		clock.updMinTrack(ballVal)
	}
	durationStats := jsonResult{
		Min:     clock.minTrack,
		FiveMin: clock.fiveMinTrack,
		Hour:    clock.hourTrack,
		Main:    clock.ballq,
	}

	result, err := json.Marshal(durationStats)
	if err != nil {
		log.Println("error: ", err)
	}
	return string(result)
}

func (clock *ballClock) updMinTrack(ballVal int) {
	if len(clock.minTrack) == minTrackCapacity {
		//Put the balls on the track in reverse order
		//back into the ballq
		for len(clock.minTrack) > 0 {
			lastVal := 0
			lastVal, clock.minTrack = slice.Pop(clock.minTrack)
			clock.ballq = slice.Push(clock.ballq, lastVal)
		}
		//Put the next ball from ballq into the next track
		clock.updFiveMinTrack(ballVal)
	} else {
		clock.minTrack = slice.Push(clock.minTrack, ballVal)

	}
}

func (clock *ballClock) updFiveMinTrack(ballVal int) {
	if len(clock.fiveMinTrack) == fiveMinTrackCapacity {
		for len(clock.fiveMinTrack) > 0 {
			lastVal := 0
			lastVal, clock.fiveMinTrack = slice.Pop(clock.fiveMinTrack)
			clock.ballq = slice.Push(clock.ballq, lastVal)
		}

		clock.updHourTrack(ballVal)
	} else {
		clock.fiveMinTrack = slice.Push(clock.fiveMinTrack, ballVal)
	}
}

func (clock *ballClock) updHourTrack(ballVal int) {
	if len(clock.hourTrack) == hourTrackCapacity {
		for len(clock.hourTrack) > 0 {
			lastVal := 0
			lastVal, clock.hourTrack = slice.Pop(clock.hourTrack)
			clock.ballq = slice.Push(clock.ballq, lastVal)
		}

		clock.ballq = slice.Push(clock.ballq, ballVal)
		clock.halfDays += 1
	} else {
		clock.hourTrack = slice.Push(clock.hourTrack, ballVal)
	}

	if slice.Same(clock.ballq, clock.orig) {
		clock.done = true
	}
}

//PrintDiagnostic can be used as a debugging tool to see the
//state of a ballClock object. This is generally turned on by
//setting the debug flag to true.
func (clock *ballClock) PrintDiagnostic() {
	if clock.debug {
		log.Printf("ballq: %s\n", clock.ballq)
		log.Printf("minTrack: %s\n", clock.minTrack)
		log.Printf("fiveMinTrack: %s\n", clock.fiveMinTrack)
		log.Printf("hourTrack: %s\n", clock.hourTrack)
	}
}
